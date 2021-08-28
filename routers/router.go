package routers

import (
	_ "canny/docs"
	v1 "canny/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "pong"})
}

func welcomePage(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to canny! Visit swagger for more details..")
}

func InitRouter() (r *gin.Engine) {
	r = gin.Default()

	swaggerUrl := ginSwagger.URL("/swagger/doc.json")

	r.GET("/", welcomePage)
	r.GET("/ping", ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))

	apiV1 := r.Group("api/v1/")
	apiV1.GET("currencies", v1.Currencies)
	return
}
