package project

import (
	"context"
	"github.com/axzed/project-api/api/rpc"
	"github.com/axzed/project-api/pkg/model"
	"github.com/axzed/project-api/pkg/model/param"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/task"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

// HandlerTask 任务处理器
type HandlerTask struct {
}

func NewTask() *HandlerTask {
	return &HandlerTask{}
}

// taskStages 任务阶段api handler
func (t *HandlerTask) taskStages(ctx *gin.Context) {
	result := &common.Result{}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 获取参数
	projectCode := ctx.PostForm("projectCode")
	page := &model.Page{}
	page.Bind(ctx)

	// 调用grpc接口
	msg := &task.TaskReqMessage{
		ProjectCode: projectCode,
		Page:        page.Page,
		PageSize:    page.PageSize,
	}
	stages, err := rpc.TaskServiceClient.TaskStages(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}

	// 处理返回参数 (copier && 特殊返回值)
	var list []*param.TaskStagesResp
	copier.Copy(&list, stages.List)
	if list == nil {
		list = []*param.TaskStagesResp{}
	}
	for _, v := range list {
		v.TasksLoading = true  //任务加载状态
		v.FixedCreator = false //添加任务按钮定位
		v.ShowTaskCard = false //是否显示创建卡片
		v.Tasks = []int{}
		v.DoneTasks = []int{}
		v.UnDoneTasks = []int{}
	}

	// 返回结果
	ctx.JSON(http.StatusOK, result.Success(gin.H{
		"list":  list,
		"total": stages.Total,
		"page":  page.Page,
	}))

}
