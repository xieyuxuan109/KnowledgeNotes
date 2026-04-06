## goctl参考文件
```
https://go-zero.dev/cn/docs/goctl/goctl/
```
## 安装goctl
```bash
go intall github.com/zeromicro/go-zero/tools/goctl@latest
```
```
顺便在vscode中安装goctl插件
```
## 验证goctl
```bash
》goctl--version
goctl version 1.4.4 darwin/amd64
》goctl-v
goctl version 1.4.4darwin/amd64
```
## 安装grpc相关依赖
```bash
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
## 验证grpc
```bash
protoc-gen-go --version
protoc-gen-go-grpc --version
protoc-gen-grpc-gateway --version
protoc-gen-openapiv2 --version
```
## 安装grpc测试工具
```bash
//grpc ui
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
//确保$GOPATH/bin添加到环境变量里面
```
```bash
//使用 端口换成自己都端口
grpcui -plaintext localhost:8080
```