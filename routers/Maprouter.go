package routers

import (
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/controllers"

	"github.com/gin-gonic/gin"
)

func MapRoutes(route *gin.RouterGroup) {
	route.GET("/map-display", controllers.MapPageDisplay)

}
