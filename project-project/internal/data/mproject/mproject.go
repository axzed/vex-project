package mproject

import (
	"github.com/axzed/project-common/encrypts"
	"github.com/axzed/project-common/tms"
	"github.com/axzed/project-project/internal/data/mtask"
	"github.com/axzed/project-project/pkg/model"
)

type Project struct {
	Id                 int64
	Cover              string
	Name               string
	Description        string
	AccessControlType  int
	WhiteList          string
	Sort              int
	Deleted            int
	TemplateCode       int
	Schedule           float64
	CreateTime         int64
	OrganizationCode   int64
	DeletedTime        string
	Private            int
	Prefix             string
	OpenPrefix         int
	Archive            int
	ArchiveTime        int64
	OpenBeginTime      int
	OpenTaskPrivate    int
	TaskBoardTheme     string
	BeginTime          int64
	EndTime            int64
	AutoUpdateSchedule int
}

func (*Project) TableName() string {
	return "vex_project"
}

type ProjectMember struct {
	Id          int64
	ProjectCode int64
	MemberCode  int64
	JoinTime    int64
	IsOwner     int64
	Authorize   string
}

func (*ProjectMember) TableName() string {
	return "vex_project_member"
}

// ProAndMember is a union of project and project member
type ProAndMember struct {
	Project
	ProjectCode int64
	MemberCode  int64
	JoinTime    int64
	IsOwner     int64
	Authorize   string
}

// CollectionProject 项目收藏表
type CollectionProject struct {
	Id          int64
	ProjectCode int64
	MemberCode  int64
	CreateTime  int64
}

// TableName get table name
func (*CollectionProject) TableName() string {
	return "vex_project_collection"
}

// GetAccessControlType get access control type
func (m *ProAndMember) GetAccessControlType() string {
	if m.AccessControlType == 0 {
		return "open"
	}
	if m.AccessControlType == 1 {
		return "private"
	}
	if m.AccessControlType == 2 {
		return "custom"
	}
	return ""
}

// ToMap convert slice to map
func ToMap(orgs []*ProAndMember) map[int64]*ProAndMember {
	m := make(map[int64]*ProAndMember)
	for _, v := range orgs {
		m[v.Id] = v
	}
	return m
}

type ProjectTemplate struct {
	Id               int
	Name             string
	Description      string
	Sort             int
	CreateTime       int64
	OrganizationCode int64
	Cover            string
	MemberCode       int64
	IsSystem         int
}

func (*ProjectTemplate) TableName() string {
	return "vex_project_template"
}

type ProjectTemplateAll struct {
	Id               int
	Name             string
	Description      string
	Sort             int
	CreateTime       string
	OrganizationCode string
	Cover            string
	MemberCode       string
	IsSystem         int
	TaskStages       []*mtask.TaskStagesOnlyName
	Code 			 string
}


func (pt ProjectTemplate) Convert(taskStages []*mtask.TaskStagesOnlyName) *ProjectTemplateAll {
	organizationCode, _ := encrypts.EncryptInt64(pt.OrganizationCode, model.AESKey)
	memberCode, _ := encrypts.EncryptInt64(pt.MemberCode, model.AESKey)
	code, _ := encrypts.EncryptInt64(int64(pt.Id), model.AESKey)
	pta := &ProjectTemplateAll{
		Id:               pt.Id,
		Name:             pt.Name,
		Description:      pt.Description,
		Sort:             pt.Sort,
		CreateTime:       tms.FormatByMill(pt.CreateTime),
		OrganizationCode: organizationCode,
		Cover:            pt.Cover,
		MemberCode:       memberCode,
		IsSystem:         pt.IsSystem,
		TaskStages:       taskStages,
		Code:             code,
	}
	return pta
}
func ToProjectTemplateIds(pts []ProjectTemplate) []int {
	var ids []int
	for _, v := range pts {
		ids = append(ids, v.Id)
	}
	return ids
}