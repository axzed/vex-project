package project_service_v1

import (
	"context"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/project"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/data/menu"
	"github.com/axzed/project-project/internal/database/interface/transaction"
	"github.com/axzed/project-project/internal/repo"
	"github.com/axzed/project-project/pkg/model"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

// ProjectService 项目服务
type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache              // 缓存
	transaction transaction.Transaction // 事务
	menuRepo    repo.MenuRepo
}

// NewProjectService 初始化页面展示服务
func NewProjectService() *ProjectService {
	return &ProjectService{
		// 为定义的接口赋上实现类
		cache:       dao.Rc,
		transaction: dao.NewTransactionImpl(),
		menuRepo:    dao.NewMenuDao(),
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
