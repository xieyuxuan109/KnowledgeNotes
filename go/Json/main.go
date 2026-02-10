package main

import (
	"encoding/json"
	"fmt"
)

type x struct {
	A int `json:"123"` //利用反射机制
	//为什么不能让字段小写，因为后面序列化要调用其他包的函数，如果小写将不能使用该字段
}

// 演示json
func main() {
	//json是一种轻量级的数据交换格式，易于人阅读和编写，同时也利于机器解析和生成
	//json是靠键值对来保存数据 key-val
	//扩展性好，灵活
	//https://www.json.cn/网站可以验证以恶json格式数据是否正确

	//json序列化
	//一般不会把基本数据类型序列化，因为没有意义
	//json序列化是指将有key-val结构的数据类型(例如结构体，map，切片)序列化成json字符串的操作
	var Jasper map[string]interface{}
	Jasper = make(map[string]interface{}, 10)
	Jasper["x"] = 1
	Jasper["i"] = 12
	Jasper["12"] = "134rs"
	data, err := json.Marshal(Jasper)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("data=", string(data)) //注意生成的json数据无序

	//反序列化：将json数据反序列化为map，切片等原本的数据类型
	var Jasper1 map[string]interface{}
	Jasper1 = make(map[string]interface{}, 10)
	str := `{"12":"134rs","i":12,"x":1}`  //在实际项目开发过程中，是通过网络传输获取到的或者读取文件获取
	json.Unmarshal([]byte(str), &Jasper1) //前面一个参数传json字符串，后面一个传对应的数据类型的变量的地址
	fmt.Println(Jasper1)
}
