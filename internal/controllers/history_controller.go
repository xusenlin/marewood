package controllers

import (
	"marewood/internal/pkg/context"
	"marewood/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HistoryController struct {
	historyService services.HistoryService
}

func NewHistoryController(historyService services.HistoryService) *HistoryController {
	return &HistoryController{historyService: historyService}
}

func (ctrl *HistoryController) FindByTaskID(c *gin.Context) {
	ctx := context.New(c)
	taskId, err := strconv.ParseUint(c.Query("taskId"), 10, 32)
	if err != nil {
		ctx.SendErr(err)
		return
	}
	histories, err := ctrl.historyService.FindByTaskID(uint(taskId))
	if err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("success", histories)
}

func (ctrl *HistoryController) Restore(c *gin.Context) {
	ctx := context.New(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		ctx.SendErr(err)
		return
	}
	err = ctrl.historyService.Restore(uint(id))
	if err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("success", nil)
}
