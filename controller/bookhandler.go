package controller

import (
	"gobookstore/dao"
	"gobookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

// 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNum := r.FormValue("pageNum")
	if pageNum == "" { //为空的判断
		pageNum = "1"
	}
	ipageNum, _ := strconv.ParseInt(pageNum, 10, 64)

	//调用bookdao
	page, _ := dao.GetPages(ipageNum)

	t := template.Must(template.ParseFiles("./views/index.html"))
	t.Execute(w, page)

}

// 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {

	//调用bookdao
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

// 获取当前页图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {

	//获取页码
	pageNum := r.FormValue("pageNum")
	if pageNum == "" { //为空的判断
		pageNum = "1"
	}

	ipageNum, _ := strconv.ParseInt(pageNum, 10, 64)

	//调用bookdao
	page, _ := dao.GetPages(ipageNum)

	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

// 范围查询图书价格
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {

	//获取页码
	pageNum := r.FormValue("pageNum")
	if pageNum == "" { //为空的判断
		pageNum = "1"
	}

	//获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")

	ipageNum, _ := strconv.ParseInt(pageNum, 10, 64)

	var page *model.Page

	if minPrice == "" || maxPrice == "" {
		//调用bookdao
		page, _ = dao.GetPages(ipageNum)
	} else {
		//调用bookdao
		page, _ = dao.GetPagesByPrice(ipageNum, minPrice, maxPrice)
		//将价格范围设置到page结构体中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}

	//调用IsLogin,判断是否已经登录
	flag, sess := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = sess.UserName
	}

	////获取Cookie
	//cookie, _ := r.Cookie("user")
	//if cookie != nil { //debuging
	//	//获取Cookie的值
	//	cookieValue := cookie.Value //即SessionID
	//	session, _ := dao.GetSession(cookieValue)
	//	if session.UserId > 0 { //查到有这个session
	//		//已经登录
	//		//设置Page中的IsLogin和Username字段
	//		page.IsLogin = true
	//		page.Username = session.UserName
	//	}
	//}

	t := template.Must(template.ParseFiles("./views/index.html"))
	t.Execute(w, page)
}

//// 新增图书
//func AddBook(w http.ResponseWriter, r *http.Request) {
//
//	//获取图书信息
//	title := r.FormValue("title")
//	author := r.FormValue("author")
//	fprice, _ := strconv.ParseFloat(r.FormValue("price"), 64)
//	istock, _ := strconv.ParseInt(r.FormValue("stock"), 10, 0)
//	isales, _ := strconv.ParseInt(r.FormValue("sales"), 10, 0)
//	book := &model.Book{
//
//		Title:   title,
//		Author:  author,
//		Stock:   int(istock),
//		Price:   fprice,
//		Sales:   int(isales),
//		ImgPath: "/static/images/default.jpg",
//	}
//
//	//调用bookdao
//	dao.AddBooks(book)
//
//	//调用GetBooks，再查询一次，=刷新
//	GetBooks(w, r)
//}

// 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	//获取图书信息
	bookId := r.FormValue("bookId")
	ibookId, _ := strconv.ParseInt(bookId, 10, 0)
	//调用bookdao
	dao.DeleteBooks(int(ibookId))

	//调用GetBooks，再查询一次，=刷新
	GetPageBooks(w, r)
}

//// 更新图书
//func ModifyBook(w http.ResponseWriter, r *http.Request) {
//
//	//获取图书信息
//	id := r.FormValue("bookId")
//	ibookId, _ := strconv.ParseInt(id, 10, 0)
//	title := r.FormValue("title")
//	author := r.FormValue("author")
//	fprice, _ := strconv.ParseFloat(r.FormValue("price"), 64)
//	istock, _ := strconv.ParseInt(r.FormValue("stock"), 10, 0)
//	isales, _ := strconv.ParseInt(r.FormValue("sales"), 10, 0)
//	book := &model.Book{
//		ID:      int(ibookId),
//		Title:   title,
//		Author:  author,
//		Stock:   int(istock),
//		Price:   fprice,
//		Sales:   int(isales),
//		ImgPath: "/static/images/default.jpg",
//	}
//
//	//调用bookdao
//	dao.ModifyBooks(book)
//
//	//调用GetBooks，再查询一次，=刷新
//	GetBooks(w, r)
//}

// 去更新图书的页面
func GetOneBook(w http.ResponseWriter, r *http.Request) {
	//获取图书信息
	id := r.FormValue("bookId")

	// 如果是添加图书，bookId 为空
	if id == "" {
		// 新增模式：传递 nil
		t := template.Must(template.ParseFiles("views/pages/manager/book_editandmodify.html"))
		t.Execute(w, nil)
		return
	}

	// 更新模式：获取图书数据
	ibookId, _ := strconv.ParseInt(id, 10, 0)
	book, _ := dao.GetOneBook(int(ibookId))

	t := template.Must(template.ParseFiles("views/pages/manager/book_editandmodify.html"))
	t.Execute(w, book)
}

// 更新或添加图书
func AddandUpateBook(w http.ResponseWriter, r *http.Request) {
	//获取图书信息
	id := r.FormValue("bookId")
	ibookId, _ := strconv.ParseInt(id, 10, 0)
	title := r.FormValue("title")
	author := r.FormValue("author")
	fprice, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	istock, _ := strconv.ParseInt(r.FormValue("stock"), 10, 0)
	isales, _ := strconv.ParseInt(r.FormValue("sales"), 10, 0)
	book := &model.Book{
		ID:      int(ibookId),
		Title:   title,
		Author:  author,
		Stock:   int(istock),
		Price:   fprice,
		Sales:   int(isales),
		ImgPath: "/static/images/default.jpg",
	}

	if book.ID > 0 { //调用bookdao
		dao.ModifyBooks(book)
	} else {
		dao.AddBooks(book)
	}

	//调用GetBooks，再查询一次，=刷新
	//GetBooks(w, r)
	GetPageBooks(w, r)
}
