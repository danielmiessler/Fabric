package restapi

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/gin-gonic/gin"
)

type OllamaModel struct {
	Models []Model `json:"models"`
}
type Model struct {
	Details    ModelDetails `json:"details"`
	Digest     string       `json:"digest"`
	Model      string       `json:"model"`
	ModifiedAt string       `json:"modified_at"`
	Name       string       `json:"name"`
	Size       int64        `json:"size"`
}

type ModelDetails struct {
	Families          []string `json:"families"`
	Family            string   `json:"family"`
	Format            string   `json:"format"`
	ParameterSize     string   `json:"parameter_size"`
	ParentModel       string   `json:"parent_model"`
	QuantizationLevel string   `json:"quantization_level"`
}

type APIConvert struct {
	registry *core.PluginRegistry
	r        *gin.Engine
	addr     *string
}

type OllamaRequestBody struct {
	Messages  []OllamaMessage   `json:"messages"`
	Model     string            `json:"model"`
	Options   map[string]any    `json:"options,omitempty"`
	Stream    bool              `json:"stream"`
	Variables map[string]string `json:"variables,omitempty"` // Fabric-specific: pattern variables (direct)
}

type OllamaMessage struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type OllamaResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	DoneReason         string `json:"done_reason,omitempty"`
	Done               bool   `json:"done"`
	TotalDuration      int64  `json:"total_duration,omitempty"`
	LoadDuration       int64  `json:"load_duration,omitempty"`
	PromptEvalCount    int64  `json:"prompt_eval_count,omitempty"`
	PromptEvalDuration int64  `json:"prompt_eval_duration,omitempty"`
	EvalCount          int64  `json:"eval_count,omitempty"`
	EvalDuration       int64  `json:"eval_duration,omitempty"`
}

type FabricResponseFormat struct {
	Type    string `json:"type"`
	Format  string `json:"format"`
	Content string `json:"content"`
}

// parseOllamaNumCtx extracts and validates the num_ctx parameter from Ollama request options.
// Returns:
//   - (0, nil) if num_ctx is not present or is null
//   - (n, nil) if num_ctx is a valid positive integer
//   - (0, error) if num_ctx is present but invalid
func parseOllamaNumCtx(options map[string]any) (int, error) {
	if options == nil {
		return 0, nil
	}

	val, exists := options["num_ctx"]
	if !exists {
		return 0, nil // Not provided, caller should use default
	}

	if val == nil {
		return 0, nil // Explicit null, treat as not provided
	}

	var contextLength int

	// Platform-specific max int value for overflow checks
	const maxInt = int64(^uint(0) >> 1)

	switch v := val.(type) {
	case float64:
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return 0, errors.New(i18n.T("ollama_num_ctx_must_be_finite"))
		}
		if math.Trunc(v) != v {
			return 0, errors.New(i18n.T("ollama_num_ctx_must_be_integer"))
		}
		// Check for overflow on 32-bit systems (negative values handled by validation at line 166)
		if v > float64(maxInt) {
			return 0, errors.New(i18n.T("ollama_num_ctx_value_out_of_range"))
		}
		contextLength = int(v)

	case float32:
		f64 := float64(v)
		if math.IsNaN(f64) || math.IsInf(f64, 0) {
			return 0, errors.New(i18n.T("ollama_num_ctx_must_be_finite"))
		}
		if math.Trunc(f64) != f64 {
			return 0, errors.New(i18n.T("ollama_num_ctx_must_be_integer"))
		}
		// Check for overflow on 32-bit systems (negative values handled by validation at line 177)
		if f64 > float64(maxInt) {
			return 0, errors.New(i18n.T("ollama_num_ctx_value_out_of_range"))
		}
		contextLength = int(v)

	case int:
		contextLength = v

	case int64:
		if v < 0 {
			return 0, fmt.Errorf(i18n.T("ollama_num_ctx_must_be_positive"), v)
		}
		if v > maxInt {
			return 0, fmt.Errorf(i18n.T("ollama_num_ctx_value_too_large"), v)
		}
		contextLength = int(v)

	case json.Number:
		i64, err := v.Int64()
		if err != nil {
			return 0, errors.New(i18n.T("ollama_num_ctx_must_be_valid_number"))
		}
		if i64 < 0 {
			return 0, fmt.Errorf(i18n.T("ollama_num_ctx_must_be_positive"), i64)
		}
		if i64 > maxInt {
			return 0, fmt.Errorf(i18n.T("ollama_num_ctx_value_too_large"), i64)
		}
		contextLength = int(i64)

	case string:
		parsed, err := strconv.Atoi(v)
		if err != nil {
			// Truncate long strings in error messages to avoid logging excessively large input
			errVal := v
			if len(v) > 50 {
				errVal = v[:50] + "..."
			}
			return 0, fmt.Errorf(i18n.T("ollama_num_ctx_must_be_valid_number_got"), errVal)
		}
		contextLength = parsed

	default:
		return 0, errors.New(i18n.T("ollama_num_ctx_invalid_type"))
	}

	if contextLength <= 0 {
		return 0, fmt.Errorf(i18n.T("ollama_num_ctx_must_be_positive"), contextLength)
	}

	const maxContextLength = 1000000
	if contextLength > maxContextLength {
		return 0, fmt.Errorf(i18n.T("ollama_num_ctx_exceeds_maximum"), maxContextLength)
	}

	return contextLength, nil
}

func ServeOllama(registry *core.PluginRegistry, address string, version string) (err error) {
	r := gin.New()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Register routes
	fabricDb := registry.Db
	NewPatternsHandler(r, fabricDb.Patterns)
	NewContextsHandler(r, fabricDb.Contexts)
	NewSessionsHandler(r, fabricDb.Sessions)
	NewChatHandler(r, registry, fabricDb)
	NewConfigHandler(r, fabricDb)
	NewModelsHandler(r, registry.VendorManager)

	typeConversion := APIConvert{
		registry: registry,
		r:        r,
		addr:     &address,
	}
	// Ollama Endpoints
	r.GET("/api/tags", typeConversion.ollamaTags)
	r.GET("/api/version", func(c *gin.Context) {
		c.Data(200, "application/json", fmt.Appendf(nil, "{\"%s\"}", version))
	})
	r.POST("/api/chat", typeConversion.ollamaChat)

	// Start server
	err = r.Run(address)
	if err != nil {
		return err
	}

	return
}

func (f APIConvert) ollamaTags(c *gin.Context) {
	patterns, err := f.registry.Db.Patterns.GetNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var response OllamaModel
	for _, pattern := range patterns {
		today := time.Now().Format("2024-11-25T12:07:58.915991813-05:00")
		details := ModelDetails{
			Families:          []string{"fabric"},
			Family:            "fabric",
			Format:            "custom",
			ParameterSize:     "42.0B",
			ParentModel:       "",
			QuantizationLevel: "",
		}
		response.Models = append(response.Models, Model{
			Details:    details,
			Digest:     "365c0bd3c000a25d28ddbf732fe1c6add414de7275464c4e4d1c3b5fcb5d8ad1",
			Model:      fmt.Sprintf("%s:latest", pattern),
			ModifiedAt: today,
			Name:       fmt.Sprintf("%s:latest", pattern),
			Size:       0,
		})
	}

	c.JSON(200, response)

}

func (f APIConvert) ollamaChat(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf(i18n.T("ollama_error_reading_body"), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": i18n.T("ollama_error_endpoint")})
		return
	}
	var prompt OllamaRequestBody
	err = json.Unmarshal(body, &prompt)
	if err != nil {
		log.Printf(i18n.T("ollama_error_unmarshalling_body"), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": i18n.T("ollama_error_endpoint")})
		return
	}

	// Extract and validate num_ctx from options
	numCtx, err := parseOllamaNumCtx(prompt.Options)
	if err != nil {
		log.Printf(i18n.T("ollama_invalid_num_ctx_in_request"), err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	var chat ChatRequest

	// Extract variables from either top-level Variables field or Options.variables
	variables := prompt.Variables
	if variables == nil && prompt.Options != nil {
		if optVars, ok := prompt.Options["variables"]; ok {
			// Options.variables can be either a JSON string or a map
			switch v := optVars.(type) {
			case string:
				// Parse JSON string into map
				if err := json.Unmarshal([]byte(v), &variables); err != nil {
					log.Printf(i18n.T("ollama_warning_parse_variables"), err)
				}
			case map[string]any:
				// Convert map[string]any to map[string]string
				variables = make(map[string]string)
				for k, val := range v {
					if s, ok := val.(string); ok {
						variables[k] = s
					}
				}
			}
		}
	}

	if len(prompt.Messages) == 1 {
		chat.Prompts = []PromptRequest{{
			UserInput:   prompt.Messages[0].Content,
			Vendor:      "",
			Model:       "",
			ContextName: "",
			PatternName: strings.Split(prompt.Model, ":")[0],
			Variables:   variables,
		}}
	} else if len(prompt.Messages) > 1 {
		var content string
		for _, msg := range prompt.Messages {
			content = fmt.Sprintf("%s%s:%s\n", content, msg.Role, msg.Content)
		}
		chat.Prompts = []PromptRequest{{
			UserInput:   content,
			Vendor:      "",
			Model:       "",
			ContextName: "",
			PatternName: strings.Split(prompt.Model, ":")[0],
			Variables:   variables,
		}}
	}

	// Set context length from parsed num_ctx
	chat.ModelContextLength = numCtx

	fabricChatReq, err := json.Marshal(chat)
	if err != nil {
		log.Printf(i18n.T("ollama_error_marshalling_body"), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var req *http.Request
	baseURL, err := buildFabricChatURL(*f.addr)
	if err != nil {
		log.Printf(i18n.T("ollama_error_building_chat_url"), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req, err = http.NewRequest("POST", fmt.Sprintf("%s/chat", baseURL), bytes.NewBuffer(fabricChatReq))
	if err != nil {
		log.Printf(i18n.T("ollama_error_creating_chat_request"), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": i18n.T("ollama_failed_create_request")})
		return
	}

	req = req.WithContext(c.Request.Context())

	fabricRes, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf(i18n.T("ollama_error_getting_chat_body"), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer fabricRes.Body.Close()

	if fabricRes.StatusCode < http.StatusOK || fabricRes.StatusCode >= http.StatusMultipleChoices {
		bodyBytes, readErr := io.ReadAll(fabricRes.Body)
		if readErr != nil {
			log.Printf(i18n.T("ollama_upstream_non_2xx_body_unreadable"), fabricRes.StatusCode, readErr)
		} else {
			log.Printf(i18n.T("ollama_upstream_non_2xx"), fabricRes.StatusCode, string(bodyBytes))
		}

		errorMessage := fmt.Sprintf(i18n.T("ollama_upstream_returned_status"), fabricRes.StatusCode)
		if prompt.Stream {
			_ = writeOllamaResponse(c, prompt.Model, fmt.Sprintf(i18n.T("ollama_error_prefix"), errorMessage), true)
		} else {
			c.JSON(fabricRes.StatusCode, gin.H{"error": errorMessage})
		}
		return
	}

	if prompt.Stream {
		c.Header("Content-Type", "application/x-ndjson")
	}

	var contentBuilder strings.Builder
	scanner := bufio.NewScanner(fabricRes.Body)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		payload := strings.TrimPrefix(line, "data: ")
		var fabricResponse FabricResponseFormat
		if err := json.Unmarshal([]byte(payload), &fabricResponse); err != nil {
			log.Printf(i18n.T("ollama_error_unmarshalling_body"), err)
			if prompt.Stream {
				// In streaming mode, send the error in the same streaming format
				_ = writeOllamaResponse(c, prompt.Model, i18n.T("ollama_error_parse_upstream_response"), true)
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": i18n.T("ollama_failed_unmarshal_fabric_response")})
			}
			return
		}
		if fabricResponse.Type == "error" {
			if prompt.Stream {
				// In streaming mode, propagate the upstream error via a final streaming chunk
				_ = writeOllamaResponse(c, prompt.Model, fmt.Sprintf(i18n.T("ollama_error_prefix"), fabricResponse.Content), true)
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fabricResponse.Content})
			}
			return
		}
		if fabricResponse.Type != "content" {
			continue
		}
		contentBuilder.WriteString(fabricResponse.Content)
		if prompt.Stream {
			if err := writeOllamaResponse(c, prompt.Model, fabricResponse.Content, false); err != nil {
				log.Printf(i18n.T("ollama_error_writing_response"), err)
				return
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf(i18n.T("ollama_error_scanning_body"), err)
		errorMsg := fmt.Sprintf(i18n.T("ollama_failed_scan_sse_stream"), err)
		// Check for buffer size exceeded error
		if strings.Contains(err.Error(), "token too long") {
			errorMsg = i18n.T("ollama_sse_buffer_limit")
		}
		if prompt.Stream {
			// In streaming mode, send the error in the same streaming format
			_ = writeOllamaResponse(c, prompt.Model, fmt.Sprintf(i18n.T("ollama_error_prefix"), errorMsg), true)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorMsg})
		}
		return
	}

	// Capture duration once for consistent timing values
	duration := time.Since(now).Nanoseconds()

	// Check if we received any content from upstream
	if contentBuilder.Len() == 0 {
		log.Printf("%s", i18n.T("ollama_warning_no_content"))
		// In non-streaming mode, treat absence of content as an error
		if !prompt.Stream {
			c.JSON(http.StatusBadGateway, gin.H{"error": i18n.T("ollama_no_content_from_upstream")})
			return
		}
	}

	if !prompt.Stream {
		response := buildFinalOllamaResponse(prompt.Model, contentBuilder.String(), duration)
		c.JSON(200, response)
		return
	}

	finalResponse := buildFinalOllamaResponse(prompt.Model, "", duration)
	if err := writeOllamaResponseStruct(c, finalResponse); err != nil {
		log.Printf(i18n.T("ollama_error_writing_response"), err)
	}
}

// buildFinalOllamaResponse constructs the final OllamaResponse with timing metrics
// and the complete message content. Used for both streaming and non-streaming final responses.
func buildFinalOllamaResponse(model string, content string, duration int64) OllamaResponse {
	return OllamaResponse{
		Model:     model,
		CreatedAt: time.Now().UTC().Format("2006-01-02T15:04:05.999999999Z"),
		Message: struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}(struct {
			Role    string
			Content string
		}{Content: content, Role: "assistant"}),
		DoneReason:         "stop",
		Done:               true,
		TotalDuration:      duration,
		LoadDuration:       duration,
		PromptEvalDuration: duration,
		EvalDuration:       duration,
	}
}

// buildFabricChatURL constructs a valid HTTP/HTTPS base URL from various address
// formats. It accepts fully-qualified URLs (http:// or https://), :port shorthand
// which is resolved to http://127.0.0.1:port, and bare host[:port] addresses. It
// returns a normalized URL string without a trailing slash, or an error if the
// address is empty, invalid, missing a host/hostname, or (for bare addresses)
// contains a path component.
func buildFabricChatURL(addr string) (string, error) {
	if addr == "" {
		return "", errors.New(i18n.T("ollama_empty_address"))
	}
	if strings.HasPrefix(addr, "http://") || strings.HasPrefix(addr, "https://") {
		parsed, err := url.Parse(addr)
		if err != nil {
			return "", fmt.Errorf(i18n.T("ollama_invalid_address"), err)
		}
		if parsed.Host == "" {
			return "", errors.New(i18n.T("ollama_invalid_address_missing_host"))
		}
		if strings.HasPrefix(parsed.Host, ":") {
			return "", errors.New(i18n.T("ollama_invalid_address_missing_hostname"))
		}
		return strings.TrimRight(parsed.String(), "/"), nil
	}
	if strings.HasPrefix(addr, ":") {
		return fmt.Sprintf("http://127.0.0.1%s", addr), nil
	}
	// Validate bare addresses (without http/https prefix)
	parsed, err := url.Parse("http://" + addr)
	if err != nil {
		return "", fmt.Errorf(i18n.T("ollama_invalid_address"), err)
	}
	if parsed.Host == "" {
		return "", errors.New(i18n.T("ollama_invalid_address_missing_host"))
	}
	if strings.HasPrefix(parsed.Host, ":") {
		return "", errors.New(i18n.T("ollama_invalid_address_missing_hostname"))
	}
	// Bare addresses should be host[:port] only - reject path components
	if parsed.Path != "" && parsed.Path != "/" {
		return "", errors.New(i18n.T("ollama_invalid_address_path_not_allowed"))
	}
	return strings.TrimRight(parsed.String(), "/"), nil
}

// writeOllamaResponse constructs an Ollama-formatted response chunk and writes it
// to the streaming output associated with the provided Gin context. The model
// parameter identifies the model, content is the assistant message text, and
// done indicates whether this is the final chunk in the stream.
func writeOllamaResponse(c *gin.Context, model string, content string, done bool) error {
	response := OllamaResponse{
		Model:     model,
		CreatedAt: time.Now().UTC().Format("2006-01-02T15:04:05.999999999Z"),
		Message: struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}(struct {
			Role    string
			Content string
		}{Content: content, Role: "assistant"}),
		Done: done,
	}
	return writeOllamaResponseStruct(c, response)
}

// writeOllamaResponseStruct marshals the provided OllamaResponse and writes it
// as newline-delimited JSON to the HTTP response stream.
func writeOllamaResponseStruct(c *gin.Context, response OllamaResponse) error {
	marshalled, err := json.Marshal(response)
	if err != nil {
		return err
	}
	if _, err := c.Writer.Write(marshalled); err != nil {
		return err
	}
	if _, err := c.Writer.Write([]byte("\n")); err != nil {
		return err
	}
	if flusher, ok := c.Writer.(http.Flusher); ok {
		flusher.Flush()
	}
	return nil
}
