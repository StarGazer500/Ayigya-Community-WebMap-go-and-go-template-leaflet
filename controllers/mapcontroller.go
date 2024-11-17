package controllers

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapPageDisplay(ctx *gin.Context) {

	if ctx.Request.Method == http.MethodGet {

		// if err := ctx.ShouldBind(&form); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		// }

		ctx.HTML(http.StatusOK, "map.html", gin.H{"profilepage": "map page opened"})
	}

}
