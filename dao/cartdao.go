package dao

import (
	"gobookstore/model"
	"gobookstore/utils"
)

// 插入购物项
func AddCart(cart *model.Cart) error {

	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?) "

	//执行
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}

	//获取购物车中的所有购物项
	cartItems := cart.CartItems
	for _, item := range cartItems {
		//插入购物项表
		AddCartItem(item)
	}
	return nil
}

// 依据userid查对应购物车
func GetCartByUserID(userId int) (*model.Cart, error) {

	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id = ?"

	//执行
	row := utils.Db.QueryRow(sqlStr, userId)
	//创建一个购物车
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}

	//获取当前购物车中的购物项
	cartItems, _ := GetCartItemByCartID(cart.CartID)
	cart.CartItems = cartItems

	return cart, nil
}

// 更新购物车总数和总金额
func UpdateCart(cart *model.Cart) error {

	sqlStr := "update carts set total_count = ?, total_amount = ? where id = ?"

	//执行
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

// 根据购物车id删除购物车
func DeleteCartByCartID(cartID string) error {
	//删除购物车前需要删除对应的所有购物项
	err := DeleteCartItemByCartID(cartID)
	if err != nil {
		return err
	}

	//删除购物车
	sqlStr := "delete from carts where id = ?"
	//执行
	_, err1 := utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err1
	}
	return nil
}
