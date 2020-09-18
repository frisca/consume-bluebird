package dbmodels

type OrderDetail struct {
	ID           uint64    `gorm:"column:id;pk"`
	OrderID      uint64    `gorm:"column:order_id"`
	Product 	 string    `gorm:"column:product"`
	Qty          int       `gorm:"column:qty"`
	Price        int       `gorm:"column:price"`
}

// TableName ...
func (OrderDetail *OrderDetail) TableName() string {
	return "orders_detail"
}