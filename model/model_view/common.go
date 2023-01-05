package model_view

// 批量删除的入参
type BatchDelReq struct {
	ID []int64 `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id" form:"id"`
}
