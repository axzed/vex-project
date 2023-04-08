package data

import "github.com/jinzhu/copier"

// ProjectMenu 项目菜单
type ProjectMenu struct {
	Id         int64
	Pid        int64
	Title      string
	Icon       string
	Url        string
	FilePath   string
	Params     string
	Node       string
	Sort       int
	Status     int
	CreateBy   int64
	IsInner    int
	Values     string
	ShowSlider int
}

func (*ProjectMenu) TableName() string {
	return "vex_project_menu"
}

// ProjectMenuChild 项目菜单树形结构
type ProjectMenuChild struct {
	ProjectMenu
	StatusText string
	InnerText  string
	FullUrl    string
	Children   []*ProjectMenuChild
}

// CovertChild 转换成树形结构
func CovertChild(pms []*ProjectMenu) []*ProjectMenuChild {
	var pmcs []*ProjectMenuChild
	copier.Copy(&pmcs, pms)
	for _, v := range pmcs {
		v.StatusText = getStatus(v.Status)
		v.InnerText = getInnerText(v.IsInner)
		v.FullUrl = getFullUrl(v.Url, v.Params, v.Values)
	}
	var childPmcs []*ProjectMenuChild
	//递归
	for _, v := range pmcs {
		if v.Pid == 0 {
			pmc := &ProjectMenuChild{}
			copier.Copy(pmc, v)
			childPmcs = append(childPmcs, pmc)
		}
	}
	toChild(childPmcs, pmcs)
	return childPmcs
}

// getFullUrl 拼接完整的url
func getFullUrl(url string, params string, values string) string {
	if (params != "" && values != "") || values != "" {
		return url + "/" + values
	}
	return url
}

// getInnerText 获取内页还是导航
func getInnerText(inner int) string {
	if inner == 0 {
		return "导航"
	}
	if inner == 1 {
		return "内页"
	}
	return ""
}

// getStatus 获取状态
func getStatus(status int) string {
	if status == 0 {
		return "禁用"
	}
	if status == 1 {
		return "使用中"
	}
	return ""
}

// toChild 递归转换
func toChild(childPmcs []*ProjectMenuChild, pmcs []*ProjectMenuChild) {
	for _, pmc := range childPmcs {
		for _, pm := range pmcs {
			if pmc.Id == pm.Pid {
				child := &ProjectMenuChild{}
				copier.Copy(child, pm)
				pmc.Children = append(pmc.Children, child)
			}
		}
		toChild(pmc.Children, pmcs)
	}
}
