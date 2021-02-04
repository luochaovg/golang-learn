# gRPC_demo
some code snippets about gRPC
- https://www.liwenzhou.com/posts/Go/gRPC/


# 1 . install protoc by mac  
- 安装protoc   brew install protobuf

# 2 .  protoc-gen-go 
- 编译插件protoc-gen-go将会安装到$GOBIN，默认是$GOPATH/bin，它必须在你的$PATH中以便协议编译器protoc能够找到它。
```shell script 
 go get -u github.com/golang/protobuf/protoc-gen-go
```


# 3. helloword.proto 生成go语言源代码
- 在根目录 gRPC_demo 下执行
```shell script
protoc -I helloworld/ helloworld/pb/helloworld.proto --go_out=plugins=grpc:helloworld
```