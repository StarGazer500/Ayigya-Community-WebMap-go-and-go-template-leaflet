package routers

import (
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.RouterGroup) {
	// route.GET("/profile", controllers.Profile)
	route.POST("/register", controllers.Register)
	route.GET("/register", controllers.Register)
	route.POST("/login", controllers.LoginUser)
	route.GET("/login", controllers.LoginUser)

}
