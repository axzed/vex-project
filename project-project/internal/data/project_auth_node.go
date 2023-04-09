package data

type ProjectAuthNode struct {
	Id   int64
	Auth int64
	Node string
}

func (*ProjectAuthNode) TableName() string {
	return "vex_project_auth_node"
}
