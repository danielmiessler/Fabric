package restapi

import (
	"net/http"

	"github.com/danielmiessler/fabric/internal/plugins/db/supadb"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SupabaseHandler struct {
	client *supadb.Client
}

type patternRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description *string  `json:"description"`
	Body        string   `json:"body" binding:"required"`
	Tags        []string `json:"tags"`
	IsSystem    bool     `json:"is_system"`
}

func NewSupabaseHandler(r *gin.Engine, client *supadb.Client) *SupabaseHandler {
	if client == nil {
		return nil
	}

	handler := &SupabaseHandler{client: client}
	group := r.Group("/supabase")
	group.GET("/health", handler.Health)
	group.GET("/sessions", handler.ListSessions)
	group.GET("/patterns", handler.ListPatterns)
	group.GET("/patterns/:id", handler.GetPattern)
	group.POST("/patterns", handler.CreatePattern)
	group.PUT("/patterns/:id", handler.UpdatePattern)
	group.DELETE("/patterns/:id", handler.DeletePattern)
	group.GET("/notes/:sessionId", handler.ListNotes)

	return handler
}

func (h *SupabaseHandler) Health(c *gin.Context) {
	if err := h.client.Ping(c.Request.Context()); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *SupabaseHandler) ListSessions(c *gin.Context) {
	repo := h.client.Sessions()
	sessions, err := repo.List(c.Request.Context(), 50)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

func (h *SupabaseHandler) ListPatterns(c *gin.Context) {
	repo := h.client.Patterns()
	patterns, err := repo.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patterns)
}

func (h *SupabaseHandler) GetPattern(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pattern id"})
		return
	}

	repo := h.client.Patterns()
	pattern, err := repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if pattern == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pattern not found"})
		return
	}
	c.JSON(http.StatusOK, pattern)
}

func (h *SupabaseHandler) CreatePattern(c *gin.Context) {
	var req patternRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo := h.client.Patterns()
	pattern, err := repo.Create(c.Request.Context(), map[string]any{
		"name":        req.Name,
		"description": req.Description,
		"body":        req.Body,
		"tags":        req.Tags,
		"is_system":   req.IsSystem,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if pattern == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create pattern"})
		return
	}
	c.JSON(http.StatusCreated, pattern)
}

func (h *SupabaseHandler) UpdatePattern(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pattern id"})
		return
	}

	var req patternRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo := h.client.Patterns()
	pattern, err := repo.UpdateByID(c.Request.Context(), id, map[string]any{
		"name":        req.Name,
		"description": req.Description,
		"body":        req.Body,
		"tags":        req.Tags,
		"is_system":   req.IsSystem,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if pattern == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pattern not found"})
		return
	}
	c.JSON(http.StatusOK, pattern)
}

func (h *SupabaseHandler) DeletePattern(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pattern id"})
		return
	}

	repo := h.client.Patterns()
	if err := repo.DeleteByID(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *SupabaseHandler) ListNotes(c *gin.Context) {
	sessionIDParam := c.Param("sessionId")
	var sessionID uuid.UUID
	if sessionIDParam != "" {
		parsed, err := uuid.Parse(sessionIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
			return
		}
		sessionID = parsed
	}

	repo := h.client.Notes()
	notes, err := repo.ListBySession(c.Request.Context(), sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notes)
}
