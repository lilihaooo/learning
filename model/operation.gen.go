// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOperation = "operation"

// Operation mapped from table <operation>
type Operation struct {
	ID       int64 `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	Type     int64 `gorm:"column:type;type:int;not null" json:"type"` // 1为下订单, 2为充值
	ActionID int64 `gorm:"column:action_id;type:int;not null" json:"action_id"`
}

// TableName Operation's table name
func (*Operation) TableName() string {
	return TableNameOperation
}