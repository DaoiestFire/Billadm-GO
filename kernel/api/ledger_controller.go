package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/billadm/constant"
	"github.com/billadm/models"
	"github.com/billadm/service"
)

func getLedger(c *gin.Context) {
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

	var ledgers []models.Ledger
	var err error
	var jsonData []byte
	if ledgerId == constant.All {
		// 返回全部的账本信息
		ledgers, err = service.GetLedgerService().ListAllLedger()
		if err != nil {
			ret.Code = -1
			ret.Msg = err.Error()
			return
		}
	} else {
		// 查询指定id的账本
		ledgerIds := strings.Split(ledgerId, ",")
		for _, id := range ledgerIds {
			id = strings.TrimSpace(id)
			ledger, err := service.GetLedgerService().QueryLedgerById(id)
			if err != nil {
				logrus.Errorf("query ledger by id: %s err: %v", id, err)
				ret.Code = -1
				ret.Msg = fmt.Sprintf("ledger not found, id: %s", id)
				return
			}
			ledgers = append(ledgers, *ledger)
		}
	}

	if len(ledgers) == 1 {
		jsonData, err = json.Marshal(ledgers[0])
		if err != nil {
			ret.Code = -1
			ret.Msg = err.Error()
			return
		}
	} else {
		jsonData, err = json.Marshal(ledgers)
		if err != nil {
			ret.Code = -1
			ret.Msg = err.Error()
			return
		}
	}

	ret.Data = string(jsonData)
	ret.Msg = "success"

	return
}

func createLedger(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)

	arg, ok := JsonArg(c, ret)
	if !ok {
		return
	}

	ledgerName, ok := arg["name"].(string)
	if !ok {
		ret.Code = -1
		ret.Msg = "name field not exist in request body"
		return
	}

	// 在指定用户下创建账本
	ledgerId, err := service.GetLedgerService().CreateLedger(ledgerName)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = ledgerId

	return
}

func updateLedger(c *gin.Context) {

}

func deleteLedger(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)

	arg, ok := JsonArg(c, ret)
	if !ok {
		return
	}

	trId, ok := arg["ledger_id"].(string)
	if !ok {
		ret.Code = -1
		ret.Msg = "ledger_id field not exist in request body"
		return
	}

	err := service.GetLedgerService().DeleteLedgerById(trId)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Msg = "success"

	return
}
