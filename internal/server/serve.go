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

	return newServeEngine(registry, apiKey, corsOrigins).Run(address)
}

// newSecuredEngine makes an engine with the security middleware that each
// server uses. noKeyWarning is the log message when apiKey is empty.
func newSecuredEngine(registry *core.PluginRegistry, apiKey string, corsOrigins []string, noKeyWarning string) *gin.Engine {
	// A client can save a pattern. Thus the server must not run a pattern
	// plugin that reads the environment or files, fetches a URL, or runs an
	// extension. Set this here, so that each server entry point gets it.
	registry.Db.Patterns.NoSystemPlugins = true

	// Debug mode logs each route and warns at startup. Use release mode,
	// unless the operator sets GIN_MODE.
	if os.Getenv(gin.EnvGinMode) == "" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// The server is not behind a known proxy. Thus c.ClientIP() in the logs
	// must not come from the X-Forwarded-For header, which a client can set.
	_ = r.SetTrustedProxies(nil) // A nil list gives no error.

	// The query string can hold template variables, for example on
	// POST /patterns/:name/apply. Do not write it to the log.
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipQueryString: true}))
	r.Use(gin.Recovery())

	if len(corsOrigins) > 0 {
		r.Use(CORSMiddleware(corsOrigins))
	}

	if apiKey != "" {
		r.Use(APIKeyMiddleware(apiKey))
	} else {
		// With no key, the server does no authentication. Reject a request
		// whose Host is not loopback or whose Origin is from a different site.
		r.Use(LoopbackSecurityMiddleware(corsOrigins))
		slog.Warn(noKeyWarning)
	}

	// The body limit comes after the API key and loopback checks. The
	// middleware can read a chunked body into memory, and it must not do
	// this for a request that the server rejects.
	r.Use(MaxBodyBytesMiddleware(MaxRequestBodyBytes))
	return r
}

// newServeEngine makes the REST API engine but does not start it, which lets
// tests operate the middleware and the routes.
func newServeEngine(registry *core.PluginRegistry, apiKey string, corsOrigins []string) *gin.Engine {
	r := newSecuredEngine(registry, apiKey, corsOrigins, i18n.T("server_no_api_key_warning"))

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
	return r
}
