package mysql

import (
	"fmt"
	"questionplatform/global"
	"questionplatform/model"
)

// InsertUser 插入用户
func InsertUser(user model.User) (err error) {
	exec := "insert into users (username,password) values (?,?)"
	_, err = global.DB.Exec(exec, user.Username, user.Password)
	if err != nil {
		err = fmt.Errorf("failed to InsertUser: %w", err)
	}
	return
}

// SearchUserByUsername 根据用户名查询用户信息
func SearchUserByUsername(username string) (user model.User, err error) {
	query := "SELECT  id,username,password FROM users WHERE username =?"
	row := global.DB.QueryRow(query, username)
	err = row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		err = fmt.Errorf("failed to SearchUserByUsername: %w", err)
	}
	return
}

// Changepass 修改密码
func Changepass(id int, newpass string) (err error) {
	exec := "update users set password =? where id = ? "
	_, err = global.DB.Exec(exec, newpass, id)
	if err != nil {
		err = fmt.Errorf("failed to Changepass: %w", err)
	}
	return
}

// InsertUserInfo 插入或者是修改用户信息
func InsertUserInfo(user model.User) (err error) {
	exec := "update users set nickname=?,signature=? where id = ? "
	_, err = global.DB.Exec(exec, user.Nickname, user.Signature, user.Id)
	if err != nil {
		err = fmt.Errorf("failed to InsertUserInfo: %w", err)
	}
	return
}

// GetUserInfo 根据 id 获取用户信息
func GetUserInfo(id int) (info model.User, err error) {
	query := "select id,username,nickname,signature from users where id =?"
	row := global.DB.QueryRow(query, id)
	err = row.Scan(&info.Id, &info.Username, &info.Nickname, &info.Signature)
	if err != nil {
		err = fmt.Errorf("failed to GetUserInfo: %w", err)
	}
	return
}
