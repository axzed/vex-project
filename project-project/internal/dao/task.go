package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data/mtask"
	"github.com/axzed/project-project/internal/database/gorm"
	"github.com/axzed/project-project/internal/database/interface/conn"
	gorm2 "gorm.io/gorm"
)

type TaskDao struct {
	conn *gorm.GormConn
}

// FindTaskMaxIdNum 查询当前任务最大id
// 注意这里可能会查出来null，所以要用指针
func (t *TaskDao) FindTaskMaxIdNum(ctx context.Context, projectCode int64) (v *int, err error) {
	session := t.conn.Session(ctx)
	//select * from -> 一定要用Scan 不能用Find(null)
	err = session.Model(&mtask.Task{}).
		Where("project_code=?", projectCode).
		Select("max(id_num)").Scan(&v).Error
	return
}

// FindTaskSort 查询当前任务最大排序id
// 注意这里可能会查出来null，所以要用指针
func (t *TaskDao) FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (v *int, err error) {
	session := t.conn.Session(ctx)
	//select * from -> 一定要用Scan 不能用Find(null)
	err = session.Model(&mtask.Task{}).
		Where("project_code=? and stage_code=?", projectCode, stageCode).
		Select("max(sort)").Scan(&v).Error
	return
}

// SaveTask 保存任务
func (t *TaskDao) SaveTask(ctx context.Context, conn conn.DbConn, ts *mtask.Task) error {
	t.conn = conn.(*gorm.GormConn) // 事务经典操作 将事务连接传递给dao的conn
	err := t.conn.Tx(ctx).Save(&ts).Error
	return err
}

// SaveTaskMember 保存创建当前任务成员
func (t *TaskDao) SaveTaskMember(ctx context.Context, conn conn.DbConn, tm *mtask.TaskMember) error {
	t.conn = conn.(*gorm.GormConn)
	err := t.conn.Tx(ctx).Save(&tm).Error
	return err
}

func (t *TaskDao) FindTaskById(ctx context.Context, taskCode int64) (ts *mtask.Task, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskDao) UpdateTaskSort(ctx context.Context, conn conn.DbConn, ts *mtask.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskDao) FindTaskByStageCodeLtSort(ctx context.Context, stageCode int, sort int) (ts *mtask.Task, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskDao) FindTaskByAssignTo(ctx context.Context, memberId int64, done int, page int64, size int64) ([]*mtask.Task, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskDao) FindTaskByMemberCode(ctx context.Context, memberId int64, done int, page int64, size int64) (tList []*mtask.Task, total int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskDao) FindTaskByCreateBy(ctx context.Context, memberId int64, done int, page int64, size int64) (tList []*mtask.Task, total int64, err error) {
	//TODO implement me
	panic("implement me")
}

// FindTaskMemberByTaskId 根据任务id查询任务成员
func (t *TaskDao) FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) (task *mtask.TaskMember, err error) {
	err = t.conn.Session(ctx).
		Where("task_code=? and member_code=?", taskCode, memberId).
		Limit(1).
		Find(&task).Error
	if err == gorm2.ErrRecordNotFound {
		return nil, nil
	}
	return
}

// FindTaskByStageCode 根据阶段id查询任务
func (t *TaskDao) FindTaskByStageCode(ctx context.Context, stageCode int) (list []*mtask.Task, err error) {
	//select * from ms_task where stage_code=77 and deleted =0 order by sort asc
	session := t.conn.Session(ctx)
	err = session.Model(&mtask.Task{}).
		Where("stage_code=? and deleted =0", stageCode).
		Order("sort asc").
		Find(&list).Error
	return
}

func NewTaskDao() *TaskDao {
	return &TaskDao{
		conn: gorm.NewGormConn(),
	}
}
