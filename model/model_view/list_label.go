package model_view

// 教师列表入参  ----分页/id排序
type ListLabelReq struct {
	Page     int `json:"page" form:"page"`
	SortType int ` json:"sort_type"  form:"sort_type"`
}
