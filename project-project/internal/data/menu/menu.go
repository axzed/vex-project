package menu

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
	Children []*ProjectMenuChild
}

// CovertChild 转换成树形结构
func CovertChild(pms []*ProjectMenu) []*ProjectMenuChild {
	var pmcs []*ProjectMenuChild
	copier.Copy(&pmcs, pms)
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
