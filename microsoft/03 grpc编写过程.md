# 第一个grpc示例
```
hello world
```
## 三个步骤
```
1.编写protobuf文件
2.生成代码
3.编写业务逻辑代码
```
```
生成代码
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
pb/hello.proto

或者
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/hello.proto//再hello_server那一级生成
```
```
go语言代码
protoc --proto_path=pb --go_out=pb --go_opt=paths=source_relative book/price.proto book/book.proto 
grpc代码

```