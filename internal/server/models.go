package restapi

import (
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/gin-gonic/gin"
)

type ModelsHandler struct {
	vendorManager *ai.VendorsManager
}

func NewModelsHandler(r *gin.Engine, vendorManager *ai.VendorsManager) {
	handler := &ModelsHandler{
		vendorManager: vendorManager,
	}

	r.GET("/models/names", handler.GetModelNames)
}

// GetModelNames godoc
// @Summary List all available models
// @Description Get a list of all available AI models grouped by vendor
// @Tags models
// @Produce json
// @Success 200 {object} map[string]interface{} "Returns models (array) and vendors (map)"
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /models/names [get]
func (h *ModelsHandler) GetModelNames(c *gin.Context) {
	vendorsModels, err := h.vendorManager.GetModels()
	if err != nil {
		c.JSON(500, gin.H{"error": "Server failed to retrieve model names"})
		return
	}

	response := make(map[string]any)
	response["models"] = h.getAllModelNames(vendorsModels)
	response["vendors"] = buildVendorsMap(vendorsModels)
	c.JSON(200, response)
}

// buildVendorsMap groups the model names by vendor name for the response.
// A vendor with no models gets an empty slice, not a nil slice, because a nil
// slice becomes null in JSON and a client that reads the list of a vendor then
// gets null in place of an array. Ollama does this when it is in the
// configuration but serves no models.
func buildVendorsMap(vendorsModels *ai.VendorsModels) map[string][]string {
	vendors := make(map[string][]string)
	for _, groupItems := range vendorsModels.GroupsItems {
		if groupItems.Items == nil {
			vendors[groupItems.Group] = []string{}
			continue
		}
		vendors[groupItems.Group] = groupItems.Items
	}
	return vendors
}

func (h *ModelsHandler) getAllModelNames(vendorsModels *ai.VendorsModels) []string {
	var allModelNames []string
	for _, groupItems := range vendorsModels.GroupsItems {
		allModelNames = append(allModelNames, groupItems.Items...)
	}
	return allModelNames
}
