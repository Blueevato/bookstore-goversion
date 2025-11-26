package model

import (
	"time"
)

// 订单
type Order struct {
	OrderID     string
	CreateTime  time.Time
	TotalCount  int64
	TotalAmount float64
	State       int64 //0未发货 1已发货 2交易完成
	UserID      int64
}

// 未发货
func (order Order) NoSend() bool {
	return order.State == 0
}

// 已发货
func (order Order) YesSend() bool {
	return order.State == 1
}

// 已收货
func (order Order) ReveiveDone() bool {
	return order.State == 2
}
