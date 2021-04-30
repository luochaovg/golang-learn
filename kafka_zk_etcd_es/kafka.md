#### kafaka
- 1.kafka 集群的架构
    ![Image text](./pic/WX20210125-161421@2x.png)
    -  1.broker （相当于服务器）
    -  2.topic  （主题）
    -  3.partition:分区，把同一个topic分成多个分区，提高负载
        -  3.1.leader:分区的主节点
        -  3.2.follower:分区的从节点（leader 副本）
    -  4.Consumer Group （消费组）
        - 某一个topic下分区数据，只能被消费者组里面的一个消费者消费
    
- 2.生产者往kafka发送数据的流程（六步）
    ![Image text](./pic/1611544184721.jpg)
    
- 3.kafka选择分区的模式（3种）
    ![Image text](./pic/WX20210125-111141@2x.png)
    -   3.1 指定往哪个分区写
    -   3.2 指定key, kafka 根据key做hash然后决定写哪个分区
    -   3.3 轮询分时

- 4.生产者往kafka发送数据的模式（3种）
    ![Image text](./pic/WX20210125-111323@2x.png)
    - 0：把数据发送给leader就成功 ，效率最高，安全性最低
    - 1：把数据发送给leader,等待leader回ACK
    - all： 把数据发送给leader,确保follower从leader拉去数据回复ack给leader，leader在回复ack，安全性高

- 5.分区存储文件原理
    ![Image text](./pic/WX20210125-162248@2x.png)

- 6.为什么kafka快？ (随机度换为了顺序读， 记录了索引位置)

- 7.消费者组消费数据的原理
    ![Image text](./pic/WX20210125-160134@2x.png)

- 8.日志收集系统架构图
    ![Image text](./pic/1611543041483.jpg)
    
- 9.offset 在0.9版本之前存在zk, 0.9版本之后存在kafka本地，存在磁盘
#### docker-composer
- https://github.com/simplesteph/kafka-stack-docker-compose

#### 参考文献
 https://blog.csdn.net/miss1181248983/article/details/90724870
 https://blog.csdn.net/valada/article/details/80892612