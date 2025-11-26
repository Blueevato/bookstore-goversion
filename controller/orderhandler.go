package controller

import (
	"gobookstore/dao"
	"gobookstore/model"
	"gobookstore/utils"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

// 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	//获取session
	_, sess := dao.IsLogin(r)
	userID := sess.UserId

	//获取购物车
	cart, _ := dao.GetCartByUserID(userID)
	//创建Order
	orderID, _ := utils.CreatUUID()
	order := &model.Order{
		OrderID:     orderID,
		UserID:      int64(userID),
		CreateTime:  time.Now(),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
	}
	//订单存回数据库
	dao.AddOrder(order)

	//保存订单项
	cartItems := cart.CartItems
	for _, item := range cartItems {
		orderItem := &model.OrderItem{
			OrderID: orderID,
			Count:   item.Count,
			Amount:  item.Amount,
			Title:   item.Book.Title,
			Author:  item.Book.Author,
			Price:   item.Book.Price,
			ImgPath: item.Book.ImgPath,
		}
		//订单项存回数据库
		dao.AddOrderItem(orderItem)
		//更新当前购物项中图书的库存和销量
		book := item.Book
		book.Sales += int(item.Count)
		book.Stock -= int(item.Count)
		dao.ModifyBooks(book) //写回数据库
	}

	//清空购物车
	dao.DeleteCartByCartID(cart.CartID)
	//将订单号设置到session中
	sess.Order = order

	//返回模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, sess)
}

// 获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {

	orders, _ := dao.GetOrders()

	//返回模板
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

// 获取我的订单
func GetMyOrders(w http.ResponseWriter, r *http.Request) {

	//获取session
	_, sess := dao.IsLogin(r)
	userId := sess.UserId

	//调方法
	orders, _ := dao.GetMyOrders(strconv.Itoa(userId))
	sess.Orders = orders

	//返回模板
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, sess)

}

// 获取订单对应的订单项
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {

	orderId := r.FormValue("orderId")
	orderItems, _ := dao.GetOrderItemsByOrderID(orderId)

	//返回模板
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

// 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {

	//获取单号
	orderId := r.FormValue("orderId")

	dao.UpdateOrderState(orderId, 1)
	//刷新
	GetOrders(w, r)

}

// 收货
func ReceiveOrder(w http.ResponseWriter, r *http.Request) {

	//获取单号
	orderId := r.FormValue("orderId")

	dao.UpdateOrderState(orderId, 2)
	//刷新
	GetMyOrders(w, r)

}
