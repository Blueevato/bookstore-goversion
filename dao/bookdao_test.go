package dao

import (
	"fmt"
	"gobookstore/model"
	"testing"
)

func TestQueryAllBook(t *testing.T) {

	books, _ := GetBooks()
	for k, v := range books {
		fmt.Println("第", k+1, "本图书信息：", v)
	}

}

func TestAddBook(t *testing.T) {

	book := &model.Book{
		Title:   "测试图书",
		Author:  "测试",
		Price:   88.88,
		Sales:   200,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	AddBooks(book)
}

func TestDeleteBooks(t *testing.T) {

	id := 18
	DeleteBooks(id)
}

func TestModifyBooks(t *testing.T) {

	book := &model.Book{
		ID:      22,
		Title:   "测试图书1",
		Author:  "测试",
		Price:   88.88,
		Sales:   200,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	ModifyBooks(book)
}

func TestGetOneBook(t *testing.T) {

	id := 18
	book, _ := GetOneBook(id)
	fmt.Println("该书信息是：", book)
}

func TestGetPages(t *testing.T) {

	pageNums := int64(1)
	pages, _ := GetPages(pageNums)
	fmt.Println("当前页：", pages.PageNum)
	fmt.Println("总页数：", pages.TotalPageNum)
	fmt.Println("总记录数：", pages.TotalRecord)
	fmt.Println("当前页图书信息：")
	for _, v := range pages.Books {
		fmt.Println(v)
	}
}

func TestGetPagesByPrice(t *testing.T) {

	pageNums := int64(2)
	pages, _ := GetPagesByPrice(pageNums, "0", "30")
	fmt.Println("当前页：", pages.PageNum)
	fmt.Println("总页数：", pages.TotalPageNum)
	fmt.Println("总记录数：", pages.TotalRecord)
	fmt.Println("当前页图书信息：")
	for _, v := range pages.Books {
		fmt.Println(v)
	}
}
