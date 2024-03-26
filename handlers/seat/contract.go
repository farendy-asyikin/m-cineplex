package seathandler

import (
	"github.com/gin-gonic/gin"
	seatservice "main.go/services/seat"
)

type SeatHandler interface {
	ListUnbookedSeatByFilmIDAndLocationID(ctx *gin.Context)
}

type seatHandler struct {
	seatService seatservice.SeatService
}

func NewSeatHandler(seatService seatservice.SeatService) *seatHandler {
	return &seatHandler{seatService}
}
