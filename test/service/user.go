package service

import (
	"test/controller"
	"test/db"
)

func Finduser(newuser controller.User) int {
	var myuser controller.User
	sqlstr := "select name,password from users where name = ?"
	err := db.Db.QueryRow(sqlstr, newuser.Name).Scan(&myuser.Name, &myuser.Password)
	if err != nil {
		return 0 //无此用户
	} else if myuser.Name == newuser.Name && myuser.Password == newuser.Password {
		return 1 //用户密码正确
	} else {
		return 2 //密码错误
	}
}

func Register(newuser controller.User) (err error) {
	tx, err := db.Db.Begin()
	if err != nil {
		tx.Rollback()
	}
	sqlstr := "insert into users (name, password) values (?,?)"
	_, err = db.Db.Exec(sqlstr, newuser.Name, newuser.Password)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return err
	}
}
