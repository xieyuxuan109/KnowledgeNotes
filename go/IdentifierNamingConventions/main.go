package main

func main() {
	// 一、基本规则
	// 1. 组成字符
	// 必须以字母或下划线 _ 开头
	// 后续可以是字母、数字、下划线
	// 区分大小写
	// 不能是 Go 关键字（25个）

	// 合法标识符
	// var name string
	// var _age int
	// var π float64  // Unicode字母允许
	// var 变量1 string  // 中文允许
	// var firstName string

	// 2. 可见性规则（重要）
	// 首字母大小写决定访问权限：
	// 首字母大写：可导出（public），其他包可访问
	// 首字母小写：不可导出（private），仅本包内可访问

	// var privateVar int    // 仅本包可见
	// var PublicVar int     // 其他包可访问
	// func privateFunc() {} // 仅本包可见
	// func PublicFunc() {}  // 其他包可访问

	// 二、命名约定（惯例）
	// 1. 驼峰命名法
	// 大驼峰（PascalCase）：类型、接口、公开成员
	// 小驼峰（camelCase）：变量、函数、非公开成员

	// // 类型和接口
	// type UserInfo struct {}
	// type ReadWriter interface {}
	// // 变量和函数
	// var maxValue int
	// func calculateTotal() {}
	// // 常量（通常全大写）
	// const MaxSize = 100
	// const DefaultTimeout = 5

	// 2. 缩写处理
	// 全大写或全小写，保持一致性
	// var userID int      // 推荐
	// var userId int      // 不推荐
	// var httpURL string  // 推荐

	// var HttpUrl string  // 不推荐
	// 3. 接口命名
	// 单个方法接口：方法名 + "er"
	// 多个方法接口：描述性名称
	// type Reader interface {
	//     Read(p []byte) (n int, err error)
	// }

	// type ReadCloser interface {
	//     Reader
	//     Closer
	// }

	// 4. 特殊命名
	// 布尔变量
	// var isReady bool
	// var hasError bool
	// var canWrite bool
	// 测试文件
	// // 测试文件以 _test.go 结尾
	// // 测试函数以 Test 开头
	// func TestAdd(t *testing.T) {}

	// 5. 包名规范
	// 简短、小写、单数名词
	// 避免下划线和驼峰
	// 避免与标准库冲突

	// // 推荐
	// package model
	// package utils
	// package handler

	// // 不推荐
	// package MyPackage
	// package my_package
	// package models  // 除非确实是多个模型

	// 6. 避免使用
	// // 避免使用内置标识符
	// var len int    // 覆盖了内置len函数
	// var copy []int // 覆盖了内置copy函数

	// // 避免单字母变量（除了常见约定）
	// for i := 0; i < 10; i++ {}      // i, j, k 循环变量可以
	// for _, v := range items {}      // v 值变量可以
	// func (u User) Name() string {}  // u 接收者可以

	// // 上下文不清晰时应使用完整单词
	// var u User       // 不推荐（除非作用域很小）
	// var user User    // 推荐

	// 三、常见实践示例
	// // 结构体
	// type UserProfile struct {
	//     ID        int
	//     UserName  string
	//     Email     string
	//     isActive  bool  // 私有字段
	// }

	// // 方法
	// func (u *UserProfile) GetEmail() string {
	//     return u.Email
	// }

	// func (u *UserProfile) setActive(status bool) {
	//     u.isActive = status
	// }

	// // 函数
	// func CalculateTotalPrice(items []Item) float64 {
	//     // ...
	// }

	// // 常量
	// const (
	//
	//	DefaultPageSize = 20
	//	MaxRetryCount   = 3
	//
	// )
}
