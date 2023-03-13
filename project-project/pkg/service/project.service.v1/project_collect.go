package project_service_v1

import (
	"context"
	"github.com/axzed/project-common/encrypts"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-grpc/project"
	"github.com/axzed/project-project/internal/data/mproject"
	"github.com/axzed/project-project/pkg/model"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// UpdateCollectProject 更新项目的是否被收藏状态
func (p *ProjectService) UpdateCollectProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.CollectProjectResponse, error) {
	projectCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	projectCode, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var err error
	if "collect" == msg.CollectType {
		pc := &mproject.CollectionProject{
			ProjectCode: projectCode,
			MemberCode:  msg.MemberId,
			CreateTime:  time.Now().UnixMilli(),
		}
		err = p.projectRepo.SaveProjectCollect(c, pc)
	}
	if "cancel" == msg.CollectType {
		err = p.projectRepo.DeleteProjectCollect(c, msg.MemberId, projectCode)
	}
	if err != nil {
		zap.L().Error("UpdateCollectProject", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	return &project.CollectProjectResponse{}, nil
}
