package ai

import (
	"context"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/plugins"

	"github.com/danielmiessler/fabric/internal/domain"
)

type Vendor interface {
	plugins.Plugin
	ListModels() ([]string, error)
	SendStream([]*chat.ChatCompletionMessage, *domain.ChatOptions, chan string) error
	Send(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions) (string, error)
	NeedsRawMode(modelName string) bool
}
