package controller

import (
	"encoding/json"
	"fmt"
	"gobookstore/dao"
	"gobookstore/model"
	"gobookstore/utils"
	"html/template"
	"net/http"
	"strconv"
)

// 获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	_, sess := dao.IsLogin(r)
	//获取userid,并查对应购物车
	userID := sess.UserId
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil { //购物车不为空
		//设置用户名信息
		cart.UserName = sess.UserName //当 cart 为 nil 时，设置 cart.UserName 会发生了空指针解引用
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, cart)
	} else { //购物车为空
		//解析模板文件
		fmt.Println("购物车为空")
		// 创建空的购物车对象
		cart = &model.Cart{
			UserID:      sess.UserId,
			UserName:    sess.UserName,
			CartItems:   []*model.CartItem{},
			TotalCount:  0,
			TotalAmount: 0,
		}
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, cart)
	}
}

// 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, sess := dao.IsLogin(r)
	if flag {

		//获取要添加的图书id
		bookID := r.FormValue("bookId")
		//fmt.Println("要添加的图书id是：", bookID) //test
		ibookID, _ := strconv.Atoi(bookID)
		book, _ := dao.GetOneBook(ibookID)
		//获取用户id
		userID := sess.UserId
		//判断数据库是否有当前用户的购物车
		cart, _ := dao.GetCartByUserID(userID)
		if cart != nil {
			//已经有购物车了
			//step1: 需要判断购物车中是否有这本图书
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			if cartItem != nil { //有这本书
				fmt.Println("当前购物车中有这本图书")
				//获取购物车中的购物项
				cts := cart.CartItems
				for _, v := range cts {
					//debug
					//fmt.Println("Book:", v)
					//fmt.Println("cartItem:", cartItem.Book) //nil

					//寻找匹配的购物项
					if v.Book.ID == cartItem.Book.ID {
						//将当前购物项数量++
						v.Count++
						//写回数据库
						err := dao.UpdateBookCount(v)
						if err != nil {
							fmt.Println(err)
							return
						}
					}
				}

			} else { //没这本书
				//添加购物项
				fmt.Println("当前购物车中还没有该图书对应的购物项")
				cartItem = &model.CartItem{
					CartID: cart.CartID,
					Book:   book,
					Count:  1,
				}
				//购物项加入当前购物车
				cart.CartItems = append(cart.CartItems, cartItem)
				//存到数据库中
				dao.AddCartItem(cartItem)
			}

			//不管有没有这本书，都要更新购物车的图书总数和总金额
			//cart中book是nil，会panic
			dao.UpdateCart(cart)

		} else {
			//没有购物车
			//step1 :创建购物车
			cartID, _ := utils.CreatUUID()
			cart := &model.Cart{
				UserID: userID,
				CartID: cartID,
			}
			//step2:添加购物项
			cartItem := &model.CartItem{
				CartID: cartID,
				Book:   book,
				Count:  1,
			}
			//step3:购物项加入购物车
			cartItems := []*model.CartItem{cartItem} //生成切片
			cart.CartItems = cartItems
			//step4：存到数据库中
			dao.AddCart(cart)
		}

		w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车中")) //响应书名
	} else { //未登录
		w.Write([]byte("请先登录!")) //响应书名
	}

}

// 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物车id
	cartID := r.FormValue("CartId")
	fmt.Println("cartID", cartID)
	//清空购物车
	dao.DeleteCartByCartID(cartID)
	//刷新
	GetCartInfo(w, r)
}

// 删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物项id
	cartItemId := r.FormValue("cartItemId")
	icartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)
	//获取session
	_, sess := dao.IsLogin(r)
	userID := sess.UserId
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemID == icartItemId {
			//找到删除的购物项并移除
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems //赋值回去
			//删除数据库里的购物项
			dao.DeleteCartItemByID(cartItemId)
		}
	}
	//更新数据库中的购物车
	dao.UpdateCart(cart)
	//刷新
	GetCartInfo(w, r)
}

// 更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要更新的购物项id,用户输入的数量
	cartItemId := r.FormValue("cartItemId")
	bookCount := r.FormValue("bookCount")
	ibookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	icartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)
	//获取session
	_, sess := dao.IsLogin(r)
	userID := sess.UserId
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if v.CartItemID == icartItemId { //找到购物项

			v.Count = ibookCount //切片是指针，修改的就是底层数组,不用赋值回去
			//更新数据库里的购物项
			dao.UpdateBookCount(v)
		}
	}
	//更新数据库中的购物车
	dao.UpdateCart(cart)
	//刷新
	//GetCartInfo(w, r)
	//获取购物车信息
	cart, _ = dao.GetCartByUserID(userID)
	fmt.Println("购物车信息", cart)
	totalCount := cart.TotalCount
	totalAmount := cart.TotalAmount
	var amount float64
	cIs := cart.CartItems
	for _, v := range cIs {
		if v.CartItemID == icartItemId {
			//找到购物项，获取当前金额小计
			amount = v.Amount
		}
	}
	//声明一个切片,转换成字符串
	data := &model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	//strTotalCount := strconv.FormatInt(totalCount, 10)
	//strTotalAmount := strconv.FormatFloat(totalAmount, 'f', 2, 64)
	//strAmount := strconv.FormatFloat(amount, 'f', 2, 64)
	//data = append(data, strTotalCount, strTotalAmount, strAmount)
	//转成json
	json, _ := json.Marshal(data)
	//响应到浏览器
	w.Write(json)
}
