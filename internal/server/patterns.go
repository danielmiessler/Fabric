package restapi

import (
	"maps"
	"net/http"
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
)

// PatternsHandler defines the handler for patterns-related operations
type PatternsHandler struct {
	*StorageHandler[fsdb.Pattern]
	patterns *fsdb.PatternsEntity
}

// NewPatternsHandler creates a new PatternsHandler
func NewPatternsHandler(r *gin.Engine, patterns *fsdb.PatternsEntity) (ret *PatternsHandler) {
	// Create a storage handler but don't register any routes yet
	storageHandler := &StorageHandler[fsdb.Pattern]{storage: patterns}
	ret = &PatternsHandler{StorageHandler: storageHandler, patterns: patterns}

	// Register routes manually - use custom Get for patterns, others from StorageHandler
	r.GET("/patterns/:name", ret.Get)                       // Custom method with variables support
	r.GET("/patterns/names", ret.GetNames)                  // From StorageHandler
	r.DELETE("/patterns/:name", ret.Delete)                 // From StorageHandler
	r.GET("/patterns/exists/:name", ret.Exists)             // From StorageHandler
	r.PUT("/patterns/rename/:oldName/:newName", ret.Rename) // From StorageHandler
	r.POST("/patterns/:name", ret.Save)                     // From StorageHandler
	// Add POST route for patterns with variables in request body
	r.POST("/patterns/:name/apply", ret.ApplyPattern)
	return
}

// Get handles the GET /patterns/:name route - returns raw pattern without variable processing
// @Summary Get a pattern
// @Description Retrieve a pattern by name
// @Tags patterns
// @Accept json
// @Produce json
// @Param name path string true "Pattern name"
// @Success 200 {object} fsdb.Pattern
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /patterns/{name} [get]
func (h *PatternsHandler) Get(c *gin.Context) {
	name := c.Param("name")

	pattern, err := h.patterns.GetRaw(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, pattern)
}

// PatternApplyRequest represents the request body for applying a pattern
type PatternApplyRequest struct {
	Input     string            `json:"input"`
	Variables map[string]string `json:"variables,omitempty"`
}

// ApplyPattern handles the POST /patterns/:name/apply route
// @Summary Apply pattern with variables
// @Description Apply a pattern with variable substitution
// @Tags patterns
// @Accept json
// @Produce json
// @Param name path string true "Pattern name"
// @Param request body PatternApplyRequest true "Pattern application request"
// @Success 200 {object} fsdb.Pattern
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /patterns/{name}/apply [post]
func (h *PatternsHandler) ApplyPattern(c *gin.Context) {
	name := c.Param("name")

	var request PatternApplyRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Merge query parameters with request body variables (body takes precedence)
	variables := make(map[string]string)
	for key, values := range c.Request.URL.Query() {
		if len(values) > 0 {
			variables[key] = values[0]
		}
	}
	maps.Copy(variables, request.Variables)

	pattern, err := h.patterns.GetApplyVariables(name, variables, request.Input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Patterns without an explicit {{input}} marker do not contain the input
	// now. Add the input at the end to keep the response contract of this
	// endpoint.
	if request.Input != "" && !strings.Contains(pattern.Pattern, request.Input) {
		pattern.Pattern = strings.TrimSuffix(pattern.Pattern, "\n") + "\n" + request.Input
	}

	c.JSON(http.StatusOK, pattern)
}
