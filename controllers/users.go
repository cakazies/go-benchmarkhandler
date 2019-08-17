package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
	"github.com/local/go-benchmark/models"
)

func GinGetProfile(c *gin.Context) {
	models.GetProfile()
}

func MuxGetProfile(w http.ResponseWriter, r *http.Request) {
	models.GetProfile()
}

func EchoGetProfile(c echo.Context) error {
	models.GetProfile()
	return nil
}
