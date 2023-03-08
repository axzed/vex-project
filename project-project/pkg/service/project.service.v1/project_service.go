package project_service_v1

import (
	"context"
	"github.com/axzed/project-common/encrypts"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-common/tms"
	"github.com/axzed/project-grpc/project"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/data/menu"
	"github.com/axzed/project-project/internal/data/mproject"
	"github.com/axzed/project-project/internal/data/mtask"
	"github.com/axzed/project-project/internal/database/interface/transaction"
	"github.com/axzed/project-project/internal/repo"
	"github.com/axzed/project-project/pkg/model"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"strconv"
)

// ProjectService 项目服务
type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache              // 缓存
	transaction transaction.Transaction // 事务
	menuRepo    repo.MenuRepo
	projectRepo repo.ProjectRepo
	projectTemplateRepo repo.ProjectTemplateRepo
	taskStagesTemplateRepo repo.TaskStagesTemplateRepo
}

// NewProjectService 初始化页面展示服务
func NewProjectService() *ProjectService {
	return &ProjectService{
		// 为定义的接口赋上实现类
		cache:       dao.Rc,
		transaction: dao.NewTransactionImpl(),
		menuRepo:    dao.NewMenuDao(),
		projectRepo: dao.NewProjectDao(),
		projectTemplateRepo: dao.NewProjectTemplateDao(),
		taskStagesTemplateRepo: dao.NewTaskStagesTemplateDao(),
	}
}

// Index 项目列表 具体rpc服务实现
func (p *ProjectService) Index(context.Context, *project.IndexMessage) (*project.IndexResponse, error) {
	menus, err := p.menuRepo.FindMenus(context.Background())
	if err != nil {
		zap.L().Error("Show Index db error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	childs := menu.CovertChild(menus)
	var menuMessages []*project.MenuMessage
	copier.Copy(&menuMessages, childs)
	return &project.IndexResponse{Menus: menuMessages}, nil
}

func (p *ProjectService) FindProjectByMemId(ctx context.Context, msg *project.ProjectRpcMessage) (*project.MyProjectResponse, error) {
	// 获取参数
	memberId := msg.MemberId
	page := msg.Page
	pageSize := msg.PageSize
	var pms []*mproject.ProAndMember
	var total int64
	var err error
	// 通过SelectBy参数判断调用哪个服务
	if msg.SelectBy == "" || msg.SelectBy == "my" {
		pms, total, err = p.projectRepo.FindProjectByMemId(ctx, memberId, "", page, pageSize)
	}
	if msg.SelectBy == "archived" {
		pms, total, err = p.projectRepo.FindProjectByMemId(ctx, memberId, "and archived = 1", page, pageSize)
	}
	if msg.SelectBy == "deleted" {
		pms, total, err = p.projectRepo.FindProjectByMemId(ctx, memberId, "and deleted = 1", page, pageSize)
	}
	if msg.SelectBy == "collect" {
		pms, total, err = p.projectRepo.FindCollectProjectByMemId(ctx, memberId, page, pageSize)
	}
	// 调用服务
	if err != nil {
		zap.L().Error("menu FindProjectByMember error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	if pms == nil {
		return &project.MyProjectResponse{Pm: []*project.ProjectMessage{}, Total: total}, nil
	}
	// 处理返回值
	var pmm []*project.ProjectMessage
	copier.Copy(&pmm, pms)
	for _, v := range pmm {
		// 加密
		v.Code, _ = encrypts.EncryptInt64(v.Id, model.AESKey)
		pam := mproject.ToMap(pms)[v.Id]
		v.AccessControlType = pam.GetAccessControlType()
		v.OrganizationCode, _ = encrypts.EncryptInt64(pam.OrganizationCode, model.AESKey)
		v.JoinTime = tms.FormatByMill(pam.JoinTime)
		v.OwnerName = msg.MemberName
		v.Order = int32(pam.Sort)
		v.CreateTime = tms.FormatByMill(pam.CreateTime)
	}
	return &project.MyProjectResponse{Pm: pmm, Total: total}, nil
}

// FindProjectTemplate 获取项目模板
func (p *ProjectService) FindProjectTemplate(ctx context.Context, msg *project.ProjectRpcMessage) (*project.ProjectTemplateResponse, error) {
	organizationCodeStr, _ := encrypts.Decrypt(msg.OrganizationCode, model.AESKey)
	organizationCode, _ := strconv.ParseInt(organizationCodeStr, 10, 64)
	page := msg.Page
	pageSize := msg.PageSize
	var pts []mproject.ProjectTemplate
	var total int64
	var err error
	// 1. 根据viewType去查询项目模板表 得到List
	if msg.ViewType == -1 { // 查询全部模板
		pts, total, err = p.projectTemplateRepo.FindProjectTemplateAll(ctx, organizationCode, page, pageSize)
	}
	if msg.ViewType == 0 { // 查询自定义模板
		pts, total, err = p.projectTemplateRepo.FindProjectTemplateCustom(ctx, msg.MemberId, organizationCode, page, pageSize)
	}
	if msg.ViewType == 1 { // 查询系统模板
		pts, total, err = p.projectTemplateRepo.FindProjectTemplateSystem(ctx, page, pageSize)
	}
	if err != nil {
		zap.L().Error("FindProjectTemplate FindProjectTemplate error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	// 2. 模型转换 拿到模板id列表去任务步骤模板表进行查询
	tsts, err := p.taskStagesTemplateRepo.FindInProTemIds(ctx, mproject.ToProjectTemplateIds(pts))
	if err != nil {
		zap.L().Error("FindProjectTemplate FindInProTemIds error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	var ptas []*mproject.ProjectTemplateAll
	for _, v := range pts {
		ptas = append(ptas, v.Convert(mtask.CovertProjectMap(tsts)[v.Id]))
	}
	// 3. 组装数据
	var pmMsgs []*project.ProjectTemplateMessage
	copier.Copy(&pmMsgs, ptas)
	return &project.ProjectTemplateResponse{Ptm: pmMsgs, Total: total}, nil
}