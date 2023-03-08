package project

import (
	"context"
	"github.com/axzed/project-api/api/rpc"
	"github.com/axzed/project-api/pkg/model"
	"github.com/axzed/project-api/pkg/model/param"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/project"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

// HandlerProject Handler接口实现类
type HandlerProject struct {
}

// NewHandlerProject NewHandlerUser
func NewHandlerProject() *HandlerProject {
	return &HandlerProject{}
}

// index 首页展示
func (p *HandlerProject) index(ctx *gin.Context) {
	result := &common.Result{}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project.IndexMessage{}
	indexResponse, err := rpc.ProjectServiceClient.Index(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}
	menus := indexResponse.Menus
	var ms []*param.Menu
	copier.Copy(&ms, menus)
	ctx.JSON(http.StatusOK, result.Success(ms))
}

// myProjectList 我的项目列表
func (p *HandlerProject) myProjectList(ctx *gin.Context) {
	// 返回结果的结构体
	result := &common.Result{}
	// 1. 获取参数
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId, _ := ctx.Get("memberId")
	memberName := ctx.GetString("memberName")
	// 分页
	page := &model.Page{}
	page.Bind(ctx)
	msg := &project.ProjectRpcMessage{MemberId: memberId.(int64), MemberName: memberName, Page: page.Page, PageSize: page.PageSize}
	myProjectResponse, err := rpc.ProjectServiceClient.FindProjectByMemId(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}
	// 若为空 设置默认值
	if myProjectResponse.Pm == nil {
		myProjectResponse.Pm = []*project.ProjectMessage{}
	}
	var pms []*param.ProjectAndMember
	err = copier.Copy(&pms, myProjectResponse.Pm)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy参数失败"))
	}
	ctx.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pms,
		"total": myProjectResponse.Total,
	}))
}
