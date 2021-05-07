> KAFKA 专题
- 视频地址： https://www.bilibili.com/video/BV1a4411B7V9?from=search&seid=5013796797653472442
- 线上资料： https://my.oschina.net/jallenkwong/blog/4449224

> 发布/订阅
```text
多个消费者
消费者的两种模式： 
- 消费者主动拉去数据 
- 队列主动推送给消费者（基于listen监听key）
```

4.生产者往kafka发送数据的模式（3种）
    ![Image text](./pic/WX20210125-111323@2x.png)
    - 0：把数据发送给leader就成功 ，效率最高，安全性最低
    - 1：把数据发送给leader,等待leader回ACK
    - all(-1)： 把数据发送给leader,确保follower从leader拉去数据回复ack给leader，leader在回复ack，安全性高

- 等于all 的时候 
也会丢失数据 原因是ISR 只有一个leader,这种情况很少
也会重复数据，follower 同步完成时候，leader挂了， product没有收到ack, 此时其中一个follower
变为leader， product重新发送数据给leader, 就会出现重复

### 如何数据一致性
```text
HW:high watermark 所有副本中最小的LEO, 指消费者能见到最大的offset， ISR队列中最小的LEO
LEO: log end offset, 每个分区，  每个副本的最后一个offset

注：这只能保证副本之间数据一致性， 并不能保证数据不丢失或者不重复
```

### Exactly once 语以
```text
ack = 0 , 会丢数据， 不重复
ack = 1 , 会丢数据
ack = -1 , isr = 1 会丢数据， isr >= 2 保证数据不丢失， 可能会重复


幂等性 + at least once = exactly once
幂等性 0.1版本后， 解决数据重复问题 
启用幂等性 product 参数中 enable.idompotence = true , ack 默认为 -1  ，只能保证单
会话 幂等， 没办法跨分区／会话的做到幂等
```
> 架构
