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
		g.POST("/GetDashboardBuildings", c.Handle_GetDashboardBuildings).
			SetOperationId("GetDashboardBuildings").
			AddParamBody(requests.GetDashboardBuildings_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.GetDashboardBuildings_Data{}, nil)
	})
}

func (a *AdapterController) Handle_GetDashboardBuildings(c echo.Context) (err error) {
	var data struct {
		Variables requests.GetDashboardBuildings_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetDashboardBuildings(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}