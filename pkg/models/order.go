package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shadheeraNa/go-ordermanagement/pkg/config"
)

var db *gorm.DB

type Order struct {
	OrderId      int       `json:"OrderId"`
	OrderPrice   string    `json:"OrderPrice"`
	OrderAddress string    `json:"OrderAddress"`
	Customer     *Customer `json:"Customer"`
}

type Customer struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	MobileNum string `json:"mobilenum"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Order{}, &Customer{})
}

func (b *Order) CreateOrder() *Order {
	db.NewRecord(b) //new record comes with the gorm and it creates a new record for the order
	db.Create(&b)
	return b
}

func GetAllOrders() []Order {
	var Orders []Order
	db.Find(&Orders)
	return Orders
}

func GetOrderById(Id int64) (*Order, *gorm.DB) {
	var getOrder Order
	db := db.Where("ID = ?", Id).Find(&getOrder)
	return &getOrder, db
}

func DeleteOrder(ID int64) Order {
	var order Order
	db.Where("ID = ?", ID).Delete(order)
	return order
}
