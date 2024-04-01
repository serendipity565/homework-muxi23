package main

import (
	"log"
	"net/http"
	"test/controller"
)

func main() {
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/login", controller.Login)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
