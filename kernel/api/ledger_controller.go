package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/billadm/kernel/constant"
	"github.com/billadm/kernel/models"
	"github.com/billadm/kernel/service"
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

	userId := constant.DefaultUUID
	userName, ok := arg["username"].(string)
	if ok {
		logrus.Infof("create ledger with username: %v", userName)
		// TODO: 检查用户是否存在，如果存在则使用指定用户的用户id
	}

	// 返回全部的账本信息
	var ledgers []models.Ledger
	var err error
	var jsonData []byte
	if ledgerId == constant.All {
		ledgers, err = service.GetLedgerService().ListAllLedger(userId)
		if err != nil {
			ret.Code = -1
			ret.Msg = err.Error()
			return
		}
	} else {
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

	userId := constant.DefaultUUID
	userName, ok := arg["username"].(string)
	if ok {
		logrus.Infof("create ledger with username: %v", userName)
		// TODO: 检查用户是否存在，如果存在则使用指定用户的用户id
	}
	ledgerId, err := service.GetLedgerService().CreateLedger(ledgerName, userId)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Msg = fmt.Sprintf("create ledger successfully, ledger id: %s", ledgerId)

	return
}

func updateLedger(c *gin.Context) {

}

func deleteLedger(c *gin.Context) {

}
