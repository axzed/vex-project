package project

type Project struct {
	Id                 int64
	Cover              string
	Name               string
	Description        string
	AccessControlType  int
	WhiteList          string
	Order              int
	Deleted            int
	TemplateCode       string
	Schedule           float64
	CreateTime         string
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
	isOwner     int64
	authorize   string
}

func (*ProjectMember) TableName() string {
	return "vex_project_member"
}

// ProjectMemberUnion is a union of project and project member
type ProjectMemberUnion struct {
	Project
	ProjectCode int64
	MemberCode  int64
	JoinTime    int64
	isOwner     int64
	authorize   string
}
