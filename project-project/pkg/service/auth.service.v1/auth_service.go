package auth_service_v1

import (
	"context"
	"github.com/axzed/project-common/encrypts"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/auth"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/database/interface/conn"
	"github.com/axzed/project-project/internal/database/interface/transaction"
	"github.com/axzed/project-project/internal/domain"
	"github.com/axzed/project-project/internal/repo"
	"github.com/jinzhu/copier"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
	cache             repo.Cache
	transaction       transaction.Transaction
	projectAuthDomain *domain.ProjectAuthDomain
}

func New() *AuthService {
	return &AuthService{
		cache:             dao.Rc,
		transaction:       dao.NewTransactionImpl(),
		projectAuthDomain: domain.NewProjectAuthDomain(),
	}
}

// AuthList 获取权限列表rpc服务
func (a *AuthService) AuthList(ctx context.Context, msg *auth.AuthReqMessage) (*auth.ListAuthMessage, error) {
	organizationCode := encrypts.DecryptNoErr(msg.OrganizationCode)
	listPage, total, err := a.projectAuthDomain.AuthListPage(organizationCode, msg.Page, msg.PageSize)
	if err != nil {
		return nil, errs.ConvertToGrpcError(err)
	}
	var prList []*auth.ProjectAuth
	copier.Copy(&prList, listPage)
	return &auth.ListAuthMessage{List: prList, Total: total}, nil
}

// Apply 保存权限节点rpc服务
func (a *AuthService) Apply(ctx context.Context, msg *auth.AuthReqMessage) (*auth.ApplyResponse, error) {
	// 根据Action字段判断是获取列表还是保存
	if msg.Action == "getnode" {
		//获取列表
		list, checkedList, err := a.projectAuthDomain.AllNodeAndAuth(msg.AuthId)
		if err != nil {
			return nil, errs.ConvertToGrpcError(err)
		}
		var prList []*auth.ProjectNodeMessage
		copier.Copy(&prList, list)
		return &auth.ApplyResponse{List: prList, CheckedList: checkedList}, nil
	}
	if msg.Action == "save" {
		//先删除 project_auth_node表 在新增  事务
		//保存
		nodes := msg.Nodes
		//先删在存 加事务
		authId := msg.AuthId
		err := a.transaction.Action(func(conn conn.DbConn) error {
			err := a.projectAuthDomain.Save(conn, authId, nodes)
			return err
		})
		if err != nil {
			return nil, errs.ConvertToGrpcError(err.(*errs.BError))
		}
	}
	return &auth.ApplyResponse{}, nil
}

// AuthNodesByMemberId 根据用户id获取权限节点列表rpc服务
func (a *AuthService) AuthNodesByMemberId(ctx context.Context, msg *auth.AuthReqMessage) (*auth.AuthNodesResponse, error) {
	list, err := a.projectAuthDomain.AuthNodes(msg.MemberId)
	if err != nil {
		return nil, errs.ConvertToGrpcError(err)
	}
	return &auth.AuthNodesResponse{List: list}, nil
}
