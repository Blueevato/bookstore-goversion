package model

// 购物项
type CartItem struct {
	CartItemID int64 //购物项ID
	Book       *Book
	Count      int64
	Amount     float64
	CartID     string //购物车id
}

// 总价计算
func (cartItem *CartItem) GetAmount() float64 {

	price := cartItem.Book.Price
	return float64(cartItem.Count) * price

}
