package restapi

import (
	"log/slog"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/gin-gonic/gin"
)

func Serve(registry *core.PluginRegistry, address string, apiKey string) (err error) {
	r := gin.New()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if apiKey != "" {
		r.Use(APIKeyMiddleware(apiKey))
	} else {
		slog.Warn("Starting REST API server without API key authentication. This may pose security risks.")
	}

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
