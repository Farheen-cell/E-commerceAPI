package controllers

import (
	"errors"
	"net/http"

	"github.com/Farheen-cell/ecom/dtos"
	"github.com/Farheen-cell/ecom/services"
	"github.com/gin-gonic/gin"
)

func RegisterPageRoutes(router *gin.RouterGroup) {
	router.GET("", Home)
	router.GET("/home", Home)

}

func Home(c *gin.Context) {

	tags, err := services.FetchAllTags()
	categories, err := services.FetchAllCategories()
	if err != nil {
		c.JSON(http.StatusNotFound, dtos.CreateDetailedErrorDto("comments", errors.New("Somethign went wrong")))
		return
	}

	c.JSON(http.StatusOK, dtos.CreateHomeResponse(tags, categories))
}
