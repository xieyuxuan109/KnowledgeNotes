## 自定义错误处理
```json
{
    "code": 10010
    "msg": "内部错误"
}
```
### api下面简历errorx文件夹
### 建立baseerror.go 自定义错误
```
const (
    defaultErrCode = 1001
    RPCErrCode = 1002
    MysqlErrCode = 1003
    RedisErrCode = 1004
)
// CodeError 自定义错误
type CodeError struct{
    Code int `json:"code"`
    Msg string `json:"msg"`
}
// CodeErrResponse 自定义错误响应
type CodeErrResponse struct{
    Code int `json:"code"`
    Msg string `json:"msg"`
}
// NewCodeError 返回自定义错误
func NewCodeError(code int,msg string)CodeError{
    return CodeError{
        Code: code,
        Msg: msg,
    }
}
// NewDefaultErr返回默认自定义错误
func NewDefaultErr(msg string)Error{
    return CodeError{
        Code: defaultErrCode,
        Msg: msg, 
    }
}
// Error CodeError实现error接口
func (e CodeError)Error()string{
    return e.Msg
} 
// Data 返回自定义类型的错误响应
func (e *CodeError)Data()*CodeErrorResponse{
    return &CodeErrorResponse{
        Code: e.Code,
        Msg: e.msg,
    }
}
```
### 使用自定义错误
```
return nil,errorx.NewCodeError(errorx.RPCErrorCode,"内部错误")
```
### 告诉go-zero框架处理自定义错误
```
// 自定义错误 钩子函数 main函数中
httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
    switch e := err.(type) {
    case *errorx.CodeError:
        return http.StatusOK, e.Data()
    default:
        return http.StatusInternalServerError, nil
    }
})
```
```
//使用
// 注册自定义错误的处理方法
httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, any) {
    switch e := err.(type) {
    case errorx.CodeError: // 自定义错误
        return http.StatusOK, e.Data()
    default:
        return http.StatusInternalServerError, nil
    }
})
```