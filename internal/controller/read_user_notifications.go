package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
	"github.com/s21toolkit/s21client/requests"
)

func init() {
	registerMethod(func(g echoswagger.ApiGroup, c *AdapterController) {
		g.POST("/ReadUserNotifications", c.Handle_ReadUserNotifications).
			SetOperationId("ReadUserNotifications").
			AddParamBody(requests.ReadUserNotifications_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.ReadUserNotifications_Data{}, nil)
	})
}

func (a *AdapterController) Handle_ReadUserNotifications(c echo.Context) (err error) {
	var data struct {
		Variables requests.ReadUserNotifications_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().ReadUserNotifications(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}