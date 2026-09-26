package restapi

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/danielmiessler/fabric/internal/plugins/db"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
)

// StorageHandler defines the handler for storage-related operations
type StorageHandler[T any] struct {
	storage db.Storage[T]
}

// setHSTS sets the Strict-Transport-Security header. Each validation
// 400 sends it, the same as the chat BindJSON 400 path.
func setHSTS(c *gin.Context) {
	c.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
}

// storageError answers err. A name-validation rejection becomes a 400,
// and its body contains only the rejected name. All other errors stay
// 500 errors with a generic body, because fsdb wraps *os.PathError
// values and err.Error() then sends absolute filesystem paths to the
// client. The full error goes to the log.
func storageError(c *gin.Context, err error) {
	if _, ok := errors.AsType[*fsdb.InvalidStorageNameError](err); ok {
		setHSTS(c)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slog.Error("storage operation failed", "error", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
}

// rejectInvalidStorageName answers a 400 when name does not obey
// storage-name validation. An empty name passes, because the fields
// that this guards are optional.
func rejectInvalidStorageName(c *gin.Context, name string) bool {
	if name == "" {
		return false
	}
	if err := fsdb.ValidateStorageName(name); err != nil {
		setHSTS(c)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
	}
	return false
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
		storageError(c, err)
		return
	}
	c.JSON(http.StatusOK, item)
}

// GetNames handles the GET /storage/names route
func (h *StorageHandler[T]) GetNames(c *gin.Context) {
	names, err := h.storage.GetNames()
	if err != nil {
		storageError(c, err)
		return
	}
	c.JSON(http.StatusOK, names)
}

// Delete handles the DELETE /storage/:name route
func (h *StorageHandler[T]) Delete(c *gin.Context) {
	name := c.Param("name")
	err := h.storage.Delete(name)
	if err != nil {
		storageError(c, err)
		return
	}
	c.Status(http.StatusOK)
}

// Exists handles the GET /storage/exists/:name route. The storage
// Exists contract cannot report an invalid name, and the handler must
// validate the name itself. An invalid name is a 400, not a "false".
func (h *StorageHandler[T]) Exists(c *gin.Context) {
	name := c.Param("name")
	if rejectInvalidStorageName(c, name) {
		return
	}
	exists := h.storage.Exists(name)
	c.JSON(http.StatusOK, exists)
}

// Rename handles the PUT /storage/rename/:oldName/:newName route
func (h *StorageHandler[T]) Rename(c *gin.Context) {
	oldName := c.Param("oldName")
	newName := c.Param("newName")
	err := h.storage.Rename(oldName, newName)
	if err != nil {
		storageError(c, err)
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
		storageError(c, err)
		return
	}

	// Save the content to storage
	err = h.storage.Save(name, content)
	if err != nil {
		storageError(c, err)
		return
	}
	c.Status(http.StatusOK)
}
