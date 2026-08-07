package restapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	registry *core.PluginRegistry
	db       *fsdb.Db
}

type PromptRequest struct {
	UserInput    string            `json:"userInput"`
	Vendor       string            `json:"vendor"`
	Model        string            `json:"model"`
	ContextName  string            `json:"contextName"`
	PatternName  string            `json:"patternName"`
	StrategyName string            `json:"strategyName"`        // Optional strategy name
	SessionName  string            `json:"sessionName"`         // Session name for multi-turn conversations
	Variables    map[string]string `json:"variables,omitempty"` // Pattern variables
}

type ChatRequest struct {
	Prompts            []PromptRequest `json:"prompts"`
	Language           string          `json:"language"`
	ModelContextLength int             `json:"modelContextLength,omitempty"` // Context window size
	domain.ChatOptions                 // Embed the ChatOptions from common package
}

type StreamResponse struct {
	Type    string                `json:"type"`             // "content", "usage", "error", "complete"
	Format  string                `json:"format,omitempty"` // "markdown", "mermaid", "plain"
	Content string                `json:"content,omitempty"`
	Usage   *domain.UsageMetadata `json:"usage,omitempty"`
}

func NewChatHandler(r *gin.Engine, registry *core.PluginRegistry, db *fsdb.Db) *ChatHandler {
	handler := &ChatHandler{
		registry: registry,
		db:       db,
	}

	r.POST("/chat", handler.HandleChat)

	return handler
}

// HandleChat godoc
// @Summary Stream chat completions
// @Description Stream AI responses using Server-Sent Events (SSE)
// @Tags chat
// @Accept json
// @Produce text/event-stream
// @Param request body ChatRequest true "Chat request with prompts and options"
// @Success 200 {object} StreamResponse "Streaming response"
// @Failure 400 {object} map[string]string
// @Security ApiKeyAuth
// @Router /chat [post]
func (h *ChatHandler) HandleChat(c *gin.Context) {
	var request ChatRequest

	if err := c.BindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(i18n.T("server_invalid_request_format"), err)})
		return
	}

	// Add log to check received language field
	log.Printf("Received chat request - Language: '%s', Prompts: %d", request.Language, len(request.Prompts))

	// Set headers for SSE
	c.Writer.Header().Set("Content-Type", "text/readystream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Writer.Header().Set("X-Accel-Buffering", "no")

	clientGone := c.Writer.CloseNotify()

	for i, prompt := range request.Prompts {
		select {
		case <-clientGone:
			log.Printf("Client disconnected")
			return
		default:
			log.Printf("Processing prompt %d: Model=%s Pattern=%s Context=%s",
				i+1, prompt.Model, prompt.PatternName, prompt.ContextName)

			streamChan := make(chan domain.StreamUpdate)
			// Send returns an error for a failure that happens before it starts
			// to stream, such as a pattern that needs a variable the request
			// does not give. Such an error never reaches streamChan, so keep it
			// here and report it after the stream ends.
			sendErrChan := make(chan error, 1)

			go func(p PromptRequest) {
				defer close(streamChan)

				chatter, err := h.registry.GetChatter(
					p.Model,
					request.ModelContextLength,
					p.Vendor,
					true,
					false,
				)
				if err != nil {
					log.Printf("Error creating chatter: %v", err)
					streamChan <- domain.StreamUpdate{Type: domain.StreamTypeError, Content: fmt.Sprintf(i18n.T("server_chat_error"), err)}
					return
				}

				chatReq := buildPromptChatRequest(p, request.Language)

				opts := &domain.ChatOptions{
					Model:            p.Model,
					Temperature:      request.Temperature,
					TopP:             request.TopP,
					FrequencyPenalty: request.FrequencyPenalty,
					PresencePenalty:  request.PresencePenalty,
					Thinking:         request.Thinking,
					Search:           request.Search,
					SearchLocation:   request.SearchLocation,
					UpdateChan:       streamChan,
					Quiet:            true,
				}

				_, err = chatter.Send(c.Request.Context(), chatReq, opts)
				if err != nil {
					log.Printf("Error from chatter.Send: %v", err)
					sendErrChan <- err
					return
				}
			}(prompt)

			sawError := false
			for update := range streamChan {
				select {
				case <-clientGone:
					return
				default:
					var response StreamResponse
					switch update.Type {
					case domain.StreamTypeContent:
						response = StreamResponse{
							Type:    "content",
							Format:  detectFormat(update.Content),
							Content: update.Content,
						}
					case domain.StreamTypeUsage:
						response = StreamResponse{
							Type:  "usage",
							Usage: update.Usage,
						}
					case domain.StreamTypeError:
						sawError = true
						response = StreamResponse{
							Type:    "error",
							Format:  "plain",
							Content: update.Content,
						}
					}

					if err := writeSSEResponse(c.Writer, response); err != nil {
						log.Printf("Error writing response: %v", err)
						return
					}
				}
			}

			// The goroutine writes sendErrChan before it closes streamChan, so
			// the value is here by the time the loop above ends.
			if response, ok := unreportedSendError(sendErrChan, sawError); ok {
				if err := writeSSEResponse(c.Writer, response); err != nil {
					log.Printf("Error writing error response: %v", err)
					return
				}
			}

			completeResponse := StreamResponse{
				Type:    "complete",
				Format:  "plain",
				Content: "",
			}
			if err := writeSSEResponse(c.Writer, completeResponse); err != nil {
				log.Printf("Error writing completion response: %v", err)
				return
			}
		}
	}
}

// unreportedSendError reads the error that Send returned and makes the response
// that tells the client about it. The second result is false when there is
// nothing to send.
//
// Send both gives a stream error to the update channel and returns it, so the
// client already has that error and must not get it twice. The sawError
// argument says whether an error went to the client during the stream. An error
// that Send makes before or after the stream, such as a pattern that needs a
// variable the request does not give, never goes to the update channel, and
// this function is the only way the client learns about it.
func unreportedSendError(sendErrChan <-chan error, sawError bool) (StreamResponse, bool) {
	select {
	case sendErr := <-sendErrChan:
		if sendErr == nil || sawError {
			return StreamResponse{}, false
		}
		return StreamResponse{
			Type:    "error",
			Format:  "plain",
			Content: fmt.Sprintf(i18n.T("server_chat_error"), sendErr),
		}, true
	default:
		return StreamResponse{}, false
	}
}

func buildPromptChatRequest(p PromptRequest, language string) *domain.ChatRequest {
	return &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: p.UserInput,
		},
		PatternName:      p.PatternName,
		ContextName:      p.ContextName,
		SessionName:      p.SessionName,
		PatternVariables: p.Variables,
		StrategyName:     p.StrategyName,
		Language:         language,
	}
}

func writeSSEResponse(w gin.ResponseWriter, response StreamResponse) error {
	data, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("%s", fmt.Sprintf(i18n.T("server_error_marshaling_response"), err))
	}

	if _, err := fmt.Fprintf(w, "data: %s\n\n", string(data)); err != nil {
		return fmt.Errorf("%s", fmt.Sprintf(i18n.T("server_error_writing_response"), err))
	}

	w.(http.Flusher).Flush()
	return nil
}

func detectFormat(content string) string {
	if strings.HasPrefix(content, "graph TD") ||
		strings.HasPrefix(content, "gantt") ||
		strings.HasPrefix(content, "flowchart") ||
		strings.HasPrefix(content, "sequenceDiagram") ||
		strings.HasPrefix(content, "classDiagram") ||
		strings.HasPrefix(content, "stateDiagram") {
		return "mermaid"
	}
	return "markdown"
}
