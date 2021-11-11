

## 面试要点

### 技术能力

> 技术面试的首要考虑点
>
> 比如说事务，对于3年以下的应试者，要知道什么是事务，事务有哪些特性；
>对于3到5年的应试者，要知道如何使用事务，知道事务隔离级别对系统设计的影响范围，知道事务传播机制在代码结构上的影响范围；
>对于5到8年的应试者，要知道数据库的锁机制，知道事务隔离级别在不同数据库中的设计原理，知道事务导致的死锁原理及应对方案
- [参考]（https://www.toutiao.com/i6930164998287589896）
```text
所谓死锁<DeadLock>：是指两个或两个以上的进程在执行过程中,因争夺资源而造成的一种互相等待的现象,若无外力作用，它们都将无法推进下去.
此时称系统处于死锁状态或系统产生了死锁，这些永远在互相等待的进程称为死锁进程。
表级锁不会产生死锁.所以解决死锁主要还是针对于最常用的InnoDB。

show processlist;
-- kill杀死进程id（id列）
SELECT * FROM INFORMATION_SCHEMA.INNODB_TRX;
-- kill杀死进程id（trx_mysql_thread_id列）
```

- 1到3年。考察基础和学习能力
- 3到5年。考察解决问题的能力，能独挡一面。做过哪些东西，最近在做哪些东西，遇到过哪些问题，如何解决
- 5到8年。基础和解决问题的能力并重。工作不在仅仅是填坑，要细化自己的知识体系，知其然知其所以然 

### 学习能力

>关注什么技术，开源项目，个人博客，有无阅读习惯

### 沟通能力

>能够清晰的描述所做的项目，面试回答口齿清晰声音洪亮，看上去性格开朗

### 自我认知

> 最近的项目，项目组成员组成结构，在项目组中是什么角色
>
> 自我评价能够达到什么水平，初级中级还是高级
>
> 优点和缺点，未来的规划






## 部分试题

### 程序基础

- 自动加载、composer
[参考](https://segmentfault.com/a/1190000014948542)
```text
PHP 自动加载函数 __autoload() php中spl_autoload spl_autoload_register() 就是我们上面所说的__autoload调用堆栈
```
- 什么是 CGI？什么是 FastCGI？php-fpm，FastCGI，Nginx 之间是什么关系？
 [参考](https://www.cnblogs.com/jianzhaojing/p/13623509.html)
 [参考](https://blog.csdn.net/weixin_34001430/article/details/92447294)
 [参考](https://blog.51cto.com/13581826/2093473)
```text
CGi全称“通用网关接口”，是一个web服务器和php应用之间用于数据传递的协议，web服务器收到用户请求，就会把请求提交给cgi程序,缺点是每一次web请求都会有启动和退出过程，也就是最为人诟病的fork-and-execute模式，这样一在大规模并发下，就死翘翘了。
FastCGI像是一个常驻(long-live)型的CGI，它可以一直执行着，只要激活后，不会每次都要花费时间去fork一次。
php-fpm 就开发出了管理php-cgi进程的程序，即php-fpm
```
- 如何查看并提升PHP代码的性能？
```text
XHProf XHProf是facebook 开发的一个测试php性能的扩展，本文记录了在PHP应用中使用XHProf对PHP进行性能优化，查找性能瓶颈的方法。
opcache Opcache是一种通过将解析的PHP脚本预编译的字节码（Operate Code）存放在共享内存中来避免每次加载和解析PHP脚本的开销，解析器可以直接从共享内存读取已经缓存的字节码（Operate Code），从而大大提高PHP的执行效率。
https://blog.csdn.net/fanghailiang2016/article/details/104749830
```

- 如何使用和处理异常。生产环境除了try..catch外，
[参考](https://www.cnblogs.com/jiangxiaobo/p/9155225.html)
```text

全局注册
set_exception_handler/set_error_handler/register_shutdown_function
注册错误和异常处理机制有三个PHP函数需要学习

1. register_shutdown_function('Bootstrap\Library\Frame::fatalError');
2. set_error_handler('Bootstrap\Library\Frame::appError');
3. set_exception_handler('Bootstrap\Library\Frame::appException');
```


### 框架

- 各种主流框架的优缺点总结
[参考](https://cloud.tencent.com/developer/article/1543966)
```text
【关于CodeIgniter】

优点：
Code Igniter推崇“简单就是美”这一原则。没有花哨的设计模式、没有华丽的对象结构，一切都是那么简单。
几行代码就能开始运行，再加几行代码就可以进行输出。可谓是“大道至简”的典范。
配置简单，全部的配置使用PHP脚本来配置，执行效率高；具有基本的路由功能，能够进行一定程度的路 由；具有初步的Layout功能，能够制作一定程度的界面外观；
数据库层封装的不错，具有基本的MVC功能. 快速简洁，代码不多，执行性能高，框架简 单，容易上手，学习成本低，文档详细；自带了很多简单好用的library，框架适合小型应用.
缺点：
本身的实现不太理想。内部结构过于混乱，虽然简单易用，但缺乏扩展能力。 把Model层简单的理解为数据库操作. 框架略显简单，只能够满足小型应用，略微不太能够满足中型应用需要.
评价：
总体来说，拿CodeIgniter来完成简单快速的应用还是值得，同时能够构造一定程度的layout，便于模板的复用，数据操作层来说封装的不 错，并且CodeIgniter没有使用很多太复杂的设计模式，执行性能和代码可读性上都不错。至于附加的library 也还不错，简洁高效。

 

Lavarel 框架
优点
Laravel 的设计思想是很先进的，非常适合应用各种开发模式TDD, DDD 和BDD，作为一个框
架，它准备好了一切，composer 是个php 的未来，没有composer，PHP 肯定要走向没落。
laravel 最大的特点和优秀之处就是集合了php 比较新的特性，以及各种各样的设计模式，
Ioc 容器，依赖注入等。
缺点
基于组件式的框架，所以比较臃肿
```
- 门面、依赖注入、IOC模式、设计模式
[参考](https://blog.csdn.net/hizzana/article/details/53212323)

```text
门面：在 Laravel 应用中，Facade就是一个为容器中对象提供访问方式的类。该机制原理由 Facade 类实现。Laravel 自带的 Facade，以及我们创建的自定义门面，都会继承自 Illuminate\Support\Facades\Facade 基类。

依赖注入
依赖注入原理其实就是利用类方法反射，取得参数类型，然后利用容器构造好实例。然后再使用回调函数调起。
注入对象构造函数不能有参数。否则会报错。Missing argument 1
依赖注入故然好，但它必须要由 Router类调起，否则直接用 new 方式是无法实现注入的。所以这就为什么只有 Controller 、Job 类才能用这个特性了。

IOC模式 
https://zhuanlan.zhihu.com/p/48412781
一个类所需要的依赖类是由我们主动实例化后传入类中的
控制反转意思是说将依赖类的控制权交出去，由主动变为被动。

设计模式
单列/工厂
```

##### Swoole的理解和应用
- [参考](https://blog.csdn.net/qq_35619227/article/details/106585596)
- [参考](http://wangzhenkai.com/article/16)
- [进程模型1](https://wiki.swoole.com/#/server/init?id=%e8%bf%9b%e7%a8%8b%e7%ba%bf%e7%a8%8b%e7%bb%93%e6%9e%84%e5%9b%be)
- [进程模型2](https://blog.csdn.net/weixin_40022980/article/details/82705109)
```text
swoole是完全的长驻内存的，长驻内存一个最大的好处就是可以性能加速。
在fpm模式下，我们处理一个请求，通常会有一些空消耗,比如框架共用文件加载，配置文件加载，
那么在swoole中，可以在on workerstart的时候提前一次性把一些必要的文件和配置加载好，
不必每次receive重复加载一遍，这样能提升不小的性能。
```
- 如果自己写过框架，有什么心得体会

### 缓存

- memcached和redis的区别
[参考](https://www.html.cn/qa/other/19404.html)
```text
1、数据操作不同
与Memcached仅支持简单的key-value结构的数据记录不同，Redis支持的数据类型要丰富得多。Memcached基本只支持简单的key-value存储，不支持枚举，不支持持久化和复制等功能。
Redis支持服务器端的数据操作相比Memcached来说，拥有更多的数据结构和并支持更丰富的数据操作，支持list、set、sorted set、hash等众多数据结构，还同时提供了持久化和复制等功能。


2、内存管理机制不同
在Redis中，并不是所有的数据都一直存储在内存中的。这是和Memcached相比一个最大的区别。当物理内存用完时，Redis可以将一些很久没用到的value交换到磁盘
从内存利用率来讲，使用简单的key-value存储的话，Memcached的内存利用率更高。而如果Redis采用hash结构来做key-value存储，由于其组合式的压缩，其内存利用率会高于Memcached。

3.性能不同
Redis只使用单核，而Memcached可以使用多核，所以平均每一个核上Redis在存储小数据时比Memcached性能更高。而在100k以上的数据中，Memcached性能要高于Redis，虽然Redis也在存储大数据的性能上进行了优化，但是比起Memcached，还是稍有逊色。


4、集群管理不同
Memcached本身并不支持分布式
Redis更偏向于在服务器端构建分布式存储。

小结：Redis和Memcached哪个更好？

Redis更多场景是作为Memcached的替代者来使用，当需要除key-value之外的更多数据类型支持或存储的数据不能被剔除时，使用Redis更合适。
如果只做缓存的话，Memcached已经足够应付绝大部分的需求，Redis 的出现只是提供了一个更加好的选择。总的来说，根据使用者自身的需求去选择才是最合适的。


```
- redis是如何对数据进行持久化存储的？常见的数据结构都有什么？redis优化
```text
redis是如何对数据进行持久化存储的？
AOF
将写命令添加到 AOF 文件（Append Only File）的末尾。
    always	每个写命令都同步
    everysec	每秒同步一次
    no	让操作系统来决定何时同步

RDB
将某个时间点的所有数据都存放到硬盘上。

支持list、set、sorted set、hash等众多数据结构

redis优化
批处理/不要存大数据/尽可能使用时间复杂度为O(1)的操作避免使用复杂度为O(N)的操作/Redis提供了Slow Log功能
SLOWLOG GET n命令，可以输出最近n条慢查询日志
缩短键值对的存储长度；
使用 lazy free（延迟删除）特性；
设置键值的过期时间；
禁用长耗时的查询命令；
使用 slowlog 优化耗时命令；
使用 Pipeline 批量操作数据；
避免大量数据同时失效；
客户端使用优化；
限制 Redis 内存大小；
使用物理机而非虚拟机安装 Redis 服务；
检查数据持久化策略；
禁用 THP 特性；
使用分布式架构来增加读写速度。
```

### 数据库

- mysql和mongo的区别
[参考](https://www.jianshu.com/p/56524b50b376)
```text
一、关系型数据库-MySQL
1、在不同的引擎上有不同的存储方式。
2、查询语句是使用传统的sql语句，拥有较为成熟的体系，成熟度很高。
3、开源数据库的份额在不断增加，mysql的份额页在持续增长。
4、缺点就是在海量数据处理的时候效率会显著变慢。

二、非关系型数据库-MongoDB
非关系型数据库(nosql ),属于文档型数据库。先解释一下文档的数据库，即可以存放xml、json、bson类型系那个的数据。这些数据具备自述性，呈现分层的树状数据结构。数据结构由键值(key=>value)对组成。

1、存储方式：虚拟内存+持久化。
2、查询语句：是独特的MongoDB的查询方式。
3、适合场景：事件的记录，内容管理或者博客平台等等。
4、架构特点：可以通过副本集，以及分片来实现高可用。
5、数据处理：数据是存储在硬盘上的，只不过需要经常读取的数据会被加载到内存中，将数据存储在物理内存中，从而达到高速读写。
6、成熟度与广泛度：新兴数据库，成熟度较低，Nosql数据库中最为接近关系型数据库，比较完善的DB之一，适用人群不断在增长。

```
- mysql 和pgsql的区别
[参考](https://zhuanlan.zhihu.com/p/46405604)
[参考](https://www.modb.pro/db/25544)

```text
MySQL与PostgreSQL的区别
MySQL是应用开发者创建出来的DBMS；而PostgreSQL是由数据库开发者创建出来的DBMS 。
换句话说，MySQL倾向于使用者的角度，回答的问题是 “你想解决的是什么问题”；而PostgreSQL倾向于理论角度，回答的问题是 “数据库应该如何来解决问题” 。

MySQL一般会将数据合法性验证交给客户；PostgreSQL在合法性难方面做得比较严格。比如MySQL里插入 “2012-02-30” 这个时间时，会成功，但结果会是 “0000-00-00”；PostgreSQL不允许插入此值。
通常，PostgreSQL 被认为特性丰富，而MySQL被认为速度更快。但这个观点基本是在 MySQL 4.x / PostgreSQL 7.x 的事情，现在情况已经变了，PostgreSQL 在9.x版本速度上有了很大的改进，而MySQL特性也在增加。
在架构上，MySQL分为两层：上层的SQL层和几个存储引擎（比如InnoDB，MyISAM）。PostgreSQL 只有一个存储引擎提供这两个功能。
这两个数据库系统都可以针对应用的情境被优化、定制，精确的说哪个性能更好很难。MySQL项目一开始焦点就在速度上，而PostgreSQL一开始焦点在特性和规范标准上。
选哪个？

可能是由于历史原因MySQL在开发者中更流行一些。至少我们上学时没听说过PostgreSQL，当时不是MS SQL Server就是MySQL，而MySQL是开源的。实事上PostgreSQL直到8.0才官方支持了Windows系统。

如果没有什么历史原因（比如系统已经基于MySQL多年了），或技术积累原因（同事中MySQL高手多），那么我觉得选择PostgreSQL不会有错。

```

- 数据库优化
[参考](https://juejin.cn/post/6844904038459244552)
[参考](http://wangzhenkai.com/article/11)
[配置优化](https://segmentfault.com/a/1190000003072283)
```text
索引优化
优化索引，sql语句，分析慢查询
设计表的时候严格按照数据库设计规范来设计数据库
使用缓存，把经常访问并且不需要经常变化的数据放在缓存中，能够节约磁盘IO
优化硬件，采用ssd，使用磁盘队列技术（RAID0， RAID1，RAID5）
采用mysql自带的表分区技术，把数据分析分成不同文件，能够磁盘的读写效率
垂直分表，把一些不经常读的数据放在一张表当中，节约磁盘IO
主从分离读写，采用主从复制把数据库的读操作和写操作分离开来
分库分表分机器，数据量特别大的时候，主要的原理是数据路由
选择合适的表引擎，参数上的优化
进行架构级别的缓存，静态化和分布式
不采用全文检索
采用更快的存储方恨少，例如nosql存储经常访问的数据

配置优化
```

- 数据库隔离级别，事务，mongo的事务
- [参考]（https://www.toutiao.com/i6930164998287589896/?tt_from=weixin&utm_campaign=client_share&wxshare_count=1&timestamp=1614831082&app=news_article&utm_source=weixin&utm_medium=toutiao_android&use_new_style=1&req_id=202103041211220101511921465B00A51E&group_id=6930164998287589896）
```text
数据库隔离级别: 读未提交/读已提交（orcal默认）/可重复度（mysql默认） mvcc /序列化
事务： ACID ， 原子性(undo log)/一致性/隔离性/持久化（redo log）

mongodb不支持事务
```

- 为什么innoDB要使用B+树
[参考](https://blog.csdn.net/xlgen157387/article/details/79450295)
```text
由于索引数据是按顺序排序的，即每次读取了数据页的时候，里面的索引数据大部分都是需要用的，所以也很好的解决了上文提到的如何存储尽量多的有效的索引数据的问题。
因为B树不管叶子节点还是非叶子节点，都会保存数据，这样导致在非叶子节点中能保存的指针数量变少（有些资料也称为扇出）
指针少的情况下要保存大量数据，只能增加树的高度，导致IO操作变多，查询性能变低；
但是b树要求每个索引后面直接跟着数据，b+树则是非叶子结点会冗余到下一层，直到叶子结点层再追加数据。
mysql一个节点16kb，如果不跟数据只存索引，一个节点可以多存更多的索引，最后达到更多索引全放内存里，加快速度。我的理解是这样的。
B+树是应文件系统所需而产生的一种B树的变形树（文件的目录一级一级索引，只有最底层的叶子节点（文件）保存数据）非叶子节点只保存索引，不保存实际的数据，数据都保存在叶子节点中，这不就是文件系统文件的查找吗?

1、 B+树的磁盘读写代价更低：B+树的内部节点并没有指向关键字具体信息的指针，因此其内部节点相对B树更小，
如果把所有同一内部节点的关键字存放在同一盘块中，那么盘块所能容纳的关键字数量也越多，
一次性读入内存的需要查找的关键字也就越多，相对IO读写次数就降低了。

2、由于B+树的数据都存储在叶子结点中，分支结点均为索引，方便扫库，只需要扫一遍叶子结点即可，
但是B树因为其分支结点同样存储着数据，我们要找到具体的数据，需要进行一次中序遍历按序来扫，
所以B+树更加适合在区间查询的情况，所以通常B+树用于数据库索引。
```

- 锁。悲观锁(行锁: 共享锁、排它锁、更新锁)、乐观锁
[参考](https://www.toutiao.com/i6779129204651852299/?tt_from=weixin&utm_campaign=client_share&wxshare_count=1&timestamp=1614842865&app=news_article&utm_source=weixin&utm_medium=toutiao_android&use_new_style=1&req_id=20210304152745010151203160330208CD&group_id=6779129204651852299)
```text
行锁/表锁
行锁：共享锁(S)/排他锁（X）
```

### 业务场景

- 秒杀
[参考](https://jishuin.proginn.com/p/763bfbd2c789)
[参考](https://cloud.tencent.com/developer/article/1542291)
```text
负载
Redis集群，主从同步、读写分离，事务 + 异步处理
Redis 集群 ： https://www.cnblogs.com/yufeng218/p/13688582.html
```
- 支付
```text
强一致性
失败重试机制
```
- 单点登录
>[参考](https://www.jianshu.com/p/75edcc05acfd)
```text

```
 -权限管理(RBAC)
 [参考](https://www.pianshen.com/article/88011888708/)
 [参考](https://www.jianshu.com/p/44bfd8d6184b)
 ```text

```
- 大文件处理
[参考](https://zhuanlan.zhihu.com/p/124940982)
```text
curl, fie_get_content 是读到变量 不合适

fopen  r 只读方式打开，将文件指针指向文件头, 在一行一行的读
fopen() 将 filename 指定的名字资源绑定到一个流上。

```
 