## goctl 安装
```
goctl工具
https://go-zero.dev/cn/docs/goctl/goctl/

安装
go intall github.com/zeromicro/go-zero/tools/goctl@latest
```bash
》goctl--version
goctl version 1.4.4 darwin/amd64
》goctl-v
goctl version 1.4.4darwin/amd64
```
```
安装依赖grpc相关
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
```
protoc-gen-go --version
protoc-gen-go-grpc --version
protoc-gen-grpc-gateway --version
protoc-gen-openapiv2 --version
```