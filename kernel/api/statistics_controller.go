package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/billadm/models"
	"github.com/billadm/models/dto"
	"github.com/billadm/service"
	"github.com/billadm/workspace"
)

func getTrStatistics(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)

	ws := workspace.Manager.OpenedWorkspace()
	if ws == nil {
		ret.Code = -1
		ret.Msg = workspace.ErrOpenedWorkspaceNotFoundMsg
		return
	}

	queryCondition, ok := dto.JsonQueryCondition(c, ret)
	if !ok {
		return
	}
	logrus.Debugf("query condition: %v", queryCondition)

	ts, err := service.GetTrService().QueryTrStatisticsOnCondition(ws, queryCondition)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = ts
}
