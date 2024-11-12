package main

import (
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/inits/db"
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/inits/tables"

	// "fmt"

	// "Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/models"
	// "fmt"
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	db.InitpgDb()
	tables.CreateAllTablesIfNotExist()
}

func deinit() {

	db.PG.Db.Close()

}

func main() {
	// Initialize Gin engine

	defer deinit()

	engine := gin.Default()

	// Set up account routes
	accountGroup := engine.Group("/account")
	routers.UserRoutes(accountGroup) // Assuming you have a UserRoutes function to define routes

	// Load HTML templates (make sure the template path is correct)
	engine.LoadHTMLGlob("./views/templates/*.html")

	// Serve static files (adjust the paths if necessary)
	engine.Static("/assets", "./views/staticfiles")

	// Start the server
	engine.Run(":8080") // This starts the server on http://localhost:8080

}
