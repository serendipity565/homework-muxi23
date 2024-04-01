package controller

import (
	"fmt"
	"net/http"
	"strings"
	"test/service"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	repassword := r.FormValue("rePassword")
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "输入用户名或密码为空！")
		return
	}
	if strings.TrimSpace(password) != strings.TrimSpace(repassword) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "两次密码不一致！")
		return
	}
	newuser := User{
		Name:     username,
		Password: password,
	}
	num := service.Finduser(newuser)
	if num == 0 {
		err := service.Register(newuser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "注册失败！%v", err)
			return
		} else {
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, "注册成功！")
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "用户已存在！")
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "输入用户名或密码为空！")
		return
	}
	newuser := User{
		Name:     username,
		Password: password,
	}
	num := service.Finduser(newuser)
	if num == 1 {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "登入成功！")
		return
	} else if num == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "此用户不存在！")
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "密码错误！")
		return
	}
}
