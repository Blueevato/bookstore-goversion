package dao

import (
	"fmt"
	"gobookstore/model"
	"testing"
	"time"
)

func TestAddOrder(t *testing.T) {
	order := &model.Order{
		OrderID:     "6688",
		CreateTime:  time.Now(),
		TotalAmount: 300,
		TotalCount:  2,
		State:       0,
		UserID:      25,
	}
	err := AddOrder(order)
	if err != nil {
		fmt.Println(err)
	}
}

func TestAddOrderItem(t *testing.T) {
	orderItem1 := &model.OrderItem{
		Count:   1,
		Amount:  200,
		Title:   "测试书1",
		Author:  "测试1",
		Price:   200,
		ImgPath: "/static/img/default.jpg",
		OrderID: "6688",
	}
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "测试书2",
		Author:  "测试2",
		Price:   100,
		ImgPath: "/static/img/default.jpg",
		OrderID: "6688",
	}

	err := AddOrderItem(orderItem1)
	if err != nil {
		fmt.Println(err)
	}
	err1 := AddOrderItem(orderItem2)
	if err1 != nil {
		fmt.Println(err1)
	}
}

func TestGetOrders(t *testing.T) {
	orders, err := GetOrders()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range orders {
		fmt.Println(v)
	}
}

func TestGetMyOrders(t *testing.T) {
	orders, err := GetMyOrders("17")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range orders {
		fmt.Println(v)
	}
}

func TestGetOrderItemsByOrderID(t *testing.T) {
	orderItems, err := GetOrderItemsByOrderID("6688")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range orderItems {
		fmt.Println(v)
	}
}

func TestUpdateOrder(t *testing.T) {
	UpdateOrderState("6688", 1)
}
