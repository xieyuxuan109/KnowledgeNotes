package main

import (
	"demo2/api"
	"fmt"

	"google.golang.org/protobuf/proto"
)

// one of demo示例
func oneofDemo() {
	//client
	// req1 := &api.NoticeReaderRequest{
	// 	Msg: "这是谢宇轩的博客",
	// 	NoticeWay: &api.NoticeReaderRequest_Email{
	// 		Email: "123@xxx.vom",
	// 	},
	// }
	req2 := &api.NoticeReaderRequest{
		Msg: "这是谢宇轩的博客",
		NoticeWay: &api.NoticeReaderRequest_Phone{
			Phone: "1234567",
		},
	}
	//server 类型断言
	req := req2
	switch req.NoticeWay.(type) {
	case *api.NoticeReaderRequest_Email: // ← 类型判断
		fmt.Println("邮件方式")
	case *api.NoticeReaderRequest_Phone: // ← 类型判断
		fmt.Println("电话方式")
	}
}

//	func wrapValueDemo() {
//		//client
//		book := api.Book{
//			Title: "跟谢宇轩学习go语言",
//			Price: &wrapperspb.Int64Value{Value: 9000},
//			Memo:  &wrapperspb.StringValue{Value: "学就完事了"},
//		}
//		if book.GetPrice() == nil {
//			//price没有被赋值
//		} else {
//			//赋值了
//			fmt.Println(book.GetPrice().GetValue())
//		}
//		if book.GetMemo() == nil {
//			//memo没有被赋值
//		} else {
//			//赋值了
//			fmt.Println(book.GetMemo().GetValue())
//		}
//	}
func optionalDemo() {
	book := &api.Book{
		Title: "跟着谢宇轩学习go语言",
		Price: proto.Int64(9000),
	}
	if book.Price == nil { //不能使用GetPrice()，因为如果为nil会返回0
		//没有赋值
	} else {
		fmt.Println("已赋值", book.GetPrice())
	}
}
func main() {
	oneofDemo()
	// wrapValueDemo()
	optionalDemo()
}
