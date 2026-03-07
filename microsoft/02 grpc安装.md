### 安装grpc
```
grpc go install google.golang.org/grpc
安装编译器protocal buffer
https://github.com/protocolbuffers/protobuf/releases
提取出后写入环境变量path（bin目录）
验证protoc --version
```
### 安装插件
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest//生成go代码
验证
PS C:\workspace\KnowledgeNotes> protoc-gen-go --version
protoc-gen-go.exe v1.36.11
该插件会根据.proto文件生成一个后缀为.pb.go的文件，包含所有.proto文件中定义的类型机器序列化方法
```
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest//生成grpc代码
验证
PS C:\workspace\KnowledgeNotes> protoc-gen-go-grpc --version
protoc-gen-go-grpc 1.6.1
该插件会生成grpc相关代码
```
```
vscode上安装Protobuf VSC
```