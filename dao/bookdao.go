package dao

import (
	"gobookstore/model"
	"gobookstore/utils"
)

// 获取数据库所有图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"

	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 重要：记得关闭 rows

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	return books, nil
}

// 添加图书
func AddBooks(b *model.Book) error {

	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"

	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// 删除图书
func DeleteBooks(id int) error {
	sqlStr := "delete from books where id=?"
	_, err := utils.Db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

// 修改图书
func ModifyBooks(b *model.Book) error {
	sqlStr := "UPDATE books SET title=?, author=?, price=?, sales=?, stock=?, img_path=? WHERE id=?"
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath, b.ID)
	if err != nil {
		return err
	}
	return nil
}

// 查一条图书 byid
func GetOneBook(id int) (*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, id)

	book := &model.Book{}
	//赋值
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)

	return book, nil

}

// 查页码及图书信息
func GetPages(pageNums int64) (*model.Page, error) {
	sqlStr := "select count(*) from books"

	//执行
	row := utils.Db.QueryRow(sqlStr)

	var totalRecord int64
	row.Scan(&totalRecord) //写入总记录数

	pageSize := int64(4) //每页只显示4条记录
	var totalPageNum int64
	if totalRecord%pageSize == 0 {
		totalPageNum = totalRecord / pageSize
	} else {
		totalPageNum = (totalRecord / pageSize) + 1
	}

	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"

	//执行
	rows, err := utils.Db.Query(sqlStr2, (pageNums-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 重要：记得关闭 rows

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	pages := &model.Page{
		TotalRecord:  totalRecord,
		TotalPageNum: totalPageNum,
		PageSize:     pageSize,
		PageNum:      pageNums, //当前页
		Books:        books,
	}

	return pages, nil
}

// 查页码及图书信息及图书价格
func GetPagesByPrice(pageNums int64, minPrice string, maxPrice string) (*model.Page, error) {
	sqlStr := "select count(*) from books where price between ? and ?"

	//执行
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)

	var totalRecord int64
	row.Scan(&totalRecord) //写入总记录数

	pageSize := int64(4) //每页只显示4条记录
	var totalPageNum int64
	if totalRecord%pageSize == 0 {
		totalPageNum = totalRecord / pageSize
	} else {
		totalPageNum = (totalRecord / pageSize) + 1
	}

	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"

	//执行
	rows, err := utils.Db.Query(sqlStr2, minPrice, maxPrice, (pageNums-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 重要：记得关闭 rows

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	pages := &model.Page{
		TotalRecord:  totalRecord,
		TotalPageNum: totalPageNum,
		PageSize:     pageSize,
		PageNum:      pageNums, //当前页
		Books:        books,
	}

	return pages, nil
}
