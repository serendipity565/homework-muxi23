package controller

import (
	"fmt"
	"net/http"
	"strings"
	"test/service"
)

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

	err := service.Register(username, password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "注册失败：%v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "注册成功！")
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "输入用户名或密码为空！")
		return
	}

	num := service.Finduser(username, password)
	switch num {
	case 1:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "登入成功！")
	case 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "此用户不存在！")
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "密码错误！")
	}
}
