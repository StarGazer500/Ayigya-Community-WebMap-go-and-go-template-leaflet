package routers

import (
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/controllers"

	"github.com/gin-gonic/gin"
)

func MapRoutes(route *gin.RouterGroup) {
	route.GET("/map-display", controllers.MapPageDisplay)
	route.GET("/featurelayers", controllers.FeatureLayers)
	route.POST("/featureattributes", controllers.FeatreAttributes)
	route.POST("/featureoperatures", controllers.SelectOperator)
	route.POST("/makeqquery", controllers.MakeQuery)
	route.POST("/simplesearch", controllers.SimpleSearch)

}
