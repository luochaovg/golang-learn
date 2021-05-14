> ETCD 专题
- https://www.bilibili.com/video/BV1sJ41127Rm?from=search&seid=352779058055681455
- 文档： https://www.kubernetes.org.cn/6226.html
- https://www.liwenzhou.com/posts/Go/go_etcd/
- [参考](https://www.infoq.cn/article/etcd-interpretation-application-scenario-implement-principle/)

###ETCD背后的Raft一致性算法原理
- https://www.jianshu.com/p/5aed73b288f7

### watch底层实现原理
- https://www.lixueduan.com/post/etcd/05-watch/

### etcdctl一些操作命令 ， 去了解！

### 图示
- ![Image text](./pic/WX20210126-194402@2x.png)
- ![Image text](./pic/WX20210126-195413@2x.png)
- ![Image text](./pic/WX20210126-195941@2x.png)
- ![Image text](./pic/WX20210126-195954@2x.png)
- ![Image text](./pic/WX20210126-200107@2x.png)
- ![Image text](./pic/WX20210127-101133@2x.png)
- ![Image text](./pic/WX20210127-101306@2x.png)
- ![Image text](./pic/WX20210127-101400@2x.png)
- ![Image text](./pic/WX20210127-101503@2x.png)
- ![Image text](./pic/WX20210127-101540@2x.png)
- ![Image text](./pic/WX20210204-120329@2x.png)

#### ETCD 数据
```text
etcd 中所有的数据都存储在一个 b+tree 中（灰色），该 b+tree 保存在磁盘中
一个数据有多个版本
通过定期的Compaction来清理历史数据

etcd lease(租约) 的概念
```


#### 分布式锁
- https://zhuanlan.zhihu.com/p/42056183