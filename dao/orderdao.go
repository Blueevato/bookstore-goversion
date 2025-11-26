package dao

import (
	"fmt"
	"gobookstore/model"
	"gobookstore/utils"
	"time"
)

// 插入订单
func AddOrder(order *model.Order) error {

	sqlStr := "insert into orders (id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}

// 获取所有订单
func GetOrders() ([]*model.Order, error) {
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		var createTimeStr string
		rows.Scan(&order.OrderID, &createTimeStr, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		// 手动解析时间字符串
		order.CreateTime, err = time.Parse("2006-01-02 15:04:05", createTimeStr)
		if err != nil {
			return nil, fmt.Errorf("parse time failed: %v", err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// 获取我的订单
func GetMyOrders(userID string) ([]*model.Order, error) {

	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id = ?"
	rows, err := utils.Db.Query(sqlStr, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		var createTimeStr string
		rows.Scan(&order.OrderID, &createTimeStr, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		// 手动解析时间字符串
		order.CreateTime, err = time.Parse("2006-01-02 15:04:05", createTimeStr)
		if err != nil {
			return nil, fmt.Errorf("parse time failed: %v", err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// 更新订单状态
func UpdateOrderState(orderID string, state int64) error {

	sqlStr := "update orders set state=? where id=?"
	_, err := utils.Db.Exec(sqlStr, state, orderID)
	if err != nil {
		return err
	}
	return nil
}
