##### RPC
- 为什么R(远程)?
- RPC与http2区别？
- 传输 序列化/反序列化
- 网络IO

```text

```

##### 微服务的定义和优缺点
- [参考](https://zhuanlan.zhihu.com/p/46459720)
```text

```

##### GRPC

```text
1,gRPC 的应用场景
总的来说，gRPC 可以高效地连接数据中心内部和跨数据中心的服务，并为负载平衡、跟踪、健康检查和身份验证提供可插拔的支持。它也适用于最后一英里的分布式计算连接设备，移动应用程序和浏览器的后端服务。主要使用场景有：

在微服务风格的体系结构中有效连接多语言服务；
连接移动设备、浏览器客户端和后端服务；
开发与云服务器通信的移动客户端；
构建高可扩展、低延迟的分布式系统和微服务；
设计一个新的协议，需要准确，高效和语言独立；
分层设计，以支持扩展，例如身份验证、负载平衡、日志记录和监控等；

2,gRPC 的主要优点
相比 RESTful 服务或其他 RPC 框架，gRPC 有哪些突出的优点？

1）极速
gRPC 使用 protocol buffers（简称 protobuf）作为消息编码格式。protobuf 的特点是语言中立、平台无关、高可扩展，它可以序列化和反序列化结构化数据。
与采用文本格式的 JSON 相比，采用二进制格式的 protobuf 在速度上可以达到前者的 5 倍！

2）使用 HTTP/2
gRPC 使用 HTTP/2 作为传输协议。我们来看看 HTTP/2 与 HTTP/1.X 相比有何优势。
二进制传输
多路复用，即可以在一个 TCP 连接上并行发送多个数据请求
双向同时通信，即可以同时发送客户端请求和服务器端响应
头部压缩技术，降低网络使用率

3）多语言支持、社区活跃
目前，gRPC 支持 11 种语言，GitHub 项目总星数近 6 万，其中 Go 实现和J ava 实现的星数最多。

Go/Java/C#/C++/Dart/Kotlin/JVM/Node/Objective-C/PHP/Python/Ruby
此外，gRPC 还支持多平台，比如 Web、Android，以及 Flutter、iOS 等。


```

- 双向认证下rpc-gatway提供grpc/http接口
- https://github.com/grpc-ecosystem/grpc-gateway

http 请求流程  浏览器->http>grpc

#### proto 
```text
-- Prod.proto
syntax="proto3";
package services;
import "google/api/annotaions/proto";
import "Models.proto"

enum ProdAreas {
    A=0;
    B=1;
}
message ProdRequest {
    int32 prod_id=1;
    ProAreas prod_area=2;
}
message ProdResponse {
    int32 prod_stock=1
}
message QuerySize{
    int32 size=1;
}
message ProdResponseList {
    repeated ProdResponse prodres=1;
}

service ProdService {
    rpc GetProdList(QuerySize) returns (ProdResponseList) {}
    
    rpc NewOrder(OrderRequest) returns (OrderResponse) {
        option (google.api.http) = {
            post:"/v1/orders"
            body:"order_main"
        }
    }
}

----------Models.proto
syntax="proto3"
package services;
import "google/protobuf/timestamp.proto";

message ProdModel {
    int32 prod_id=1;
    string prod_name=2;
    float prod_price=3;
}

message OrderRequest {
    OrderMain order_main=1;
}

message OrderMain {
    int32 order_id=1;
    google.protobuf.Timestamp order_time=2;
}

message OrderResponse {
    string status=1;
    string message=2;
}
```
- 生成代码
```shell 
cd pb && protoc --go_out=plugins=grpc:../services Prod.proto
cd pb && protoc --go_out=plugins=grpc:../services Models.proto

cd pb && protoc --grpc-gateway_out=logtostderr=true:../services Orders.proto
```

#### 流模式
- 传输数据大，