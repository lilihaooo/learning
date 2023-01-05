package model_view

type Teacher struct {
	ID      int64  `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Source  string `json:"source" form:"source"`
	City    string `json:"city" form:"city"`
	Created int64  `json:"created" form:"created"`
}
