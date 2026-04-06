## 生成命令
```
goctl api go -api user.api -dir . -style=goZero
```
```
goctl是go-zero框架的代码生成工具，“api”表示处理api文件，“go”是生成Go代码。“-api”指定源api文件，“-dir”指定输出目录，“-style”指定命名风格
```
## DSL 变更后重新生成
### 每次修改 user.api 后，重新生成但不会覆盖你的逻辑文件：
```
goctl api go -api user.api -dir .
```
## 文件结构
```
syntax = "v1"         // 必须的版本声明
info (                // 可选元数据块
    title: "用户 API"
    version: "1.0"
)
import "shared.api"   // 导入其他 .api 文件
type (...)             // 类型定义
service name-api {     // 服务块
    @server (...)
    @handler HandlerName
    method /path (RequestType) returns (ResponseType)
}
```
## 类型定义
```
类型直接映射为 Go 结构体，使用标准的 Go 结构体标签：
type (
    //body
    LoginReq {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    // 路径参数：/user/:id
    UserReq {
        Id int64 `path:"id"`
    }
    // 查询参数：/search?keyword=foo&page=1
    SearchReq {
        Keyword string `form:"keyword"`
        Page    int    `form:"page,default=1"`
    }
)
```
## 标签类型对应的字段
|标签	|来源|	示例|
|---|---|---|
|json|	请求/响应体（POST/PUT）|	json:"username"|
path|	URL 路径参数|	path:"id"|
|form|	URL 查询字符串（GET）|	form:"page,default=1"|
header|	HTTP 请求头	|header:"Authorization"|
## 可选字段
```
type SearchReq {
    Keyword string `form:"keyword"`
    Page    int    `form:"page,optional"`
    Size    int    `form:"size,default=20"`
}
```
## 示例
```
syntax = "v1"
info (
	title:   "mall"
	desc:    "go-zero学习项目"
	author:  "xieyuxuan"
	email:   "@xieyuxuan.cn"
	version: "1.0"
)
type SignupRequest {
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   int    `json:"gender,options=0|1|2,default=0"`
}

type SignupResponse {
	Message string `json:"message"`
}

@server (
	prefix: api
)
service user-api {
	@handler SignupHandler
	post /user/signup (SignupRequest) returns (SignupResponse)
}
```