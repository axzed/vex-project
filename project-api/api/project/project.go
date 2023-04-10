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
	"strconv"
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
	selectBy := ctx.PostForm("selectBy")
	// 构造grpc请求参数
	msg := &project.ProjectRpcMessage{
		MemberId:   memberId.(int64),
		MemberName: memberName,
		SelectBy:   selectBy,
		Page:       page.Page,
		PageSize:   page.PageSize,
	}
	// 调用grpc服务
	myProjectResponse, err := rpc.ProjectServiceClient.FindProjectByMemId(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}

	var pam []*param.ProjectAndMember
	err = copier.Copy(&pam, myProjectResponse.Pm)
	// 若为空 设置默认值
	if pam == nil {
		pam = []*param.ProjectAndMember{}
	}
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy参数失败"))
	}
	ctx.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pam, // 不能返回null nil 前端无法解析 -> []
		"total": myProjectResponse.Total,
	}))
}

func (p *HandlerProject) projectTemplate(ctx *gin.Context) {
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
	viewTypeStr := ctx.PostForm("viewType")
	viewType, _ := strconv.ParseInt(viewTypeStr, 10, 64)
	msg := &project.ProjectRpcMessage{
		MemberId:         memberId.(int64),
		MemberName:       memberName,
		ViewType:         int32(viewType),
		Page:             page.Page,
		PageSize:         page.PageSize,
		OrganizationCode: ctx.GetString("organizationCode"),
	}
	templateResponse, err := rpc.ProjectServiceClient.FindProjectTemplate(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}

	var pts []*param.ProjectTemplate
	err = copier.Copy(&pts, templateResponse.Ptm)
	// 若为空 设置默认值
	if pts == nil {
		pts = []*param.ProjectTemplate{}
	}
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy参数失败"))
	}
	for _, v := range pts {
		// 若为空 设置默认值
		if v.TaskStages == nil {
			v.TaskStages = []*param.TaskStagesOnlyName{}
		}
	}
	ctx.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pts, // 不能返回null nil 前端无法解析 -> []
		"total": templateResponse.Total,
	}))
}

// projectSave 项目保存(新建项目)
func (p *HandlerProject) projectSave(ctx *gin.Context) {
	// 返回结果的结构体
	result := &common.Result{}
	// 1. 获取参数
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId, _ := ctx.Get("memberId")
	organizationCodeStr := ctx.GetString("organizationCode")
	var req *param.SaveProjectRequest
	ctx.ShouldBind(&req)
	msg := &project.ProjectRpcMessage{
		MemberId:         memberId.(int64),
		OrganizationCode: organizationCodeStr,
		TemplateCode:     req.TemplateCode,
		Name:             req.Name,
		Id:               int64(req.Id),
		Description:      req.Description,
	}
	saveProject, err := rpc.ProjectServiceClient.SaveProject(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var rsp *param.SaveProject
	copier.Copy(&rsp, saveProject)
	ctx.JSON(http.StatusOK, result.Success(rsp))
}

// readProject 项目详情
func (p *HandlerProject) readProject(ctx *gin.Context) {
	result := &common.Result{}
	projectCode := ctx.PostForm("projectCode")
	memberId := ctx.GetInt64("memberId")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	detail, err := rpc.ProjectServiceClient.FindProjectDetail(c, &project.ProjectRpcMessage{
		ProjectCode: projectCode,
		MemberId:    memberId,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}
	pd := &param.ProjectDetail{}
	copier.Copy(&pd, detail)
	ctx.JSON(http.StatusOK, result.Success(pd))
}

// recycleProject 项目回收
func (p *HandlerProject) recycleProject(ctx *gin.Context) {
	result := &common.Result{}
	projectCode := ctx.PostForm("projectCode")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 调用 rpc 的回收项目服务
	_, err := rpc.ProjectServiceClient.UpdateDeletedProject(c, &project.ProjectRpcMessage{
		ProjectCode: projectCode,
		Deleted:     true, // 逻辑删除
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}
	ctx.JSON(http.StatusOK, result.Success([]int{}))
}

// recoveryProject 项目恢复(从回收站恢复)
func (p *HandlerProject) recoveryProject(ctx *gin.Context) {
	result := &common.Result{}
	projectCode := ctx.PostForm("projectCode")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 调用 rpc 的回收项目服务
	_, err := rpc.ProjectServiceClient.UpdateDeletedProject(c, &project.ProjectRpcMessage{
		ProjectCode: projectCode,
		Deleted:     false, // 不删除
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}
	ctx.JSON(http.StatusOK, result.Success([]int{}))
}

func (p *HandlerProject) collectProject(ctx *gin.Context) {
	result := &common.Result{}
	projectCode := ctx.PostForm("projectCode")
	collectType := ctx.PostForm("type")
	memberId := ctx.GetInt64("memberId")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 调用 rpc 的服务
	_, err := rpc.ProjectServiceClient.UpdateCollectProject(c, &project.ProjectRpcMessage{
		ProjectCode: projectCode,
		CollectType: collectType,
		MemberId:    memberId,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}
	ctx.JSON(http.StatusOK, result.Success([]int{}))
}

// editProject 编辑项目API
func (p *HandlerProject) editProject(ctx *gin.Context) {
	// 获取请求参数
	result := &common.Result{}
	var req *param.ProjectReq
	_ = ctx.ShouldBind(&req)
	memberId := ctx.GetInt64("memberId")
	// 设置rpc服务超时时间
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 构造rpc请求参数
	msg := &project.UpdateProjectMessage{}
	copier.Copy(msg, req)
	msg.MemberId = memberId
	// 调用rpc服务
	_, err := rpc.ProjectServiceClient.UpdateProject(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
	}
	// 返回响应
	ctx.JSON(http.StatusOK, result.Success([]int{}))
}

// getLogBySelfProject 获取自己的项目日志
func (p *HandlerProject) getLogBySelfProject(c *gin.Context) {
	result := &common.Result{}
	var page = &model.Page{}
	page.Bind(c)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project.ProjectRpcMessage{
		MemberId: c.GetInt64("memberId"),
		Page:     page.Page,
		PageSize: page.PageSize,
	}
	projectLogResponse, err := rpc.ProjectServiceClient.GetLogBySelfProject(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var list []*model.ProjectLog
	copier.Copy(&list, projectLogResponse.List)
	if list == nil {
		list = []*model.ProjectLog{}
	}
	c.JSON(http.StatusOK, result.Success(list))
}

// nodeList 获取访问节点
func (p *HandlerProject) nodeList(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := rpc.ProjectServiceClient.NodeList(ctx, &project.ProjectRpcMessage{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var list []*model.ProjectNodeTree
	copier.Copy(&list, response.Nodes)
	c.JSON(http.StatusOK, result.Success(gin.H{
		"nodes": list,
	}))
}

func (p *HandlerProject) FindProjectByMemberId(memberId int64, projectCode string, taskCode string) (*param.Project, bool, bool, *errs.BError) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project.ProjectRpcMessage{
		MemberId:    memberId,
		ProjectCode: projectCode,
		TaskCode:    taskCode,
	}
	projectResponse, err := rpc.ProjectServiceClient.FindProjectByMemberId(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		return nil, false, false, errs.NewError(errs.ErrorCode(code), msg)
	}
	if projectResponse.Project == nil {
		return nil, false, false, nil
	}
	pr := &param.Project{}
	copier.Copy(pr, projectResponse.Project)
	return pr, true, projectResponse.IsOwner, nil
}
