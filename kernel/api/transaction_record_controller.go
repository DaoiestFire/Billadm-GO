package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/billadm/kernel/api/dto"
	"github.com/billadm/kernel/models"
	"github.com/billadm/kernel/service"
)

func getTransactionRecord(c *gin.Context) {

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

	ret.Msg = "success"
	ret.Data = trId

	return
}

func updateTransactionRecord(c *gin.Context) {

}

func deleteTransactionRecord(c *gin.Context) {

}
