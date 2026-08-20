## 根据需要组合请求body
```
targetUrl := "https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1"
payload := strings.NewReader("{\"name\":\"xxx\"}")
req, _ := http.NewRequest("POST", targetUrl, payload)
req.Header.Add("Content-Type", "application/json")
response, err := http.DefaultClient.Do(req)
```
## 创建请求对象
```
req, _ := http.NewRequest("DELETE", targetUrl, nil)
```
```
执行动作：构造（创建） 一个 HTTP 请求对象（*http.Request），但并不真正发送网络数据。

入参详解：
"DELETE"：指定 HTTP 方法（此处为删除操作）。
targetUrl：请求的目标地址（例如 "http://example.com/api/resource/1"）。
nil：请求体（Body）。对于 DELETE、GET 等请求，通常不携带数据，故传 nil。

返回值：
req：新创建的请求对象，包含了方法、URL、Header 等元数据，可以后续修改（如添加认证头）。
本质：这一步只是在内存中构建了一个请求结构体，没有任何网络 I/O 发生。
```
## 根据需要增加header
```
req.Header.Add("Authorization", "xxxx") //header主要用来认证鉴权 以及一些apikey
```
## 执行请求体
```
response, err := http.DefaultClient.Do(req)
```