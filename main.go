package main

import (
	"log"
	"net/http"
)

// index 创建一个响应函数
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("heelo world"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9999",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
