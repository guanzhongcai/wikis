# WAL

## 1. WAL机制

> WAL即 Write Ahead Log，WAL的主要意思是说在将元数据的变更操作写入磁盘之前，先预先写入到一个log文件中



##  磁盘、SSD、内存随机读写与顺序读写性能对比

![random-serial-compare](images/random-serial-compare.jpg)



### 1.2 WAL原理

- 通过cache合并多条写操作为一条，减少IO次数
- 日志顺序追加性能远高于数据随机写.
- 随机内存处理性能远高于数据随机处理.

**性能:顺序的日志磁盘处理+随机的数据内存处理>随机的数据磁盘处理**



![WAL](images/WAL.png)

WAL还有一点很重要的帮助是可以在disaster recovery过程中起到状态恢复的作用，系统在load完元数据db后，再把未来得及提交的WAL apply进来，就能恢复成和之前最终一致的状态。



PGSQL架构：

![Proces and Memory Architecture ](images/Proces and Memory Architecture .png)





# Checkpoint

Checkpoint，数据库检查点。 “检查点”会创建一个已知的正常点，在意外关闭或崩溃后进行恢复的过程中，数据库引擎 可以从该点开始应用日志（WAL）中所包含的更改。

出于性能方面的考虑，数据库引擎 对内存（缓冲区缓存）中的数据库页进行修改，但在每次更改后不将这些页写入磁盘。 相反， 数据库引擎 定期发出对每个数据库的检查点命令。 检查点将当前内存中已修改的页面（称为“脏页”）和事务日志信息从内存写入磁盘，并记录事务日志中的信息 。



**checkpoints相关参数：**
checkpoint_segments:
WAL log的最大数量，系统默认值是3。该值越大，在执行介质恢复时处理的数据量也越大，时间相对越长。
checkpoint_timeout:
系统自动执行checkpoint之间的最大时间间隔，同样间隔越大介质恢复的时间越长。系统默认值是5分钟。
checkpoint_completion_target:
该参数表示checkpoint的完成目标，系统默认值是0.5,也就是说每个checkpoint需要在checkpoints间隔时间的50%内完成。
checkpoint_warning:
系统默认值是30秒，如果checkpoints的实际发生间隔小于该参数，将会在server log中写入写入一条相关信息。可以通过设置为0禁用信息写入。

**checkpoint执行控制:**
1,数据量达到checkpoint_segments*16M时，系统自动触发；
2,时间间隔达到checkpoint_timeout参数值时；
3,用户发出checkpoint命令时。

**checkpoints参数调整：**
正确合适的参数值总能够给系统带来益处，checkpoints参数合理的配置不仅能够减少系统IO写入的阻塞，同时还会减少高峰时IO给系统带来的压力。
首先可以通过观察checkpoint_warning参数写入的日志，来估算系统写入的数据量：一般情况下checkpoint_warning参数值小于checkpoint_timeout；
估算公式：checkpoint_segments*16M*(60s/m)/checkpoint_warning=大致每分钟数据量,得到每分钟写入的数据量(这里全部是估算，建立在warning参数的合理设置上)。
合理配置情况：checkpoint_segments*16M*checkpoint_timeout(m)略大于上述值.
以上述公式为依据，配置checkpoint_segments与checkpoint_timeout，两个参数应该尽量平衡为一个足够大和足够小的值。
在数据量异常高的情况下应该考虑，磁盘带宽与checkpoint时数据量的关系。





## 参考资料

- [WAL(Write Ahead Log)-知乎](https://zhuanlan.zhihu.com/p/228335203)

- [postgres-wal-doc](https://pgadminedb.readthedocs.io/en/latest/module_02/#write-ahead-logging-wal)

- [Database Checkpoints (SQL Server) - SQL Server | Microsoft Docs](https://docs.microsoft.com/en-us/sql/relational-databases/logs/database-checkpoints-sql-server?view=sql-server-ver15)

- [postgresql之WAL(Write Ahead Log) - 大肚熊 - 博客园 (cnblogs.com)](https://www.cnblogs.com/daduxiong/archive/2010/09/30/1839533.html)

- [postgresql之checkpoints - 大肚熊 - 博客园 (cnblogs.com)](https://www.cnblogs.com/daduxiong/archive/2010/09/28/1837682.html)

