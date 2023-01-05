package model_view

// 登陆注册入参
type UserReq struct {
	Username string `gorm:"column:username;type:varchar(20);not null" json:"username" form:"username"`
	Password string `gorm:"column:password;type:varchar(10);not null" json:"password"  form:"password"`
}
