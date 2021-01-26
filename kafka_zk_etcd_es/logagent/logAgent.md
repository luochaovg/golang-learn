![Image text](../pic/1611543041483.jpg)

### LogAgent 的构造流程
1.读日志 - tailf 第三方库
2.写日志sarama , sarama v1.20之后版本加入zstd压缩算法， 需要用到cgo,在windows平台编译时报错（所以在windows平台使用v1.19版本的sarama）
```shell script
go get github.com/Shopify/sarama
```