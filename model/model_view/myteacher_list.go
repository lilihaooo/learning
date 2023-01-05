package model_view

import "learning2.0/model"

// 教师列表入参  ----分页/id排序
type ListMyteacherReq struct {
	UserID int64 `json:"user_id" form:"user_id"`
	Page   int   `json:"page" form:"page"`
	//todo
	SortType int `json:"sort_type"  form:"sort_type"`
}

type ListMyteacherResp struct {
	TeacherInfo []*model.Teacher `json:"teacher_info"`
	Total       int64            `json:"total"`
}
