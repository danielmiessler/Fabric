package restapi

import (
	"net/http"

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
func (h *PatternsHandler) Get(c *gin.Context) {
	name := c.Param("name")

	// Get the raw pattern content without any variable processing
	content, err := h.patterns.Load(name + "/" + h.patterns.SystemPatternFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Return raw pattern in the same format as the processed patterns
	pattern := &fsdb.Pattern{
		Name:        name,
		Description: "",
		Pattern:     string(content),
	}
	c.JSON(http.StatusOK, pattern)
}

// PatternApplyRequest represents the request body for applying a pattern
type PatternApplyRequest struct {
	Input     string            `json:"input"`
	Variables map[string]string `json:"variables,omitempty"`
}

// ApplyPattern handles the POST /patterns/:name/apply route
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
	for key, value := range request.Variables {
		variables[key] = value
	}

	pattern, err := h.patterns.GetApplyVariables(name, variables, request.Input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, pattern)
}
