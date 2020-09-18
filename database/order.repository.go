package database

import (
	"consume-bluebird/dbmodels"
	"fmt"
	"consume-bluebird/worker/consume/models"
	"strings"
	"strconv"
)


func SaveOrder(order dbmodels.Order, detail []models.Products) {
	db := GetDbConnect()
	var dataOrder dbmodels.Order
	
	err := db.Save(&order).Scan(&dataOrder).Error
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	if len(detail) > 0 {
		SaveOrderDetail(dataOrder.ID, detail)
	}

	fmt.Println("berhasil simpan ke database")
}

func SaveOrderDetail(orderID uint64, detail []models.Products) {
	db := GetDbConnect()
	
	for _, data := range detail {
		var orderdetail dbmodels.OrderDetail

		orderdetail.OrderID = orderID
		orderdetail.Product = data.Product
		orderdetail.Qty = data.Qty
		orderdetail.Price = int(data.Price)

		errDet := db.Save(&orderdetail).Error
		if errDet != nil {
			fmt.Println("error: ", errDet)
			return
		}
	}
	fmt.Println("Data berhasil di save")
}

func GenerateOrderNo() string {
	db := GetDbConnect()

	var order []dbmodels.Order
	err := db.Model(&dbmodels.Order{}).Order("order_no desc").First(&order).Error

	if err != nil {
		return "INV0001"
	}

	if len(order) > 0 {
		orderprefix := strings.TrimPrefix(order[0].OrderNo, "INV")
		latestCode, err := strconv.Atoi(orderprefix)
		if err != nil {
			fmt.Printf("error")
			return "INV0001"
		}
		orderpadding := fmt.Sprintf("%04s", strconv.Itoa(latestCode+1))
		return "INV" + orderpadding
	}
	return "INV0001"

}