package menu_service_v1

import (
	"context"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/menu"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/database/interface/transaction"
	"github.com/axzed/project-project/internal/domain"
	"github.com/axzed/project-project/internal/repo"
	"github.com/jinzhu/copier"
)

type MenuService struct {
	menu.UnimplementedMenuServiceServer
	cache       repo.Cache
	transaction transaction.Transaction
	menuDomain  *domain.MenuDomain
}

func New() *MenuService {
	return &MenuService{
		cache:       dao.Rc,
		transaction: dao.NewTransactionImpl(),
		menuDomain:  domain.NewMenuDomain(),
	}
}

// MenuList 查询菜单路由rpc服务
func (d *MenuService) MenuList(ctx context.Context, msg *menu.MenuReqMessage) (*menu.MenuResponseMessage, error) {
	list, err := d.menuDomain.MenuTreeList()
	if err != nil {
		return nil, errs.ConvertToGrpcError(err)
	}
	var mList []*menu.MenuMessage
	copier.Copy(&mList, list)
	return &menu.MenuResponseMessage{List: mList}, nil
}
