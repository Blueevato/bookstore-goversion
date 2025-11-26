package model

type Page struct {
	Books        []*Book //图书切片
	PageSize     int64   //页大小
	PageNum      int64   //当前页
	TotalPageNum int64   //总页数
	TotalRecord  int64   //总记录数
	MinPrice     string
	MaxPrice     string
	IsLogin      bool   //是否登录状态
	Username     string //登录用户名
}

// 是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNum > 1
}

// 是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNum < p.TotalPageNum
}

// 获取上一页
func (p *Page) GetPre() int64 {
	if p.IsHasPrev() {
		return p.PageNum - 1
	}
	return 1
}

// 获取下一页
func (p *Page) GetNext() int64 {
	if p.IsHasNext() {
		return p.PageNum + 1
	}
	return p.PageNum
}
