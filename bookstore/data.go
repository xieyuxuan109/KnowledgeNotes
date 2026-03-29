package main

import (
	"context"
	"errors"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	defaultShelfSize = 5
)

// "gorm.io/driver/sqlite"
// 使用Gorm
func NewDB(dsn string) (*gorm.DB, error) {
	// 	 dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) //go get -u gorm.io/driver/sqlite
	if err != nil {
		panic(err)
	}
	// 自动建表
	db.AutoMigrate(&Shelf{}, &Book{})
	return db, nil
}

// 定义模型
// 书架
type Shelf struct {
	ID       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}
type Book struct {
	ID       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}

// 数据库操作
type bookstore struct {
	db *gorm.DB
}

// CreateShelf创建书架
func (b *bookstore) CreateShelf(ctx context.Context, data Shelf) (*Shelf, error) {
	if len(data.Theme) <= 0 {
		return nil, errors.New("invalid theme")
	}
	if data.Size <= 0 {
		data.Size = defaultShelfSize
	}
	v := Shelf{Theme: data.Theme, Size: data.Size, CreateAt: time.Now(), UpdateAt: time.Now()}
	err := b.db.WithContext(ctx).Create(&v).Error
	return &v, err
}

// GetShelf获取书架
func (b *bookstore) GetShelf(ctx context.Context, id int64) (*Shelf, error) {
	v := Shelf{}
	err := b.db.WithContext(ctx).First(&v, id).Error
	return &v, err
}

// ListShelves书架列表
func (b *bookstore) ListShelves(ctx context.Context) ([]*Shelf, error) {
	v := []*Shelf{}
	err := b.db.WithContext(ctx).Find(&v).Error
	return v, err
}

// DeleteShelves删除暑假
func (b *bookstore) DeleteShelves(ctx context.Context, id int64) error {
	return b.db.WithContext(ctx).Delete(&Shelf{}, id).Error
}
