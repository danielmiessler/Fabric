package gemini_openai

import (
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
)

func NewClient() (ret *Client) {
	ret = &Client{}
	ret.Client = openai.NewClientCompatible("GeminiOpenAI", "https://generativelanguage.googleapis.com/v1beta", nil)
	return
}

type Client struct {
	*openai.Client
}
