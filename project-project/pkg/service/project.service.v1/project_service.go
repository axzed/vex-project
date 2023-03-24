package project_service_v1

import (
	"context"
	"github.com/axzed/project-common/encrypts"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-common/tms"
	"github.com/axzed/project-grpc/project"
	"github.com/axzed/project-grpc/user/login"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/data/menu"
	"github.com/axzed/project-project/internal/data/mproject"
	"github.com/axzed/project-project/internal/database/interface/conn"
	"github.com/axzed/project-project/internal/database/interface/transaction"
	"github.com/axzed/project-project/internal/repo"
	"github.com/axzed/project-project/internal/rpc"
	"github.com/axzed/project-project/pkg/model"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"log"
	"strconv"
	"time"
)

// ProjectService 项目服务
type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache                  repo.Cache              // 缓存
	transaction            transaction.Transaction // 事务
	menuRepo               repo.MenuRepo
	projectRepo            repo.ProjectRepo
	projectTemplateRepo    repo.ProjectTemplateRepo
	taskStagesTemplateRepo repo.TaskStagesTemplateRepo
	taskStagesRepo         repo.TaskStagesRepo
}

// NewProjectService 初始化页面展示服务
func NewProjectService() *ProjectService {
	return &ProjectService{
		// 为定义的接口赋上实现类
		cache:                  dao.Rc,
		transaction:            dao.NewTransactionImpl(),
		menuRepo:               dao.NewMenuDao(),
		projectRepo:            dao.NewProjectDao(),
		projectTemplateRepo:    dao.NewProjectTemplateDao(),
		taskStagesTemplateRepo: dao.NewTaskStagesTemplateDao(),
		taskStagesRepo:         dao.NewTaskStagesDao(),
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

// FindProjectByMemId 通过用户id获取项目列表
func (p *ProjectService) FindProjectByMemId(ctx context.Context, msg *project.ProjectRpcMessage) (*project.MyProjectResponse, error) {
	// 获取参数
	memberId := msg.MemberId
	page := msg.Page
	pageSize := msg.PageSize
	var pms []*data.ProAndMember
	var total int64
	var err error
	// 通过SelectBy参数判断调用哪个服务
	if msg.SelectBy == "" || msg.SelectBy == "my" {
		pms, total, err = p.projectRepo.FindProjectByMemId(ctx, memberId, "and deleted = 0", page, pageSize)
	}
	if msg.SelectBy == "archived" {
		pms, total, err = p.projectRepo.FindProjectByMemId(ctx, memberId, "and archived = 1", page, pageSize)
	}
	if msg.SelectBy == "deleted" {
		pms, total, err = p.projectRepo.FindProjectByMemId(ctx, memberId, "and deleted = 1", page, pageSize)
	}
	// 显示收藏项目
	// 用户点击收藏 将该项目设为收藏
	if msg.SelectBy == "collect" {
		pms, total, err = p.projectRepo.FindCollectProjectByMemId(ctx, memberId, page, pageSize)
		// 将收藏的项目标记为已收藏
		for _, v := range pms {
			v.Collected = model.Collected
		}
	} else {
		// 查询全部的收藏项目
		collectPms, _, err := p.projectRepo.FindCollectProjectByMemId(ctx, memberId, page, pageSize)
		if err != nil {
			zap.L().Error("project FindProjectByMember error", zap.Error(err))
			return nil, errs.ConvertToGrpcError(model.ErrDBFail)
		}
		// 将收藏的项目放入map中 cmap[项目id] = 项目
		var cMap = make(map[int64]*data.ProAndMember)
		// 遍历收藏的项目
		for _, v := range collectPms {
			// v.Id 为项目id (collectPms)
			cMap[v.Id] = v
		}
		// 将查出来的项目集在cmap中查找 如果存在则标记为已收藏
		for _, v := range pms {
			if cMap[v.ProjectCode] != nil {
				// v.ProjectCode 为项目code (pms)
				v.Collected = model.Collected
			}
		}
	}
	// 调用服务
	if err != nil {
		zap.L().Error("project FindProjectByMember error", zap.Error(err))
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
		if msg.SelectBy == "collect" {
			v.Code, _ = encrypts.EncryptInt64(v.Id, model.AESKey)
		} else {
			v.Code, _ = encrypts.EncryptInt64(v.ProjectCode, model.AESKey)
		}
		pam := data.ToMap(pms)[v.Id]
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
	var pts []data.ProjectTemplate
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
	var ptas []*data.ProjectTemplateAll
	for _, v := range pts {
		ptas = append(ptas, v.Convert(data.CovertProjectMap(tsts)[v.Id]))
	}
	// 3. 组装数据
	var pmMsgs []*project.ProjectTemplateMessage
	copier.Copy(&pmMsgs, ptas)
	return &project.ProjectTemplateResponse{Ptm: pmMsgs, Total: total}, nil
}

// SaveProject 保存项目
func (p *ProjectService) SaveProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.SaveProjectMessage, error) {
	// 获取参数
	organizationCodeStr, _ := encrypts.Decrypt(msg.OrganizationCode, model.AESKey)
	organizationCode, _ := strconv.ParseInt(organizationCodeStr, 10, 64)
	templateCodeStr, _ := encrypts.Decrypt(msg.TemplateCode, model.AESKey)
	templateCode, _ := strconv.ParseInt(templateCodeStr, 10, 64)

	// 通过项目模板获取对应的任务流程模板信息
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	stageTemplateList, err := p.taskStagesTemplateRepo.FindByProjectTemplateId(c, int(templateCode))
	if err != nil {
		zap.L().Error("SaveProject taskStagesTemplateRepo.FindByProjectTemplateId error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}

	pr := &data.Project{
		Name:              msg.Name,
		Description:       msg.Description,
		TemplateCode:      int(templateCode),
		CreateTime:        time.Now().UnixMilli(),
		Cover:             "https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500",
		Deleted:           model.NoDeleted,
		Archive:           model.NoArchive,
		OrganizationCode:  organizationCode,
		AccessControlType: model.Open,
		TaskBoardTheme:    model.Simple,
	}

	// 通过事务控制保存项目相关的写入操作
	err = p.transaction.Action(func(conn conn.DbConn) error {
		// 1.保存项目表
		err := p.projectRepo.SaveProject(conn, ctx, pr)
		if err != nil {
			zap.L().Error("SaveProject SaveProject error", zap.Error(err))
			return errs.ConvertToGrpcError(model.ErrDBFail)
		}
		pm := &data.ProjectMember{
			ProjectCode: pr.Id,
			MemberCode:  msg.MemberId,
			JoinTime:    time.Now().UnixMilli(),
			IsOwner:     msg.MemberId,
			Authorize:   "",
		}

		// 2.保存项目成员关系表
		err = p.projectRepo.SaveProjectMember(conn, ctx, pm)
		if err != nil {
			zap.L().Error("SaveProject SaveProjectMember error", zap.Error(err))
			return errs.ConvertToGrpcError(model.ErrDBFail)
		}

		// 3.生成任务的步骤并存入表中
		for index, v := range stageTemplateList {
			taskStage := &data.TaskStages{
				Name:        v.Name,
				ProjectCode: pr.Id,
				Sort:        index + 1,
				Description: "",
				CreateTime:  time.Now().UnixMilli(),
				Deleted:     model.NoDeleted,
			}
			// 操作db
			err := p.taskStagesRepo.SaveTaskStages(c, conn, taskStage)
			if err != nil {
				zap.L().Error("SaveProject taskStagesRepo.SaveTaskStages error", zap.Error(err))
				return errs.ConvertToGrpcError(model.ErrDBFail)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// 生成项目code
	code, _ := encrypts.EncryptInt64(pr.Id, model.AESKey)
	rsp := &project.SaveProjectMessage{
		Id:               pr.Id,
		Code:             code,
		OrganizationCode: organizationCodeStr,
		Name:             pr.Name,
		Cover:            pr.Cover,
		CreateTime:       tms.FormatByMill(pr.CreateTime),
		TaskBoardTheme:   pr.TaskBoardTheme,
	}

	return rsp, nil
}

// FindProjectDetail 查看项目详情
func (p *ProjectService) FindProjectDetail(ctx context.Context, msg *project.ProjectRpcMessage) (*project.ProjectDetailMessage, error) {
	//1. 查项目表
	//2. 项目和成员的关联表 查到项目的拥有者 去member表查名字
	//3. 查收藏表 判断收藏状态
	projectCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	projectCode, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	memberId := msg.MemberId
	c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	projectAndMember, err := p.projectRepo.FindProjectByPIdAndMemId(c, projectCode, memberId)
	if err != nil {
		zap.L().Error("FindProjectDetail FindProjectByPIdAndMemId error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	ownerId := projectAndMember.IsOwner
	// 从user rpc 中通过id获取用户信息
	member, err := rpc.LoginServiceClient.FindMemberInfoById(c, &login.UserMessage{MemId: ownerId})
	if err != nil {
		zap.L().Error("FindProjectDetail FindMemberInfoById error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	log.Println(ownerId)
	// TODO 优化 收藏的时候 可以放入redis
	isCollected, err := p.projectRepo.FindCollectByPIdAndMemId(c, projectCode, memberId)
	if err != nil {
		zap.L().Error("FindProjectDetail FindCollectByPIdAndMemId error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	if isCollected {
		projectAndMember.Collected = model.Collected
	}
	// ProjectDetailMessage 项目详情
	var detailMsg = &project.ProjectDetailMessage{}
	copier.Copy(&detailMsg, projectAndMember) // 复制属性
	// 将当前项目的owner的信息放入要返回的detailMsg中
	detailMsg.OwnerAvatar = member.Avatar
	detailMsg.OwnerName = member.Name
	// 将ProjectAndMember的信息放入要返回的detailMsg中
	detailMsg.Code, _ = encrypts.EncryptInt64(projectAndMember.Id, model.AESKey)
	detailMsg.AccessControlType = projectAndMember.GetAccessControlType()
	detailMsg.OrganizationCode, _ = encrypts.EncryptInt64(projectAndMember.OrganizationCode, model.AESKey)
	detailMsg.Order = int32(projectAndMember.Sort)
	detailMsg.CreateTime = tms.FormatByMill(projectAndMember.CreateTime)
	// 返回项目详情
	return detailMsg, nil
}

// UpdateDeletedProject 更新项目的是否被删除状态
func (p *ProjectService) UpdateDeletedProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.DeletedProjectResponse, error) {
	projectCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	projectCode, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err := p.projectRepo.UpdateDeleteProject(c, projectCode, msg.Deleted)
	if err != nil {
		zap.L().Error("RecycleProject DeleteProject error", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	return &project.DeletedProjectResponse{}, nil
}

// UpdateCollectProject 更新项目的是否被收藏状态
func (p *ProjectService) UpdateCollectProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.CollectProjectResponse, error) {
	projectCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	projectCode, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var err error
	if "collect" == msg.CollectType {
		pc := &data.CollectionProject{
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

// UpdateProject 更新项目rpc服务实现
func (p *ProjectService) UpdateProject(ctx context.Context, msg *project.UpdateProjectMessage) (*project.UpdateProjectResponse, error) {
	// FIXME: 项目的图片无法修改
	// 获取rpc传递过来的项目code
	projectCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	projectCode, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// 构造update项目实体用于和dao打交道
	proj := &data.Project{
		Id:                 projectCode,
		Name:               msg.Name,
		Description:        msg.Description,
		Cover:              msg.Cover,
		TaskBoardTheme:     msg.TaskBoardTheme,
		Prefix:             msg.Prefix,
		Private:            int(msg.Private),
		OpenPrefix:         int(msg.OpenPrefix),
		OpenBeginTime:      int(msg.OpenBeginTime),
		OpenTaskPrivate:    int(msg.OpenTaskPrivate),
		Schedule:           msg.Schedule,
		AutoUpdateSchedule: int(msg.AutoUpdateSchedule),
	}
	// 调用接口更新项目
	err := p.projectRepo.UpdateProject(c, proj)
	if err != nil {
		zap.L().Error("UpdateProject", zap.Error(err))
		return nil, errs.ConvertToGrpcError(model.ErrDBFail)
	}
	return &project.UpdateProjectResponse{}, nil
}
