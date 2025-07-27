package restapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/core"
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
	Messages []OllamaMessage `json:"messages"`
	Model    string          `json:"model"`
	Options  struct {
	} `json:"options"`
	Stream bool `json:"stream"`
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
	LoadDuration       int    `json:"load_duration,omitempty"`
	PromptEvalCount    int    `json:"prompt_eval_count,omitempty"`
	PromptEvalDuration int    `json:"prompt_eval_duration,omitempty"`
	EvalCount          int    `json:"eval_count,omitempty"`
	EvalDuration       int64  `json:"eval_duration,omitempty"`
}

type FabricResponseFormat struct {
	Type    string `json:"type"`
	Format  string `json:"format"`
	Content string `json:"content"`
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
		c.Data(200, "application/json", []byte(fmt.Sprintf("{\"%s\"}", version)))
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
		log.Printf("Error reading body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "testing endpoint"})
		return
	}
	var prompt OllamaRequestBody
	err = json.Unmarshal(body, &prompt)
	if err != nil {
		log.Printf("Error unmarshalling body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "testing endpoint"})
		return
	}
	now := time.Now()
	var chat ChatRequest

	if len(prompt.Messages) == 1 {
		chat.Prompts = []PromptRequest{{
			UserInput:   prompt.Messages[0].Content,
			Vendor:      "",
			Model:       "",
			ContextName: "",
			PatternName: strings.Split(prompt.Model, ":")[0],
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
		}}
	}
	fabricChatReq, err := json.Marshal(chat)
	if err != nil {
		log.Printf("Error marshalling body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx := context.Background()
	var req *http.Request
	if strings.Contains(*f.addr, "http") {
		req, err = http.NewRequest("POST", fmt.Sprintf("%s/chat", *f.addr), bytes.NewBuffer(fabricChatReq))
	} else {
		req, err = http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1%s/chat", *f.addr), bytes.NewBuffer(fabricChatReq))
	}
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx)

	fabricRes, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error getting /chat body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	body, err = io.ReadAll(fabricRes.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "testing endpoint"})
		return
	}
	var forwardedResponse OllamaResponse
	var forwardedResponses []OllamaResponse
	var fabricResponse FabricResponseFormat
	err = json.Unmarshal([]byte(strings.Split(strings.Split(string(body), "\n")[0], "data: ")[1]), &fabricResponse)
	if err != nil {
		log.Printf("Error unmarshalling body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "testing endpoint"})
		return
	}
	for _, word := range strings.Split(fabricResponse.Content, " ") {
		forwardedResponse = OllamaResponse{
			Model:     "",
			CreatedAt: "",
			Message: struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			}(struct {
				Role    string
				Content string
			}{Content: fmt.Sprintf("%s ", word), Role: "assistant"}),
			Done: false,
		}
		forwardedResponses = append(forwardedResponses, forwardedResponse)
	}
	forwardedResponse.Model = prompt.Model
	forwardedResponse.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05.999999999Z")
	forwardedResponse.Message.Role = "assistant"
	forwardedResponse.Message.Content = ""
	forwardedResponse.DoneReason = "stop"
	forwardedResponse.Done = true
	forwardedResponse.TotalDuration = time.Since(now).Nanoseconds()
	forwardedResponse.LoadDuration = int(time.Since(now).Nanoseconds())
	forwardedResponse.PromptEvalCount = 42
	forwardedResponse.PromptEvalDuration = int(time.Since(now).Nanoseconds())
	forwardedResponse.EvalCount = 420
	forwardedResponse.EvalDuration = time.Since(now).Nanoseconds()
	forwardedResponses = append(forwardedResponses, forwardedResponse)

	var res []byte
	for _, response := range forwardedResponses {
		marshalled, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error marshalling body: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		res = append(res, marshalled...)
		res = append(res, '\n')
	}
	c.Data(200, "application/json", res)

	//c.JSON(200, forwardedResponse)
}
