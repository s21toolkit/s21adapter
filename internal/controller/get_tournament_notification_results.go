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
		g.POST("/GetTournamentNotificationResults", c.Handle_GetTournamentNotificationResults).
			SetOperationId("GetTournamentNotificationResults").
			AddParamBody(requests.GetTournamentNotificationResults_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.GetTournamentNotificationResults_Data{}, nil)
	})
}

func (a *AdapterController) Handle_GetTournamentNotificationResults(c echo.Context) (err error) {
	var data struct {
		Variables requests.GetTournamentNotificationResults_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetTournamentNotificationResults(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}