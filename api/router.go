package api

import (
	"Kitchenizer/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start(s service.Service) error {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Yummy yummy in my tummy !")
	})

	router.POST("/recipe", insertRecipe(s))

	return router.Run("localhost:8080")
}
