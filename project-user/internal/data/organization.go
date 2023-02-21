package data

// Organization 组织表
type Organization struct {
	Id          int64
	Name        string
	Avatar      string
	Description string
	MemberId    int64
	CreateTime  int64
	Personal    int32
	Address     string
	Province    int32
	City        int32
	Area        int32
}

func (*Organization) TableName() string {
	return "vex_organization"
}
