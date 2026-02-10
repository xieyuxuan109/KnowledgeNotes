package utils

import "testing"

// 注意运行该文件不用go run而是用go test -v显示详细信息
// 编写测试文件时候 必须以xxx_test.go结尾，该文件必须包含TestXxx的函数
// 一般将该文件放在与被测试的包相同的包中，该文件将被排除在正常程序包之外，但在运行go test命令时候将被包含
// 一个测试文件可以有多个测试用例函数
func TestSub(t *testing.T) { //Test后面第一个字母不能为a-z字母，可以为数字和大写字母
	res := Sub(21, 1)
	if res != 20 {
		t.Fatalf("Sub(21,1)执行错误，期望值%v,实际值%v\n", 55, res) //输出错误日志并终止程序
	}
	t.Logf("Sub执行正确...") //输出日志
}
