package domain

import (
	"context"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/database/interface/conn"
	"github.com/axzed/project-project/internal/repo"
	"github.com/axzed/project-project/pkg/model"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type ProjectAuthDomain struct {
	projectAuthRepo       repo.ProjectAuthRepo
	userRpcDomain         *UserRpcDomain
	projectNodeDomain     *ProjectNodeDomain
	projectAuthNodeDomain *ProjectAuthNodeDomain
	accountDomain         *AccountDomain
}

// AuthList 查询所有权限
func (d *ProjectAuthDomain) AuthList(orgCode int64) ([]*data.ProjectAuthDisplay, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, err := d.projectAuthRepo.FindAuthList(c, orgCode)
	if err != nil {
		zap.L().Error("project AuthList projectAuthRepo.FindAuthList error", zap.Error(err))
		return nil, model.ErrDBFail
	}
	var pdList []*data.ProjectAuthDisplay
	for _, v := range list {
		display := v.ToDisplay()
		pdList = append(pdList, display)
	}
	return pdList, nil
}

// AuthListPage 查询所有权限 分页
func (d *ProjectAuthDomain) AuthListPage(orgCode int64, page int64, pageSize int64) ([]*data.ProjectAuthDisplay, int64, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, total, err := d.projectAuthRepo.FindAuthListPage(c, orgCode, page, pageSize)
	if err != nil {
		zap.L().Error("project AuthList projectAuthRepo.FindAuthList error", zap.Error(err))
		return nil, 0, model.ErrDBFail
	}
	var pdList []*data.ProjectAuthDisplay
	for _, v := range list {
		display := v.ToDisplay()
		pdList = append(pdList, display)
	}
	return pdList, total, nil
}

// AllNodeAndAuth 查询所有节点和权限
func (d *ProjectAuthDomain) AllNodeAndAuth(authId int64) ([]*data.ProjectNodeAuthTree, []string, *errs.BError) {
	nodeList, err := d.projectNodeDomain.NodeList()
	if err != nil {
		return nil, nil, err
	}
	checkedList, err := d.projectAuthNodeDomain.AuthNodeList(authId)
	if err != nil {
		return nil, nil, err
	}
	list := data.ToAuthNodeTreeList(nodeList, checkedList)
	return list, checkedList, nil
}

// Save 保存权限
func (d *ProjectAuthDomain) Save(conn conn.DbConn, authId int64, nodes []string) *errs.BError {
	err := d.projectAuthNodeDomain.Save(conn, authId, nodes)
	if err != nil {
		return err
	}
	return nil
}

// AuthNodes 查询用户权限节点
func (d *ProjectAuthDomain) AuthNodes(memberId int64) ([]string, *errs.BError) {
	account, err := d.accountDomain.FindAccount(memberId)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, model.ParamsError
	}
	//c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	authorize := account.Authorize
	authId, _ := strconv.ParseInt(authorize, 10, 64)
	authNodeList, dbErr := d.projectAuthNodeDomain.AuthNodeList(authId)
	if dbErr != nil {
		return nil, model.ErrDBFail
	}
	return authNodeList, nil
}

func NewProjectAuthDomain() *ProjectAuthDomain {
	return &ProjectAuthDomain{
		projectAuthRepo:       dao.NewProjectAuthDao(),
		userRpcDomain:         NewUserRpcDomain(),
		projectNodeDomain:     NewProjectNodeDomain(),
		projectAuthNodeDomain: NewProjectAuthNodeDomain(),
		accountDomain:         NewAccountDomain(),
	}
}
