package dao

import (
	"context"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/database/gorm"
	"github.com/axzed/project-project/internal/database/interface/conn"
	gorm2 "gorm.io/gorm"
)

type TaskDao struct {
	conn *gorm.GormConn
}

// FindTaskMemberPage 查询任务成员列表
func (t *TaskDao) FindTaskMemberPage(ctx context.Context, taskCode int64, page int64, size int64) (list []*data.TaskMember, total int64, err error) {
	session := t.conn.Session(ctx)
	offset := (page - 1) * size
	err = session.Model(&data.TaskMember{}).
		Where("task_code=?", taskCode).
		Limit(int(size)).Offset(int(offset)).
		Find(&list).Error
	err = session.Model(&data.TaskMember{}).
		Where("task_code=?", taskCode).
		Count(&total).Error
	return
}

// FindTaskByIds 根据id列表查询任务
func (t *TaskDao) FindTaskByIds(background context.Context, taskIdList []int64) (list []*data.Task, err error) {
	session := t.conn.Session(background)
	err = session.Model(&data.Task{}).Where("id in (?)", taskIdList).Find(&list).Error
	return
}

// FindTaskMaxIdNum 查询当前任务最大id
// 注意这里可能会查出来null，所以要用指针
func (t *TaskDao) FindTaskMaxIdNum(ctx context.Context, projectCode int64) (v *int, err error) {
	session := t.conn.Session(ctx)
	//select * from -> 一定要用Scan 不能用Find(null)
	err = session.Model(&data.Task{}).
		Where("project_code=?", projectCode).
		Select("max(id_num)").Scan(&v).Error
	return
}

// FindTaskSort 查询当前任务最大排序id
// 注意这里可能会查出来null，所以要用指针
func (t *TaskDao) FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (v *int, err error) {
	session := t.conn.Session(ctx)
	//select * from -> 一定要用Scan 不能用Find(null)
	err = session.Model(&data.Task{}).
		Where("project_code=? and stage_code=?", projectCode, stageCode).
		Select("max(sort)").Scan(&v).Error
	return
}

// SaveTask 保存任务
func (t *TaskDao) SaveTask(ctx context.Context, conn conn.DbConn, ts *data.Task) error {
	t.conn = conn.(*gorm.GormConn) // 事务经典操作 将事务连接传递给dao的conn
	err := t.conn.Tx(ctx).Save(&ts).Error
	return err
}

// SaveTaskMember 保存创建当前任务成员
func (t *TaskDao) SaveTaskMember(ctx context.Context, conn conn.DbConn, tm *data.TaskMember) error {
	t.conn = conn.(*gorm.GormConn)
	err := t.conn.Tx(ctx).Save(&tm).Error
	return err
}

// FindTaskById 根据任务id查询任务
func (t *TaskDao) FindTaskById(ctx context.Context, taskCode int64) (ts *data.Task, err error) {
	session := t.conn.Session(ctx)
	err = session.Where("id=?", taskCode).Find(&ts).Error
	return
}

// UpdateTaskSort 更新任务排序
func (t *TaskDao) UpdateTaskSort(ctx context.Context, conn conn.DbConn, ts *data.Task) error {
	t.conn = conn.(*gorm.GormConn)
	err := t.conn.Tx(ctx).
		Where("id=?", ts.Id).
		Select("sort", "stage_code").
		Updates(&ts).
		Error
	return err
}

func (t *TaskDao) FindTaskByStageCodeLtSort(ctx context.Context, stageCode int, sort int) (ts *data.Task, err error) {
	session := t.conn.Session(ctx)
	err = session.Where("stage_code=? and sort < ?", stageCode, sort).Order("sort desc").Limit(1).Find(&ts).Error
	if err == gorm2.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (t *TaskDao) FindTaskByAssignTo(ctx context.Context, memberId int64, done int, page int64, size int64) (tsList []*data.Task, total int64, err error) {
	session := t.conn.Session(ctx)
	offset := (page - 1) * size
	err = session.Model(&data.Task{}).Where("assign_to=? and deleted=0 and done=?", memberId, done).Limit(int(size)).Offset(int(offset)).Find(&tsList).Error
	err = session.Model(&data.Task{}).Where("assign_to=? and deleted=0 and done=?", memberId, done).Count(&total).Error
	return
}

func (t *TaskDao) FindTaskByMemberCode(ctx context.Context, memberId int64, done int, page int64, size int64) (tList []*data.Task, total int64, err error) {
	session := t.conn.Session(ctx)
	offset := (page - 1) * size
	sql := "select a.* from ms_task a,ms_task_member b where a.id=b.task_code and member_code=? and a.deleted=0 and a.done=? limit ?,?"
	raw := session.Model(&data.Task{}).Raw(sql, memberId, done, offset, size)
	err = raw.Scan(&tList).Error
	if err != nil {
		return nil, 0, err
	}
	sqlCount := "select count(*) from ms_task a,ms_task_member b where a.id=b.task_code and member_code=? and a.deleted=0 and a.done=?"
	rawCount := session.Model(&data.Task{}).Raw(sqlCount, memberId, done)
	err = rawCount.Scan(&total).Error
	return
}

func (t *TaskDao) FindTaskByCreateBy(ctx context.Context, memberId int64, done int, page int64, size int64) (tList []*data.Task, total int64, err error) {
	session := t.conn.Session(ctx)
	offset := (page - 1) * size
	err = session.Model(&data.Task{}).Where("create_by=? and deleted=0 and done=?", memberId, done).Limit(int(size)).Offset(int(offset)).Find(&tList).Error
	err = session.Model(&data.Task{}).Where("create_by=? and deleted=0 and done=?", memberId, done).Count(&total).Error
	return
}

// FindTaskMemberByTaskId 根据任务id查询任务成员
func (t *TaskDao) FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) (task *data.TaskMember, err error) {
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
func (t *TaskDao) FindTaskByStageCode(ctx context.Context, stageCode int) (list []*data.Task, err error) {
	//select * from ms_task where stage_code=77 and deleted =0 order by sort asc
	session := t.conn.Session(ctx)
	err = session.Model(&data.Task{}).
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
