package dao

import (
	"gobookstore/model"
	"gobookstore/utils"
	"net/http"
)

// 添加
func AddSession(sess *model.Session) error {

	sqlStr := "INSERT INTO sessions values (?,?,?)"

	//执行
	_, err := utils.Db.Exec(sqlStr, sess.SessionId, sess.UserName, sess.UserId)
	if err != nil {
		return err
	}
	return nil
}

// 删除
func DeleteSession(sessID string) error {

	sqlStr := "DELETE FROM sessions WHERE session_id = ?"

	//执行
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

// 删除
func DeleteSessionByuserID(sessID string) error {

	sqlStr := "DELETE FROM sessions WHERE user_id = ?"

	//执行
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

// 查询
func GetSession(sessID string) (*model.Session, error) {

	sqlStr := "SELECT session_id,username,user_id FROM sessions WHERE session_id = ?"
	//预编译，防止SQL注入
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessID)

	sess := &model.Session{}
	//扫描数据库字段并为sess赋值
	row.Scan(&sess.SessionId, &sess.UserName, &sess.UserId)
	return sess, nil
}

// 判断用户是否已经登录
func IsLogin(r *http.Request) (bool, *model.Session) {

	//依据cookie的name获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {

		//获取value
		cookieValue := cookie.Value
		//查session
		session, _ := GetSession(cookieValue)
		if session.UserId > 0 { //已登录
			return true, session
		}

	}
	return false, nil
}
