```
这使得 HTTP REST 到 gRPC 的映射如下所示:
```

|HTTP	|gRPC|
|---|---|
|GET /v1/messages/123456	|GetMessage(name: "messages/123456")|
```
如果没有 HTTP 请求正文，则请求消息中没有被路径模板绑定的任何字段都会自动成为 HTTP 查询参数
```
```
service Messaging {
    rpc GetMessage(GetMessageRequest) returns (Message) {
    option (google.api.http) = {
        get:"/v1/messages/{message_id}"
    };
    }
}
message GetMessageRequest {
    message SubMessage {
    string subfield = 1;
    }
    string message_id = 1; // Mapped to URL path.
    int64 revision = 2;    // Mapped to URL query parameter `revision`.
    SubMessage sub = 3;    // Mapped to URL query parameter `sub.subfield`.
}
```
```
这使得 HTTP JSON 到 RPC 的映射如下所示
```
|HTTP|	gRPC|
|---|---|
|GET /v1/messages/123456?revision=2&sub.subfield=foo|	GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield:"foo"))|
```
service Messaging {
    rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
    option (google.api.http) = {
        patch: "/v1/messages/{message_id}"
        body: "message"
    };
    }
}
message UpdateMessageRequest {
    string message_id = 1; // mapped to the URL
    Message message = 2;   // mapped to the body
}
```
|HTTP|	gRPC|
|---|---|
|PATCH /v1/messages/123456 { "text": "Hi!" }	|UpdateMessage(message_id:"123456" message { text: "Hi!" })|
```
特殊名称 * 可用于主体映射来定义不受路径模板绑定的每个字段都应映射到请求正文。更新方法可以替换为以下定义
```
```
service Messaging {
    rpc UpdateMessage(Message) returns (Message) {
    option (google.api.http) = {
        patch: "/v1/messages/{message_id}"
        body: "*"
    };
    }
}
message Message {
    string message_id = 1;
    string text = 2;
}
```
```
此时 HTTP JSON 到 RPC 的映射:
```
|HTTP|	gRPC|
|---|---|
|PATCH /v1/messages/123456 { "text": "Hi!" }	|UpdateMessage(message_id:"123456" text: "Hi!")|
```
可以使用 additional_bindings 选项为一个 RPC 定义多个 HTTP 方法。例如：
```
```
service Messaging {
    rpc GetMessage(GetMessageRequest) returns (Message) {
    option (google.api.http) = {
        get: "/v1/messages/{message_id}"
        additional_bindings {
        get: "/v1/users/{user_id}/messages/{message_id}"
        }
    };
    }
}
message GetMessageRequest {
    string message_id = 1;
    string user_id = 2;
}
```
```
这启用了以下两种可选的 HTTP JSON 到 RPC 映射：
```
|HTTP|	gRPC|
|---|---|
|GET /v1/messages/123456|	GetMessage(message_id: "123456")|
|GET /v1/users/me/messages/123456	|GetMessage(user_id: "me" message_id:"123456")|
```
总结
叶请求字段（请求消息中的递归扩展嵌套消息）分为三类
由路径模板引用的字段。它们通过 URL 路径传递。
[HttpRule.body][google.api.HttpRule.body] 引用的字段。 它们通过 HTTP 请求正文传递。
所有其他字段都是通过 URL 查询参数传递的，参数名称是请求消息中的字段路径。 一个重复的字段可以表示为同名的多个查询参数。
如果 [HttpRule.body][google.api.HttpRule.body] 为“*”，则没有 URL 查询参数，所有字段都通过 URL 路径和 HTTP 请求正文传递。
如果 [HttpRule.body][google.api.HttpRule.body] 省略，则没有 HTTP 请求正文，所有字段都通过 URL 路径和 URL 查询参数传递。
```
