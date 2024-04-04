package controller

import (
	"fmt"
	"net/http"
	"strings"
	"test/service"
)

func Find(w http.ResponseWriter, r *http.Request) {
	// 展示书籍
	fmt.Fprintf(w, "%v", service.Listbooks())

	// 查找书籍
	bookname := r.PostFormValue("bookname")
	if strings.TrimSpace(bookname) == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "请重新输入书名！")
		return
	}

	book := service.Findbooks(bookname)
	if strings.TrimSpace(book) == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "找不到书籍：%s", bookname)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", book)
}
