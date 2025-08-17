package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/billadm/models"
	"github.com/billadm/models/dto"
	"github.com/billadm/service"
)

func queryAllTrs(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)

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

	trs, err := service.GetTrService().ListAllTrsByLedgerId(ledgerId)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = trs
}

func queryTrsByPage(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)

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

	offset, ok := arg["offset"].(float64)
	if !ok {
		ret.Code = -1
		ret.Msg = "offset field not exist in request body"
		return
	}

	limit, ok := arg["limit"].(float64)
	if !ok {
		ret.Code = -1
		ret.Msg = "limit field not exist in request body"
		return
	}

	trs, err := service.GetTrService().QueryTrsByPage(ledgerId, int(offset), int(limit))
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = trs
}

func createTransactionRecord(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)

	trDto, ok := dto.JsonTransactionRecordDto(c, ret)
	if !ok {
		return
	}
	logrus.Debugf("tr dto: %v", trDto)

	if !trDto.Validate(ret) {
		logrus.Errorf("validate transaction record error, err: %v", ret.Msg)
		return
	}

	trId, err := service.GetTrService().CreateTr(trDto)
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

	err := service.GetTrService().DeleteTrById(trId)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}
}
