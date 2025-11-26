package dao

import (
	"gobookstore/model"
	"gobookstore/utils"
)

// 插入订单项
func AddOrderItem(orderItem *model.OrderItem) error {

	sqlStr := "insert into order_items (count,amount,title,author,price,img_path,order_id) values (?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil
}

// 根据订单号查询该订单号所有订单项
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	sqlStr := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id = ?"
	rows, err := utils.Db.Query(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orderItems := make([]*model.OrderItem, 0)
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
