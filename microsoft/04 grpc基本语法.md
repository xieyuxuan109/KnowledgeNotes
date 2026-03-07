## 一、Protocol Buffers 基础
```
1.1 语法版本
protobuf
// 指定使用 proto3 语法
syntax = "proto3";
// 如果不指定，默认使用 proto2
1.2 基本结构
protobuf
// 包名定义
package hello;
// 导入其他 proto 文件
import "google/protobuf/timestamp.proto";
import "myproject/other_protos.proto";
// 定义服务
service Greeter {
  // 方法定义
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}
```
## 二、消息类型定义
```
2.1 基本消息格式
protobuf
message SearchRequest {
  string query = 1;      // 字段类型 字段名 = 字段编号
  int32 page_number = 2;
  int32 result_per_page = 3;
}
2.2 字段类型
类型	说明
double	双精度浮点数
float	单精度浮点数
int32	变长编码，负数效率低
int64	变长编码，负数效率低
uint32	无符号整数
uint64	无符号长整数
sint32	有符号整数，负数效率高
sint64	有符号长整数，负数效率高
fixed32	4字节，值大于2^28时效率高
fixed64	8字节，值大于2^56时效率高
sfixed32	有符号4字节
sfixed64	有符号8字节
bool	布尔值
string	字符串（UTF-8编码）
bytes	字节数组
2.3 字段修饰符
protobuf
message Example {
  string normal_field = 1;              // 默认字段
  
  repeated string list_field = 2;        // 重复字段（数组）
  
  map<string, string> map_field = 3;     // Map类型
  
  reserved 4, 5, 10 to 12;                // 保留字段编号
  reserved "foo", "bar";                   // 保留字段名
  
  oneof test_oneof {                       // Oneof 字段
    string name = 6;
    int64 id = 7;
  }
}
```
## 三、枚举类型
```
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
  enum Corpus {
    option allow_alias = true;//允许设置别名
    UNIVERSAL = 0;//从零开始
    WEB = 1;
    IMAGES = 2;
    LOCAL = 3;
    NEWS = 4;
    PRODUCTS = 5;
    VIDEO = 6;
  }
  Corpus corpus = 4;
}
```
## 四、服务定义
```
4.1 基本服务
protobuf
service Greeter {
  // 一元 RPC
  rpc SayHello (HelloRequest) returns (HelloReply) {}
  
  // 服务器流式 RPC
  rpc ListFeatures (Rectangle) returns (stream Feature) {}
  
  // 客户端流式 RPC
  rpc RecordRoute (stream Point) returns (RouteSummary) {}
  
  // 双向流式 RPC
  rpc RouteChat (stream RouteNote) returns (stream RouteNote) {}
}
4.2 流式关键字
stream - 表示流式数据传输
```
## 五、高级特性
```
5.1 嵌套类型
protobuf
message Outer {
  message MiddleAA {
    message Inner {
      int64 ival = 1;
      bool booly = 2;
    }
  }
  
  message MiddleBB {
    message Inner {
      int32 ival = 1;
      bool booly = 2;
    }
  }
  
  MiddleAA.Inner aa = 1;
  MiddleBB.Inner bb = 2;
}
5.2 Any 类型
protobuf
import "google/protobuf/any.proto";

message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
5.3 选项（Options）
protobuf
// 文件级别选项
option java_package = "com.example.foo";
option java_multiple_files = true;
option java_outer_classname = "Ponycopter";
option optimize_for = CODE_SIZE;  // SPEED, CODE_SIZE, LITE_RUNTIME

// 消息级别选项
message MyMessage {
  option message_set_wire_format = false;
  option no_standard_descriptor_accessor = false;
  option deprecated = false;
  
  int32 field1 = 1;
  int32 field2 = 2 [deprecated=true];  // 字段级别选项
}

// 服务级别选项
service MyService {
  option deprecated = false;
  
  rpc MyMethod (MyRequest) returns (MyResponse) {
    option idempotency_level = NO_SIDEEFFECTS;  // 方法级别选项
  }
}
```
## 六、常用导入类型
```
protobuf
// 时间戳
import "google/protobuf/timestamp.proto";
google.protobuf.Timestamp timestamp = 1;

// 持续时间
import "google/protobuf/duration.proto";
google.protobuf.Duration duration = 2;

// 空消息
import "google/protobuf/empty.proto";
google.protobuf.Empty empty = 3;

// 结构体
import "google/protobuf/struct.proto";
google.protobuf.Struct struct = 4;

// 包装类型
import "google/protobuf/wrappers.proto";
google.protobuf.StringValue string_value = 5;
google.protobuf.Int32Value int32_value = 6;
```
## 七、完整示例
```
protobuf
syntax = "proto3";

package helloworld;

import "google/protobuf/timestamp.proto";
import "google/protobuf/wrappers.proto";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
  rpc SayHelloStream (stream HelloRequest) returns (stream HelloReply) {}
}

message HelloRequest {
  string name = 1;
  google.protobuf.Int32Value age = 2;  // 可选字段
  repeated string hobbies = 3;          // 数组
  map<string, string> attributes = 4;   // Map类型
  
  enum Gender {
    UNSPECIFIED = 0;
    MALE = 1;
    FEMALE = 2;
  }
  
  Gender gender = 5;
  google.protobuf.Timestamp reg_time = 6;
}

message HelloReply {
  string message = 1;
  int32 code = 2;
}
```
## 八、编译命令
```
bash
# 生成 gRPC 代码
protoc --proto_path=./proto \
       --go_out=./go \
       --go-grpc_out=./go \
       ./proto/*.proto

# 参数说明
--proto_path: 指定 proto 文件搜索路径
--go_out: 生成 Go 代码的输出目录
--go-grpc_out: 生成 gRPC Go 代码的输出目录
```
## 九、最佳实践
```
9.1 字段编号规则
1-15：使用1字节编码，用于频繁出现的字段

16-2047：使用2字节编码，用于不常出现的字段

19000-19999：保留编号，不可使用

最大可到 2^29-1
9.2 命名规范
消息名: 驼峰命名法（CamelCase）

字段名: 下划线命名法（underscore_separated）

枚举名: 驼峰命名法

枚举值: 全大写加下划线（SCREAMING_SNAKE_CASE）

服务名: 驼峰命名法

方法名: 驼峰命名法
```
```
repeated 想要一个字段接受多个值
reserved 保留字段
```