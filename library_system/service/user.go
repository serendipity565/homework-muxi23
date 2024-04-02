package service

import (
	"test/db"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 查找是否有该用户
func Finduser(username string, password string) int {
	var myuser User
	sqlstr := "select name,password from users where name = ?"
	err := db.Db.QueryRow(sqlstr, username).Scan(&myuser.Name, &myuser.Password)
	if err != nil {
		return 0 //无此用户
	} else if myuser.Name == username && myuser.Password == password {
		return 1 //用户密码正确
	} else {
		return 2 //密码错误
	}
}

// 注册用户
func Register(username string, password string) (err error) {
	tx, err := db.Db.Begin()
	if err != nil {
		tx.Rollback()
	}
	sqlstr := "insert into users (name, password) values (?,?)"
	_, err = db.Db.Exec(sqlstr, username, password)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return err
	}
}
