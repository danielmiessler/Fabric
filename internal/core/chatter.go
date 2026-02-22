package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"

	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/plugins/strategy"
	"github.com/danielmiessler/fabric/internal/plugins/template"
)

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
	// Use o.model (normalized) for NeedsRawMode check instead of opts.Model
	// This ensures case-insensitive model names work correctly (e.g., "GPT-5" → "gpt-5")
	if o.vendor.NeedsRawMode(o.model) {
		opts.Raw = true
	}
	if session, err = o.BuildSession(request, opts.Raw); err != nil {
		return
	}

	vendorMessages := session.GetVendorMessages()

	if debuglog.GetLevel() >= debuglog.Wire {
		debuglog.Debug(debuglog.Wire, "FABRIC->LLM request messages (%d)\n", len(vendorMessages))
		for i, msg := range vendorMessages {
			debuglog.Debug(debuglog.Wire, "FABRIC->LLM [%d] role=%s content=%q\n", i, msg.Role, msg.Content)
			if len(msg.MultiContent) > 0 {
				debuglog.Debug(debuglog.Wire, "FABRIC->LLM [%d] parts=%d\n", i, len(msg.MultiContent))
			}
		}
	}
	if len(vendorMessages) == 0 {
		if session.Name != "" {
			err = o.db.Sessions.SaveSession(session)
			if err != nil {
				return
			}
		}
		err = errors.New(i18n.T("chatter_error_no_messages_provided"))
		return
	}

	// Always use the normalized model name from the Chatter
	// This handles cases where user provides "GPT-5" but we've normalized it to "gpt-5"
	opts.Model = o.model

	if opts.ModelContextLength == 0 {
		opts.ModelContextLength = o.modelContextLength
	}

	message := ""

	if o.Stream {
		responseChan := make(chan domain.StreamUpdate)
		errChan := make(chan error, 1)
		done := make(chan struct{})
		printedStream := false

		go func() {
			defer close(done)
			if streamErr := o.vendor.SendStream(session.GetVendorMessages(), opts, responseChan); streamErr != nil {
				errChan <- streamErr
			}
		}()

		for update := range responseChan {
			if debuglog.GetLevel() >= debuglog.Wire {
				debuglog.Debug(debuglog.Wire, "LLM->FABRIC stream update type=%s content=%q\n", update.Type, update.Content)
				if update.Usage != nil {
					debuglog.Debug(debuglog.Wire, "LLM->FABRIC stream usage input=%d output=%d total=%d\n", update.Usage.InputTokens, update.Usage.OutputTokens, update.Usage.TotalTokens)
				}
			}
			if opts.UpdateChan != nil {
				opts.UpdateChan <- update
			}
			switch update.Type {
			case domain.StreamTypeContent:
				message += update.Content
				if !opts.SuppressThink && !opts.Quiet {
					fmt.Print(update.Content)
					printedStream = true
				}
			case domain.StreamTypeUsage:
				if opts.ShowMetadata && update.Usage != nil && !opts.Quiet {
					fmt.Fprintf(
						os.Stderr,
						"\n%s\n",
						fmt.Sprintf(
							i18n.T("chatter_log_stream_usage_metadata"),
							update.Usage.InputTokens,
							update.Usage.OutputTokens,
							update.Usage.TotalTokens,
						),
					)
				}
			case domain.StreamTypeError:
				if !opts.Quiet {
					fmt.Fprintf(os.Stderr, "%s\n", fmt.Sprintf(i18n.T("chatter_error_stream_update"), update.Content))
				}
				errChan <- errors.New(update.Content)
			}
		}

		if printedStream && !opts.SuppressThink && !strings.HasSuffix(message, "\n") && !opts.Quiet {
			fmt.Println()
		}

		// Wait for goroutine to finish
		<-done

		// Check for errors in errChan
		select {
		case streamErr := <-errChan:
			if streamErr != nil {
				err = streamErr
				return
			}
		default:
			// No errors, continue
		}
	} else {
		if message, err = o.vendor.Send(context.Background(), session.GetVendorMessages(), opts); err != nil {
			return
		}
		if debuglog.GetLevel() >= debuglog.Wire {
			debuglog.Debug(debuglog.Wire, "LLM->FABRIC response content=%q\n", message)
		}
	}

	if opts.SuppressThink && !o.DryRun {
		message = domain.StripThinkBlocks(message, opts.ThinkStartTag, opts.ThinkEndTag)
	}

	if message == "" {
		session = nil
		err = errors.New(i18n.T("chatter_error_empty_response"))
		return
	}

	// Process file changes for create_coding_feature pattern
	if request.PatternName == "create_coding_feature" {
		summary, fileChanges, parseErr := domain.ParseFileChanges(message)
		if parseErr != nil {
			fmt.Printf("%s\n", fmt.Sprintf(i18n.T("chatter_warning_parse_file_changes_failed"), parseErr))
		} else if len(fileChanges) > 0 {
			projectRoot, err := os.Getwd()
			if err != nil {
				fmt.Printf("%s\n", fmt.Sprintf(i18n.T("chatter_warning_get_current_directory_failed"), err))
			} else {
				if applyErr := domain.ApplyFileChanges(projectRoot, fileChanges); applyErr != nil {
					fmt.Printf("%s\n", fmt.Sprintf(i18n.T("chatter_warning_apply_file_changes_failed"), applyErr))
				} else {
					fmt.Println(i18n.T("chatter_info_file_changes_applied_successfully"))
					fmt.Printf("%s\n\n", i18n.T("chatter_help_review_changes_with_git_diff"))
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

func (o *Chatter) BuildSession(request *domain.ChatRequest, raw bool) (session *fsdb.Session, err error) {
	if request.SessionName != "" {
		var sess *fsdb.Session
		if sess, err = o.db.Sessions.Get(request.SessionName); err != nil {
			err = fmt.Errorf(i18n.T("chatter_error_find_session"), request.SessionName, err)
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
			err = fmt.Errorf(i18n.T("chatter_error_find_context"), request.ContextName, err)
			return
		}
		contextContent = ctx.Content
	}

	// Process template variables in message content
	// Double curly braces {{variable}} indicate template substitution
	// Ensure we have a message before processing
	if request.Message == nil {
		request.Message = &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "",
		}
	}

	// Now we know request.Message is not nil, process template variables
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
			return nil, fmt.Errorf(i18n.T("chatter_error_get_pattern"), request.PatternName, err)
		}
		patternContent = pattern.Pattern
		inputUsed = true
	}

	systemMessage := strings.TrimSpace(contextContent) + strings.TrimSpace(patternContent)

	if request.StrategyName != "" {
		strategy, err := strategy.LoadStrategy(request.StrategyName)
		if err != nil {
			return nil, fmt.Errorf(i18n.T("chatter_error_load_strategy"), request.StrategyName, err)
		}
		if strategy != nil && strategy.Prompt != "" {
			// prepend the strategy prompt to the system message
			systemMessage = fmt.Sprintf("%s\n%s", strategy.Prompt, systemMessage)
		}
	}

	// Apply refined language instruction if specified
	if request.Language != "" && request.Language != "en" {
		// Refined instruction: Execute pattern using user input, then translate the entire response.
		systemMessage = fmt.Sprintf(i18n.T("chatter_prompt_enforce_response_language"), systemMessage, request.Language)
	}

	if raw {
		var finalContent string
		if systemMessage != "" {
			if request.PatternName != "" {
				finalContent = systemMessage
			} else {
				finalContent = fmt.Sprintf("%s\n\n%s", systemMessage, request.Message.Content)
			}

			// Handle MultiContent properly in raw mode
			if len(request.Message.MultiContent) > 0 {
				// When we have attachments, add the text as a text part in MultiContent
				newMultiContent := []chat.ChatMessagePart{
					{
						Type: chat.ChatMessagePartTypeText,
						Text: finalContent,
					},
				}
				// Add existing non-text parts (like images)
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
				// No attachments, use regular Content field
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
		// If multi-part content, it is in the user message, and should be added.
		// Otherwise, we should only add it if we have not already used it in the systemMessage.
		if len(request.Message.MultiContent) > 0 || (request.Message != nil && !inputUsed) {
			session.Append(request.Message)
		}
	}

	if session.IsEmpty() {
		session = nil
		err = errors.New(i18n.T("chatter_error_no_session_pattern_user_messages"))
	}
	return
}
