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
		g.POST("/GetDismissInfoByStudentId", c.Handle_GetDismissInfoByStudentId).
			SetOperationId("GetDismissInfoByStudentId").
			AddParamBody(requests.GetDismissInfoByStudentId_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.GetDismissInfoByStudentId_Data{}, nil)
	})
}

func (a *AdapterController) Handle_GetDismissInfoByStudentId(c echo.Context) (err error) {
	var data struct {
		Variables requests.GetDismissInfoByStudentId_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetDismissInfoByStudentId(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}