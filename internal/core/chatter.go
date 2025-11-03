package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/plugins/strategy"
	"github.com/danielmiessler/fabric/internal/plugins/template"
)

const NoSessionPatternUserMessages = "no session, pattern or user messages provided"

type Chatter struct {
	db *fsdb.Db

	Stream bool
	DryRun bool

	model              string
	modelContextLength int
	vendor             ai.Vendor
	strategy           string
}

// Send processes a chat request and applies file changes for create_coding_feature pattern
func (o *Chatter) Send(request *domain.ChatRequest, opts *domain.ChatOptions) (session *fsdb.Session, err error) {
	modelToUse := opts.Model
	if modelToUse == "" {
		modelToUse = o.model
	}
	if o.vendor.NeedsRawMode(modelToUse) {
		opts.Raw = true
	}
	if session, err = o.BuildSession(request, opts.Raw); err != nil {
		return
	}

	vendorMessages := session.GetVendorMessages()
	if len(vendorMessages) == 0 {
		if session.Name != "" {
			err = o.db.Sessions.SaveSession(session)
			if err != nil {
				return
			}
		}
		err = fmt.Errorf("no messages provided")
		return
	}

	if opts.Model == "" {
		opts.Model = o.model
	}

	if opts.ModelContextLength == 0 {
		opts.ModelContextLength = o.modelContextLength
	}

	message := ""

	// --- FIXED STREAM HANDLING START ---
	if o.Stream {
		responseChan := make(chan string)
		errChan := make(chan error, 1)

		go func() {
			defer close(responseChan)
			if streamErr := o.vendor.SendStream(session.GetVendorMessages(), opts, responseChan); streamErr != nil {
				errChan <- streamErr
			}
			close(errChan)
		}()

		for response := range responseChan {
			message += response
			if !opts.SuppressThink {
				fmt.Print(response)
			}
		}

		// Wait for potential errors after streaming finishes
		if streamErr, ok := <-errChan; ok && streamErr != nil {
			err = streamErr
			return
		}
	} else {
		if message, err = o.vendor.Send(context.Background(), session.GetVendorMessages(), opts); err != nil {
			return
		}
	}
	// --- FIXED STREAM HANDLING END ---

	if opts.SuppressThink && !o.DryRun {
		message = domain.StripThinkBlocks(message, opts.ThinkStartTag, opts.ThinkEndTag)
	}

	if message == "" {
		session = nil
		err = fmt.Errorf("empty response")
		return
	}

	// Process file changes for create_coding_feature pattern
	if request.PatternName == "create_coding_feature" {
		summary, fileChanges, parseErr := domain.ParseFileChanges(message)
		if parseErr != nil {
			fmt.Printf("Warning: Failed to parse file changes: %v\n", parseErr)
		} else if len(fileChanges) > 0 {
			projectRoot, err := os.Getwd()
			if err != nil {
				fmt.Printf("Warning: Failed to get current directory: %v\n", err)
			} else {
				if applyErr := domain.ApplyFileChanges(projectRoot, fileChanges); applyErr != nil {
					fmt.Printf("Warning: Failed to apply file changes: %v\n", applyErr)
				} else {
					fmt.Println("Successfully applied file changes.")
					fmt.Printf("You can review the changes with 'git diff' if you're using git.\n\n")
				}
			}
		}
		message = summary
	}

	session.Append(&chat.ChatCompletionMessage{Role: chat.ChatMessageRoleAssistant, Content: message})

	if session.Name != "" {
		err = o.db.Sessions.SaveSession(session)
	}
	return
}

// BuildSession builds the chat session from the provided request
func (o *Chatter) BuildSession(request *domain.ChatRequest, raw bool) (session *fsdb.Session, err error) {
	if request.SessionName != "" {
		var sess *fsdb.Session
		if sess, err = o.db.Sessions.Get(request.SessionName); err != nil {
			err = fmt.Errorf("could not find session %s: %v", request.SessionName, err)
			return
		}
		session = sess
	} else {
		session = &fsdb.Session{}
	}

	if request.Meta != "" {
		session.Append(&chat.ChatCompletionMessage{Role: domain.ChatMessageRoleMeta, Content: request.Meta})
	}

	// if a context name is provided, retrieve it from the database
	var contextContent string
	if request.ContextName != "" {
		var ctx *fsdb.Context
		if ctx, err = o.db.Contexts.Get(request.ContextName); err != nil {
			err = fmt.Errorf("could not find context %s: %v", request.ContextName, err)
			return
		}
		contextContent = ctx.Content
	}

	// Process template variables in message content
	if request.Message == nil {
		request.Message = &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: " ",
		}
	}

	if request.InputHasVars && !request.NoVariableReplacement {
		request.Message.Content, err = template.ApplyTemplate(request.Message.Content, request.PatternVariables, "")
		if err != nil {
			return nil, err
		}
	}

	var patternContent string
	inputUsed := false
	if request.PatternName != "" {
		var pattern *fsdb.Pattern
		if request.NoVariableReplacement {
			pattern, err = o.db.Patterns.GetWithoutVariables(request.PatternName, request.Message.Content)
		} else {
			pattern, err = o.db.Patterns.GetApplyVariables(request.PatternName, request.PatternVariables, request.Message.Content)
		}

		if err != nil {
			return nil, fmt.Errorf("could not get pattern %s: %v", request.PatternName, err)
		}
		patternContent = pattern.Pattern
		inputUsed = true
	}

	systemMessage := strings.TrimSpace(contextContent) + strings.TrimSpace(patternContent)

	if request.StrategyName != "" {
		strategy, err := strategy.LoadStrategy(request.StrategyName)
		if err != nil {
			return nil, fmt.Errorf("could not load strategy %s: %v", request.StrategyName, err)
		}
		if strategy != nil && strategy.Prompt != "" {
			systemMessage = fmt.Sprintf("%s\n%s", strategy.Prompt, systemMessage)
		}
	}

	if request.Language != "" && request.Language != "en" {
		systemMessage = fmt.Sprintf("%s\n\nIMPORTANT: First, execute the instructions provided in this prompt using the user's input. Second, ensure your entire final response, including any section headers or titles generated as part of executing the instructions, is written ONLY in the %s language.", systemMessage, request.Language)
	}

	if raw {
		var finalContent string
		if systemMessage != "" {
			if request.PatternName != "" {
				finalContent = systemMessage
			} else {
				finalContent = fmt.Sprintf("%s\n\n%s", systemMessage, request.Message.Content)
			}

			if len(request.Message.MultiContent) > 0 {
				newMultiContent := []chat.ChatMessagePart{
					{
						Type: chat.ChatMessagePartTypeText,
						Text: finalContent,
					},
				}
				for _, part := range request.Message.MultiContent {
					if part.Type != chat.ChatMessagePartTypeText {
						newMultiContent = append(newMultiContent, part)
					}
				}
				request.Message = &chat.ChatCompletionMessage{
					Role:         chat.ChatMessageRoleUser,
					MultiContent: newMultiContent,
				}
			} else {
				request.Message = &chat.ChatCompletionMessage{
					Role:    chat.ChatMessageRoleUser,
					Content: finalContent,
				}
			}
		}
		if request.Message != nil {
			session.Append(request.Message)
		}
	} else {
		if systemMessage != "" {
			session.Append(&chat.ChatCompletionMessage{Role: chat.ChatMessageRoleSystem, Content: systemMessage})
		}
		if len(request.Message.MultiContent) > 0 || (request.Message != nil && !inputUsed) {
			session.Append(request.Message)
		}
	}

	if session.IsEmpty() {
		session = nil
		err = errors.New(NoSessionPatternUserMessages)
	}
	return
}
