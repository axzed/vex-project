package mtask

// TaskStages 任务阶段
type TaskStages struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ProjectCode int64  `json:"project_code"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
	CreateTime  int64  `json:"create_time"`
	Deleted     int    `json:"deleted"`
}

func (*TaskStages) TableName() string {
	return "vex_task_stages"
}

// ToTaskStagesMap 转换为map
func ToTaskStagesMap(tsList []*TaskStages) map[int]*TaskStages {
	m := make(map[int]*TaskStages)
	for _, v := range tsList {
		m[v.Id] = v
	}
	return m
}