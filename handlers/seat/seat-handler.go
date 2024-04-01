package seathandler

import (
	"github.com/gin-gonic/gin"
	"main.go/schemas"
	"main.go/utils"
	"net/http"
)

func (h *seatHandler) ListUnbookedSeatByFilmIDAndLocationID(ctx *gin.Context) {

	var req schemas.ListSeatRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	filmID := req.FilmID
	locationID := req.LocationID

	seat, err := h.seatService.ListUnbookedSeatByFilmIDAndLocationID(filmID, locationID)
	if err != nil {
		return
	}
	utils.ApiResponse(ctx, http.StatusOK, "success", seat, nil)
}
