package mtask

// VexTaskStagesTemplate 任务阶段模板
type VexTaskStagesTemplate struct {
	Id                  int
	Name                string
	ProjectTemplateCode int
	CreateTime          int64
	Sort                int
}

func (*VexTaskStagesTemplate) TableName() string {
	return "vex_task_stages_template"
}

// TaskStagesOnlyName 任务阶段名称
type TaskStagesOnlyName struct {
	Name string
}

// CovertProjectMap 转换成map
// 模板id->任务步骤
func CovertProjectMap(tsts []VexTaskStagesTemplate) map[int][]*TaskStagesOnlyName {
	var tss = make(map[int][]*TaskStagesOnlyName)
	for _, v := range tsts {
		ts := &TaskStagesOnlyName{}
		ts.Name = v.Name
		tss[v.ProjectTemplateCode] = append(tss[v.ProjectTemplateCode], ts)
	}
	return tss
}