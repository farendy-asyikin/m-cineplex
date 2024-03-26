package locationhandler

import (
	"github.com/gin-gonic/gin"
	locationservice "main.go/services/location"
)

type LocationHandler interface {
	ListLocation(ctx *gin.Context)
}

type locationHandler struct {
	locationService locationservice.LocationService
}

func NewLocationHandler(locationService locationservice.LocationService) *locationHandler {
	return &locationHandler{locationService}
}
