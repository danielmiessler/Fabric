package restapi

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
)

// ConfigHandler defines the handler for configuration-related operations
type ConfigHandler struct {
	db *fsdb.Db
	// configurations *fsdb.EnvFilePath("$HOME/.config/fabric/.env")
}

func NewConfigHandler(r *gin.Engine, db *fsdb.Db) *ConfigHandler {
	handler := &ConfigHandler{
		db: db,
		// configurations: db.Configurations,
	}

	r.GET("/config", handler.GetConfig)
	r.POST("/config/update", handler.UpdateConfig)

	return handler
}

func (h *ConfigHandler) GetConfig(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": ".env file not found"})
		return
	}

	if !h.db.IsEnvFileExists() {
		c.JSON(http.StatusOK, gin.H{
			"openai":     "",
			"anthropic":  "",
			"groq":       "",
			"mistral":    "",
			"gemini":     "",
			"ollama":     "",
			"openrouter": "",
			"silicon":    "",
			"deepseek":   "",
			"grokai":     "",
		})
		return
	}

	err := h.db.LoadEnvFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config := map[string]string{
		"openai":                    os.Getenv("OPENAI_API_KEY"),
		"anthropic":                 os.Getenv("ANTHROPIC_API_KEY"),
		"anthropic_use_oauth_login": os.Getenv("ANTHROPIC_USE_OAUTH_LOGIN"),
		"groq":                      os.Getenv("GROQ_API_KEY"),
		"mistral":                   os.Getenv("MISTRAL_API_KEY"),
		"gemini":                    os.Getenv("GEMINI_API_KEY"),
		"ollama":                    os.Getenv("OLLAMA_URL"),
		"openrouter":                os.Getenv("OPENROUTER_API_KEY"),
		"silicon":                   os.Getenv("SILICON_API_KEY"),
		"deepseek":                  os.Getenv("DEEPSEEK_API_KEY"),
		"grokai":                    os.Getenv("GROKAI_API_KEY"),
		"lmstudio":                  os.Getenv("LM_STUDIO_API_BASE_URL"),
	}

	c.JSON(http.StatusOK, config)
}

func (h *ConfigHandler) UpdateConfig(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
		return
	}

	var config struct {
		OpenAIApiKey          string `json:"openai_api_key"`
		AnthropicApiKey       string `json:"anthropic_api_key"`
		AnthropicUseAuthToken string `json:"anthropic_use_auth_token"`
		GroqApiKey            string `json:"groq_api_key"`
		MistralApiKey         string `json:"mistral_api_key"`
		GeminiApiKey          string `json:"gemini_api_key"`
		OllamaURL             string `json:"ollama_url"`
		OpenRouterApiKey      string `json:"openrouter_api_key"`
		SiliconApiKey         string `json:"silicon_api_key"`
		DeepSeekApiKey        string `json:"deepseek_api_key"`
		GrokaiApiKey          string `json:"grokai_api_key"`
		LMStudioURL           string `json:"lm_studio_base_url"`
	}

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	envVars := map[string]string{
		"OPENAI_API_KEY":            config.OpenAIApiKey,
		"ANTHROPIC_API_KEY":         config.AnthropicApiKey,
		"ANTHROPIC_USE_OAUTH_LOGIN": config.AnthropicUseAuthToken,
		"GROQ_API_KEY":              config.GroqApiKey,
		"MISTRAL_API_KEY":           config.MistralApiKey,
		"GEMINI_API_KEY":            config.GeminiApiKey,
		"OLLAMA_URL":                config.OllamaURL,
		"OPENROUTER_API_KEY":        config.OpenRouterApiKey,
		"SILICON_API_KEY":           config.SiliconApiKey,
		"DEEPSEEK_API_KEY":          config.DeepSeekApiKey,
		"GROKAI_API_KEY":            config.GrokaiApiKey,
		"LM_STUDIO_API_BASE_URL":    config.LMStudioURL,
	}

	var envContent strings.Builder
	for key, value := range envVars {
		if value != "" {
			envContent.WriteString(fmt.Sprintf("%s=%s\n", key, value))
			os.Setenv(key, value)
		}
	}

	// Save configuration to file
	if err := h.db.SaveEnv(envContent.String()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.LoadEnvFile(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Configuration updated successfully"})
}
