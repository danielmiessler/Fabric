package ai

import (
	"context"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/plugins"

	"github.com/danielmiessler/fabric/internal/domain"
)

type Vendor interface {
	plugins.Plugin
	ListModels(context.Context) ([]string, error)
	SendStream(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions, chan domain.StreamUpdate) error
	Send(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions) (string, error)
	NeedsRawMode(modelName string) bool
}
