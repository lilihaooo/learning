package model_view

// 教师添加标签入参
type MyTeacherAddLabelReq struct {
	MyTeacherID int64   `gorm:"column:myteacher_id;type:int;not null" json:"myteacher_id" form:"myteacher_id"`
	LabelID     []int64 `gorm:"column:label_id;type:int;not null" json:"label_id" form:"label_id"`
}
