package dao

import (
	"fmt"
	"gobookstore/model"
	"gobookstore/utils"
)

// 根据用户名和密码查一条记录，是否注册在籍
func CheckLogin(username string, password string) (*model.User, error) {
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

// 根据用户名查一条记录,是否重名
func CheckRegister(username string) (*model.User, error) {
	sqlStr := "select id,username,password,email from users where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

// 插入信息
func RegisterUser(username, password string, email string) error {
	//1 写sql语句
	sqlStr := "insert into users (username, password, email) values (?, ?, ?)"

	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		fmt.Println("插入异常2", err)
		return err
	}
	return nil
}
