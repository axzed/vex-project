package data

// Member 成员表
type Member struct {
	Id              int64  `gorm:"id"`
	Account         string `gorm:"account"`
	Password        string `gorm:"password"`
	Name            string `gorm:"name"`
	Mobile          string `gorm:"mobile"`
	Realname        string `gorm:"realname"`
	CreateTime      int64  `gorm:"create_time"`
	Status          int    `gorm:"status"`
	LastLoginTime   int64  `gorm:"last_login_time"`
	Sex             int    `gorm:"sex"`
	Avatar          string `gorm:"avatar"`
	Idcard          string `gorm:"idcard"`
	Province        int    `gorm:"province"`
	City            int    `gorm:"city"`
	Area            int    `gorm:"area"`
	Address         string `gorm:"address"`
	Description     string `gorm:"description"`
	Email           string `gorm:"email"`
	DingtalkOpenid  string `gorm:"dingtalk_openid"`
	DingtalkUnionid string `gorm:"dingtalk_unionid"`
	DingtalkUserid  string `gorm:"dingtalk_userid"`
}

// TableName 映射表名
func (*Member) TableName() string {
	return "vex_member"
}
