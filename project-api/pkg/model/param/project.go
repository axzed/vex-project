package param

type Project struct {
	Id                 int64   `json:"id"`
	Cover              string  `json:"cover"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	AccessControlType  string  `json:"access_control_type"`
	WhiteList          string  `json:"white_list"`
	Order              int     `json:"order"`
	Deleted            int     `json:"deleted"`
	TemplateCode       string  `json:"template_code"`
	Schedule           float64 `json:"schedule"`
	CreateTime         string  `json:"create_time"`
	OrganizationCode   string  `json:"organization_code"`
	DeletedTime        string  `json:"deleted_time"`
	Private            int     `json:"private"`
	Prefix             string  `json:"prefix"`
	OpenPrefix         int     `json:"open_prefix"`
	Archive            int     `json:"archive"`
	ArchiveTime        int64   `json:"archive_time"`
	OpenBeginTime      int     `json:"open_begin_time"`
	OpenTaskPrivate    int     `json:"open_task_private"`
	TaskBoardTheme     string  `json:"task_board_theme"`
	BeginTime          int64   `json:"begin_time"`
	EndTime            int64   `json:"end_time"`
	AutoUpdateSchedule int     `json:"auto_update_schedule"`
	Code               string  `json:"code"`
}

type ProjectMember struct {
	Id          int64  `json:"id"`
	ProjectCode int64  `json:"project_code"`
	MemberCode  int64  `json:"member_code"`
	JoinTime    string `json:"join_time"`
	IsOwner     int64  `json:"is_owner"`
	Authorize   string `json:"authorize"`
}

// ProjectAndMember is a union of project and project member
type ProjectAndMember struct {
	Project
	ProjectCode int64  `json:"project_code"`
	MemberCode  int64  `json:"member_code"`
	JoinTime    int64  `json:"join_time"`
	IsOwner     int64  `json:"is_owner"`
	Authorize   string `json:"authorize"`
	OwnerName   string `json:"owner_name"`
	Collected   int    `json:"collected"`
}


type ProjectTemplate struct {
	Id               int                   `json:"id"`
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	Sort             int                   `json:"sort"`
	CreateTime       string                `json:"create_time"`
	OrganizationCode string                `json:"organization_code"`
	Cover            string                `json:"cover"`
	MemberCode       string                `json:"member_code"`
	IsSystem         int                   `json:"is_system"`
	TaskStages       []*TaskStagesOnlyName `json:"task_stages"`
	Code             string                `json:"code"`
}

// SaveProjectRequest is a request of save project
type SaveProjectRequest struct {
	Name         string `json:"name" form:"name"`
	TemplateCode string `json:"templateCode" form:"templateCode"`
	Description  string `json:"description" form:"description"`
	Id           int    `json:"id" form:"id"`
}

// SaveProject is a response of save project
type SaveProject struct {
	Id               int64  `json:"id"`
	Cover            string `json:"cover"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Code             string `json:"code"`
	CreateTime       string `json:"create_time"`
	TaskBoardTheme   string `json:"task_board_theme"`
	OrganizationCode string `json:"organization_code"`
}


// TaskStagesOnlyName is a task stage with only name
type TaskStagesOnlyName struct {
	Name string `json:"name"`
}

