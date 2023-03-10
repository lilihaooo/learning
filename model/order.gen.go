// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOrder = "order"

// Order mapped from table <order>
type Order struct {
	ID      int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	Title   string `gorm:"column:title;type:varchar(20);not null" json:"title"`
	UserID  int64  `gorm:"column:user_id;type:int;not null" json:"user_id"`
	PayeeID int64  `gorm:"column:payee_id;type:int;not null" json:"payee_id"`
	Status  int64  `gorm:"column:status;type:int;not null" json:"status"`
	Updated int64  `gorm:"column:updated;type:int" json:"updated"`
	Created int64  `gorm:"column:created;type:int" json:"created"`
	Deleted int64  `gorm:"column:deleted;type:int" json:"deleted"`
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}
