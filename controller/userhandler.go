package controller

import (
	"fmt"
	"gobookstore/dao"
	"gobookstore/model"
	"gobookstore/utils"
	"html/template"
	"net/http"
)

// Login handler处理用户登录
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	//判断是否已登录
	flag, _ := dao.IsLogin(r)
	if flag {
		//已登录
		//去首页
		GetPageBooksByPrice(w, r)
		return
	}

	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckLogin(username, password)
	if user.ID <= 0 { //go结构体有默认值 0
		//不正确
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	} else {
		//正确
		uuid, err := utils.CreatUUID()
		if err != nil {
			fmt.Println(err)
		}

		//// 先删除该用户的所有旧 Session
		//dao.DeleteSessionByuserID(strconv.Itoa(user.ID))

		//创建一个Session
		sess := &model.Session{
			SessionId: uuid, //uuid
			UserName:  user.Username,
			UserId:    user.ID,
		}

		//Session保存到数据库
		dao.AddSession(sess)
		//创建Cookie，关联Session
		cookie := http.Cookie{
			Name:     "user",
			Value:    uuid,
			HttpOnly: true,
		}
		//将Cookie发送给浏览器
		http.SetCookie(w, &cookie)

		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, user)
	}

}

// Logout handler处理用户注销
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	//获取Cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		fmt.Println(err)
	}
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		dao.DeleteSession(cookieValue)
		//设置Cookie失效
		cookie.MaxAge = 0 // 0立即失效，负数浏览器关闭失效

		//将修改后的Cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//去首页
	GetPageBooksByPrice(w, r)
}

// Register 处理用户注册
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	// 验证输入是否为空
	if username == "" || password == "" || email == "" {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名、密码和邮箱不能为空！")
		return
	}

	//调用userdao中验证用户名的方法
	user, _ := dao.CheckRegister(username)
	if user.ID > 0 { //go结构体有默认值 0
		//用户已存在,重新注册
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		//正确注册
		dao.RegisterUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

// Ajax 检查注册用户名
func CheckRegisterHandler(w http.ResponseWriter, r *http.Request) {

	//获取输入的用户名
	username := r.PostFormValue("username")
	fmt.Println("CheckRegisterHandler")

	// 验证输入是否为空
	if username == "" {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名不能为空！")
		return
	}

	//调用userdao中验证用户名的方法
	user, _ := dao.CheckRegister(username)
	if user.ID > 0 { //go结构体有默认值 0
		//用户已存在,响应
		w.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		//用户已存在,响应
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}

}
