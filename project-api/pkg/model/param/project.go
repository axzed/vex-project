package param

type Project struct {
	Id                 int64   `json:"id"`
	Cover              string  `json:"cover"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	AccessControlType  int     `json:"accessControlType"`
	WhiteList          string  `json:"whiteList"`
	Order              int     `json:"order"`
	Deleted            int     `json:"deleted"`
	TemplateCode       string  `json:"templateCode"`
	Schedule           float64 `json:"schedule"`
	CreateTime         string  `json:"createTime"`
	OrganizationCode   int64   `json:"organizationCode"`
	DeletedTime        string  `json:"deletedTime"`
	Private            int     `json:"private"`
	Prefix             string  `json:"prefix"`
	OpenPrefix         int     `json:"openPrefix"`
	Archive            int     `json:"archive"`
	ArchiveTime        int64   `json:"archiveTime"`
	OpenBeginTime      int     `json:"openBeginTime"`
	OpenTaskPrivate    int     `json:"openTaskPrivate"`
	TaskBoardTheme     string  `json:"taskBoardTheme"`
	BeginTime          int64   `json:"beginTime"`
	EndTime            int64   `json:"endTime"`
	AutoUpdateSchedule int     `json:"autoUpdateSchedule"`
	Code               string  `json:"code"`
}

type ProjectMember struct {
	Id          int64  `json:"id"`
	ProjectCode int64  `json:"projectCode"`
	MemberCode  int64  `json:"memberCode"`
	JoinTime    int64  `json:"joinTime"`
	IsOwner     int64  `json:"isOwner"`
	Authorize   string `json:"authorize"`
}

// ProjectAndMember is a union of project and project member
type ProjectAndMember struct {
	Project
	ProjectCode int64  `json:"projectCode"`
	MemberCode  int64  `json:"memberCode"`
	JoinTime    int64  `json:"joinTime"`
	IsOwner     int64  `json:"isOwner"`
	Authorize   string `json:"authorize"`
}
