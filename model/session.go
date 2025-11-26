package model

// 会话
type Session struct {
	SessionId string
	UserName  string
	UserId    int //外键
	Order     *Order
	Orders    []*Order
}
