package filmhandler

import (
	"github.com/gin-gonic/gin"
	"main.go/schemas"
	"main.go/utils"
	"net/http"
)

func (h *filmHandler) CreateFilm(ctx *gin.Context) {
	var request schemas.CreateFilmRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	film, err := h.filmService.CreateFilm(request)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", film, nil)
}

func (h *filmHandler) GetFilmByID(ctx *gin.Context) {
	id := ctx.Param("id")

	film, err := h.filmService.GetFilmByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", film, nil)
}

func (h *filmHandler) ListFilm(ctx *gin.Context) {
	film, err := h.filmService.ListFilm()
	if err != nil {
		return
	}
	utils.ApiResponse(ctx, http.StatusOK, "success", film, nil)
}

func (h *filmHandler) UpdateFilm(ctx *gin.Context) {
	id := ctx.Param("id")
	var request schemas.UpdateFilmRequest

	err := ctx.ShouldBind(&request)
	if err != nil {

		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	film, err := h.filmService.GetFilmByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	film, err = h.filmService.UpdateFilm(request, *film)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", film, nil)
}

func (h *filmHandler) DeleteFilmByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	_, err := h.filmService.GetFilmByID(ID)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	err = h.filmService.DeleteFilmByID(ID)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", nil, nil)
}
