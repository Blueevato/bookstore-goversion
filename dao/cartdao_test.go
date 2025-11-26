package dao

import (
	"fmt"
	"gobookstore/model"
	"testing"
)

func TestAddCart(t *testing.T) {
	book1 := &model.Book{
		ID:    3,
		Price: 44.50,
	}

	book2 := &model.Book{
		ID:    5,
		Price: 19.30,
	}
	//创建两个购物项
	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  2,
		CartID: "6688",
	}
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  2,
		CartID: "6688",
	}

	cartItems := []*model.CartItem{cartItem1, cartItem2}
	//创建购物车
	cart := &model.Cart{
		CartID:    "6688",
		CartItems: cartItems,
		UserID:    16,
	}
	AddCart(cart)
}

func TestGetCartItemByBookID(t *testing.T) {
	id, err := GetCartItemByBookID("3")
	if err != nil {
		return
	}
	fmt.Println(id)
}

func TestGetCartItemByBookIDAndCartID(t *testing.T) {
	id, err := GetCartItemByBookIDAndCartID("3", "6688")
	if err != nil {
		return
	}
	fmt.Println("查询到的数据：", id)
}

func TestGetCartItemByCartID(t *testing.T) {
	id, err := GetCartItemByCartID("6688")
	if err != nil {
		return
	}
	for k, v := range id {
		fmt.Println(k+1, v)
	}

}

func TestGetCartByUserID(t *testing.T) {
	cart, err := GetCartByUserID(16)
	if err != nil {
		return
	}
	for k, v := range cart.CartItems {
		fmt.Println(k+1, v)
	}
	fmt.Println(cart)
}

func TestDeleteCartByCartID(t *testing.T) {
	err := DeleteCartByCartID("6688")
	fmt.Println(err)
}
