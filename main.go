package main

import (
	"gobookstore/controller"
	"net/http"
)

func main() {

	//设置处理静态资源
	//http.StripPrefix 是一个中间件函数，用于在处理HTTP请求前移除URL路径中的指定前缀
	//会匹配以/static开头的路径,替换成views/static，去对应目录找.css文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/main", controller.GetPageBooksByPrice)

	//登录
	http.HandleFunc("/login", controller.LoginHandler)
	//注销
	http.HandleFunc("/logout", controller.LogoutHandler)

	//去注册
	http.HandleFunc("/regist", controller.RegisterHandler)

	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/CheckRegister", controller.CheckRegisterHandler)

	//获取所有图书
	http.HandleFunc("/getBooks", controller.GetBooks)

	//获取当前页图书
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)

	//添加图书
	//http.HandleFunc("/addBooks", controller.AddBook)

	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)

	//更新图书
	//http.HandleFunc("/modifyBook", controller.ModifyBook)

	//更新或添加图书
	http.HandleFunc("/addandupdateBooks", controller.AddandUpateBook)

	//添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)

	//获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)

	//清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)

	//删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)

	//更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)

	//去更新图书的页面
	http.HandleFunc("/getOneBook", controller.GetOneBook)

	//范围查询图书价格
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)

	//去结账
	http.HandleFunc("/checkout", controller.Checkout)

	//获取所有订单
	http.HandleFunc("/getOrders", controller.GetOrders)

	//获取我的订单
	http.HandleFunc("/getMyOrders", controller.GetMyOrders)

	//获取订单对应的订单项
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)

	//发货
	http.HandleFunc("/sendOrder", controller.SendOrder)

	//收货
	http.HandleFunc("/receiveOrder", controller.ReceiveOrder)

	http.ListenAndServe(":8080", nil)
}
