package tables

import(
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/models"
)




func CreateAllTablesIfNotExist() {
	// create user table
    // CreateTable(models.UserSQLModel)
	models.CreateUserTable()
	
}

