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

// maskAPIKey redacts all but the last 4 characters of a secret key (CWE-200).
// An empty value (key not configured) is returned unchanged so the UI can
// distinguish "not set" from "set but redacted".
func maskAPIKey(key string) string {
	const visible = 4
	if len(key) <= visible {
		return key
	}
	return strings.Repeat("*", len(key)-visible) + key[len(key)-visible:]
}

// isRedacted returns true when a submitted value looks like a masked key
// returned by maskAPIKey, signalling that the user did not change the field.
func isRedacted(value string) bool {
	return strings.Contains(value, "*")
}

func (h *ConfigHandler) GetConfig(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": ".env file not found"})
		return
	}

	if !h.db.IsEnvFileExists() {
		c.JSON(http.StatusOK, gin.H{
			"openai":        "",
			"anthropic":     "",
			"groq":          "",
			"mistral":       "",
			"gemini":        "",
			"ollama":        "",
			"openrouter":    "",
			"trustedrouter": "",
			"silicon":       "",
			"deepseek":      "",
			"grokai":        "",
		})
		return
	}

	err := h.db.LoadEnvFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// API keys are masked to their last 4 characters (CWE-200).
	// URLs are not secrets and are returned as-is so the UI can display them.
	config := map[string]string{
		"openai":        maskAPIKey(os.Getenv("OPENAI_API_KEY")),
		"anthropic":     maskAPIKey(os.Getenv("ANTHROPIC_API_KEY")),
		"groq":          maskAPIKey(os.Getenv("GROQ_API_KEY")),
		"mistral":       maskAPIKey(os.Getenv("MISTRAL_API_KEY")),
		"gemini":        maskAPIKey(os.Getenv("GEMINI_API_KEY")),
		"ollama":        os.Getenv("OLLAMA_URL"),
		"openrouter":    maskAPIKey(os.Getenv("OPENROUTER_API_KEY")),
		"trustedrouter": maskAPIKey(os.Getenv("TRUSTEDROUTER_API_KEY")),
		"silicon":       maskAPIKey(os.Getenv("SILICON_API_KEY")),
		"deepseek":      maskAPIKey(os.Getenv("DEEPSEEK_API_KEY")),
		"grokai":        maskAPIKey(os.Getenv("GROKAI_API_KEY")),
		"lmstudio":      os.Getenv("LM_STUDIO_API_BASE_URL"),
	}

	c.JSON(http.StatusOK, config)
}

func (h *ConfigHandler) UpdateConfig(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
		return
	}

	var config struct {
		OpenAIApiKey        string `json:"openai_api_key"`
		AnthropicApiKey     string `json:"anthropic_api_key"`
		GroqApiKey          string `json:"groq_api_key"`
		MistralApiKey       string `json:"mistral_api_key"`
		GeminiApiKey        string `json:"gemini_api_key"`
		OllamaURL           string `json:"ollama_url"`
		OpenRouterApiKey    string `json:"openrouter_api_key"`
		TrustedRouterApiKey string `json:"trustedrouter_api_key"`
		SiliconApiKey       string `json:"silicon_api_key"`
		DeepSeekApiKey      string `json:"deepseek_api_key"`
		GrokaiApiKey        string `json:"grokai_api_key"`
		LMStudioURL         string `json:"lm_studio_base_url"`
	}

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	envVars := map[string]string{
		"OPENAI_API_KEY":         config.OpenAIApiKey,
		"ANTHROPIC_API_KEY":      config.AnthropicApiKey,
		"GROQ_API_KEY":           config.GroqApiKey,
		"MISTRAL_API_KEY":        config.MistralApiKey,
		"GEMINI_API_KEY":         config.GeminiApiKey,
		"OLLAMA_URL":             config.OllamaURL,
		"OPENROUTER_API_KEY":     config.OpenRouterApiKey,
		"TRUSTEDROUTER_API_KEY":  config.TrustedRouterApiKey,
		"SILICON_API_KEY":        config.SiliconApiKey,
		"DEEPSEEK_API_KEY":       config.DeepSeekApiKey,
		"GROKAI_API_KEY":         config.GrokaiApiKey,
		"LM_STUDIO_API_BASE_URL": config.LMStudioURL,
	}

	var envContent strings.Builder
	for key, value := range envVars {
		// Skip empty values and redacted placeholders returned by GET /config.
		// Writing a masked value back would corrupt the stored key.
		if value != "" && !isRedacted(value) {
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
