package restapi

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/danielmiessler/fabric/docs" // swagger docs
)

// @title Fabric REST API
// @version 1.0
// @description REST API for Fabric AI augmentation framework. Provides endpoints for chat completions, pattern management, contexts, sessions, and more.
// @contact.name Fabric Support
// @contact.url https://github.com/danielmiessler/fabric
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
func Serve(registry *core.PluginRegistry, address string, apiKey string) (err error) {
	r := gin.New()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS — required for browser and WebView clients connecting from a different
	// origin (e.g. localhost:5173 in dev, or tauri://localhost in a desktop app).
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-API-Key")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	if apiKey != "" {
		r.Use(APIKeyMiddleware(apiKey))
	} else {
		slog.Warn("Starting REST API server without API key authentication. This may pose security risks.")
	}

	// Swagger UI and documentation endpoint with custom YAML handler
	r.GET("/swagger/*any", func(c *gin.Context) {
		// Check if request is for swagger.yaml
		if c.Param("any") == "/swagger.yaml" {
			// Try to find swagger.yaml relative to current directory or executable
			yamlPath := "docs/swagger.yaml"
			if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
				// Try relative to executable
				if exePath, err := os.Executable(); err == nil {
					yamlPath = filepath.Join(filepath.Dir(exePath), "docs", "swagger.yaml")
				}
			}

			if _, err := os.Stat(yamlPath); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "swagger.yaml not found - generate it with: swag init -g internal/server/serve.go -o docs"})
				return
			}

			c.File(yamlPath)
			return
		}

		// For all other swagger paths, use the default handler
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})

	// Register routes
	fabricDb := registry.Db
	NewPatternsHandler(r, fabricDb.Patterns)
	NewContextsHandler(r, fabricDb.Contexts)
	NewSessionsHandler(r, fabricDb.Sessions)
	NewChatHandler(r, registry, fabricDb)
	NewYouTubeHandler(r, registry)
	NewConfigHandler(r, fabricDb)
	NewModelsHandler(r, registry.VendorManager)
	NewStrategiesHandler(r)

	// Start server
	err = r.Run(address)
	if err != nil {
		return err
	}

	return
}
