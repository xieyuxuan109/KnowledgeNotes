# Bookstore
gRPC&gRPC-Gateaway小练习
## bookstore介绍
书店里面有很多暑假。每个暑假有自己的主题和大小，分别表示摆放的图书图书的主题和数量
## 要点
1.数据库
2.proto
3.业务逻辑
## proto文件
pb/bookstore.proto
## 依赖安装
```
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
## 生成代码
```shell
protoc -I=pb \
  --go_out=pb --go_opt=paths=source_relative \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
  bookstore.proto
```