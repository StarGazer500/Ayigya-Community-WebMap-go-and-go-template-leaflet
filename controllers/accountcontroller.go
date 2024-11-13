package controllers

import (
	// "Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/inits/db"
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/inits/db"
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/middlewares"
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserDetails struct {
	Firstname string `form:"firstname" json:"firstname"  xml:"firstname" binding:"required"`
	Surname   string `form:"surname" json:"surname"  xml:"surname" binding:"required"`
	Password1 string `form:"password1" json:"password1"  xml:"password1" binding:"required"`
	Password2 string `form:"password2" json:"password2"  xml:"password2" binding:"required"`
	Email     string `form:"email" json:"email"  xml:"email" binding:"required"`
}

type ProfileDeTails struct {
	Firstname string `form:"firstname" json:"firstname"  xml:"firstname" binding:"required"`
	Surname   string `form:"surname" json:"surname"  xml:"surname" binding:"required"`
	Password1 string `form:"password1" json:"password1"  xml:"password1" binding:"required"`
	Email     string `form:"email" json:"email"  xml:"email" binding:"required"`
}

func Profile(ctx *gin.Context) {

	if ctx.Request.Method == http.MethodGet {

		// if err := ctx.ShouldBind(&form); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		// }

		ctx.HTML(http.StatusOK, "profile.html", gin.H{"profilepage": "opened"})
	}

}
func SaveUser(db *sql.DB, user UserDetails) (sql.Result, error) {

	tableName := models.UserModel.TableName

	columns := []string{"firstname", "surname", "email", "password1"}

	data, err := models.InsertOne(db, tableName, columns, user.Firstname, user.Surname, user.Email, user.Password1)

	if err != nil {

		fmt.Println("Insertion Error Occured", err)
	}

	return data, nil
}

// Binding from JSON
type Login struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	// Default empty data for the registration form
	var form UserDetails
	defaultData := UserDetails{
		Firstname: "",
		Surname:   "",
		Password1: "",
		Password2: "",
		Email:     "",
	}

	if ctx.Request.Method == http.MethodPost {

		if err := ctx.ShouldBind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		res, err := SaveUser(db.PG.Db, form)
		if err != nil {
			fmt.Println("db error", err)
			return
		}
		fmt.Println("saved data", res)

		ctx.JSON(http.StatusOK, gin.H{"status": "You have registered Successfully"})

	} else {
		// If the request is not a POST, show the registration form with empty data
		ctx.HTML(http.StatusOK, "register.html", gin.H{
			"data": defaultData,
		})
	}
}

func LoginUser(ctx *gin.Context) {
	// Default empty data for the registration form
	var form Login
	defaultData := Login{
		Email:    "",
		Password: "",
	}

	if ctx.Request.Method == http.MethodPost {

		if err := ctx.ShouldBind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		fmt.Println(db.PG.Db, models.UserModel.TableName, form.Email)

		result, err := models.FindOne(db.PG.Db, models.UserModel.TableName, "email", form.Email)

		if err != nil {
			fmt.Println("Querying errror occured", err)
		}

		for _, user := range result {

			middlewares.Claim = middlewares.TokenClaimStruct{
				MyAuthServer:    "AuthServer",
				AuthUserEmail:   user["email"].(string),
				AuthUserSurname: user["surname"].(string),
				AuthUserId:      user["id"].(string),
				// AuthExp:         middlewares.GenerateExpiryTime(120),
			}

		}

		access_token, aerr := middlewares.GenerateAccessToken(middlewares.Claim)

		refresh_token, rerr := middlewares.GenerateRefreshToken(middlewares.Claim)

		if aerr != nil {
			fmt.Println("token generation error")
			return
		} else if rerr != nil {
			fmt.Println("token generation error")
			return
		}

		// fmt.Println("access", access_token)
		// fmt.Println("refresh", refresh_token)

		ctx.SetCookie("access", access_token, 3600, "/", "localhost", false, true)
		ctx.SetCookie("refresh", refresh_token, 3600, "/", "localhost", false, true)

		// fmt.Println(result, "returned succesfully")

		// ctx.JSON(http.StatusOK, gin.H{"status": "You have logged in Successfully"})
		// ctx.HTML(http.StatusOK, "Profile.html", gin.H{})
		ctx.Redirect(http.StatusFound, "/account/profile")

	} else {
		// If the request is not a POST, show the registration form with empty data
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"data": defaultData,
		})
	}
}
