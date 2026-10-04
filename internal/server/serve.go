package restapi

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
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
func Serve(registry *core.PluginRegistry, address string, apiKey string, corsOrigins []string) (err error) {
	if err = requireAPIKeyForBind(address, apiKey); err != nil {
		return err
	}
	if corsOrigins, err = cleanCORSOrigins(corsOrigins, apiKey); err != nil {
		return err
	}

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if len(corsOrigins) > 0 {
		r.Use(CORSMiddleware(corsOrigins))
	}

	if apiKey != "" {
		r.Use(APIKeyMiddleware(apiKey))
	} else {
		slog.Warn(i18n.T("server_no_api_key_warning"))
	}

	// Swagger UI and documentation endpoint with custom YAML handler
	r.GET("/swagger/*any", func(c *gin.Context) {
		if c.Param("any") == "/swagger.yaml" {
			yamlPath := "docs/swagger.yaml"
			if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
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

	fabricDb := registry.Db
	NewPatternsHandler(r, fabricDb.Patterns)
	NewContextsHandler(r, fabricDb.Contexts)
	NewSessionsHandler(r, fabricDb.Sessions)
	NewChatHandler(r, registry, fabricDb)
	NewYouTubeHandler(r, registry)
	NewConfigHandler(r, fabricDb)
	NewModelsHandler(r, registry.VendorManager)
	NewStrategiesHandler(r)

	err = r.Run(address)
	if err != nil {
		return err
	}

	return
}
