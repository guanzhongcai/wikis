# memo



## etcd的raft算法

- 节点会有三个状态：follower，candidate，leader
- 算法中有两个超时设置，一个是选举超时，一个是心跳超时
- 数据的强一致性依靠日志的复制机制
- 初始时，大家都是follwer，这时选举超时机制触发leader选举，超时是随机的一个时间区间150-300ms，最先发起选举的节点，自己就会成为candidate，如果后续的选举得到大多数节点的通过，自己就成为leader；如果同时存在各一半节点数的选举，那么此次选举不算，将重新来一次leader选举；
- 以后所有的数据修改，都提交到leader，开始时先产生一条未提交的修改日志，当leader通知其他的大多数节点都统一修改，就执行这次修改，并提交此次的修改。
- 脑裂
  - 当发生网路分区时，因为节点数是奇数个，多数的节点继续提供服务，少数的节点不可用
  - 到网络分区恢复正常，少数的节点同步多数节点的日志，数据状态实现最终一致

## goroutine 实现机制

- golang语言层面控制的，运行在线程上，队列式的
- 每个协程最多运行10ms



## go的slice

**`append()`这个函数在 `cap` 不够用的时候就会重新分配内存以扩大容量，而如果够用的时候不不会重新分享内存！**

```go
dir1 := path[:sepIndex:sepIndex]
```

代码使用了 Full Slice Expression，其最后一个参数叫“Limited Capacity”，于是，后续的 `append()` 操作将会导致重新分配内存。



## 接口完整性检查

```go
type Shape interface {
    Sides() int
    Area() int
}
type Square struct {
    len int
}
func (s* Square) Sides() int {
    return 4
}
func main() {
    s := Square{len: 5}
    fmt.Printf("%d\n",s.Sides())
}

// 声明一个 _ 变量（没人用），其会把一个 nil 的空指针，从 Square 转成 Shape，这样，如果没有实现完相关的接口方法，编译器就会报错：
var _ Shape = (*Square)(nil)
```

这样就做到了个强验证的方法。



## go的gc

> go语言内置运行时Runtime，抛弃传统的内存分配策略，改为自主分配；使用google的Thread-Cache Malloc算法，自己管理内存池和预分配，不用每次内存分配都需要向进行系统调用。

> 该算法核心思想就是把内存做分级管理，从而降低锁的粒度。它将可用的堆内存采用二级分配（全局内存池、线程内存池）的方式进行管理。每个线程都会自行维护一个独立的内存池，进行内存分配时优先从某个线程中的内存池分配；当内存不足时，才会向全局内存池申请，已避免不同线程对全局内存池的频繁竞争。

- 基本策略：
  - 每次从操作系统申请一块大内存，以减少系统调用
  - 将申请的大内存切分成不同大小的小块，构成链表，供后续使用
  - 为对象分配内存是，只需要从大小合适的链表中提取一个即可
  - 回收对象内存时，将该小块归还给原链表，以便复用
  - 如果闲置内存过多，则尝试归还部分内存给操作系统，降低整体开销。

- go程序启动时，向操作系统申请一块内存空间，切成小块然后自己进行管理。
- 申请到的内存会被分成3个区域，分别为：
  - 512M的span区域
  - 16G的bitmap区域
  - 512G的arena区域
- 这些只是虚拟的地址空间，并不会真正地分配内存
- 内存管理组件，分为3部分组成：
  - cache：每个运行期的工作线程都会绑定一个cache，用于无锁object的分配
  - central：为所有cache提供切分好的的后备span资源
  - heap：管理闲置span，需要时想操作系统申请内存
  - 
  - 分别为并启动多个线程管理，每个线程管理一部分被切割为不同大小的内存片，以后的使用直接向这些线程申请，避免锁粒度的性能消耗，使用完再返回给内存调度器



## sync.Pool

这个类设计的目的是用来短时间（2分钟内）保存和复用临时对象，以减少内存分配，降低CG压力。

这个清理过程是在每次垃圾回收之前做的。垃圾回收是固定两分钟触发一次。

```go
package main
 
import(
    "sync"
    "log"
)
 
func main(){
    // 建立对象
    var pipe = &sync.Pool{New:func()interface{}{return "Hello, BeiJing"}}
     
    // 准备放入的字符串
    val := "Hello,World!"
     
    // 放入
    pipe.Put(val)
     
    // 取出
    log.Println(pipe.Get())
     
    // 再取就没有了,会自动调用NEW
    log.Println(pipe.Get())
}
 
// 输出
2014/09/30 15:43:30 Hello, World!
2014/09/30 15:43:30 Hello, BeiJing
```

- 每次清理会将 Pool 中的所有对象都清理掉！

- sync.Pool 的定位不是做类似连接池的东西，它的用途仅仅是增加对象重用的几率，减少 gc 的负担，而开销方面也不是很便宜的。

[参考资料](https://www.cnblogs.com/-wenli/p/12325248.html)



#### 性能提示

Go 语言是一个高性能的语言，但并不是说这样我们就不用关心性能了，我们还是需要关心的。下面是一个在编程方面和性能相关的提示。

- 如果需要把数字转字符串，使用 `strconv.Itoa()` 会比 `fmt.Sprintf()` 要快一倍左右
- 尽可能地避免把`String`转成`[]Byte` 。这个转换会导致性能下降。
- 如果在for-loop里对某个slice 使用 `append()`请先把 slice的容量很扩充到位，这样可以避免内存重新分享以及系统自动按2的N次方幂进行扩展但又用不到，从而浪费内存。
- 使用`StringBuffer` 或是`StringBuild` 来拼接字符串，会比使用 `+` 或 `+=` 性能高三到四个数量级。
- 尽可能的使用并发的 go routine，然后使用 `sync.WaitGroup` 来同步分片操作
- 避免在热代码中进行内存分配，这样会导致gc很忙。尽可能的使用 `sync.Pool` 来重用对象。
- 使用 lock-free的操作，避免使用 mutex，尽可能使用 `sync/Atomic`包。 （关于无锁编程的相关话题，可参看《[无锁队列实现](https://coolshell.cn/articles/8239.html)》或《[无锁Hashmap实现](https://coolshell.cn/articles/9703.html)》）
- **使用 I/O缓冲，I/O是个非常非常慢的操作，使用 `bufio.NewWrite()` 和 `bufio.NewReader()` 可以带来更高的性能。**
- 对于在for-loop里的固定的正则表达式，一定要使用 `regexp.Compile()` 编译正则表达式。性能会得升两个数量级。
- 如果你需要更高性能的协议，你要考虑使用 [protobuf](https://github.com/golang/protobuf) 或 [msgp](https://github.com/tinylib/msgp) 而不是JSON，因为JSON的序列化和反序列化里使用了反射。
- 你在使用map的时候，使用整型的key会比字符串的要快，因为整型比较比字符串比较要快。

> bufio 是通过缓冲来提高效率
>
> 简单的说就是，把文件读取进缓冲（内存）之后再读取的时候就可以避免文件系统的io 从而提高速度。同理，在进行写操作时，先把文件写入缓冲（内存），然后由缓冲写入文件系统。看完以上解释有人可能会表示困惑了，直接把 内容->文件 和 内容->缓冲->文件相比， 缓冲区好像没有起到作用嘛。其实缓冲区的设计是为了存储多次的写入，最后一口气把缓冲区内容写入文件。下面会详细解释
>
>
> 
> 链接：[golang中bufio包的实现原理](https://blog.csdn.net/LiangWenT/article/details/78995468)

## go的chann

- chann的数据结构是hchan：

  ```go
  type hchan struct {
  	//channel队列里面总的数据量
  	qcount   uint           // total data in the queue
  	// 循环队列的容量，如果是非缓冲的channel就是0
  	dataqsiz uint           // size of the circular queue
  	// 缓冲队列，数组类型。
  	buf      unsafe.Pointer // points to an array of dataqsiz elements
  	// 元素占用字节的size
  	elemsize uint16
  	// 当前队列关闭标志位，非零表示关闭
  	closed   uint32
  	// 队列里面元素类型
  	elemtype *_type // element type
  	// 队列send索引
  	sendx    uint   // send index
  	// 队列索引
  	recvx    uint   // receive index
  	// 等待channel的G队列。
  	recvq    waitq  // list of recv waiters
  	// 向channel发送数据的G队列。
  	sendq    waitq  // list of send waiters
  
  	// lock protects all fields in hchan, as well as several
  	// fields in sudogs blocked on this channel.
  	//
  	// Do not change another G's status while holding this lock
  	// (in particular, do not ready a G), as this can deadlock
  	// with stack shrinking.
  	// 全局锁
  	lock mutex
  }
  ```

  - 全局的mutex锁
  - 接收协程队列recvq和发送协程队列sendq
  - 数组实现的环形队列circlebuffer，对于有缓冲的channel，sendx和recvx表示读写的缓冲区索引

- 在接收协程接收到新的缓冲消息时，会顺便触发阻塞读协程的重新运行，反之亦然。
- 思考：通过通信来实现共享内存，而不是通过共享内存来实现通信。（CSP）

## mysql索引

- B+树
- 是由二叉树演变而来的N叉数，即子节点数不是2个，是n个
- 叶子节点数目和数据节点的数目一样多，便于范围查找
- 聚簇索引是索引和数据都存储在一起，都在叶子节点，一般主键索引都是聚簇索引
- 非聚簇索引的索引和数据是分开存放的
- 二级索引又名辅助索引，存储的是主键值，而非数据地址，通过二级索引查询时先找到二级索引存储的主键值，然后再通过主键索引查找到存储的数据。唯一索引、普通索引、前缀索引都是二级索引。
- 因为B+树有最左原则，所以复合索引会依赖第一个字段索引排序，每个叶子节点对应的数据是已经排序好的



## SQL的事务隔离级别

- Read Uncommitted: 读未提交，事务做的操作，即使没有提交，其他事务也是可见的，所有会有脏读。脏读即读取了错误的数据，因为可能数据操作会需要undo回滚。
- Read Committed: 读已提交，只能读到已提交的事务，大多数数据库的默认隔离级别，又名不可重复读，因为会出现幻读，幻读即前后两次查询的结果可能不一致，因为这个前后两次间隔了某个事务的操作完成。
- Repeatable Read: 可重复读，mysql的默认隔离级别
- Serializable: 可串行化，事务的最高级别，让所有事务串行执行
- MVCC
  -  Multi-Version Concurrency Control，InnoDB的多版本并发控制，解决了幻读中的幻行问题。
  - InnoDB的MVCC是对每行记录增加两个隐藏的列实现的，一列是行的创建时间，一列是行的删除时间，列值存储的实际是系统版本号
  - 事务的版本号是事务开始的系统版本号，用来和查询到的每行记录的版本号进行比较
  - 事务只查找小于或等于事务版本号的数据行，小于事务版本号的是事务开始前就存在的，等于事务版本号的是事务自身修改过的。
  - 行的删除版本，要么未定义，要么比事务版本号大，以确保事务读取到的数据，在事务开始前未被删除

## mongo事务

- start transaction...commit
- 

## redis



### 事务

- mulit
- watch
- exec
- discard
- 如果没有watch，就会都执行，有语法错误也都顺序执行，不会停下。若有watch，则会都执行，如果无语法错误，则都成功

### 分片

### 备份方式

- RDB：二进制方式

- AOF：记录修改指令

  

## 架构设计的弹力设计

## 数据库异步写

## websocket



## mongo分片

## nodejs中的stream流

## Docker

### 镜像制作

- 将 Dockerfile 置于一个空目录下，或者项目根目录下