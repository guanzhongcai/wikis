# 高性能MySQL

【高性能MySQL】书的笔记。



## MySQL架构和历史

### 架构

![](mysql/logical-image.jpeg)

- MySQL架构将 **查询处理（Query Processing）、其他系统任务（Server Task）数据的存储/提取** 相分离。

- 每个客户端连接都会在服务器进程中拥有一个线程。
- MySQL会解析查询，并创建内部数据结构（解析树），然后对其进行各种优化，包括：
  - 重写查询
  - 决定表的读取顺序
  - 选择合适的索引等



### 锁

- 并发读写锁：
  - 共享锁（shared lock）和排它锁（exclusive lock），也叫读锁（read lock）和写锁（write lock）
  - 写锁比读锁有更高的优先级，可以插队到读锁前面
- 锁粒度
  - 行级锁（row-level lock）：只在存储引擎层实现
  - 表锁（table lock）：例如 `ALTER TABLE`



### 事务

- 事务
  - 事务就是一组原子性的SQL查询，或者说一个独立的工作单元
  - 事务内的语句，要么全部执行成功，要么全部执行失败，例如银行转账
- 事务的ACID特征
  - 原子性（atomicity）：一个事务必须被视为一个不可分割的最小工作单元。
  - 一致性（consistency）：数据库总是从一个一致性的状态转换到另外一个一致性的状态。
  - 隔离性（isolation）：隔离级别（isolation level）
  - 持久性（durability）：一旦事务提交，则其所作的修改就会永久保存到数据库中。当然也有持久性级别的概念。

- 事务隔离级别
  - 未提交读（read uncommitted）：事务中的修改，即使没有提交，对其他事务也是可见的，俗称脏读（dirty read）
  - 提交读（read committed）：又称不可重复读（nonrepeatable read），因为重复读就得到不一样的结果了
  - 可重复读（repeatable read）：MySQL默认隔离级别，解决了脏读问题。该级别保证了在同一个事务中，多次读取同样记录的结果是一致的！但理论上无法解决幻读（phantom read)。幻读指的是当某个事务在读取某个范围内的记录时，另外一个事务又在该范围内插入了新的记录，当之前的事务再次读取该范围的记录时，会产生幻行（phantom row）。InnoDB存储引擎通过多版本并发控制（MVCC，Multiversion Concurrency Control）解决了幻读的问题。
  - 可串行化（serialization）：最高的隔离级别。通过强制事务串行执行，避免了幻读。serialization会在读取的每一行数据上都加锁（所以可能导致大量的超时和锁争用问题）