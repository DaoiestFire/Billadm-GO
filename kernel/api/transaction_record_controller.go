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

func queryAllTrs(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)

	ws := workspace.Manager.OpenedWorkspace()
	if ws == nil {
		ret.Code = -1
		ret.Msg = workspace.ErrOpenedWorkspaceNotFoundMsg
		return
	}

	arg, ok := JsonArg(c, ret)
	if !ok {
		return
	}

	ledgerId, ok := arg["ledger_id"].(string)
	if !ok {
		ret.Code = -1
		ret.Msg = "ledger_id field not exist in request body"
		return
	}

	trs, err := service.GetTrService().ListAllTrsByLedgerId(ws, ledgerId)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = trs
}

func queryTrsOnCondition(c *gin.Context) {
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

	trs, err := service.GetTrService().QueryTrsOnCondition(ws, queryCondition)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = trs
}

func queryTrCountOnCondition(c *gin.Context) {
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

	cnt, err := service.GetTrService().QueryTrCountOnCondition(ws, queryCondition)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = cnt
}

func createTransactionRecord(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)

	ws := workspace.Manager.OpenedWorkspace()
	if ws == nil {
		ret.Code = -1
		ret.Msg = workspace.ErrOpenedWorkspaceNotFoundMsg
		return
	}

	trDto, ok := dto.JsonTransactionRecordDto(c, ret)
	if !ok {
		return
	}
	logrus.Debugf("tr dto: %v", trDto)

	if !trDto.Validate(ret) {
		logrus.Errorf("validate transaction record error, err: %v", ret.Msg)
		return
	}

	trId, err := service.GetTrService().CreateTr(ws, trDto)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = trId
}

func updateTransactionRecord(c *gin.Context) {

}

func deleteTransactionRecord(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)

	ws := workspace.Manager.OpenedWorkspace()
	if ws == nil {
		ret.Code = -1
		ret.Msg = workspace.ErrOpenedWorkspaceNotFoundMsg
		return
	}

	arg, ok := JsonArg(c, ret)
	if !ok {
		return
	}

	trId, ok := arg["tr_id"].(string)
	if !ok {
		ret.Code = -1
		ret.Msg = "tr_id field not exist in request body"
		return
	}

	err := service.GetTrService().DeleteTrById(ws, trId)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}
}
