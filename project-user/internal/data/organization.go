package data

// Organization 组织表
type Organization struct {
	Id          int64  `gorm:"id"`
	Name        string `gorm:"name"`
	Avatar      string `gorm:"avatar"`
	Description string `gorm:"description"`
	MemberId    int64  `gorm:"member_id"`
	CreateTime  int64  `gorm:"create_time"`
	Personal    int32  `gorm:"personal"`
	Address     string `gorm:"address"`
	Province    int32  `gorm:"province"`
	City        int32  `gorm:"city"`
	Area        int32  `gorm:"area"`
}

// TableName 映射表名
func (*Organization) TableName() string {
	return "vex_organization"
}

// ToMap 转换为map
func ToMap(orgs []*Organization) map[int64]*Organization {
	m := make(map[int64]*Organization)
	for _, v := range orgs {
		m[v.Id] = v
	}
	return m
}
