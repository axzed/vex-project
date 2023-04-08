package project

import (
	"context"
	"github.com/axzed/project-api/api/rpc"
	"github.com/axzed/project-api/pkg/model"
	common "github.com/axzed/project-common"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/menu"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

type HandlerMenu struct {
}

// menuList 查询所有路由菜单
func (m HandlerMenu) menuList(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := rpc.MenuServiceClient.MenuList(ctx, &menu.MenuReqMessage{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	var list []*model.Menu
	copier.Copy(&list, res.List)
	if list == nil {
		list = []*model.Menu{}
	}

	c.JSON(http.StatusOK, result.Success(list))
}

func NewMenu() *HandlerMenu {
	return &HandlerMenu{}
}
