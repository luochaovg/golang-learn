 > MySQL实战34讲
 ##### 课程链接
 - https://e.naixuejiaoyu.com/detail/p_61726279e4b039fe30f05d84/6?fromH5=true
  
 ##### 课件地址 
 - https://h5.tanma.tech/#/file?hjType=1&corp_id=ww81977112b716d7c8&sales_id=439513751017792&snapshot_id=956d27b0001e11ec983400163e342bad&content_id=956d27b0001e11ec983400163e342bad&content_title=%E6%9E%B6%E6%9E%84%E5%B8%88%E7%BA%A7MySQL%E6%B5%B7%E9%87%8F%E6%95%B0%E6%8D%AE%E8%AE%BE%E8%AE%A1%E4%B8%8E%E5%AE%9E%E6%88%98%EF%BC%88%E4%B8%8A%EF%BC%89.pdf&content_type=7&source_type=1&source=439513751017792&source_title=&connectFlag=1&showType=1,2,3,4&trace_id=3eee3843-5c17-46a1-b61b-792f25817589&trace_deep=0
 - https://h5.tanma.tech/#/file?hjType=1&corp_id=ww81977112b716d7c8&sales_id=439513751017792&snapshot_id=9a1245c1001e11ec983400163e342bad&content_id=9a1245c1001e11ec983400163e342bad&content_title=%E6%9E%B6%E6%9E%84%E5%B8%88%E7%BA%A7MySQL%E6%B5%B7%E9%87%8F%E6%95%B0%E6%8D%AE%E8%AE%BE%E8%AE%A1%E4%B8%8E%E5%AE%9E%E6%88%98%EF%BC%88%E4%B8%AD%EF%BC%89.pdf&content_type=7&source_type=1&source=439513751017792&source_title=&connectFlag=1&showType=1,2,3,4&trace_id=88894145-bee6-4548-8707-7cf07ff31c47&trace_deep=0
 - https://h5.tanma.tech/#/file?hjType=1&corp_id=ww81977112b716d7c8&sales_id=439513751017792&snapshot_id=9f15140b001e11ec8b8b00163e3442f2&content_id=9f15140b001e11ec8b8b00163e3442f2&content_title=%E6%9E%B6%E6%9E%84%E5%B8%88%E7%BA%A7MySQL%E6%B5%B7%E9%87%8F%E6%95%B0%E6%8D%AE%E8%AE%BE%E8%AE%A1%E4%B8%8E%E5%AE%9E%E6%88%98%EF%BC%88%E4%B8%8B%EF%BC%89.pdf&content_type=7&source_type=1&source=439513751017792&source_title=&connectFlag=1&showType=1,2,3,4&trace_id=f9e67256-6a2f-4a23-b610-c4d34bbc84fc&trace_deep=0
 
 ##### 预防死锁
 ```text
1,如果使用 insert ... select 语句备份表格且数据量较大，在单独的时间点操作，
避免与其他 sql 语句争夺资源，或使用 select into outfile 加上 load data infile 代替 insert ... select，这样不仅快，而且不会要求锁定。

2,一个锁定记录集的事务，其操作结果集应尽量简短，以免一次占用太多资源，与其他事务处理的记录冲突。(尤其注意 DDL 语句)。

3,更新或者删除表格数据，sql 语句的 where 条件都是主键或都是索引，避免两种情况交叉，造成死锁。对于 where 子句较复杂的情况，将其单独通过 sql 得到后，再在更新语句中使用。

4,sql 语句的嵌套表不要太多，能拆分就拆分，避免占有资源同时等待资源，导致与其他事务冲突。

5,对定点运行脚本的情况，避免在同一时间点运行多个对同一表进行读写的脚本，特别注意加锁且操作数据量比较大的语句。

6,应用程序中增加对死锁的判断，如果事务意外结束，重新运行该事务，减少对功能的影响。

7,mysqldump 默认是锁表的，请使用参数 --single-transaction 或者 --skip-opt。
```

##### 如果有一列含有 NULL 值那么这个组合索引都将失效，一般需要给默认值 0 或者 "" 字符串。 （重要不要忘记）

##### 索引失效场景
```text
1、like 以 % 开头，索引无效；当 like 前缀没有 %，后缀有 % 时，索引有效。

2、当 or 左右查询字段只有一个是索引，该索引失效，只有当 or 左右查询字段均为索引时，才会生效。

3、组合索引，不遵循最左原则，索引失效。

4、数据类型出现隐式转化。如 varchar 不加单引号的话可能会自动转换为 int 型，使索引无效，产生全表扫描。

5、在索引列上使用 IS NULL 或 IS NOT NULL 操作。在回表操作成本大于全表扫描的时候不使用索引 (并排除覆盖索引)。

6、在索引字段上使用 not，<>，!=。不等于操作符是永远不会用到索引的，因此对它的处理只会产生全表扫描。优化方法：key <> 0 改为 key> 0 or key < 0。

7、对索引字段使用运算符或函数操作时索引失效。

8、当全表扫描速度比索引速度快时，MySQL 会使用全表扫描，此时索引失效。(同第5点)

9、order by 和 group by 中的字段违反最左前缀原则或含有非索引字段，则分别会产生文件排序和临时表。

10、范围查询之后的列索引会失效，包含 <、>、between。
```

##### 使用索引的优点
```text
1,可以通过建立唯一索引或者主键索引，保证数据库表中每一行数据的唯一性。

2,建立索引可以大大提高检索的速度，以及减少检索的行数。

3,在表连接的条件字段创建索引后，可以加快表与表之间的连接速度 (否则以笛卡尔积数量级进行表扫描，表和扫描行数的增长成指数级上升)。

4,在分组和排序字段创建索引，可以减少查询中分组和排序所消耗的时间 (数据库的记录会重新排序)。
```

##### 使用索引的缺点
```text
1,每个索引会占用一定的物理空间 (如果你在一个大表上创建了多个组合索引，索引文件则会膨胀非常快)。

2,当对表的数据进行 INSERT，UPDATE，DELETE 的时候，索引也要动态的维护，这样就会降低数据库的写速度。
```

##### 使用索引需要注意的地方
```text
在经常需要搜索的列上，可以加快索引的速度。

在表与表的连接条件上加上索引，可以加快连接查询的速度。

在经常需要 order by、group by 和 distinct 的列上加索引可以加快排序的时间，(单独 order by 用不了索引，索引考虑加 where 或加组合索引)。

在 <、<=、>、>=、BETWEEN、IN 之后的相关字段可能不走索引。

like 语句中如果你对 nickname 字段建立了一个索引，当查询语句是 nickname like '%abc%' 时索引不会被使用；而查询语句是 nickname like 'abc%' 时则会被使用。

索引不会包含 NULL 列，如果列中包含 NULL 值查询时索引失效。

使用短索引，如果你的一个字段是 char(32) 或者 int(32)，在创建索引的时候建议指定前缀长度，比如前 10 个字符 (前提是多数值是唯一的) 那么短索引可以提高查询速度，并且可以减少磁盘空间的占用，也可以减少 I/0 消耗。

不要在列上使用运算或函数，这样会使得 MySQL 索引失效，并进行全表扫描。

选择越小的数据类型越好，因为通常越小的数据类型通常在磁盘、内存、cpu 和缓存中占用的空间就越小，处理起来则会更快。

查询中很少使用到的列不应该创建索引，如果建立了索引反而会降低 MySQL 的性能和加大磁盘的占用。

表数据、或列去重后数据量较小的情况下 (如性别字段) 都不应该创建索引，因为几乎没有意义。

定义为 image 和 bit 的数据类型列不应该创建索引。

当表的 UPDATE、INSERT 和 DELETE 操作远多于 SELECT 操作时不应该创建索引，这两类操作是互斥的关系。

count(*) 和 count(1) 一样无任何差别，但是建议用 count(*)，因为是 SQL92 标准语法并且做过很多优化，至于 count(column) 需要逐行判断 IS NOT NULL 才会被计数，不建议使用。

```


##### where in 是否走索引
 - [参考](https://www.cnblogs.com/taotaozhuanyong/p/14812561.html)
 
##### SQL语句
```text
查询参考(重点)： 
https://www.cnblogs.com/chiangchou/p/mysql-3.html#_label1_1

CASE...WHEN...
（参考）https://juejin.cn/post/6971040309065187342

```