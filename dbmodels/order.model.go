package dbmodels

import (
	"time"
)

// Order ...
type Order struct {
	ID                     	uint64    `gorm:"column:id;pk"`
	OrderNo                	string    `gorm:"column:order_no"`
	CreatedBy           	string    `gorm:"column(created_by);null"`
	CreatedDate             time.Time `gorm:"column(created_date);type(timestamp without time zone);null"`
	Total                  	int64     `gorm:"column:total"`
	ReffNo					string	  `gorm:"column:reff_no"`
}

// TableName ...
func (Order *Order) TableName() string {
	return "orders"
}