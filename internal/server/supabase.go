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

func NewSupabaseHandler(r *gin.Engine, client *supadb.Client) *SupabaseHandler {
	if client == nil {
		return nil
	}

	handler := &SupabaseHandler{client: client}
	group := r.Group("/supabase")
	group.GET("/health", handler.Health)
	group.GET("/sessions", handler.ListSessions)
	group.GET("/patterns", handler.ListPatterns)
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
