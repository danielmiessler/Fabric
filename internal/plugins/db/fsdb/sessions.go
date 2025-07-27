package fsdb

import (
	"fmt"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

type SessionsEntity struct {
	*StorageEntity
}

func (o *SessionsEntity) Get(name string) (session *Session, err error) {
	session = &Session{Name: name}

	if o.Exists(name) {
		err = o.LoadAsJson(name, &session.Messages)
	} else {
		fmt.Printf("Creating new session: %s\n", name)
	}
	return
}

func (o *SessionsEntity) PrintSession(name string) (err error) {
	if o.Exists(name) {
		var session Session
		if err = o.LoadAsJson(name, &session.Messages); err == nil {
			fmt.Println(session.String())
		}
	}
	return
}

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
		ret += fmt.Sprintf("\n--- \n[%v]\n%v", message.Role, message.Content)
		if message.MultiContent != nil {
			for _, part := range message.MultiContent {
				switch part.Type {
				case chat.ChatMessagePartTypeImageURL:
					ret += fmt.Sprintf("\n%v: %v", part.Type, *part.ImageURL)
				case chat.ChatMessagePartTypeText:
					ret += fmt.Sprintf("\n%v: %v", part.Type, part.Text)
				}
			}
		}
	}
	return
}
