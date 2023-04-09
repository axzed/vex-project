package domain

import (
	"context"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/database/interface/conn"
	"github.com/axzed/project-project/internal/repo"
	"github.com/axzed/project-project/pkg/model"
)

type ProjectAuthNodeDomain struct {
	projectAuthNodeRepo repo.ProjectAuthNodeRepo
}

func NewProjectAuthNodeDomain() *ProjectAuthNodeDomain {
	return &ProjectAuthNodeDomain{
		projectAuthNodeRepo: dao.NewProjectAuthNodeDao(),
	}
}

// AuthNodeList 获取权限节点列表
func (d *ProjectAuthNodeDomain) AuthNodeList(authId int64) ([]string, *errs.BError) {
	list, err := d.projectAuthNodeRepo.FindNodeStringList(context.Background(), authId)
	if err != nil {
		return nil, model.ErrDBFail
	}
	return list, nil
}

// Save 保存权限节点
func (d *ProjectAuthNodeDomain) Save(conn conn.DbConn, authId int64, nodes []string) *errs.BError {
	err := d.projectAuthNodeRepo.DeleteByAuthId(context.Background(), conn, authId)
	if err != nil {
		return model.ErrDBFail
	}
	err = d.projectAuthNodeRepo.Save(context.Background(), conn, authId, nodes)
	if err != nil {
		return model.ErrDBFail
	}
	return nil
}
