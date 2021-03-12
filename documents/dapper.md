# dapper



### 低能耗

##### （1）`span`创建耗时，降低`span`占用空间

创建和销毁`span`，并将它们记录到本地磁盘以进行后续收集是生成开销的最重要来源。根`span`的创建和销毁平均需要204纳秒，而非根`span`的相同操作则需要176纳秒。 区别在于为根`span`分配全局唯一的`trace-id`会增加成本。

存储库中的每个`span`平均仅对应426个字节。 Dapper跟踪数据收集仅占我们正在监视的应用程序中网络活动的一小部分，不足Google生产环境中网络流量的0.01％。

##### （2）磁盘写入昂贵，合并磁盘写入操作

在Dapper的运行时库中，对本地磁盘的写入是最昂贵的操作

每次磁盘写入都会合并多个日志文件写入操作，并且相对于被跟踪的应用程序异步执行，因此减少了可见的开销。 但是，日志写入活动可能会对高吞吐量应用程序性能产生可察觉的影响，尤其是在跟踪所有请求的情况下。

##### （3）降低守护程序优先级

Dapper守护程序限制为内核调度程序中的最低优先级，以防CPU争用在高负载主机中出现

##### （4）采用较低的采样频率

使应用程序可以轻松使用注释API的全部带宽，而不必担心性能损失。

允许数据在被垃圾回收之前在主机的本地磁盘上保留的时间更长，这为收集基础结构提供了更大的灵活性。

##### （5）针对吞吐量，采用不同追踪采样策略

- 吞吐量大：均匀采样概率，每1024个采样一个，因为吞吐量大， 所以也不会丢失关键事件
- 吞吐量小：覆盖默认的采样概率，并避免用户手动干预

##### （6）控制数据文件到中央存储库的总大小，采用二次采样

将`trace-id`进行散列函数映射到参数z `0<z<1`，若参数 z 小于配置参数，则保留到中央数据库，否则丢弃



### 应用级别透明

这是最具挑战性的目标，即应用程序级开发人员不需要知道监控系统的存在

使用小部分线程，控制流，RPC库

达到语言无关



### 可扩展性

自适应采样



### 即时性

一分钟以内数据收集分析，对异常作出更快反应。和低能耗目标有相似部分。

Dapper 运行时最耗时的部分在 创建和销毁 spans 和 注释，以及将他们记录到本地磁盘。根 spans 比非根 span 花费更多时间，区别在于为根 spans 分配全局唯一的跟踪ID会增加成本。

磁盘IO 为最耗时的部分：合并多次写，异步执行



### Dapper 监控流程

1. `span`跨度数据写入本地日志文件

2. Dapper 后台程序拉取收集信息

3. 最后写入`Bigtable`中的一个单元中

   其中一个`trace`视为`Bigtable`中的一行，每列对应一个`span`

效率：数据从写入二进制文件到中央存储库所需时间，一般不到15秒，98%小于2分钟，其他25%时间里，延迟可能会到一小时





## 参考资料

- [论文笔记《Dapper, a_Large-Scale_Distributed_Systems_Tracing_Infrastructure》](https://blog.csdn.net/qq_39446719/article/details/109381939?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.control&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.control)

- [google dapper论文中文版](https://blog.csdn.net/qq_41618510/article/details/86500806)