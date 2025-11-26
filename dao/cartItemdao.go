package dao

import (
	"gobookstore/model"
	"gobookstore/utils"
	"strconv"
)

// 插入购物项
func AddCartItem(cartItem *model.CartItem) error {

	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?) "

	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

// 依据图书id获取对应购物项
func GetCartItemByBookID(bookID string) (*model.CartItem, error) {

	sqlStr := "select id,COUNT,amount,cart_id from cart_items where book_id = ?"

	//执行
	row := utils.Db.QueryRow(sqlStr, bookID)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

// 依据购物车id获取购物车中所有购物项
func GetCartItemByCartID(cartID string) ([]*model.CartItem, error) {

	sqlStr := "select id,COUNT,amount,book_id,cart_id from cart_items where cart_id = ?"

	//执行
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	cartItems := make([]*model.CartItem, 0)
	for rows.Next() {
		//设置一个变量接收bookid
		var bookID string
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err2 != nil {
			return nil, err
		}
		//根据bookID获取图书信息
		ibookID, _ := strconv.Atoi(bookID)
		book, _ := GetOneBook(ibookID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

// 依据图书id和购物车id获取对应购物项
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	sqlStr := "select id,COUNT,amount,cart_id from cart_items where book_id = ? and cart_id = ?"

	//执行
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	cartItem := &model.CartItem{}

	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	//查询图书信息
	ibookID, _ := strconv.Atoi(bookID)
	book, _ := GetOneBook(ibookID)
	cartItem.Book = book
	return cartItem, nil
}

// 根据图书id和购物车id 更新购物车购物项的图书数量及金额小计
func UpdateBookCount(cartItem *model.CartItem) error {
	sqlStr := "update cart_items set COUNT = ?,amount = ? where book_id = ? and cart_id = ?"

	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

// 根据购物车id删除所有购物项
func DeleteCartItemByCartID(cartID string) error {
	sqlStr := "delete from cart_items where cart_id = ?"

	//执行
	_, err := utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}

// 根据购物项id删除购物项
func DeleteCartItemByID(cartItemID string) error {
	sqlStr := "delete from cart_items where id = ?"

	//执行
	_, err := utils.Db.Exec(sqlStr, cartItemID)
	if err != nil {
		return err
	}
	return nil
}
