package fsdb

import (
	"fmt"
	"sync"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
)

type SessionsEntity struct {
	*StorageEntity

	// locks has one mutex for each session name. Without it, two /chat
	// calls on the same session can both read the file, append and save.
	// Then the second save removes the messages of the first call.
	// core.Chatter.Send holds the lock from the read until the save, thus
	// for the full model call. A second call on the same session waits
	// for that time. The lock is only in this process. Two fabric
	// processes on the same session can still lose messages. The map
	// keeps one mutex for each session name until the process stops.
	locks sync.Map // name -> *sync.Mutex
}

// Lock locks the named session and returns the function that unlocks it.
// Hold the lock from Get until SaveSession.
func (o *SessionsEntity) Lock(name string) (unlock func()) {
	m, _ := o.locks.LoadOrStore(name, &sync.Mutex{})
	mu := m.(*sync.Mutex)
	mu.Lock()
	return mu.Unlock
}

// Save, Delete and Rename take the session lock. Thus a REST call cannot
// write a session file while a /chat call on that session is in progress.
// A call with an invalid name does not get a lock, so that a client cannot
// add a mutex to the map for each invalid name.

func (o *SessionsEntity) Save(name string, content []byte) (err error) {
	if err = ValidateStorageName(name); err != nil {
		return
	}
	defer o.Lock(name)()
	return o.StorageEntity.Save(name, content)
}

func (o *SessionsEntity) Delete(name string) (err error) {
	if err = ValidateStorageName(name); err != nil {
		return
	}
	defer o.Lock(name)()
	return o.StorageEntity.Delete(name)
}

func (o *SessionsEntity) Rename(oldName, newName string) (err error) {
	if err = ValidateStorageName(oldName); err != nil {
		return
	}
	if err = ValidateStorageName(newName); err != nil {
		return
	}
	// Lock the two names in the same order each time, so that two renames
	// in opposite directions cannot wait for each other.
	first, second := min(oldName, newName), max(oldName, newName)
	defer o.Lock(first)()
	if second != first {
		defer o.Lock(second)()
	}
	return o.StorageEntity.Rename(oldName, newName)
}

func (o *SessionsEntity) Get(name string) (session *Session, err error) {
	return o.GetWithNotice(name, true)
}

func (o *SessionsEntity) GetWithNotice(name string, announceNewSession bool) (session *Session, err error) {
	// Reject invalid names here. Exists reports false for them, and the
	// missing-session branch then answers with a new empty session and
	// no error.
	if err = ValidateStorageName(name); err != nil {
		return nil, err
	}
	session = &Session{Name: name}

	if o.Exists(name) {
		err = o.LoadAsJson(name, &session.Messages)
	} else if announceNewSession {
		fmt.Printf(i18n.T("sessions_creating_new"), name)
	}
	return
}

func (o *SessionsEntity) PrintSession(name string) (err error) {
	if o.Exists(name) {
		var session Session
		if err = o.LoadAsJson(name, &session.Messages); err == nil {
			// The session keeps the full model replies. Remove terminal
			// control sequences from the displayed copy only.
			fmt.Println(domain.SanitizeTerminalOutput(session.String()))
		}
	}
	return
}

// SaveSession writes the session to its file. The caller must hold Lock
// for the session from Get until SaveSession. Lock uses the session name
// as its key. Thus two names for one file, for example "Work" and "work"
// on a file system that ignores case, do not share a lock.
func (o *SessionsEntity) SaveSession(session *Session) (err error) {
	return o.SaveAsJson(session.Name, session.Messages)
}

type Session struct {
	Name     string
	Messages []*chat.ChatCompletionMessage

	vendorMessages []*chat.ChatCompletionMessage
}

func (o *Session) IsEmpty() bool {
	return len(o.Messages) == 0
}

func (o *Session) Append(messages ...*chat.ChatCompletionMessage) {
	if o.vendorMessages != nil {
		for _, message := range messages {
			o.Messages = append(o.Messages, message)
			o.appendVendorMessage(message)
		}
	} else {
		o.Messages = append(o.Messages, messages...)
	}
}

func (o *Session) GetVendorMessages() (ret []*chat.ChatCompletionMessage) {
	if len(o.vendorMessages) == 0 {
		for _, message := range o.Messages {
			o.appendVendorMessage(message)
		}
	}
	ret = o.vendorMessages
	return
}

func (o *Session) appendVendorMessage(message *chat.ChatCompletionMessage) {
	// A session file can contain a JSON null entry. It loads as a nil
	// message. Skip it, because message.Role on nil stops the server.
	if message == nil {
		return
	}
	if message.Role != domain.ChatMessageRoleMeta {
		o.vendorMessages = append(o.vendorMessages, message)
	}
}

func (o *Session) GetLastMessage() (ret *chat.ChatCompletionMessage) {
	if len(o.Messages) > 0 {
		ret = o.Messages[len(o.Messages)-1]
	}
	return
}

func (o *Session) String() (ret string) {
	for _, message := range o.Messages {
		if message == nil {
			continue
		}
		ret += fmt.Sprintf("\n--- \n[%v]\n%v", message.Role, message.Content)
		if message.MultiContent != nil {
			for _, part := range message.MultiContent {
				switch part.Type {
				case chat.ChatMessagePartTypeImageURL:
					if part.ImageURL != nil {
						ret += fmt.Sprintf("\n%v: %v", part.Type, *part.ImageURL)
					}
				case chat.ChatMessagePartTypeText:
					ret += fmt.Sprintf("\n%v: %v", part.Type, part.Text)
				}
			}
		}
	}
	return
}
