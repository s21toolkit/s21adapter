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
		g.POST("/GetCampusWorkstation", c.Handle_GetCampusWorkstation).
			SetOperationId("GetCampusWorkstation").
			AddParamBody(requests.GetCampusWorkstation_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.GetCampusWorkstation_Data{}, nil)
	})
}

func (a *AdapterController) Handle_GetCampusWorkstation(c echo.Context) (err error) {
	var data struct {
		Variables requests.GetCampusWorkstation_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetCampusWorkstation(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}