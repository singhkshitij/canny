package service

import (
	_ "canny/docs"
	v1 "canny/service/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @Summary health check endpoint
// @Produce json
// @Success 200 {object} meta.Response
// @Failure 500 {object} meta.Response
// @Router /ping [get]
// @tags health-check
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
	apiV1.GET("currencies/price", v1.AllCurrencyData)
	apiV1.GET("currencies/:currency", v1.CurrencyData)
	return
}
