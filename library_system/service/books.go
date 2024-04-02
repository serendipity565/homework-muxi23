package service

import (
	"fmt"
	"test/db"
)

type Books struct {
	Id        int    `json:"id"`
	Book_name string `json:"book_name"`
	Author    string `json:"author"`
}

// 查找数据
func Findbooks(bookname string) string {
	var book Books
	sqlstr := "SELECT * FROM books WHERE book_name LIKE ?"
	rows, err := db.Db.Query(sqlstr, "%"+bookname+"%")
	if err != nil {
		return fmt.Sprintf("查找失败: %v", err)
	}
	defer rows.Close()
	var s string
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Book_name, &book.Author)
		if err != nil {
			return fmt.Sprintf("无法输出: %v", err)
		}
		s += fmt.Sprintf("id:%d name:%s author:%s\n", book.Id, book.Book_name, book.Author)
	}
	return s
}

// 展示部分书籍
func Listbooks() string {
	var book Books
	sqlstr := "SELECT * FROM books"
	rows, err := db.Db.Query(sqlstr)
	if err != nil {
		return fmt.Sprintf("显示失败: %v", err)
	}
	defer rows.Close()
	var s string
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Book_name, &book.Author)
		if err != nil {
			return fmt.Sprintf("显示失败: %v", err)
		}
		s += fmt.Sprintf("id:%d name:%s author:%s\n", book.Id, book.Book_name, book.Author)
	}
	return s
}
