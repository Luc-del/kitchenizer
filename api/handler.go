package api

import (
	"Kitchenizer/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func insertRecipe(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, nil)
	}
}
