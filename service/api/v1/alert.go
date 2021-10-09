package v1

import (
	"canny/domain"
	"canny/model"
	"canny/pkg/err"
	"canny/pkg/validate"
	"github.com/gin-gonic/gin"
)

// @Summary Get coin data
// @Produce json
// @Param alert body model.CreateAlertRequest true "Alert Data"
// @Success 200 {object} model.CreateAlertResponse
// @Failure 500 {object} model.Response
// @Router /api/v1/alert [post]
// @tags alert
func CreateAlert(c *gin.Context) {
	appG := model.Gin{C: c}
	var requestBody model.CreateAlertRequest
	_ = c.BindJSON(&requestBody)
	validationErr := validate.Validate(requestBody)
	if validationErr != nil {
		appG.Response(400, err.BadRequest, validationErr.Error())
	} else {
		savedData, er := domain.CreateAlert(requestBody)
		if er != nil {
			appG.Response(500, err.Error, map[string]string{})
		} else {
			appG.Response(200, err.Success, savedData)
		}
	}
}
