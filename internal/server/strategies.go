package restapi

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// StrategyMeta represents the minimal info about a strategy
type StrategyMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Prompt      string `json:"prompt"`
}

// NewStrategiesHandler registers the /strategies GET endpoint
func NewStrategiesHandler(r *gin.Engine) {
	r.GET("/strategies", func(c *gin.Context) {
		strategiesDir := filepath.Join(os.Getenv("HOME"), ".config", "fabric", "strategies")

		files, err := os.ReadDir(strategiesDir)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read strategies directory"})
			return
		}

		var strategies []StrategyMeta

		for _, file := range files {
			if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
				continue
			}

			fullPath := filepath.Join(strategiesDir, file.Name())
			data, err := os.ReadFile(fullPath)
			if err != nil {
				continue
			}

			var s struct {
				Description string `json:"description"`
				Prompt      string `json:"prompt"`
			}
			if err := json.Unmarshal(data, &s); err != nil {
				continue
			}

			strategies = append(strategies, StrategyMeta{
				Name:        strings.TrimSuffix(file.Name(), ".json"),
				Description: s.Description,
				Prompt:      s.Prompt,
			})
		}

		c.JSON(http.StatusOK, strategies)
	})
}
