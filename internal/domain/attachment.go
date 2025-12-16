package domain

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

type Attachment struct {
	Type    *string `json:"type,omitempty"`
	Path    *string `json:"path,omitempty"`
	URL     *string `json:"url,omitempty"`
	Content []byte  `json:"content,omitempty"`
	ID      *string `json:"id,omitempty"`
}

func (a *Attachment) GetId() (ret string, err error) {
	if a.ID == nil {
		var hash string
		if a.Content != nil {
			hash = fmt.Sprintf("%x", sha256.Sum256(a.Content))
		} else if a.Path != nil {
			var content []byte
			if content, err = os.ReadFile(*a.Path); err != nil {
				return ret, err
			}
			hash = fmt.Sprintf("%x", sha256.Sum256(content))
		} else if a.URL != nil {
			data := map[string]string{"url": *a.URL}
			var jsonData []byte
			if jsonData, err = json.Marshal(data); err != nil {
				return ret, err
			}
			hash = fmt.Sprintf("%x", sha256.Sum256(jsonData))
		}
		a.ID = &hash
	}
	ret = *a.ID
	return ret, err
}

func (a *Attachment) ResolveType() (ret string, err error) {
	if a.Type != nil {
		ret = *a.Type
		return ret, err
	}
	if a.Path != nil {
		var mime *mimetype.MIME
		if mime, err = mimetype.DetectFile(*a.Path); err != nil {
			return ret, err
		}
		ret = mime.String()
		return ret, err
	}
	if a.URL != nil {
		var resp *http.Response
		if resp, err = http.Head(*a.URL); err != nil {
			return ret, err
		}
		defer resp.Body.Close()
		ret = resp.Header.Get("Content-Type")
		return ret, err
	}
	if a.Content != nil {
		ret = mimetype.Detect(a.Content).String()
		return ret, err
	}
	err = fmt.Errorf("attachment has no type and no content to derive it from")
	return ret, err
}

func (a *Attachment) ContentBytes() (ret []byte, err error) {
	if a.Content != nil {
		ret = a.Content
		return ret, err
	}
	if a.Path != nil {
		if ret, err = os.ReadFile(*a.Path); err != nil {
			return ret, err
		}
		return ret, err
	}
	if a.URL != nil {
		var resp *http.Response
		if resp, err = http.Get(*a.URL); err != nil {
			return ret, err
		}
		defer resp.Body.Close()
		if ret, err = io.ReadAll(resp.Body); err != nil {
			return ret, err
		}
		return ret, err
	}
	err = fmt.Errorf("no content available")
	return ret, err
}

func (a *Attachment) Base64Content() (ret string, err error) {
	var content []byte
	if content, err = a.ContentBytes(); err != nil {
		return ret, err
	}
	ret = base64.StdEncoding.EncodeToString(content)
	return ret, err
}

func NewAttachment(value string) (ret *Attachment, err error) {
	if isURL(value) {
		var mimeType string
		if mimeType, err = detectMimeTypeFromURL(value); err != nil {
			return ret, err
		}
		ret = &Attachment{
			Type: &mimeType,
			URL:  &value,
		}
		return ret, err
	}

	var absPath string
	if absPath, err = filepath.Abs(value); err != nil {
		return ret, err
	}
	if _, err = os.Stat(absPath); os.IsNotExist(err) {
		err = fmt.Errorf("file %s does not exist", value)
		return ret, err
	}

	var mimeType string
	if mimeType, err = detectMimeTypeFromFile(absPath); err != nil {
		return ret, err
	}
	ret = &Attachment{
		Type: &mimeType,
		Path: &absPath,
	}
	return ret, err
}

func detectMimeTypeFromURL(url string) (string, error) {
	resp, err := http.Head(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	mimeType := resp.Header.Get("Content-Type")
	if mimeType == "" {
		return "", fmt.Errorf("could not determine mimetype of URL")
	}
	return mimeType, nil
}

func detectMimeTypeFromFile(path string) (string, error) {
	mime, err := mimetype.DetectFile(path)
	if err != nil {
		return "", err
	}
	return mime.String(), nil
}

func isURL(value string) bool {
	return bytes.Contains([]byte(value), []byte("://"))
}
