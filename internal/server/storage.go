package restapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/danielmiessler/fabric/internal/plugins/db"
	"github.com/gin-gonic/gin"
)

// StorageHandler defines the handler for storage-related operations
type StorageHandler[T any] struct {
	storage db.Storage[T]
}

// NewStorageHandler creates a new StorageHandler
func NewStorageHandler[T any](r *gin.Engine, entityType string, storage db.Storage[T]) (ret *StorageHandler[T]) {
	ret = &StorageHandler[T]{storage: storage}
	r.GET(fmt.Sprintf("/%s/:name", entityType), ret.Get)
	r.GET(fmt.Sprintf("/%s/names", entityType), ret.GetNames)
	r.DELETE(fmt.Sprintf("/%s/:name", entityType), ret.Delete)
	r.GET(fmt.Sprintf("/%s/exists/:name", entityType), ret.Exists)
	r.PUT(fmt.Sprintf("/%s/rename/:oldName/:newName", entityType), ret.Rename)
	r.POST(fmt.Sprintf("/%s/:name", entityType), ret.Save)
	return
}

// Get handles the GET /storage/:name route
func (h *StorageHandler[T]) Get(c *gin.Context) {
	name := c.Param("name")
	item, err := h.storage.Get(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

// GetNames handles the GET /storage/names route
func (h *StorageHandler[T]) GetNames(c *gin.Context) {
	names, err := h.storage.GetNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, names)
}

// Delete handles the DELETE /storage/:name route
func (h *StorageHandler[T]) Delete(c *gin.Context) {
	name := c.Param("name")
	err := h.storage.Delete(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

// Exists handles the GET /storage/exists/:name route
func (h *StorageHandler[T]) Exists(c *gin.Context) {
	name := c.Param("name")
	exists := h.storage.Exists(name)
	c.JSON(http.StatusOK, exists)
}

// Rename handles the PUT /storage/rename/:oldName/:newName route
func (h *StorageHandler[T]) Rename(c *gin.Context) {
	oldName := c.Param("oldName")
	newName := c.Param("newName")
	err := h.storage.Rename(oldName, newName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

// Save handles the POST /storage/save/:name route
func (h *StorageHandler[T]) Save(c *gin.Context) {
	name := c.Param("name")

	// Read the request body
	body := c.Request.Body
	defer body.Close()

	content, err := io.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Save the content to storage
	err = h.storage.Save(name, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
