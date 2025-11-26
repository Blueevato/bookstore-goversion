package dao

import (
	"fmt"
	"gobookstore/model"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("开始测试：")
	t.Run("验证用户:", testCheckLogin)
	t.Run("用户注册：", testRegisterUser)
}

func testCheckLogin(t *testing.T) {
	user, _ := CheckLogin("admin2", "666")
	fmt.Println("用户信息:", user)
}

func TestCheckLogin(t *testing.T) {
	user, _ := CheckLogin("admin2", "666")
	fmt.Println("用户信息:", user)
}

func TestRegister(t *testing.T) {
	user, _ := CheckRegister("admin2")
	fmt.Println("用户信息:", user)
}

func testRegisterUser(t *testing.T) {
	RegisterUser("admin66", "admin77", "admin88")
}

func TestRegisterUser(t *testing.T) {
	RegisterUser("admin661", "admin77", "admin88")
}

func TestAddSession(t *testing.T) {
	sess := &model.Session{
		SessionId: "123",
		UserName:  "123",
		UserId:    15, //外键
	}
	AddSession(sess)
}

func TestDeleteSession(t *testing.T) {
	DeleteSession("123")
}
func TestDeleteSessionByuserID(t *testing.T) {
	DeleteSessionByuserID("16")
}

func TestGetSession(t *testing.T) {
	usersession, err := GetSession("dd122de8-f103-46b3-be71-a0742ef432be")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usersession)
}
