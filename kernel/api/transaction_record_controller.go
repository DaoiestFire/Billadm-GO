package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/billadm/api/dto"
	"github.com/billadm/models"
	"github.com/billadm/service"
)

func getTransactionRecord(c *gin.Context) {
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

	//TODO 当前仅能查询账本中的所有的账单
	trs, err := service.GetTrService().ListAllTrByLedgerId(ledgerId)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	jsonData, err := json.Marshal(trs)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = string(jsonData)
	ret.Msg = "success"

	return
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

	tr, ok := trDto.Vo(ret)
	if !ok {
		logrus.Errorf("convert transaction record dto to vo error, err: %v", ret.Msg)
		return
	}

	logrus.Debugf("tr vo: %v", tr)

	trId, err := service.GetTrService().CreateTr(tr)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = trId
	ret.Msg = "success"

	return
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

	ret.Msg = "success"

	return
}
