package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/billadm/models"
	"github.com/billadm/service"
	"github.com/billadm/workspace"
)

func queryAllCategory(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)

	ws := workspace.Manager.OpenedWorkspace()
	if ws == nil {
		ret.Code = -1
		ret.Msg = workspace.ErrOpenedWorkspaceNotFoundMsg
		return
	}

	categories, err := service.GetCategoryService().QueryAllCategory(ws)
	if err != nil {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}

	ret.Data = categories
}
