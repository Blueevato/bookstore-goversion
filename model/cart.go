package model

// 购物车
type Cart struct {
	CartID      string      //购物车ID
	CartItems   []*CartItem //购物车中的所有购物项
	TotalCount  int64       //购物车中图书总数量
	TotalAmount float64     //总金额
	UserID      int         //绑定的外键
	UserName    string      //用户名
}

// 总数计算
func (cart *Cart) GetTotalCount() int64 {

	var res int64 = 0
	for _, v := range cart.CartItems {
		res += v.Count
	}
	return res
}

// 总价计算
func (cart *Cart) GetTotalAmount() float64 {

	var res float64 = 0.0
	for _, v := range cart.CartItems {
		//res += v.Amount //不能用v.Amount ,默认0
		res += v.GetAmount()
	}
	return res

}
