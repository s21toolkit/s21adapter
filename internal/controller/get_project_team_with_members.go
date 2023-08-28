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
		g.POST("/GetProjectTeamWithMembers", c.Handle_GetProjectTeamWithMembers).
			AddParamBody(requests.Variables_GetProjectTeamWithMembers{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_GetProjectTeamWithMembers{}, nil)
	})
}

func (a *AdapterController) Handle_GetProjectTeamWithMembers(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_GetProjectTeamWithMembers `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetProjectTeamWithMembers(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}