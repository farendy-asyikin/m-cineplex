package bookinghandler

import (
	"github.com/gin-gonic/gin"
	"main.go/schemas"
	"main.go/utils"
	"net/http"
)

func (h *bookingHandler) CreateBooking(ctx *gin.Context) {
	var request schemas.CreateBookingRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	booking, err := h.bookingService.CreateBooking(request)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", booking, nil)
}

func (h *bookingHandler) GetBookingByID(ctx *gin.Context) {
	id := ctx.Param("id")

	booking, err := h.bookingService.GetBookingByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", booking, nil)
}
