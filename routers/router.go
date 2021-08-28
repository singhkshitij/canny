package routers

import (
	v1 "canny/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "pong"})
}

func welcomePage(c *gin.Context) {
	c.String(http.StatusOK,"Welcome to canny! Visit swagger for more details..")
}

func InitRouter() (r *gin.Engine) {
	r = gin.Default()
	r.GET("/", welcomePage)
	r.GET("/ping" , ping)

	apiv1 := r.Group("api/v1/")
	apiv1.GET("currencies", v1.Currencies)
	return
}
