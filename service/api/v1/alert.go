package v1

import (
	"canny/domain"
	"canny/model"
	"canny/pkg/err"
	"canny/pkg/validate"
	"github.com/gin-gonic/gin"
)

// @Summary Create an alert for user
// @Produce json
// @Param alert body model.CreateAlertRequest true "Alert Alerts"
// @Success 200 {object} model.CreateAlertResponse
// @Failure 500 {object} model.Response
// @Router /api/v1/alerts [post]
// @tags alerts
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

// @Summary Get all alerts for user
// @Produce json
// @Success 200 {object} model.GetAllAlertResponse
// @Failure 500 {object} model.Response
// @Router /api/v1/alerts [get]
// @tags alerts
func GetAlerts(c *gin.Context) {
	appG := model.Gin{C: c}
	data := domain.GetAllAlerts()
	appG.Response(200, err.Success, data)
}

// @Summary Get alert details
// @Produce json
// @Param id path string true "alert id"
// @Success 200 {object} model.CreateAlertResponse
// @Failure 500 {object} model.Response
// @Router /api/v1/alerts/{id} [get]
// @tags alerts
func GetAlert(c *gin.Context) {
	appG := model.Gin{C: c}
	alertId := c.Param("id")
	data, er, code := domain.GetAlert(alertId)
	if code == err.NotFound {
		appG.Response(404, err.NotFound, map[string]string{})
	} else if er != nil {
		appG.Response(500, err.Error, map[string]string{})
	} else {
		appG.Response(200, err.Success, data)
	}
}

// @Summary Delete active alert
// @Produce json
// @Param id path string true "alert id"
// @Success 204
// @Failure 500 {object} model.Response
// @Router /api/v1/alerts/{id} [delete]
// @tags alerts
func DeleteAlert(c *gin.Context) {
	appG := model.Gin{C: c}
	alertId := c.Param("id")
	er, code := domain.DeleteAlert(alertId)
	if code == 404 {
		appG.Response(404, err.NotFound, map[string]string{})
	} else if er != nil {
		appG.Response(500, err.Error, map[string]string{})
	} else {
		appG.Response(204, err.Success, nil)
	}
}

// @Summary Dry run an alert
// @Produce json
// @Param alert body model.CreateAlertRequest true "Alert Data"
// @Success 200 {object} model.DryRunAlertResponse
// @Failure 500 {object} model.Response
// @Router /api/v1/alerts/dry-run [post]
// @tags test
func DryRun(c *gin.Context) {
	appG := model.Gin{C: c}
	var requestBody model.CreateAlertRequest
	_ = c.BindJSON(&requestBody)
	validationErr := validate.Validate(requestBody)
	if validationErr != nil {
		appG.Response(400, err.BadRequest, validationErr.Error())
	} else {
		status := domain.DryRunAlert(requestBody)
		appG.Response(200, err.Success, &model.DryRunAlertStatus{Passed: status})
	}
}
