package controller

import (
	"fmt"
	"net/http"
	"strings"
	"test/service"
)

func Find(w http.ResponseWriter, r *http.Request) {
	//展示书籍
	fmt.Fprintf(w, "%v", service.Listbooks())
	//查找书籍
	bookname := r.PostFormValue("bookname")
	if strings.TrimSpace(bookname) == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "请重新输入书名！")
		return
	} else {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "%v", service.Findbooks(bookname))
		return
	}
}
