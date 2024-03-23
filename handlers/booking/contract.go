package bookinghandler

import (
	"github.com/gin-gonic/gin"
	bookingservice "main.go/services/booking"
)

type BookingHandler interface {
	CreateBooking(ctx *gin.Context)
	GetBookingByID(ctx *gin.Context)
}

type bookingHandler struct {
	bookingService bookingservice.BookingService
}

func NewBookingHandler(bookingService bookingservice.BookingService) *bookingHandler {
	return &bookingHandler{bookingService}
}
