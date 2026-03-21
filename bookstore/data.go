package main

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// "gorm.io/driver/sqlite"
// 使用Gorm
func NewDB(dsn string) (*gorm.DB, error) {
	// 	 dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) //go get -u gorm.io/driver/sqlite
	if err != nil {
		panic("连接数据库失败")
	}
	// 自动建表
	db.AutoMigrate(&Shelf{}, &Book{})
	return db, nil
}

// 定义模型
// 书架
type Shelf struct {
	id       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}
type Book struct {
	id       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}
