package locationhandler

import (
	"github.com/gin-gonic/gin"
	"main.go/utils"
	"net/http"
)

func (h *locationHandler) ListLocation(ctx *gin.Context) {
	location, err := h.locationService.ListLocation()
	if err != nil {
		return
	}
	utils.ApiResponse(ctx, http.StatusOK, "success", location, nil)
}
