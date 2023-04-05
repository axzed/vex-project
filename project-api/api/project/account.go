package project

import (
	"context"
	"github.com/axzed/project-api/api/rpc"
	"github.com/axzed/project-api/pkg/model"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/account"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

type HandlerAccount struct {
}

func (a *HandlerAccount) account(c *gin.Context) {
	//接收请求参数  一些参数的校验 可以放在api这里
	result := &common.Result{}
	var req *model.AccountReq
	_ = c.ShouldBind(&req)
	memberId := c.GetInt64("memberId")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//调用project模块 查询账户列表
	msg := &account.AccountReqMessage{
		MemberId:         memberId,
		OrganizationCode: c.GetString("organizationCode"),
		Page:             int64(req.Page),
		PageSize:         int64(req.PageSize),
		SearchType:       int32(req.SearchType),
		DepartmentCode:   req.DepartmentCode,
	}
	response, err := rpc.AccountServiceClient.Account(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	//返回数据
	var list []*model.MemberAccount
	copier.Copy(&list, response.AccountList)
	if list == nil {
		list = []*model.MemberAccount{}
	}
	var authList []*model.ProjectAuth
	copier.Copy(&authList, response.AuthList)
	if authList == nil {
		authList = []*model.ProjectAuth{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"total":    response.Total,
		"page":     req.Page,
		"list":     list,
		"authList": authList,
	}))
}

func NewAccount() *HandlerAccount {
	return &HandlerAccount{}
}
