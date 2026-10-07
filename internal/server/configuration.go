package restapi

import (
	"fmt"
	"maps"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
)

// ConfigHandler defines the handler for configuration-related operations
type ConfigHandler struct {
	db *fsdb.Db
}

func NewConfigHandler(r *gin.Engine, db *fsdb.Db) *ConfigHandler {
	handler := &ConfigHandler{
		db: db,
	}

	r.GET("/config", handler.GetConfig)
	r.POST("/config/update", requireJSON, handler.UpdateConfig)

	return handler
}

// maskedValue is the fixed mask that GET /config returns for a set key.
const maskedValue = "********"

// maskAPIKey replaces a secret key with a fixed mask (CWE-200).
// An empty value (key not configured) is returned unchanged so the UI can
// distinguish "not set" from "set but redacted".
func maskAPIKey(key string) string {
	if key == "" {
		return ""
	}
	// The mask has a fixed length. It does not show the key length or
	// characters from the key. It only shows that a key is set. isRedacted
	// finds the mask when the UI sends it back.
	return maskedValue
}

// isRedacted returns true when a submitted value is the mask that
// maskAPIKey returns. This shows that the user did not change the field.
// A value that only contains '*' is a new value.
func isRedacted(value string) bool {
	return value == maskedValue
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
			"lmstudio":      "",
		})
		return
	}

	err := h.db.LoadEnvFile()
	if err != nil {
		storageError(c, err)
		return
	}

	// API keys are replaced with a fixed mask (CWE-200).
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

	updates := make(map[string]string, len(envVars))
	// Sorted keys make the key in the error message the same each time.
	for _, key := range slices.Sorted(maps.Keys(envVars)) {
		value := envVars[key]
		// Skip empty values and redacted placeholders returned by GET /config.
		// Writing a masked value back would corrupt the stored key.
		// UpdateEnvVars also skips a value of only spaces, thus skip it here.
		if strings.TrimSpace(value) == "" || isRedacted(value) {
			continue
		}
		// A CR or LF in a value can add a line to the .env file. A NUL
		// makes os.Setenv fail. Refuse these characters.
		if strings.ContainsAny(value, "\r\n\x00") {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid value for %s: must not contain CR, LF or NUL", key)})
			return
		}
		updates[key] = value
	}

	// Merge the updates into the current .env file. Keep the keys that
	// this form does not know, for example tokens from other tools.
	// UpdateEnvVars reads, merges and writes the file under a lock.
	if err := h.db.UpdateEnvVars(updates); err != nil {
		storageError(c, err)
		return
	}
	for key, value := range updates {
		if err := os.Setenv(key, value); err != nil {
			storageError(c, err)
			return
		}
	}

	if err := h.db.LoadEnvFile(); err != nil {
		storageError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Configuration updated successfully"})
}
