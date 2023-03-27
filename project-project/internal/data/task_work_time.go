package data

import (
	"github.com/axzed/project-common/encrypts"
	"github.com/axzed/project-common/tms"
	"github.com/jinzhu/copier"
)

// TaskWorkTime dto vo
type TaskWorkTime struct {
	Id         int64
	TaskCode   int64
	MemberCode int64
	CreateTime int64
	Content    string
	BeginTime  int64
	Num        int
}

func (*TaskWorkTime) TableName() string {
	return "vex_task_work_time"
}

// TaskWorkTimeDisplay  dto vo
type TaskWorkTimeDisplay struct {
	Id         int64
	TaskCode   string
	MemberCode string
	CreateTime string
	Content    string
	BeginTime  string
	Num        int
	Member     Member
}

func (t *TaskWorkTime) ToDisplay() *TaskWorkTimeDisplay {
	td := &TaskWorkTimeDisplay{}
	copier.Copy(td, t)
	td.MemberCode = encrypts.EncryptNoErr(t.MemberCode)
	td.TaskCode = encrypts.EncryptNoErr(t.TaskCode)
	td.CreateTime = tms.FormatByMill(t.CreateTime)
	td.BeginTime = tms.FormatByMill(t.BeginTime)
	return td
}
