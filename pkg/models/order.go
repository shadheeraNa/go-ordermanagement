package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shadheeraNa/go-ordermanagement/pkg/config"
)

var db *gorm.DB

type Order struct {
	gorm.Model
	OrderId      uint   `json:"OrderId"`
	OrderPrice   string `json:"OrderPrice"`
	OrderAddress string `json:"OrderAddress"`
	// CustomerID   uint     `json:"CustomerID"`
	// Customer     Customer `json:"Customer" gorm:"foreignKey:CustomerID"`
}

// type Customer struct {
// 	gorm.Model
// 	CustomerID uint   `json:"CustomerID"`
// 	FirstName  string `json:"FirstName"`
// 	LastName   string `json:"LastName"`
// 	MobileNum  string `json:"MobileNum"`
// }

func init() {
	config.Connect()
	db = config.GetDB()
	// db.AutoMigrate(&Order{}, &Customer{})
	db.AutoMigrate(&Order{})
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

//customer

// func (c *Customer) CreateCustomer() *Customer {
// 	db.NewRecord(c) //new record comes with the gorm and it creates a new record for the order
// 	db.Create(&c)
// 	return c
// }

// func GetAllCustomers() []Customer {
// 	var Customers []Customer
// 	db.Find(&Customers)
// 	return Customers
// }

// func GetCustomerById(Id int64) (*Customer, *gorm.DB) {
// 	var getCustomer Customer
// 	db := db.Where("ID = ?", Id).Find(&getCustomer)
// 	return &getCustomer, db
// }

// func DeleteCustomer(ID int64) Customer {
// 	var customer Customer
// 	db.Where("ID = ?", ID).Delete(customer)
// 	return customer
// }
