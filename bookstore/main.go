package main

import "fmt"

//BookStore
func main() {
	//链连接数据库
	db, err := NewDB("bookstore.db")
	if err != nil {
		fmt.Printf("connect to db failed,err:%v\n")
		return
	}
	//db操作
}
