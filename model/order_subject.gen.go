// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOrderSubject = "order_subject"

// OrderSubject mapped from table <order_subject>
type OrderSubject struct {
	ID        int64 `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	OrderID   int64 `gorm:"column:order_id;type:int;not null" json:"order_id"`
	SubjectID int64 `gorm:"column:subject_id;type:int;not null" json:"subject_id"`
	Num       int64 `gorm:"column:num;type:int;not null;default:1" json:"num"`
}

// TableName OrderSubject's table name
func (*OrderSubject) TableName() string {
	return TableNameOrderSubject
}
