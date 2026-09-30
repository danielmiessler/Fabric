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
}

// recordFirstStreamError sends err to errChan when the channel has space. It discards later errors.
func recordFirstStreamError(errChan chan error, err error) {
	if err == nil {
		return
	}

	select {
	case errChan <- err:
	default:
		debuglog.Debug(debuglog.Wire, "additional stream error discarded: %v\n", err)
	}
}

// joinPromptSections trims each part, drops empty ones, and joins the rest with newline separators.
func joinPromptSections(parts ...string) string {
	sections := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			sections = append(sections, trimmed)
		}
	}

	return strings.Join(sections, "\n")
}

// Send processes a chat request and applies file changes for create_coding_feature pattern
func (o *Chatter) Send(ctx context.Context, request *domain.ChatRequest, opts *domain.ChatOptions) (session *fsdb.Session, err error) {
	// Test o.model, not opts.Model. GetChatter set o.model to the vendor's spelling of the name.
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

	// Send the vendor's spelling of the model name, not the one the user typed.
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
			if streamErr := o.vendor.SendStream(ctx, session.GetVendorMessages(), opts, responseChan); streamErr != nil {
				recordFirstStreamError(errChan, streamErr)
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
				if !opts.SuppressThink && !opts.BufferStream && !opts.Quiet {
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
				recordFirstStreamError(errChan, errors.New(update.Content))
			}
		}

		if printedStream && !opts.SuppressThink && !strings.HasSuffix(message, "\n") && !opts.Quiet {
			fmt.Println()
		}

		<-done

		select {
		case streamErr := <-errChan:
			if streamErr != nil {
				err = streamErr
				return
			}
		default:
		}
	} else {
		if message, err = o.vendor.Send(ctx, session.GetVendorMessages(), opts); err != nil {
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

	var contextContent string
	if request.ContextName != "" {
		var ctx *fsdb.Context
		if ctx, err = o.db.Contexts.Get(request.ContextName); err != nil {
			err = fmt.Errorf(i18n.T("chatter_error_find_context"), request.ContextName, err)
			return
		}
		contextContent = ctx.Content
	}

	if request.Message == nil {
		request.Message = &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "",
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
			return nil, fmt.Errorf(i18n.T("chatter_error_get_pattern"), request.PatternName, err)
		}
		patternContent = pattern.Pattern
		inputUsed = pattern.InputUsed
	}

	systemMessage := joinPromptSections(contextContent, patternContent)

	if request.StrategyName != "" {
		strategy, err := strategy.LoadStrategy(request.StrategyName)
		if err != nil {
			return nil, fmt.Errorf(i18n.T("chatter_error_load_strategy"), request.StrategyName, err)
		}
		if strategy != nil && strategy.Prompt != "" {
			systemMessage = joinPromptSections(strategy.Prompt, systemMessage)
		}
	}

	if request.Language != "" && request.Language != "en" {
		// The prompt tells the model to run the instructions first, then write the full response in request.Language.
		systemMessage = fmt.Sprintf(i18n.T("chatter_prompt_enforce_response_language"), systemMessage, request.Language)
	}

	// The request must end with a user message: some backends reject a
	// request with system messages only. The input goes to the model one time.
	msg := request.Message
	hasInput := msg.Content != "" || len(msg.MultiContent) > 0
	if !raw && !inputUsed && hasInput {
		// The usual shape: instructions in the system message, input in the user message.
		if systemMessage != "" {
			session.Append(&chat.ChatCompletionMessage{Role: chat.ChatMessageRoleSystem, Content: systemMessage})
		}
		session.Append(msg)
	} else {
		// Raw mode, a pattern that contains the input, or no input:
		// send all of the text in one user message.
		text := systemMessage
		if !inputUsed {
			text = joinPromptSections(systemMessage, msg.Content)
		}
		merged := &chat.ChatCompletionMessage{Role: chat.ChatMessageRoleUser, Content: text}
		if len(msg.MultiContent) > 0 {
			// Attachments: the text goes first, then the parts from the request.
			merged.Content = ""
			if text != "" {
				merged.MultiContent = []chat.ChatMessagePart{{Type: chat.ChatMessagePartTypeText, Text: text}}
			}
			merged.MultiContent = append(merged.MultiContent, msg.MultiContent...)
		}
		session.Append(merged)
	}

	if session.IsEmpty() {
		session = nil
		err = errors.New(i18n.T("chatter_error_no_session_pattern_user_messages"))
	}
	return
}
