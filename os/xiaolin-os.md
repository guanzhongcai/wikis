# 小林coding



## 一致性hash

> [微信一面：什么是一致性哈希？用在什么场景？解决了什么问题？ - 小林coding - 博客园 (cnblogs.com)](https://www.cnblogs.com/xiaolincoding/p/15918321.html)

- 假设总数据条数为 M，哈希算法在面对节点数量变化时，**最坏情况下所有数据都需要迁移，所以它的数据迁移规模是 O(M)**，这样数据的迁移成本太高了。
- 一致哈希算法也用了取模运算，但与哈希算法不同的是，哈希算法是对节点的数量进行取模运算，而**一致哈希算法是对 2^32 进行取模运算，是一个固定的值**。
- 哈希环
- 一致性哈希要进行两步哈希：
  - 第一步：对存储节点进行哈希计算，也就是对存储节点做哈希映射，比如根据节点的 IP 地址进行哈希；
  - 第二步：当对数据进行存储或访问时，对数据进行哈希映射；

- **一致性哈希是指将「存储节点」和「数据」都映射到一个首尾相连的哈希环上**。

- 对「数据」进行哈希映射得到一个结果要怎么找到存储该数据的节点呢？

  答案是，映射的结果值往**顺时针的方向的找到第一个节点**，就是存储该数据的节点。

- 加入**虚拟节点**，也就是对一个真实节点做多个副本。

  具体做法是，**不再将真实节点映射到哈希环上，而是将虚拟节点映射到哈希环上，并将虚拟节点映射到实际节点，所以这里有「两层」映射关系。**

- 比如 Nginx 的一致性哈希算法，每个权重为 1 的真实节点就含有160 个虚拟节点。



## SQL的count效率

- count() 是一个聚合函数，函数的参数不仅可以是字段名，也可以是其他任意表达式，该函数作用是**统计符合查询条件的记录中，函数指定的参数不为 NULL 的记录有多少个**。

  ```sql
  # 统计「 t_order 表中，name 字段不为 NULL 的记录」有多少个
  select count(name) from t_order;
  ```

- InnoDB 是通过 B+ 树来保持记录的，根据索引的类型又分为聚簇索引和二级索引，它们区别在于，聚簇索引的叶子节点存放的是实际数据，而二级索引的叶子节点存放的是主键值，而不是实际数据。

- 使用 MyISAM 引擎时，执行 count 函数只需要 O(1 )复杂度，这是因为每张 MyISAM 的数据表都有一个 meta 信息有存储了row_count值，由表级锁保证一致性，所以直接读取 row_count 值就是 count 函数的执行结果。

- 而 InnoDB 存储引擎是支持事务的，同一个时刻的多个查询，由于多版本并发控制（MVCC）的原因，InnoDB 表“应该返回多少行”也是不确定的，所以无法像 MyISAM一样，只维护一个 row_count 变量。

- 面对大表的记录统计，如果你的业务对于统计个数不需要很精确，比如搜索引擎在搜索关键词的时候，给出的搜索结果条数是一个大概值。我们就可以使用 show table status 或者 explain 命令来表进行估算。执行 explain 命令效率是很高的，因为它并不会真正的去查询



## nginx文件传输配置

来根据⽂件的⼤⼩来使⽤不同的⽅式：

```nginx
location /video/ {
  sendfile on;
  aio on;
  directio 1024m;
}
```

当⽂件⼤⼩⼤于 directio 值后，使⽤「异步 I/O + 直接 I/O」，否则使⽤「零拷⻉技术」。



- 为了提⾼⽂件传输的性能，于是就出现了零拷⻉技术，它通过⼀次系统调⽤（ sendfile ⽅

  法）合并了磁盘读取与⽹络发送两个操作，降低了上下⽂切换次数。另外，拷⻉数据都是发

  ⽣在内核中的，天然就降低了数据拷⻉的次数。

- Kafka 和 Nginx 都有实现零拷⻉技术，这将⼤⼤提⾼⽂件传输的性能。

- 零拷⻉技术是基于 PageCache 的，PageCache 会缓存最近访问的数据，提升了访问缓存数

  据的性能，同时，为了解决机械硬盘寻址慢的问题，它还协助 I/O 调度算法实现了 IO 合并与

  预读，这也是顺序读⽐随机读性能好的原因。这些优势，进⼀步提升了零拷⻉的性能。



## Linux一切皆文件

- 其实是一种面向对象的设计思想。
- 串口是文件，内存是文件，usb是文件，进程信息是文件，网卡是文件，建立的每个网络通讯都是文件，蓝牙设备也是文件，等等等等。
- 所有外设都是文件，本质上就是说他们都支持用来访问文件的那些接口，可以被当做文件来访问。这个原理与子类都能当做基类访问是一样的，就是操作系统层面的oop思想。

- 任何东西都挂在文件系统之上，即使它们不是文件，也以文件的形式来呈现。

- 有一个很关键的点，这个一切是单向的，也即所有的东西都单向通过文件系统呈现，反向不一定可行。比如通过新建文件的方式来创建磁盘设备是行不通的。

- 比如我们经常会讲的进程(/proc)、设备(/dev)、Socket等等，实际上都不是文件，但是你可以以文件系统的规范来访问它，修改属主和属性。

- Linux下有`lsof`命令，可以查看所有已经打开的文件，你使用`lsof -p [pid]`的方式就可以查看对应的进程都打开了什么文件，而其中的`type`字段就是表明它是什么类型，通过`man losf` 命令你可以查看到它有下面这么多种。

  ```bash
  # lsof siege-server
  ubuntu@test1:~$ lsof -p 1213429
  lsof: WARNING: can't stat() tracefs file system /sys/kernel/debug/tracing
        Output information may be incomplete.
  COMMAND       PID   USER   FD      TYPE             DEVICE SIZE/OFF     NODE NAME
  node\x20/ 1213429 ubuntu  cwd       DIR              252,2     4096   652074 /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server
  node\x20/ 1213429 ubuntu  rtd       DIR              252,2     4096        2 /
  node\x20/ 1213429 ubuntu  txt       REG              252,2    14416   133994 /usr/bin/node
  node\x20/ 1213429 ubuntu  mem       REG              252,2   101320   152459 /usr/lib/x86_64-linux-gnu/libresolv-2.31.so
  node\x20/ 1213429 ubuntu  mem       REG              252,2    31176   152452 /usr/lib/x86_64-linux-gnu/libnss_dns-2.31.so
  node\x20/ 1213429 ubuntu  mem       REG              252,2    51832   152453 /usr/lib/x86_64-linux-gnu/libnss_files-2.31.so
  node\x20/ 1213429 ubuntu  mem       REG              252,2    20176   274288 /usr/share/zoneinfo-icu/44/le/timezoneTypes.res
  node\x20/ 1213429 ubuntu  mem       REG              252,2   156000   274290 /usr/share/zoneinfo-icu/44/le/zoneinfo64.res
  node\x20/ 1213429 ubuntu  mem       REG              252,2 28046896   131860 /usr/lib/x86_64-linux-gnu/libicudata.so.66.1
  node\x20/ 1213429 ubuntu  mem       REG              252,2   104984   130941 /usr/lib/x86_64-linux-gnu/libgcc_s.so.1
  node\x20/ 1213429 ubuntu  mem       REG              252,2  1369352   152447 /usr/lib/x86_64-linux-gnu/libm-2.31.so
  node\x20/ 1213429 ubuntu  mem       REG              252,2  1956992   130956 /usr/lib/x86_64-linux-gnu/libstdc++.so.6.0.28
  node\x20/ 1213429 ubuntu  mem       REG              252,2    18816   152445 /usr/lib/x86_64-linux-gnu/libdl-2.31.so
  node\x20/ 1213429 ubuntu  mem       REG              252,2  1976648   131879 /usr/lib/x86_64-linux-gnu/libicuuc.so.66.1
  node\x20/ 1213429 ubuntu  mem       REG              252,2  3132040   131862 /usr/lib/x86_64-linux-gnu/libicui18n.so.66.1
  node\x20/ 1213429 ubuntu  mem       REG              252,2   598104   134055 /usr/lib/x86_64-linux-gnu/libssl.so.1.1
  node\x20/ 1213429 ubuntu  mem       REG              252,2  2954080   134054 /usr/lib/x86_64-linux-gnu/libcrypto.so.1.1
  node\x20/ 1213429 ubuntu  mem       REG              252,2   162160   138528 /usr/lib/x86_64-linux-gnu/libnghttp2.so.14.19.0
  node\x20/ 1213429 ubuntu  mem       REG              252,2    75792   133990 /usr/lib/x86_64-linux-gnu/libcares.so.2.3.0
  node\x20/ 1213429 ubuntu  mem       REG              252,2   194648   131743 /usr/lib/x86_64-linux-gnu/libuv.so.1.0.0
  node\x20/ 1213429 ubuntu  mem       REG              252,2   108936   134327 /usr/lib/x86_64-linux-gnu/libz.so.1.2.11
  node\x20/ 1213429 ubuntu  mem       REG              252,2  2029224   152443 /usr/lib/x86_64-linux-gnu/libc-2.31.so
  node\x20/ 1213429 ubuntu  mem       REG              252,2   157224   152458 /usr/lib/x86_64-linux-gnu/libpthread-2.31.so
  node\x20/ 1213429 ubuntu  mem       REG              252,2 22721208   133992 /usr/lib/x86_64-linux-gnu/libnode.so.64
  node\x20/ 1213429 ubuntu  mem       REG              252,2   191472   152435 /usr/lib/x86_64-linux-gnu/ld-2.31.so
  node\x20/ 1213429 ubuntu    0u     unix 0x0000000000000000      0t0 88567492 type=STREAM
  node\x20/ 1213429 ubuntu    1u     unix 0x0000000000000000      0t0 88567494 type=STREAM
  node\x20/ 1213429 ubuntu    2u     unix 0x0000000000000000      0t0 88567496 type=STREAM
  node\x20/ 1213429 ubuntu    3u     unix 0x0000000000000000      0t0 88567498 type=STREAM
  node\x20/ 1213429 ubuntu    4u  a_inode               0,14        0    10385 [eventpoll]
  node\x20/ 1213429 ubuntu    5r     FIFO               0,13      0t0 88567500 pipe
  node\x20/ 1213429 ubuntu    6w     FIFO               0,13      0t0 88567500 pipe
  node\x20/ 1213429 ubuntu    7r     FIFO               0,13      0t0 88567501 pipe
  node\x20/ 1213429 ubuntu    8w     FIFO               0,13      0t0 88567501 pipe
  node\x20/ 1213429 ubuntu    9u  a_inode               0,14        0    10385 [eventfd]
  node\x20/ 1213429 ubuntu   10u  a_inode               0,14        0    10385 [eventpoll]
  node\x20/ 1213429 ubuntu   11r     FIFO               0,13      0t0 88567502 pipe
  node\x20/ 1213429 ubuntu   12w     FIFO               0,13      0t0 88567502 pipe
  node\x20/ 1213429 ubuntu   13u  a_inode               0,14        0    10385 [eventfd]
  node\x20/ 1213429 ubuntu   14u  a_inode               0,14        0    10385 [eventpoll]
  node\x20/ 1213429 ubuntu   15r     FIFO               0,13      0t0 88566591 pipe
  node\x20/ 1213429 ubuntu   16w     FIFO               0,13      0t0 88566591 pipe
  node\x20/ 1213429 ubuntu   17u  a_inode               0,14        0    10385 [eventfd]
  node\x20/ 1213429 ubuntu   18r      CHR                1,3      0t0        6 /dev/null
  node\x20/ 1213429 ubuntu   19u     IPv6           88567693      0t0      TCP *:40160 (LISTEN)
  node\x20/ 1213429 ubuntu   20w      REG              252,2        0  1443133 /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/con-log-master-server-1.log
  node\x20/ 1213429 ubuntu   21w      REG              252,2        0  1443134 /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/rpc-log-master-server-1.log
  node\x20/ 1213429 ubuntu   22w      REG              252,2        0  1443135 /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/forward-log-master-server-1.log
  node\x20/ 1213429 ubuntu   23w      REG              252,2        0  1443136 /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/rpc-debug-master-server-1.log
  node\x20/ 1213429 ubuntu   24w      REG              252,2        0  1443137 /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/crash.log
  node\x20/ 1213429 ubuntu   25w      REG              252,2        0  1443138 /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/admin.log
  node\x20/ 1213429 ubuntu   26w      REG              252,2    41405  1443139 /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/pomelo-master-server-1.log
  node\x20/ 1213429 ubuntu   27w      REG              252,2     8171  1443140 /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/pomelo-admin.log
  node\x20/ 1213429 ubuntu   28w      REG              252,2        0  1443141 /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/pomelo-rpc-master-server-1.log
  node\x20/ 1213429 ubuntu   29u     unix 0x0000000000000000      0t0 88567694 type=STREAM
  node\x20/ 1213429 ubuntu   30u     IPv4           88567710      0t0      TCP localhost:47802->localhost:40160 (ESTABLISHED)
  node\x20/ 1213429 ubuntu   31u     unix 0x0000000000000000      0t0 88567696 type=STREAM
  node\x20/ 1213429 ubuntu   32u     IPv6           88567711      0t0      TCP localhost:40160->localhost:47802 (ESTABLISHED)
  node\x20/ 1213429 ubuntu   33u     unix 0x0000000000000000      0t0 88567698 type=STREAM
  node\x20/ 1213429 ubuntu   34u     IPv6           88569152      0t0      TCP localhost:40160->localhost:47848 (ESTABLISHED)
  ```

  ```bash
  ubuntu@test1:/proc/1213429$ cat limits 
  Limit                     Soft Limit           Hard Limit           Units     
  Max cpu time              unlimited            unlimited            seconds   
  Max file size             unlimited            unlimited            bytes     
  Max data size             unlimited            unlimited            bytes     
  Max stack size            8388608              unlimited            bytes     
  Max core file size        0                    unlimited            bytes     
  Max resident set          unlimited            unlimited            bytes     
  Max processes             29416                29416                processes 
  Max open files            1048576              1048576              files     
  Max locked memory         67108864             67108864             bytes     
  Max address space         unlimited            unlimited            bytes     
  Max file locks            unlimited            unlimited            locks     
  Max pending signals       29416                29416                signals   
  Max msgqueue size         819200               819200               bytes     
  Max nice priority         0                    0                    
  Max realtime priority     0                    0                    
  Max realtime timeout      unlimited            unlimited            us        
  
  ubuntu@test1:/proc/1213429$ sudo cat stack 
  [<0>] ep_poll+0x400/0x450
  [<0>] do_epoll_wait+0xb8/0xd0
  [<0>] __x64_sys_epoll_wait+0x1e/0x30
  [<0>] do_syscall_64+0x57/0x190
  [<0>] entry_SYSCALL_64_after_hwframe+0x44/0xa9
  
  ubuntu@test1:/proc/1213429$ ll fd
  total 0
  dr-x------ 2 ubuntu ubuntu  0 Apr  1 17:55 ./
  dr-xr-xr-x 9 ubuntu ubuntu  0 Mar 30 21:33 ../
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 0 -> 'socket:[88567492]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 1 -> 'socket:[88567494]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 10 -> 'anon_inode:[eventpoll]'
  lr-x------ 1 ubuntu ubuntu 64 Apr  1 17:55 11 -> 'pipe:[88567502]'
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 12 -> 'pipe:[88567502]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 13 -> 'anon_inode:[eventfd]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 14 -> 'anon_inode:[eventpoll]'
  lr-x------ 1 ubuntu ubuntu 64 Apr  1 17:55 15 -> 'pipe:[88566591]'
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 16 -> 'pipe:[88566591]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 17 -> 'anon_inode:[eventfd]'
  lr-x------ 1 ubuntu ubuntu 64 Apr  1 17:55 18 -> /dev/null
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 19 -> 'socket:[88567693]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 2 -> 'socket:[88567496]'
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 20 -> /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/con-log-master-server-1.log
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 21 -> /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/rpc-log-master-server-1.log
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 22 -> /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/forward-log-master-server-1.log
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 23 -> /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/rpc-debug-master-server-1.log
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 24 -> /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/crash.log
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 25 -> /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/admin.log
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 26 -> /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/pomelo-master-server-1.log
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 27 -> /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/pomelo-admin.log
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 28 -> /home/ubuntu/zone1/wj_slgsanguo2_server/server/siege-server/logs/pomelo-rpc-master-server-1.log
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 29 -> 'socket:[88567694]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 3 -> 'socket:[88567498]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 30 -> 'socket:[88567710]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 31 -> 'socket:[88567696]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 32 -> 'socket:[88567711]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 33 -> 'socket:[88567698]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 34 -> 'socket:[88569152]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 4 -> 'anon_inode:[eventpoll]'
  lr-x------ 1 ubuntu ubuntu 64 Apr  1 17:55 5 -> 'pipe:[88567500]'
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 6 -> 'pipe:[88567500]'
  lr-x------ 1 ubuntu ubuntu 64 Apr  1 17:55 7 -> 'pipe:[88567501]'
  l-wx------ 1 ubuntu ubuntu 64 Apr  1 17:55 8 -> 'pipe:[88567501]'
  lrwx------ 1 ubuntu ubuntu 64 Apr  1 17:55 9 -> 'anon_inode:[eventfd]'
  
  ```

  

- https://www.zhihu.com/question/422144033

- `man 5 proc`命令会输出/proc目录的帮助信息，它里面包含/proc/[pid]目录中每个目录和文件的说明信息

- 举例声卡：open()打开声卡设备， read() 就是录音， write() 就是播放。



## 僵尸进程

- 僵尸进程是当子进程比父进程先结束，而父进程又没有回收子进程，释放子进程占用的资源，此时子进程将成为一个僵尸进程。如果父进程先退出 ，子进程被init接管，子进程退出后init会回收其占用的相关资源
- 一个进程在调用exit命令结束自己的生命的时候，其实它并没有真正的被销毁， 而是留下一个称为僵尸进程（Zombie）的数据结构（系统调用exit，它的作用是 使进程退出，但也仅仅限于将一个正常的进程变成一个僵尸进程，并不能将其完全销毁）
- 在UNIX 系统中，一个进程结束了，但是他的[父进程](https://baike.baidu.com/item/父进程/614062)没有等待(调用`wait / waitpid`)他， 那么他将变成一个僵尸进程。 但是如果该进程的父进程已经先结束了，那么该进程就不会变成僵尸进程， 因为每个进程结束的时候，系统都会扫描当前系统中所运行的所有进程， 看有没有哪个进程是刚刚结束的这个进程的子进程，如果是的话，就由Init 来接管他，成为他的父进程
- 危  害：导致系统不能产生新的进程
- 解决方式：可用`kill-SIGKILL`父进程ID来解决



## 网络传输

⽂件描述符的作⽤是什么？每⼀个进程都有⼀个数据结构 **task_struct** ，该结构体⾥有⼀个指向「⽂件描述符数组」的成员指针。该数组⾥列出这个进程打开的所有⽂件的⽂件描述符。数组的下标是⽂件描述符，是⼀个整数，⽽数组的内容是⼀个指针，指向内核中所有打开的⽂件的列表，也就是说内核可以通过⽂件描述符找到对应打开的⽂件。

然后每个⽂件都有⼀个 inode，Socket ⽂件的 inode 指向了内核中的 Socket 结构，在这个结构体⾥有两个队列，分别是发送队列和接收队列，这个两个队列⾥⾯保存的是⼀个个struct sk_buff ，⽤链表的组织形式串起来。

sk_buff 可以表示各个层的数据包，在应⽤层数据包叫 data，在 TCP 层我们称为 segment， 在 IP 层我们叫 packet，在数据链路层称为 frame。



- 为什么全部数据包只⽤⼀个结构体来描述呢？协议栈采⽤的是分层结构，上层向下层传递数据时需要增加包头，下层向上层数据时⼜需要去掉包头，如果每⼀层都⽤⼀个结构体，那在层之间传递数据的时候，就要发⽣多次拷⻉，这将⼤⼤降低 CPU 效率。

- 以太网是现实世界中最普遍的一种计算机网络。以太网有两类：第一类是经典以太网，第二类是[交换式以太网](https://baike.baidu.com/item/交换式以太网/188374)，使用了一种称为[交换机](https://baike.baidu.com/item/交换机/103532)的设备连接不同的计算机。经典以太网是以太网的原始形式，运行速度从3~10 Mbps不等；而交换式以太网正是广泛应用的以太网，可运行在100、1000和10000Mbps那样的高速率，分别以[快速以太网](https://baike.baidu.com/item/快速以太网/2796711)、千兆以太网和万兆以太网的形式呈现。 [1] 
- **进程的上下⽂切换**: 不仅包含了虚拟内存、栈、全局变量等⽤户空间的资源，还包括了内核堆栈、寄存器等内核空间的资源。



## CPU Cache 伪共享

CPU 从内存中读取数据到 Cache 的时候，并不是⼀个字节⼀个字节读取，⽽是⼀块⼀块的⽅式来读取数据的，这⼀块⼀块的数据被称为 CPU Line（缓存⾏），所以 **CPU Line** 是 **CPU**从内存读取数据到 **Cache** 的单位。

⾄于 CPU Line ⼤⼩，在 Linux 系统可以⽤下⾯的⽅式查看到，你可以看我服务器的 L1 Cache Line ⼤⼩是 64 字节，也就意味着 **L1 Cache** ⼀次载⼊数据的⼤⼩是 **64** 字节。



那么对数组的加载， CPU 就会加载数组⾥⾯连续的多个数据到 Cache ⾥，因此我们应该按照物理内存地址分布的顺序去访问元素，这样访问数组元素的时候，Cache 命中率就会很⾼，于是就能减少从内存读取数据的频率， 从⽽可提⾼程序的性能。



但是，在我们不使⽤数组，⽽是使⽤单独的变量的时候，则会有 Cache 伪共享的问题，Cache 伪共享问题上是⼀个性能杀⼿，我们应该要规避它。



现在假设有⼀个双核⼼的 CPU，这两个 CPU 核⼼并⾏运⾏着两个不同的线程，它们同时从内存中读取两个不同的数据，分别是类型为 long 的变量 A 和 B，这个两个数据的地址在物理内存上是连续的，如果 Cahce Line 的⼤⼩是 64 字节，并且变量 A 在 Cahce Line 的开头位置，那么这两个数据是位于同⼀个 **Cache Line** 中，⼜因为 CPU Line 是 CPU 从内存读取数据到 Cache 的单位，所以这两个数据会被同时读⼊到了两个 CPU 核⼼中各⾃ Cache 中。

```c
struct test {
  int a;
  int b __cacheline_aligned_in_smp;
}
```

这样 a 和 b 变量就不会在同⼀个 Cache Line 中了，如下图：

<img src="xiaolincoding.assets/image-20220407072618044.png" alt="image-20220407072618044" style="zoom: 25%;" />

因此，这种因为多个线程同时读写同⼀个 Cache Line 的不同变量时，⽽导致 CPU Cache 失效的现象称为伪共享（**False Sharing**）。



所谓的 Cache Line 伪共享问题就是，多个线程同时读写同⼀个 Cache Line 的不同变量时，⽽导致 CPUCache 失效的现象。那么对于多个线程共享的热点数据，即经常会修改的数据，应该避免这些数据刚好在同⼀个 Cache Line 中，避免的⽅式⼀般有 Cache Line ⼤⼩字节对⻬，以及字节填充等⽅法。



## CPU如何选择线程

在 Linux 内核中，进程和线程都是⽤ task_struct 结构体表示的，区别在于线程的 task_struct 结构体⾥部

分资源是共享了进程已创建的资源，⽐如内存地址空间、代码段、⽂件描述符等，所以 Linux 中的线程也

被称为轻量级进程，因为线程的 task_struct 相⽐进程的 task_struct 承载的 资源⽐较少，因此以「轻」得

名。

Linux 内核⾥的调度器，调度的对象就是 tark_struct ，接下来我们就把这个数据结构统称为任务。 

在 Linux 系统中，根据任务的优先级以及响应要求，主要分为两种，其中优先级的数值越⼩，优先级越

⾼：

实时任务，对系统的响应时间要求很⾼，也就是要尽可能快的执⾏实时任务，优先级在 0~99 范围

内的就算实时任务；

普通任务，响应时间没有很⾼的要求，优先级在 100~139 范围内都是普通任务级别；

<img src="xiaolincoding.assets/image-20220408204225748.png" alt="image-20220408204225748" style="zoom:50%;" />

如果我们启动任务的时候，没有特意去指定优先级的话，默认情况下都是普通任务，普通任务的调度类是

Fail，由 CFS 调度器来进⾏管理。CFS 调度器的⽬的是实现任务运⾏的公平性，也就是保障每个任务的运

⾏的时间是差不多的。

如果你想让某个普通任务有更多的执⾏时间，可以调整任务的 nice 值，从⽽让优先级⾼⼀些的任务执⾏

更多时间。nice 的值能设置的范围是 -20～19 ， 值越低，表明优先级越⾼，因此 -20 是最⾼优先级，19

则是最低优先级，默认优先级是 0。

是不是觉得 nice 值的范围很诡异？事实上，nice 值并不是表示优先级，⽽是表示优先级的修正数值，它与

优先级（priority）的关系是这样的：priority(new) = priority(old) + nice。内核中，priority 的范围是

0~139，值越低，优先级越⾼，其中前⾯的 0~99 范围是提供给实时任务使⽤的，⽽ nice 值是映射到

100~139，这个范围是提供给普通任务⽤的，因此 nice 值调整的是普通任务的优先级。

```bash
# 我们可以在启动任务的时候，可以指定 nice 的值，⽐如将 mysqld 以 -3 优先级：
nice -n -3 /usr/sbin/mysqld

# 重新调整nice值
renice -n 10 <pid>
```

调整为实时任务：

```bash
# 修改调度策略为SCHED_FIFO, 并且优先级为 1
chrt -f 1 -p 1996
```



## 中断

那 Linux 系统为了解决中断处理程序执⾏过⻓和中断丢失的问题，将中断过程分成了两个阶段，分别是

「上半部和下半部分」。上半部⽤来快速处理中断，⼀般会暂时关闭中断请求，主要负责处理跟硬件紧密相关或者时间敏感的

事情。

下半部⽤来延迟处理上半部未完成的⼯作，⼀般以「内核线程」的⽅式运⾏。



⽹卡收到⽹络包后，会通过硬件中断通知内核有新的数据到了，于是内核就会调⽤对应的中断处理程序来

响应该事件，这个事件的处理也是会分成上半部和下半部。

上部分要做到快速处理，所以只要把⽹卡的数据读到内存中，然后更新⼀下硬件寄存器的状态，⽐如把状

态更新为表示数据已经读到内存中的状态值。

接着，内核会触发⼀个软中断，把⼀些处理⽐较耗时且复杂的事情，交给「软中断处理程序」去做，也就

是中断的下半部，其主要是需要从内存中找到⽹络数据，再按照⽹络协议栈，对⽹络数据进⾏逐层解析和

处理，最后把数据送给应⽤程序。

所以，中断处理程序的上部分和下半部可以理解为：

上半部直接处理硬件请求，也就是硬中断，主要是负责耗时短的⼯作，特点是快速执⾏；

下半部是由内核触发，也就说软中断，主要是负责上半部未完成的⼯作，通常都是耗时⽐较⻓的事

情，特点是延迟执⾏；



## 小数

- ⼗进制数转⼆进制采⽤的是**`除 2 取余法`**

- ⼩数部分的转换不同于整数部分，它采⽤的是**`乘 2 取整法`**
- 负数在计算机中是以「补码」表示的，所谓的**`补码就是把正数的⼆进制全部取反再加 1`**，⽐如 -1 的⼆进制是把数字 1 的⼆进制取反后再加 1

![image-20220409063901034](xiaolincoding.assets/image-20220409063901034.png)

可以发现， 0.1 的⼆进制表示是⽆限循环的。

`由于计算机的资源是有限的，所以是没办法⽤⼆进制精确的表示 **0.1**，只能⽤「近似值」来表示，就是在有`

`限的精度情况下，最⼤化接近 **0.1** 的⼆进制数，于是就会造成精度缺失的情况。`



对于⼆进制⼩数转⼗进制时，需要注意⼀点，⼩数点后⾯的指数幂是`负数`。

⽐如，`⼆进制 0.1` 转成⼗进制就是 `2^(-1)` ，也就是`⼗进制 0.5` ，⼆进制 0.01 转成⼗进制就是

2^-2 ，也就是⼗进制 0.25 ，以此类推。

![image-20220409064706848](xiaolincoding.assets/image-20220409064706848.png)



1000.101 这种⼆进制⼩数是「定点数」形式，代表着⼩数点是定死的，不能移动，如果你移动了它的⼩数点，这个数就变了， 就不再是它原来的值了。

然⽽，计算机并不是这样存储的⼩数的，计算机存储⼩数的采⽤的是`浮点数`，名字⾥的「浮点」表示⼩数点是可以浮动的。

⽐如 1000.101 这个⼆进制数，可以表示成 `1.000101 x 2^3` ，类似于数学上的科学记数法。



该⽅法在⼩数点左边只有⼀个数字，⽽且把这种整数部分没有前导 0 的数字称为`规格化`，⽐如 **1.0 x10^(-9)** 是规格化的科学记数法，⽽ **0.1 x 10^(-9)** 和 **10.0 x 10^(-9)** 就不是了。



所以通常将 **1000.101** 这种⼆进制数，规格化表示成 **1.000101 x 2^3** ，其中，最为关键的是 000101 和

3 这两个东⻄，它就可以包含了这个⼆进制⼩数的所有信息：

- **000101** 称为`尾数`，即⼩数点后⾯的数字；

- **3** 称为`指数`，指定了⼩数点在数据中的位置；

```
符号位|指数位|尾数
```

- `符号位`：表示数字是正数还是负数，为 0 表示正数，为 1 表示负数；

- `指数位`：指定了⼩数点在数据中的位置，指数可以是负数，也可以是正数，`指数位的⻓度越⻓则数值的表达范围就越⼤`；

- `尾数位`：⼩数点右侧的数字，也就是⼩数部分，⽐如⼆进制 1.0011 x 2^(-2)，尾数部分就是 0011，`⽽且尾数的⻓度决定了这个数的精度`，因此如果要表示精度更⾼的⼩数，则就要提⾼尾数位的⻓度；



⽤ 32 位来表示的浮点数，则称为`单精度浮点数`，也就是我们编程语⾔中的 float 变量，⽽⽤ 64 位来表示的浮点数，称为`双精度浮点数`，也就是 double 变量，

<img src="xiaolincoding.assets/image-20220409071310354.png" alt="image-20220409071310354" style="zoom:50%;" />





## NAT

`NAT（Network Address Translation）`，是指网络地址转换，1994年提出的。

当在[专用网](https://baike.baidu.com/item/专用网/1006818)内部的一些主机本来已经分配到了本地IP地址（即仅在本专用网内使用的专用地址），但又想和因特网上的主机通信（并不需要加密）时，可使用NAT方法。

这种方法需要在专用网（私网IP）连接到因特网（公网IP）的路由器上安装NAT软件。**装有NAT软件的路由器叫做NAT路由器，它至少有一个有效的外部全球IP地址（公网IP地址）**。这样，所有使用本地地址（私网IP地址）的主机在和外界通信时，都要在NAT路由器上将其本地地址转换成全球IP地址，才能和因特网连接。



借助于NAT，私有（保留）地址的"内部"网络通过[路由器](https://baike.baidu.com/item/路由器)发送[数据包](https://baike.baidu.com/item/数据包)时，[私有地址](https://baike.baidu.com/item/私有地址)被转换成合法的IP地址，一个局域网只需使用少量IP地址（甚至是1个）即可实现私有地址网络内所有计算机与Internet的通信需求。



NAT将自动修改IP[报文](https://baike.baidu.com/item/报文)的源IP地址和目的IP地址，Ip地址校验则在NAT处理过程中自动完成。有些应用程序将源IP地址嵌入到IP[报文](https://baike.baidu.com/item/报文)的数据部分中，所以还需要同时对报文的数据部分进行修改，以匹配IP头中已经修改过的源IP地址。否则，在[报文](https://baike.baidu.com/item/报文)数据部分嵌入IP地址的应用程序就不能正常工作。



## 虚拟内存

对于单页表的实现方式，在 32 位和页大小 `4KB` 的环境下，一个进程的页表需要装下 100 多万个「页表项」，并且每个页表项是占用 4 字节大小的，于是相当于每个页表需占用 4MB 大小的空间。

我们把这个 100 多万个「页表项」的单级页表再分页，将页表（一级页表）分为 `1024` 个页表（二级页表），每个表（二级页表）中包含 `1024` 个「页表项」，形成**二级分页**。如下图所示：

![](xiaolincoding.assets/13-二级页表.jpeg)

- 分了二级表，映射 4GB 地址空间就需要 4KB（一级页表）+ 4MB（二级页表）的内存

- 如果使用了二级分页，一级页表就可以覆盖整个 4GB 虚拟地址空间，但**如果某个一级页表的页表项没有被用到，也就不需要创建这个页表项对应的二级页表了，即可以在需要时才创建二级页表**。

  > 做个简单的计算，假设只有 20% 的一级页表项被用到了，那么页表占用的内存空间就只有 4KB（一级页表） + 20% * 4MB（二级页表）= `0.804MB`，这对比单级页表的 `4MB` 是不是一个巨大的节约？

- 对于 64 位的系统，两级分页肯定不够了，就变成了四级目录，分别是：

  - 全局页目录项 PGD（*Page Global Directory*）；
  - 上层页目录项 PUD（*Page Upper Directory*）；
  - 中间页目录项 PMD（*Page Middle Directory*）；
  - 页表项 PTE（*Page Table Entry*）；

- 把最常访问的几个页表项存储到访问速度更快的硬件，于是计算机科学家们，就在 CPU 芯片中，加入了一个专门存放程序最常访问的页表项的 Cache，这个 Cache 就是 `TLB`（*Translation Lookaside Buffer*） ，通常称为页表缓存、转址旁路缓存、快表等。

![地址转换](xiaolincoding.assets/16-TLB.jpg)



## Linux内存管理

**页式内存管理的作用是在由段式内存管理所映射而成的地址上再加上一层地址映射。**

**Linux 内存主要采用的是页式内存管理，但同时也不可避免地涉及了段机制**。

> **Linux 系统中的每个段都是从 0 地址开始的整个 4GB 虚拟空间（32 位环境下），也就是所有的段的起始地址都是一样的。这意味着，Linux 系统中的代码，包括操作系统本身的代码和应用程序代码，所面对的地址空间都是线性地址空间（虚拟地址），这种做法相当于屏蔽了处理器中的逻辑地址概念，段只被用于访问控制和内存保护。**



### 用户空间与内存空间

![用户空间与内存空间](xiaolincoding.assets/20-内核空间和用户空间.jpg)

内核空间与用户空间的区别：

- 进程在用户态时，只能访问用户空间内存；

- 只有进入内核态后，才可以访问内核空间的内存；

  

### 每个进程的内核空间都是一致的

![每个进程的内核空间都是一致的](xiaolincoding.assets/21-多个进程共享内核空间.jpg)

### 虚拟内存空间划分

![虚拟内存空间划分](xiaolincoding.assets/22-进程空间结构.jpg)

通过这张图你可以看到，用户空间内存，从**低到高**分别是 6 种不同的内存段：

- 程序文件段（.text），包括二进制可执行代码；
- 已初始化数据段（.data），包括静态常量；
- 未初始化数据段（.bss），包括未初始化的静态变量；
- 堆段，包括动态分配的内存，从低地址开始向上增长；
- 文件映射段，包括动态库、共享内存等，从低地址开始向上增长（[跟硬件和内核版本有关 (opens new window)](http://lishiwen4.github.io/linux/linux-process-memory-location)）；
- 栈段，包括局部变量和函数调用的上下文等。栈的大小是固定的，一般是 `8 MB`。当然系统也提供了参数，以便我们自定义大小；

在这 7 个内存段中，堆和文件映射段的内存是动态分配的。比如说，使用 C 标准库的 `malloc()` 或者 `mmap()` ，就可以分别在堆和文件映射段动态分配内存。



## malloc

- malloc() 并不是系统调用，而是 C 库里的函数，用于动态分配内存。

- malloc 申请内存的时候，会有两种方式向操作系统申请堆内存。

  - 方式一：通过 brk() 系统调用从堆分配内存
  - 方式二：通过 mmap() 系统调用在文件映射区域分配内存；

- **malloc() 分配的是虚拟内存**。

- malloc() 源码里默认定义了一个阈值：

  - 如果用户分配的内存小于 128 KB，则通过 brk() 申请内存；
  - 如果用户分配的内存大于 128 KB，则通过 mmap() 申请内存；

  注意，不同的 glibc 版本定义的阈值也是不同的。

![图片](xiaolincoding.assets/0dd0e2c1eb32b8b7cabfb95392a36f82.png)



![图片](xiaolincoding.assets/f8425aa73ca7e5ac8e3a46c2e3eb9188.png)

内存释放：

![图片](xiaolincoding.assets/cb6e3ce4532ff0a6bfd60fe3e52a806e.png)

- malloc 返回给用户态的内存起始地址比进程的堆空间起始地址多了 16 字节

- 这样当执行 free() 函数时，free 会对传入进来的内存地址向左偏移 16 字节，然后从这个 16 字节的分析出当前的内存块的大小，自然就知道要释放多大的内存了。



## 进程与线程

### 并发与并行

![并发与并行](xiaolincoding.assets/5-并发与并行.jpg)

### 进程状态

- **描述进程没有占用实际的物理内存空间的情况，这个状态就是挂起状态。**这跟阻塞状态是不一样，阻塞状态是等待某个事件的返回。
- 另外，挂起状态可以分为两种：
  - 阻塞挂起状态：进程在外存（硬盘）并等待某个事件的出现；
  - 就绪挂起状态：进程在外存（硬盘），但只要进入内存，即刻立刻运行；
- 导致进程挂起的原因不只是因为进程所使用的内存空间不在物理内存，还包括如下情况：
  - 通过 sleep 让进程间歇性挂起，其工作原理是设置一个定时器，到期后唤醒进程。
  - 用户希望挂起一个程序的执行，比如在 Linux 中用 `Ctrl+Z` 挂起进程；

![七种状态变迁](xiaolincoding.assets/10-进程七中状态.jpg)



> 每个 PCB 是如何组织的呢？

通常是通过**链表**的方式进行组织，把具有**相同状态的进程链在一起，组成各种队列**。比如：

- 将所有处于就绪状态的进程链在一起，称为**就绪队列**；
- 把所有因等待某事件而处于等待状态的进程链在一起就组成各种**阻塞队列**；
- 另外，对于运行队列在单核 CPU 系统中则只有一个运行指针了，因为单核 CPU 在某个时间，只能运行一个程序。

![就绪队列和阻塞队列](xiaolincoding.assets/12-PCB状态链表组织.jpg)



- 进程是由内核管理和调度的，所以进程的切换只能发生在内核态。

- 所以，**进程的上下文切换不仅包含了虚拟内存、栈、全局变量等用户空间的资源，还包括了内核堆栈、寄存器等内核空间的资源。**



### 线程

- **线程是进程当中的一条执行流程。**

- 同一个进程内多个线程之间可以共享代码段、数据段、打开的文件等资源，但每个线程各自都有一套独立的寄存器和栈，这样可以确保线程的控制流是相对独立的。
- 线程与进程最大的区别在于：**线程是调度的基本单位，而进程则是资源拥有的基本单位**。

![多线程](xiaolincoding.assets/16-多线程内存结构.jpg)

- 线程栈：是个线程独有的，保存其运行状态和局部自动变量的。 栈在线程开始的时候初始化，每个线程的栈互相独立，因此，栈是 thread safe的 。 操作系统在切换线程的时候会自动的切换栈 ，就是切换 ＳＳ／ＥＳＰ寄存器。 栈空间不需要在高级语言里面显式的分配和释放。
- 

![img](xiaolincoding.assets/28541347_1589621708V9Nh.png)



用户线程是基于用户态的线程管理库来实现的，那么**线程控制块（\*Thread Control Block, TCB\*）** 也是在库里面来实现的，对于操作系统而言是看不到这个 TCB 的，它只能看到**整个进程的 PCB**。

所以，**用户线程的整个线程管理和调度，操作系统是不直接参与的，而是由用户级线程库函数来完成线程的管理，包括线程的创建、终止、同步和调度等。**

![用户级线程模型](xiaolincoding.assets/20-线程PCB-一对多关系.jpg)

- 当一个线程开始运行后，**除非它主动地交出 CPU 的使用权**，否则它所在的进程当中的其他线程无法运行，因为用户态的线程没法打断当前运行中的线程，它没有这个特权，只有操作系统才有，但是用户线程不是由操作系统管理的。



- **抢占式调度算法**挑选一个进程，然后让该进程只运行某段时间，如果在该时段结束时，该进程仍然在运行时，则会把它挂起，接着调度程序从就绪队列挑选另外一个进程。这种抢占式调度处理，**需要在时间间隔的末端发生时钟中断**，**以便把 CPU 控制返回给调度程序进行调度**，也就是常说的**时间片机制**。
- **为了提高 CPU 利用率，在这种发送 `I/O 事件`致使 CPU 空闲的情况下，调度程序需要从就绪队列中选择一个进程来运行。**



## 进程间通信

### 管道

- **管道传输数据是单向的**

- 匿名管道

  ```bash
  ps auxf | grep mysql
  ```

- 管道还有另外一个类型是**命名管道**，也被叫做 `FIFO`，因为数据是先进先出的传输方式。

  ```bash
  $ mkfifo myPipe
  $ ls -l
  prw-r--r--. 1 root    root         0 Jul 17 02:45 myPipe
  
  # 将数据写进管道
  $ echo "hello" > myPipe 
  # 停住了..
  
  # 读取管道里的数据
  $ cat < myPipe
  hello
  ```

  

- **所谓的管道，就是内核里面的一串缓存**。从管道的一段写入的数据，实际上是缓存在内核中的，另一端读取，也就是从内核中读取这段数据。另外，管道传输的数据是无格式的流且大小受限。

- 我们可以使用 `fork` 创建子进程，**创建的子进程会复制父进程的文件描述符**，这样就做到了两个进程各有两个「 `fd[0]` 与 `fd[1]`」，两个进程就可以通过各自的 fd 写入和读取同一个管道文件实现跨进程通信了。

  - 父进程关闭读取的 fd[0]，只保留写入的 fd[1]；
  - 子进程关闭写入的 fd[1]，只保留读取的 fd[0]；
  - 所以说如果需要双向通信，则应该创建两个管道。

  ![img](xiaolincoding.assets/6-管道-pipe-fork.jpg)



- 在 shell 里面执行 `A | B`命令的时候，A 进程和 B 进程都是 shell 创建出来的子进程，A 和 B 之间不存在父子关系，它俩的父进程都是 shell。

  ![img](xiaolincoding.assets/8-管道-pipe-shell.jpg)

- 所以说，在 shell 里通过「`|`」匿名管道将多个命令连接在一起，实际上也就是创建了多个子进程
- 不管是匿名管道还是命名管道，进程写入的数据都是缓存在内核中，另一个进程读取数据时候自然也是从内核中获取，同时通信数据都遵循**先进先出**原则，不支持 lseek 之类的文件定位操作。



### **消息队列**

- **消息队列是保存在内核中的消息链表**，在发送数据时，会分成一个一个独立的数据单元，也就是消息体（数据块），消息体是用户自定义的数据类型，消息的发送方和接收方要约定好消息体的数据类型，所以每个消息体都是固定大小的存储块，不像管道是无格式的字节流数据。
- 消息队列生命周期随内核，如果没有释放消息队列或者没有关闭操作系统，消息队列会一直存在，而前面提到的匿名管道的生命周期，是随进程的创建而建立，随进程的结束而销毁。
- *消息队列通信过程中，存在用户态与内核态之间的数据拷贝开销**



### 共享内存

- **共享内存的机制，就是拿出一块虚拟地址空间来，映射到相同的物理内存中**。这样这个进程写入的东西，另外一个进程马上就能看到了，都不需要拷贝来拷贝去，传来传去，大大提高了进程间通信的速度。

  ![img](xiaolincoding.assets/9-共享内存.jpg)



### 信号量

- **信号量其实是一个整型的计数器，主要用于实现进程间的互斥与同步，而不是用于缓存进程间通信的数据**。

- 信号量表示资源的数量，控制信号量的方式有两种原子操作：

  - 一个是 **P 操作**，这个操作会把信号量减去 1，相减后如果信号量 < 0，则表明资源已被占用，进程需阻塞等待；相减后如果信号量 >= 0，则表明还有资源可使用，进程可正常继续执行。
  - 另一个是 **V 操作**，这个操作会把信号量加上 1，相加后如果信号量 <= 0，则表明当前有阻塞中的进程，于是会将该进程唤醒运行；相加后如果信号量 > 0，则表明当前没有阻塞中的进程；

- 可以发现，信号初始化为 `1`，就代表着是**互斥信号量**，它可以保证共享内存在任何时刻只有一个进程在访问，这就很好的保护了共享内存。

  ![img](xiaolincoding.assets/10-信号量-互斥.jpg)

- 可以发现，信号初始化为 `0`，就代表着是**同步信号量**，它可以保证进程 A 应在进程 B 之前执行。

  ![img](xiaolincoding.assets/11-信号量-同步.jpg)

### Socket

针对 UDP 协议通信的 socket 编程模型：

![img](xiaolincoding.assets/13-UDP编程模型.jpg)

- UDP 是没有连接的，所以不需要三次握手，也就不需要像 TCP 调用 listen 和 connect，但是 UDP 的交互仍然需要 IP 地址和端口号，因此也需要 bind。
- 另外，每次通信时，调用 sendto 和 recvfrom，都要传入目标主机的 IP 地址和端口。

```c
int socket(int domain, int type, int protocal)
```

- domain 参数用来指定协议族，比如 AF_INET 用于 IPV4、AF_INET6 用于 IPV6、AF_LOCAL/AF_UNIX 用于本机；
- type 参数用来指定通信特性，比如 SOCK_STREAM 表示的是字节流，对应 TCP、SOCK_DGRAM 表示的是数据报，对应 UDP、SOCK_RAW 表示的是原始套接字；
- protocal 参数原本是用来指定通信协议的，但现在基本废弃。因为协议已经通过前面两个参数指定完成，protocol 目前一般写成 0 即可；

本地字节流 socket 和 本地数据报 socket 在 bind 的时候，不像 TCP 和 UDP 要绑定 IP 地址和端口，而是**绑定一个本地文件**，这也就是它们之间的最大区别。



## 利用工具排查死锁问题

如果你想排查你的 Java 程序是否死锁，则可以使用 `jstack` 工具，它是 jdk 自带的线程堆栈分析工具。

由于小林的死锁代码例子是 C 写的，在 Linux 下，我们可以使用 `pstack` + `gdb` 工具来定位死锁问题。

pstack 命令可以显示每个线程的栈跟踪信息（函数调用过程），它的使用方式也很简单，只需要 `pstack <pid>` 就可以了。

那么，在定位死锁问题时，我们可以多次执行 pstack 命令查看线程的函数调用过程，多次对比结果，确认哪几个线程一直没有变化，且是因为在等待锁，那么大概率是由于死锁问题导致的。

```bash
$ pstack 87746
Thread 3 (Thread 0x7f60a610a700 (LWP 87747)):
#0  0x0000003720e0da1d in __lll_lock_wait () from /lib64/libpthread.so.0
#1  0x0000003720e093ca in _L_lock_829 () from /lib64/libpthread.so.0
#2  0x0000003720e09298 in pthread_mutex_lock () from /lib64/libpthread.so.0
#3  0x0000000000400725 in threadA_proc ()
#4  0x0000003720e07893 in start_thread () from /lib64/libpthread.so.0
#5  0x00000037206f4bfd in clone () from /lib64/libc.so.6
Thread 2 (Thread 0x7f60a5709700 (LWP 87748)):
#0  0x0000003720e0da1d in __lll_lock_wait () from /lib64/libpthread.so.0
#1  0x0000003720e093ca in _L_lock_829 () from /lib64/libpthread.so.0
#2  0x0000003720e09298 in pthread_mutex_lock () from /lib64/libpthread.so.0
#3  0x0000000000400792 in threadB_proc ()
#4  0x0000003720e07893 in start_thread () from /lib64/libpthread.so.0
#5  0x00000037206f4bfd in clone () from /lib64/libc.so.6
Thread 1 (Thread 0x7f60a610c700 (LWP 87746)):
#0  0x0000003720e080e5 in pthread_join () from /lib64/libpthread.so.0
#1  0x0000000000400806 in main ()

....

$ pstack 87746
Thread 3 (Thread 0x7f60a610a700 (LWP 87747)):
#0  0x0000003720e0da1d in __lll_lock_wait () from /lib64/libpthread.so.0
#1  0x0000003720e093ca in _L_lock_829 () from /lib64/libpthread.so.0
#2  0x0000003720e09298 in pthread_mutex_lock () from /lib64/libpthread.so.0
#3  0x0000000000400725 in threadA_proc ()
#4  0x0000003720e07893 in start_thread () from /lib64/libpthread.so.0
#5  0x00000037206f4bfd in clone () from /lib64/libc.so.6
Thread 2 (Thread 0x7f60a5709700 (LWP 87748)):
#0  0x0000003720e0da1d in __lll_lock_wait () from /lib64/libpthread.so.0
#1  0x0000003720e093ca in _L_lock_829 () from /lib64/libpthread.so.0
#2  0x0000003720e09298 in pthread_mutex_lock () from /lib64/libpthread.so.0
#3  0x0000000000400792 in threadB_proc ()
#4  0x0000003720e07893 in start_thread () from /lib64/libpthread.so.0
#5  0x00000037206f4bfd in clone () from /lib64/libc.so.6
Thread 1 (Thread 0x7f60a610c700 (LWP 87746)):
#0  0x0000003720e080e5 in pthread_join () from /lib64/libpthread.so.0
#1  0x0000000000400806 in main ()
```

可以看到，Thread 2 和 Thread 3 一直阻塞获取锁（*pthread_mutex_lock*）的过程，而且 pstack 多次输出信息都没有变化，那么可能大概率发生了死锁。



整个 gdb 调试过程，如下：

```bash
// gdb 命令
$ gdb -p 87746

// 打印所有的线程信息
(gdb) info thread
  3 Thread 0x7f60a610a700 (LWP 87747)  0x0000003720e0da1d in __lll_lock_wait () from /lib64/libpthread.so.0
  2 Thread 0x7f60a5709700 (LWP 87748)  0x0000003720e0da1d in __lll_lock_wait () from /lib64/libpthread.so.0
* 1 Thread 0x7f60a610c700 (LWP 87746)  0x0000003720e080e5 in pthread_join () from /lib64/libpthread.so.0
//最左边的 * 表示 gdb 锁定的线程，切换到第二个线程去查看

// 切换到第2个线程
(gdb) thread 2
[Switching to thread 2 (Thread 0x7f60a5709700 (LWP 87748))]#0  0x0000003720e0da1d in __lll_lock_wait () from /lib64/libpthread.so.0 

// bt 可以打印函数堆栈，却无法看到函数参数，跟 pstack 命令一样 
(gdb) bt
#0  0x0000003720e0da1d in __lll_lock_wait () from /lib64/libpthread.so.0
#1  0x0000003720e093ca in _L_lock_829 () from /lib64/libpthread.so.0
#2  0x0000003720e09298 in pthread_mutex_lock () from /lib64/libpthread.so.0
#3  0x0000000000400792 in threadB_proc (data=0x0) at dead_lock.c:25
#4  0x0000003720e07893 in start_thread () from /lib64/libpthread.so.0
#5  0x00000037206f4bfd in clone () from /lib64/libc.so.6

// 打印第三帧信息，每次函数调用都会有压栈的过程，而 frame 则记录栈中的帧信息
(gdb) frame 3
#3  0x0000000000400792 in threadB_proc (data=0x0) at dead_lock.c:25
27    printf("thread B waiting get ResourceA \n");
28    pthread_mutex_lock(&mutex_A);

// 打印mutex_A的值 ,  __owner表示gdb中标示线程的值，即LWP
(gdb) p mutex_A
$1 = {__data = {__lock = 2, __count = 0, __owner = 87747, __nusers = 1, __kind = 0, __spins = 0, __list = {__prev = 0x0, __next = 0x0}}, 
  __size = "\002\000\000\000\000\000\000\000\303V\001\000\001", '\000' <repeats 26 times>, __align = 2}

// 打印mutex_B的值 ,  __owner表示gdb中标示线程的值，即LWP
(gdb) p mutex_B
$2 = {__data = {__lock = 2, __count = 0, __owner = 87748, __nusers = 1, __kind = 0, __spins = 0, __list = {__prev = 0x0, __next = 0x0}}, 
  __size = "\002\000\000\000\000\000\000\000\304V\001\000\001", '\000' <repeats 26 times>, __align = 2}  
```

解释下，上面的调试过程（LWP：Light Weight Process 轻量级进程）：

1. 通过 `info thread` 打印了所有的线程信息，可以看到有 3 个线程，一个是主线程（LWP 87746），另外两个都是我们自己创建的线程（LWP 87747 和 87748）；
2. 通过 `thread 2`，将切换到第 2 个线程（LWP 87748）；
3. 通过 `bt`，打印线程的调用栈信息，可以看到有 threadB_proc 函数，说明这个是线程 B 函数，也就说 LWP 87748 是线程 B;
4. 通过 `frame 3`，打印调用栈中的第三个帧的信息，可以看到线程 B 函数，在获取互斥锁 A 的时候阻塞了；
5. 通过 `p mutex_A`，打印互斥锁 A 对象信息，可以看到它被 LWP 为 87747（线程 A） 的线程持有着；
6. 通过 `p mutex_B`，打印互斥锁 A 对象信息，可以看到他被 LWP 为 87748 （线程 B） 的线程持有着；

因为线程 B 在等待线程 A 所持有的 mutex_A, 而同时线程 A 又在等待线程 B 所拥有的mutex_B, 所以可以断定该程序发生了死锁。

> 轻量级进程 (LWP, light weight process) 是一种由[内核](https://baike.baidu.com/item/内核)支持的用户线程。它是基于[内核](https://baike.baidu.com/item/内核)线程的高级抽象，因此只有先支持[内核](https://baike.baidu.com/item/内核)线程，才能有 LWP 。每一个LWP可以支持一个或多个用户线程，每个 LWP 由一个[内核](https://baike.baidu.com/item/内核)线程支持。内核线程与LWP之间的模型实际上就是《[操作系统概念](https://baike.baidu.com/item/操作系统概念/8259031)》上所提到的一对一线程模型。在这种实现的操作系统中， LWP 相当于[用户线程](https://baike.baidu.com/item/用户线程)。 由于每个 LWP 都与一个特定的内核线程关联，因此每个 LWP 都是一个独立的[线程调度](https://baike.baidu.com/item/线程调度)单元。即使有一个 LWP 在[系统调用](https://baike.baidu.com/item/系统调用)中阻塞，也不会影响整个进程的执行。



死锁问题的产生是由两个或者以上线程并行执行的时候，争夺资源而互相等待造成的。

死锁只有同时满足互斥、持有并等待、不可剥夺、环路等待这四个条件的时候才会发生。

所以要避免死锁问题，就是要破坏其中一个条件即可，最常用的方法就是使用资源有序分配法来破坏环路等待条件。



## 锁

**对于互斥锁加锁失败而阻塞的现象，是由操作系统内核实现的**。当加锁失败时，内核会将线程置为「睡眠」状态，等到锁被释放后，内核会在合适的时机唤醒线程，当这个线程成功获取到锁后，于是就可以继续执行。

![img](xiaolin-os.assets/互斥锁工作流程.png)

所以，互斥锁加锁失败时，会从用户态陷入到内核态，让内核帮我们切换线程，虽然简化了使用锁的难度，但是存在一定的性能开销成本。

那这个开销成本是什么呢？会有**两次线程上下文切换的成本**：

- 当线程加锁失败时，内核会把线程的状态从「运行」状态设置为「睡眠」状态，然后把 CPU 切换给其他线程运行；
- 接着，当锁被释放时，之前「睡眠」状态的线程会变为「就绪」状态，然后内核会在合适的时间，把 CPU 切换给该线程运行。



线程的上下文切换的是什么？当两个线程是属于同一个进程，**因为虚拟内存是共享的，所以在切换时，虚拟内存这些资源就保持不动，只需要切换线程的私有数据、寄存器等不共享的数据。**



**上下切换的耗时**有大佬统计过，大概在几十纳秒到几微秒之间，如果你锁住的代码执行时间比较短，那可能上下文切换的时间都比你锁住的代码执行时间还要长。



所以，**如果你能确定被锁住的代码执行时间很短，就不应该用互斥锁，而应该选用自旋锁。否则使用互斥锁。**



自旋锁是通过 CPU 提供的 `CAS` 函数（*Compare And Swap*），在「用户态」完成加锁和解锁操作，不会主动产生线程上下文切换，所以相比互斥锁来说，会快一些，开销也小一些。



比如，设锁为变量 lock，整数 0 表示锁是空闲状态，整数 `pid 表示线程 ID`，那么 `CAS(lock, 0, pid)` 就表示自旋锁的加锁操作，`CAS(lock, pid, 0)` 则表示解锁操作。



使用自旋锁的时候，当发生多线程竞争锁的情况，加锁失败的线程会「忙等待」，直到它拿到锁。这里的「忙等待」可以用 `while` 循环等待实现，不过最好是使用 CPU 提供的 `PAUSE` 指令来实现「忙等待」，因为可以减少循环等待时的耗电量。



自旋锁是最比较简单的一种锁，一直自旋，利用 CPU 周期，直到锁可用。**需要注意，在单核 CPU 上，需要抢占式的调度器（即不断通过时钟中断一个线程，运行其他线程）。否则，自旋锁在单 CPU 上无法使用，因为一个自旋的线程永远不会放弃 CPU。**



自旋锁开销少，在多核系统下一般不会主动产生线程切换，适合异步、协程等在用户态切换请求的编程方式，但如果被锁住的代码执行时间过长，自旋的线程会长时间占用 CPU 资源，所以自旋的时间和被锁住的代码执行的时间是成「正比」的关系，我们需要清楚的知道这一点。



自旋锁与互斥锁使用层面比较相似，但实现层面上完全不同：**当加锁失败时，互斥锁用「线程切换」来应对，自旋锁则用「忙等待」来应对**。



它俩是锁的最基本处理方式，更高级的锁都会选择其中一个来实现，比如读写锁既可以选择互斥锁实现，也可以基于自旋锁实现。



### 读写锁

而写优先锁是优先服务写线程，其工作方式是：当读线程 A 先持有了读锁，写线程 B 在获取写锁的时候，会被阻塞，并且在阻塞过程中，后续来的读线程 C 获取读锁时会失败，于是读线程 C 将被阻塞在获取读锁的操作，这样只要读线程 A 释放读锁后，写线程 B 就可以成功获取读锁。

读优先锁对于读线程并发性更好，但也不是没有问题。我们试想一下，如果一直有读线程获取读锁，那么写线程将永远获取不到写锁，这就造成了写线程「饥饿」的现象。

写优先锁可以保证写线程不会饿死，但是如果一直有写线程获取写锁，读线程也会被「饿死」。

**公平读写锁比较简单的一种方式是：用队列把获取锁的线程排队，不管是写线程还是读线程都按照先进先出的原则加锁即可，这样读线程仍然可以并发，也不会出现「饥饿」的现象。**



### 乐观锁与悲观锁

悲观锁做事比较悲观，它认为**多线程同时修改共享资源的概率比较高，于是很容易出现冲突，所以访问共享资源前，先要上锁**。

乐观锁做事比较乐观，它假定冲突的概率很低，它的工作方式是：**先修改完共享资源，再验证这段时间内有没有发生冲突，如果没有其他线程在修改资源，那么操作完成，如果发现有其他线程已经修改过这个资源，就放弃本次操作**。

放弃后如何重试，这跟业务场景息息相关，虽然重试的成本很高，但是冲突的概率足够低的话，还是可以接受的。

**乐观锁全程并没有加锁，所以它也叫无锁编程**。

乐观锁使用场景：

- 在线文档
- SVN
- git

乐观锁虽然去除了加锁解锁的操作，但是一旦发生冲突，重试的成本非常高，所以**只有在冲突概率非常低，且加锁成本非常高的场景时，才考虑使用乐观锁。**





## 调度算法

### 缺页中断

![缺页中断的处理流程](xiaolin-os.assets/缺页异常流程.png)

### 页表项的内部字段

![img](xiaolin-os.assets/页表项字段.png)

- *状态位*：用于表示该页是否有效，也就是说是否在物理内存中，供程序访问时参考。
- *访问字段*：用于记录该页在一段时间被访问的次数，供页面置换算法选择出页面时参考。
- *修改位*：表示该页在调入内存后是否有被修改过，由于内存中的每一页都在磁盘上保留一份副本，因此，如果没有修改，在置换该页时就不需要将该页写回到磁盘上，以减少系统的开销；如果已经被修改，则将该页重写到磁盘上，以保证磁盘中所保留的始终是最新的副本。
- *硬盘地址*：用于指出该页在硬盘上的地址，通常是物理块号，供调入该页时使用。



### 虚拟内存的管理流程

![虚拟内存的流程](xiaolin-os.assets/虚拟内存管理流程.png)



### 磁盘调度

![磁盘的结构](xiaolin-os.assets/磁盘结构.jpg)

常见的机械磁盘是上图：

左边的样子，中间圆的部分是磁盘的盘片，一般会有多个盘片，每个盘面都有自己的磁头。

右边的图就是一个盘片的结构，盘片中的每一层分为多个磁道，每个磁道分多个扇区，每个扇区是 `512` 字节。

那么，多个具有相同编号的磁道形成一个圆柱，称之为磁盘的柱面，如上图里中间的样子。



## 文件系统

Linux 最经典的一句话是：「**一切皆文件**」，不仅普通的文件和目录，就连块设备、管道、socket 等，也都是统一交给文件系统管理的。

Linux 文件系统会为每个文件分配两个数据结构：**索引节点（\*index node\*）和目录项（\*directory entry\*）**，它们主要用来记录文件的元信息和目录层次结构。

- 索引节点，也就是 *inode*，`用来记录文件的元信息`，比如 inode 编号、文件大小、访问权限、创建时间、修改时间、**数据在磁盘的位置**等等。索引节点是文件的**唯一**标识，它们之间一一对应，也同样都会被存储在硬盘中，所以**索引节点同样占用磁盘空间**。
- 目录项，也就是 *dentry*，`用来记录文件的名字、**索引节点指针**以及与其他目录项的层级关联关系。`多个目录项关联起来，就会形成目录结构，但它与索引节点不同的是，**目录项是由内核维护的一个数据结构，不存放于磁盘，而是缓存在内存**。



如果查询目录频繁从磁盘读，效率会很低，所以**内核会把已经读过的目录用目录项这个数据结构缓存在内存**，下次再次读到相同的目录时，只需从内存读就可以，大大提高了文件系统的效率。



注意，目录项这个数据结构不只是表示目录，也是可以表示文件的。

磁盘读写的最小单位是**扇区**，扇区的大小只有 `512B` 大小，很明显，如果每次读写都以这么小为单位，那这读写的效率会非常低。

所以，文件系统把多个扇区组成了一个**逻辑块**，每次读写的最小单位就是逻辑块（数据块），Linux 中的逻辑块大小为 `4KB`，也就是一次性读写 8 个扇区，这将大大提高了磁盘的读写的效率。



![img](xiaolin-os.assets/目录项+索引节点+超级块+数据块 (1).png)

索引节点是存储在硬盘上的数据，那么为了加速文件的访问，通常会把索引节点加载到内存中。

另外，磁盘进行格式化的时候，会被分成三个存储区域，分别是超级块、索引节点区和数据块区。

- *超级块*，用来存储文件系统的详细信息，比如块个数、块大小、空闲块等等。
- *索引节点区*，用来存储索引节点；
- *数据块区*，用来存储文件或目录数据；



- 超级块：当文件系统挂载时进入内存；
- 索引节点区：当文件被访问时进入内存；



在 Linux 文件系统中，用户空间、系统调用、虚拟机文件系统、缓存、文件系统以及存储之间的关系如下图：

![img](xiaolin-os.assets/虚拟文件系统.png)

读文件和写文件的过程：

- 当用户进程从文件读取 1 个字节大小的数据时，文件系统则需要获取字节所在的数据块，再返回数据块对应的用户进程所需的数据部分。
- 当用户进程把 1 个字节大小的数据写进文件时，文件系统则找到需要写入数据的数据块的位置，然后修改数据块中对应的部分，最后再把数据块写回磁盘。

所以说，**文件系统的基本操作单位是数据块**。



### Unix 文件的实现方式

![img](xiaolin-os.assets/文件存储方式比较.png)



那早期 Unix 文件系统是组合了前面的文件存放方式的优点，如下图：



![早期 Unix 文件系统](xiaolin-os.assets/Unix 多级索引.png)

它是根据文件的大小，存放的方式会有所变化：

- 如果存放文件所需的数据块小于 10 块，则采用直接查找的方式；
- 如果存放文件所需的数据块超过 10 块，则采用一级间接索引方式；
- 如果前面两种方式都不够存放大文件，则采用二级间接索引方式；
- 如果二级间接索引也不够存放大文件，这采用三级间接索引方式；



### 位图法

位图是利用二进制的一位来表示磁盘中一个盘块的使用情况，磁盘上所有的盘块都有一个二进制位与之对应。

当值为 0 时，表示对应的盘块空闲，值为 1 时，表示对应的盘块已分配。它形式如下：

```text
1111110011111110001110110111111100111 ...
```

在 Linux 文件系统就采用了位图的方式来管理空闲空间，不仅用于数据空闲块的管理，还用于 inode 空闲块的管理，因为 inode 也是存储在磁盘的，自然也要有对其管理。



### 目录

和普通文件不同的是，**普通文件的块里面保存的是文件数据，而目录文件的块里面保存的是目录里面一项一项的文件信息。**

列表中每一项就代表该目录下的文件的文件名和对应的 inode，通过这个 inode，就可以找到真正的文件。

![目录格式哈希表](xiaolin-os.assets/目录哈希表.png)

目录查询是通过在磁盘上反复搜索完成，需要不断地进行 I/O 操作，开销较大。所以，为了减少 I/O 操作，把当前使用的文件目录缓存在内存，以后要使用该文件时只要在内存中操作，从而降低了磁盘操作次数，提高了文件系统的访问速度。



### 软链接和硬链接

有时候我们希望给某个文件取个别名，那么在 Linux 中可以通过**硬链接（\*Hard Link\*）** 和**软链接（\*Symbolic Link\*）** 的方式来实现，它们都是比较特殊的文件，但是实现方式也是不相同的。

硬链接是**多个目录项中的「索引节点」指向一个文件**，也就是指向同一个 inode，但是 inode 是不可能跨越文件系统的，每个文件系统都有各自的 inode 数据结构和列表，所以**硬链接是不可用于跨文件系统的**。由于多个目录项都是指向一个 inode，那么**只有删除文件的所有硬链接以及源文件时，系统才会彻底删除该文件。**

![硬链接](xiaolin-os.assets/硬链接-2.png)

软链接相当于重新创建一个文件，这个文件有**独立的 inode**，但是这个**文件的内容是另外一个文件的路径**，所以访问软链接的时候，实际上相当于访问到了另外一个文件，所以**软链接是可以跨文件系统的**，甚至**目标文件被删除了，链接文件还是在的，只不过指向的文件找不到了而已。**

![软链接](xiaolin-os.assets/软链接.png)



## 文件 I/O

文件的读写方式各有千秋，对于文件的 I/O 分类也非常多，常见的有

- 缓冲与非缓冲 I/O
- 直接与非直接 I/O
- 阻塞与非阻塞 I/O VS 同步与异步 I/O



### 缓冲与非缓冲 I/O

文件操作的标准库是可以实现数据的缓存，那么**根据「是否利用标准库缓冲」，可以把文件 I/O 分为缓冲 I/O 和非缓冲 I/O**：

- 缓冲 I/O，利用的是标准库的缓存实现文件的加速访问，而标准库再通过系统调用访问文件。
- 非缓冲 I/O，直接通过系统调用访问文件，不经过标准库缓存。

这里所说的「缓冲」特指标准库内部实现的缓冲。



### 直接与非直接 I/O

我们都知道磁盘 I/O 是非常慢的，所以 Linux 内核为了减少磁盘 I/O 次数，在系统调用后，会把用户数据拷贝到内核中缓存起来，这个内核缓存空间也就是「页缓存」，只有当缓存满足某些条件的时候，才发起磁盘 I/O 的请求。

那么，**根据是「否利用操作系统的缓存」，可以把文件 I/O 分为直接 I/O 与非直接 I/O**：

- 直接 I/O，不会发生内核缓存和用户程序之间数据复制，而是直接经过文件系统访问磁盘。
- 非直接 I/O，读操作时，数据从内核缓存中拷贝给用户程序，写操作时，数据从用户程序拷贝给内核缓存，再由内核决定什么时候写入数据到磁盘。

如果你在使用文件操作类的系统调用函数时，指定了 `O_DIRECT` 标志，则表示使用直接 I/O。如果没有设置过，默认使用的是非直接 I/O。

> 如果用了非直接 I/O 进行写数据操作，内核什么情况下才会把缓存数据写入到磁盘？

以下几种场景会触发内核缓存的数据写入磁盘：

- 在调用 `write` 的最后，当发现内核缓存的数据太多的时候，内核会把数据写到磁盘上；
- 用户主动调用 `sync`，内核缓存会刷到磁盘上；
- 当内存十分紧张，无法再分配页面时，也会把内核缓存的数据刷到磁盘上；
- 内核缓存的数据的缓存时间超过某个时间时，也会把数据刷到磁盘上；



### 阻塞与非阻塞 I/O VS 同步与异步 I/O

注意，**阻塞等待的是「内核数据准备好」和「数据从内核态拷贝到用户态」这两个过程**。过程如下图：

![阻塞 I/O](xiaolin-os.assets/阻塞 I_O.png)

知道了阻塞 I/O ，来看看**非阻塞 I/O**，非阻塞的 read 请求在数据未准备好的情况下立即返回，可以继续往下执行，此时应用程序不断轮询内核，直到数据准备好，内核将数据拷贝到应用程序缓冲区，`read` 调用才可以获取到结果。过程如下图：

![非阻塞 I/O](xiaolin-os.assets/非阻塞 I_O .png)

举个例子，访问管道或 socket 时，如果设置了 `O_NONBLOCK` 标志，那么就表示使用的是非阻塞 I/O 的方式访问，而不做任何设置的话，默认是阻塞 I/O。



为了解决这种傻乎乎轮询方式，于是 **I/O 多路复用**技术就出来了，如 select、poll，它是通过 I/O 事件分发，当内核数据准备好时，再以事件通知应用程序进行操作。

![I/O 多路复用](xiaolin-os.assets/基于非阻塞 I_O 的多路复用.png)

**read都是同步调用。因为它们在 read 调用时，内核将数据从内核空间拷贝到应用程序空间，过程都是需要等待的，也就是说这个过程是同步的，如果内核实现的拷贝效率不高，read 调用就会在这个同步过程中等待比较长的时间。**



而真正的**异步 I/O** 是「内核数据准备好」和「数据从内核态拷贝到用户态」这两个过程都不用等待。

当我们发起 `aio_read` 之后，就立即返回，内核自动将数据从内核空间拷贝到应用程序空间，这个拷贝过程同样是异步的，内核自动完成的，和前面的同步操作不一样，应用程序并不需要主动发起拷贝动作。过程如下图：

![异步 I/O](xiaolin-os.assets/异步 I_O.png)

对比：

![img](xiaolin-os.assets/同步VS异步IO.png)



## 设备管理

输入输出设备可分为两大类 ：**块设备（\*Block Device\*）\**和\**字符设备（\*Character Device\*）**。

- *块设备*，把数据存储在固定大小的块中，每个块有自己的地址，硬盘、USB 是常见的块设备。
- *字符设备*，以字符为单位发送或接收一个字符流，字符设备是不可寻址的，也没有任何寻道操作，鼠标是常见的字符设备。

![计算机 I/O 系统结构](xiaolin-os.assets/I_O系统结构.png)

块设备通常传输的数据量会非常大，于是控制器设立了一个可读写的**数据缓冲区**。

- CPU 写入数据到控制器的缓冲区时，当缓冲区的数据囤够了一部分，才会发给设备。
- CPU 从控制器的缓冲区读取数据时，也需要缓冲区囤够了一部分，才拷贝到内存。



因为这些**设备控制器（Device Controller）**都很清楚的知道对应设备的用法和功能，所以 CPU 是通过设备控制器来和设备打交道的。



![img](xiaolin-os.assets/驱动程序.png)

通常，设备驱动程序初始化的时候，要先注册一个该设备的中断处理函数。

![img](xiaolin-os.assets/中断工作过程.png)



DMA：Direct Memory Access

![img](xiaolin-os.assets/DMA工作原理.png)



可以把 Linux 存储系统的 I/O 由上到下可以分为三个层次，分别是文件系统层、通用块层、设备层。他们整个的层次关系如下图：

![img](xiaolin-os.assets/I_O软件分层.png)



# 网络系统

I_O 中断

![img](xiaolin-os.assets/I_O 中断.png)

计算机科学家们发现了事情的严重性后，于是就发明了 DMA 技术，也就是**直接内存访问（\*Direct Memory Access\*）** 技术。

什么是 DMA 技术？简单理解就是，**在进行 I/O 设备和内存的数据传输的时候，数据搬运的工作全部交给 DMA 控制器，而 CPU 不再参与任何与数据搬运相关的事情，这样 CPU 就可以去处理别的事务**。

![img](xiaolin-os.assets/DRM I_O 过程.png)

- **DMA 收到磁盘的信号，将磁盘控制器缓冲区中的数据拷贝到内核缓冲区中，此时不占用 CPU，CPU 可以执行其他任务**；



## 传统的文件传输有多糟糕？

```c
read(file, tmp_buf, len);
write(socket, tmp_buf, len);
```

![img](xiaolin-os.assets/传统文件传输.png)

首先，期间共**发生了 4 次用户态与内核态的上下文切换**，因为发生了两次系统调用，一次是 `read()` ，一次是 `write()`，每次系统调用都得先从用户态切换到内核态，等内核完成任务后，再从内核态切换回用户态。

上下文切换到成本并不小，一次切换需要耗时几十纳秒到几微秒，虽然时间看上去很短，但是在高并发的场景下，这类时间容易被累积和放大，从而影响系统的性能。

其次，还**发生了 4 次数据拷贝**，其中两次是 DMA 的拷贝，另外两次则是通过 CPU 拷贝的

所以，**要想提高文件传输的性能，就需要减少「用户态与内核态的上下文切换」和「内存拷贝」的次数**。



## 如何优化文件传输的性能？

所以，**要想减少上下文切换到次数，就要减少系统调用的次数**。

因为文件传输的应用场景中，在用户空间我们并不会对数据「再加工」，所以数据实际上可以不用搬运到用户空间，因此**用户的缓冲区是没有必要存在的**。



## 如何实现零拷贝？

零拷贝技术实现的方式通常有 2 种：

- mmap + write
- sendfile

###  mmap + write

在前面我们知道，`read()` 系统调用的过程中会把内核缓冲区的数据拷贝到用户的缓冲区里，于是为了减少这一步开销，我们可以用 `mmap()` 替换 `read()` 系统调用函数。

```c
buf = mmap(file, len);
write(sockfd, buf, len);
```

 `mmap()` 系统调用函数会直接把内核缓冲区里的数据「**映射**」到用户空间，这样，操作系统内核与用户空间就不需要再进行任何的数据拷贝操作。

![img](xiaolin-os.assets/mmap + write 零拷贝.png)

通过使用 `mmap()` 来代替 `read()`， 可以减少一次数据拷贝的过程。



### sendfile

在 Linux 内核版本 2.1 中，提供了一个专门发送文件的系统调用函数 `sendfile()`，函数形式如下：

```c
#include <sys/socket.h>
ssize_t sendfile(int out_fd, int in_fd, off_t *offset, size_t count);
```

它的前两个参数分别是目的端和源端的文件描述符，后面两个参数是源端的偏移量和复制数据的长度，返回值是实际复制数据的长度。

![img](xiaolin-os.assets/senfile-3次拷贝.png)



但是这还不是真正的零拷贝技术，如果网卡支持 SG-DMA（*The Scatter-Gather Direct Memory Access*）技术（和普通的 DMA 有所不同），我们可以进一步减少通过 CPU 把内核缓冲区里的数据拷贝到 socket 缓冲区的过程。

你可以在你的 Linux 系统通过下面这个命令，查看网卡是否支持 scatter-gather 特性：

```bash
$ ethtool -k eth0 | grep scatter-gather
scatter-gather: on
```

于是，从 Linux 内核 `2.4` 版本开始起，对于支持网卡支持 SG-DMA 技术的情况下， `sendfile()` 系统调用的过程发生了点变化，具体过程如下：

![img](xiaolin-os.assets/senfile-零拷贝.png)

这就是所谓的**零拷贝（\*Zero-copy\*）技术，因为我们没有在内存层面去拷贝数据，也就是说全程没有通过 CPU 来搬运数据，所有的数据都是通过 DMA 来进行传输的。**。

零拷贝技术的文件传输方式相比传统文件传输的方式，减少了 2 次上下文切换和数据拷贝次数，**只需要 2 次上下文切换和数据拷贝次数，就可以完成文件的传输，而且 2 次的数据拷贝过程，都不需要通过 CPU，2 次都是由 DMA 来搬运。**

所以，总体来看，**零拷贝技术可以把文件传输的性能提高至少一倍以上**。



### 使用零拷贝技术的项目

事实上，Kafka 这个开源项目，就利用了「零拷贝」技术，从而大幅提升了 I/O 的吞吐率，这也是 Kafka 在处理海量数据为什么这么快的原因之一。

如果你追溯 Kafka 文件传输的代码，你会发现，最终它调用了 Java NIO 库里的 `transferTo` 方法：

```java
@Overridepublic 
long transferFrom(FileChannel fileChannel, long position, long count) throws IOException { 
    return fileChannel.transferTo(position, count, socketChannel);
}
```

如果 Linux 系统支持 `sendfile()` 系统调用，那么 `transferTo()` 实际上最后就会使用到 `sendfile()` 系统调用函数。

曾经有大佬专门写过程序测试过，在同样的硬件条件下，传统文件传输和零拷拷贝文件传输的性能差异，你可以看到下面这张测试数据图，使用了零拷贝能够缩短 `65%` 的时间，大幅度提升了机器传输数据的吞吐量。



另外，Nginx 也支持零拷贝技术，一般默认是开启零拷贝技术，这样有利于提高文件传输的效率，是否开启零拷贝技术的配置如下：

```nginx
http {
...
    sendfile on
...
}
```

sendfile 配置的具体意思:

- 设置为 on 表示，使用零拷贝技术来传输文件：sendfile ，这样只需要 2 次上下文切换，和 2 次数据拷贝。
- 设置为 off 表示，使用传统的文件传输技术：read + write，这时就需要 4 次上下文切换，和 4 次数据拷贝。

当然，要使用 sendfile，Linux 内核版本必须要 2.1 以上的版本。





## PageCache 有什么作用？

回顾前面说道文件传输过程，其中第一步都是先需要先把磁盘文件数据拷贝「内核缓冲区」里，这个「内核缓冲区」实际上是**磁盘高速缓存（\*PageCache\*）**。

我们都知道程序运行的时候，具有「局部性」，所以通常，刚被访问的数据在短时间内再次被访问的概率很高，于是我们可以用 **PageCache 来缓存最近被访问的数据**，当空间不足时淘汰最久未被访问的缓存。

PageCache 的优点主要是两个：

- 缓存最近被访问的数据；
- 预读功能；

这两个做法，将大大提高读写磁盘的性能。

**但是，在传输大文件（GB 级别的文件）的时候，PageCache 会不起作用，那就白白浪费 DMA 多做的一次数据拷贝，造成性能的降低，即使使用了 PageCache 的零拷贝也会损失性能**



普通read：

![img](xiaolin-os.assets/阻塞 IO 的过程.png)



对于阻塞的问题，可以用异步 I/O 来解决，它工作方式如下图：

![img](xiaolin-os.assets/异步 IO 的过程.png)

它把读操作分为两部分：

- 前半部分，内核向磁盘发起读请求，但是可以**不等待数据就位就可以返回**，于是进程此时可以处理其他任务；
- 后半部分，当内核将磁盘中的数据拷贝到进程缓冲区后，进程将接收到内核的**通知**，再去处理数据；

而且，我们可以发现，异步 I/O 并没有涉及到 PageCache，所以使用异步 I/O 就意味着要绕开 PageCache。



**绕开 PageCache 的 I/O 叫直接 I/O，使用 PageCache 的 I/O 则叫缓存 I/O。通常，对于磁盘，异步 I/O 只支持直接 I/O。**

于是，**在高并发的场景下，针对大文件的传输的方式，应该使用「异步 I/O + 直接 I/O」来替代零拷贝技术**。



另外，由于直接 I/O 绕过了 PageCache，就无法享受内核的这两点的优化：

- 内核的 I/O 调度算法会缓存尽可能多的 I/O 请求在 PageCache 中，最后「**合并**」成一个更大的 I/O 请求再发给磁盘，这样做是为了减少磁盘的寻址操作；
- 内核也会「**预读**」后续的 I/O 请求放在 PageCache 中，一样是为了减少对磁盘的操作；

所以，传输文件的时候，我们要根据文件的大小来使用不同的方式：

- 传输大文件的时候，使用「异步 I/O + 直接 I/O」；
- 传输小文件的时候，则使用「零拷贝技术」；

在 nginx 中，我们可以用如下配置，来根据文件的大小来使用不同的方式：

```nginx
location /video/ { 
    sendfile on; 
    aio on; 
    directio 1024m; 
}
```

当文件大小大于 `directio` 值后，使用「异步 I/O + 直接 I/O」，否则使用「零拷贝技术」。



## 高性能网络模式：Reactor & Proactor

nginx、redis用的是Reactor。



select/poll/epoll 是如何获取网络事件的呢？

在获取事件时，先把我们要关心的连接传给内核，再由内核检测：

- 如果没有事件发生，线程只需阻塞在这个系统调用，而无需像前面的线程池方案那样轮训调用 read 操作来判断是否有数据。
- 如果有事件发生，内核会返回产生了事件的连接，线程就会从阻塞状态返回，然后在用户态中再处理这些连接对应的业务即可。

Reactor 翻译过来的意思是「反应堆」，可能大家会联想到物理学里的核反应堆，实际上并不是的这个意思。

这里的反应指的是「**对事件反应**」，也就是**来了一个事件，Reactor 就有相对应的反应/响应**。

事实上，Reactor 模式也叫 `Dispatcher` 模式，我觉得这个名字更贴合该模式的含义，即 **I/O 多路复用监听事件，收到事件后，根据事件类型分配（Dispatch）给某个进程 / 线程**。

Reactor 模式主要由 Reactor 和处理资源池这两个核心部分组成，它俩负责的事情如下：

- Reactor 负责监听和分发事件，事件类型包含连接事件、读写事件；
- 处理资源池负责处理事件，如 read -> 业务逻辑 -> send；

Reactor 模式是灵活多变的，可以应对不同的业务场景，灵活在于：

- Reactor 的数量可以只有一个，也可以有多个；
- 处理资源池可以是单个进程 / 线程，也可以是多个进程 /线程；



## Reactor

「**单 Reactor 单进程**」的方案示意图：

![img](xiaolin-os.assets/单Reactor单进程.png)

单 Reactor 单进程的方案因为全部工作都在同一个进程内完成，所以实现起来比较简单，不需要考虑进程间通信，也不用担心多进程竞争。

但是，这种方案存在 2 个缺点：

- 第一个缺点，因为只有一个进程，**无法充分利用 多核 CPU 的性能**；
- 第二个缺点，Handler 对象在业务处理时，整个进程是无法处理其他连接的事件的，**如果业务处理耗时比较长，那么就造成响应的延迟**；

所以，单 Reactor 单进程的方案**不适用计算机密集型的场景，只适用于业务处理非常快速的场景**。

`Redis` 是由 C 语言实现的，它采用的正是「**单 Reactor 单进程**」的方案，因为 Redis 业务处理主要是在内存中完成，操作的速度是很快的，性能瓶颈不在 CPU 上，所以 Redis 对于命令的处理是单进程的方案。



![img](xiaolin-os.assets/单Reactor多线程.png)

要避免多线程由于竞争共享资源而导致数据错乱的问题，就需要在操作共享资源前加上**互斥锁**，以保证任意时间里只有一个线程在操作共享资源，待该线程操作完释放互斥锁后，其他线程才有机会操作共享数据。

「单 Reactor」的模式还有个问题，**因为一个 Reactor 对象承担所有事件的监听和响应，而且只在主线程中运行，在面对瞬间高并发的场景时，容易成为性能的瓶颈的地方**。



###  多 Reactor 多进程 / 线程

![img](xiaolin-os.assets/主从Reactor多线程.png)

- 主线程和子线程分工明确，主线程只负责接收新连接，子线程负责完成后续的业务处理。
- 主线程和子线程的交互很简单，主线程只需要把新连接传给子线程，子线程无须返回数据，直接就可以在子线程将处理结果发送给客户端。

大名鼎鼎的两个开源软件 Netty 和 Memcache 都采用了「多 Reactor 多线程」的方案。

采用了「多 Reactor 多进程」方案的开源软件是 Nginx，不过方案与标准的多 Reactor 多进程有些差异。

Nginx具体差异表现在主进程中仅仅用来初始化 socket，并没有创建 mainReactor 来 accept 连接，而是由子进程的 Reactor 来 accept 连接，**通过锁来控制一次只有一个子进程进行 accept（防止出现惊群现象）**，子进程 accept 新连接后就放到自己的 Reactor 进行处理，不会再分配给其他子进程。



## Proactor

前面提到的 Reactor 是非阻塞同步网络模式，而 **Proactor 是异步网络模式**。

非阻塞 I_O:

![非阻塞 I/O](xiaolin-os.assets/非阻塞 I_O -20220421085959150.png)



注意，**这里最后一次 read 调用，获取数据的过程，是一个同步的过程，是需要等待的过程。这里的同步指的是内核态的数据拷贝到用户程序的缓存区这个过程。**

举个例子，如果 socket 设置了 `O_NONBLOCK` 标志，那么就表示使用的是非阻塞 I/O 的方式访问，而不做任何设置的话，默认是阻塞 I/O。



真正的**异步 I/O** 是「内核数据准备好」和「数据从内核态拷贝到用户态」这**两个过程都不用等待**。

当我们发起 `aio_read` （异步 I/O） 之后，就立即返回，内核自动将数据从内核空间拷贝到用户空间，这个拷贝过程同样是异步的，内核自动完成的，和前面的同步操作不一样，**应用程序并不需要主动发起拷贝动作**。过程如下图：

![异步 I/O](xiaolin-os.assets/异步 I_O-20220421210249765.png)



现在我们再来理解 Reactor 和 Proactor 的区别，就比较清晰了。

- **Reactor 是非阻塞同步网络模式，感知的是就绪可读写事件**。
- **Proactor 是异步网络模式， 感知的是已完成的读写事件**。

因此，**Reactor 可以理解为「来了事件操作系统通知应用进程，让应用进程来处理」**，而 **Proactor 可以理解为「来了事件操作系统来处理，处理完再通知应用进程」**。

> 举个实际生活中的例子，Reactor 模式就是快递员在楼下，给你打电话告诉你快递到你家小区了，你需要自己下楼来拿快递。而在 Proactor 模式下，快递员直接将快递送到你家门口，然后通知你。

无论是 Reactor，还是 Proactor，都是一种基于「事件分发」的网络编程模式，区别在于 **Reactor 模式是基于「待完成」的 I/O 事件，而 Proactor 模式则是基于「已完成」的 I/O 事件**。

![img](xiaolin-os.assets/Proactor.png)

可惜的是，在 Linux 下的异步 I/O 是不完善的， `aio` 系列函数是由 POSIX 定义的异步操作接口，不是真正的操作系统级别支持的，而是在用户空间模拟出来的异步，并且仅仅支持基于本地文件的 aio 异步操作，网络编程中的 socket 是不支持的，这也使得基于 Linux 的高性能网络程序都是使用 Reactor 方案。

而 Windows 里实现了一套完整的支持 socket 的异步编程接口，这套接口就是 `IOCP`，是由操作系统级别实现的异步 I/O，真正意义上异步 I/O，因此在 Windows 里实现高性能网络程序可以使用效率更高的 Proactor 方案。



# Linux头

![img](xiaolin-os.assets/封装.png)



## 网络配置如何看？

要想知道网络的配置和状态，我们可以使用 `ifconfig` 或者 `ip` 命令来查看。

这两个命令功能都差不多，不过它们属于不同的软件包，`ifconfig` 属于 `net-tools` 软件包，`ip` 属于 `iproute2` 软件包，我的印象中 `net-tools` 软件包没有人继续维护了，而 `iproute2` 软件包是有开发者依然在维护，所以更推荐你使用 `ip` 工具。

![img](xiaolin-os.assets/showeth0.png)



虽然这两个命令输出的格式不尽相同，但是输出的内容基本相同，比如都包含了 IP 地址、子网掩码、MAC 地址、网关地址、MTU 大小、网口的状态以及网路包收发的统计信息，下面就来说说这些信息，它们都与网络性能有一定的关系。

第一，网口的连接状态标志。其实也就是表示对应的网口是否连接到交换机或路由器等设备，如果 `ifconfig` 输出中看到有 `RUNNING`，或者 `ip` 输出中有 `LOWER_UP`，则说明物理网路是连通的，如果看不到，则表示网口没有接网线。

第二，MTU 大小。默认值是 `1500` 字节，其作用主要是限制网络包的大小，如果 IP 层有一个数据报要传，而且数据帧的长度比链路层的 MTU 还大，那么 IP 层就需要进行分片，即把数据报分成干片，这样每一片就都小于 MTU。事实上，每个网络的链路层 MTU 可能会不一样，所以你可能需要调大或者调小 MTU 的数值。

第三，网口的 IP 地址、子网掩码、MAC 地址、网关地址。这些信息必须要配置正确，网络功能才能正常工作。

第四，网路包收发的统计信息。通常有网络收发的字节数、包数、错误数以及丢包情况的信息，如果 `TX`（发送） 和 `RX`（接收） 部分中 errors、dropped、overruns、carrier 以及 collisions 等指标不为 0 时，则说明网络发送或者接收出问题了，这些出错统计信息的指标意义如下：

- *errors* 表示发生错误的数据包数，比如校验错误、帧同步错误等；
- *dropped* 表示丢弃的数据包数，即数据包已经收到了 Ring Buffer（这个缓冲区是在内核内存中，更具体一点是在网卡驱动程序里），但因为系统内存不足等原因而发生的丢包；
- *overruns* 表示超限数据包数，即网络接收/发送速度过快，导致 Ring Buffer 中的数据包来不及处理，而导致的丢包，因为过多的数据包挤压在 Ring Buffer，这样 Ring Buffer 很容易就溢出了；
- *carrier* 表示发生 carrirer 错误的数据包数，比如双工模式不匹配、物理电缆出现问题等；
- *collisions* 表示冲突、碰撞数据包数；

`ifconfig` 和 `ip` 命令只显示的是网口的配置以及收发数据包的统计信息，而看不到协议栈里的信息，那接下来就来看看如何查看协议栈里的信息。



## socket 信息如何查看？

我们可以使用 `netstat` 或者 `ss`，这两个命令查看 socket、网络协议栈、网口以及路由表的信息。

虽然 `netstat` 与 `ss` 命令查看的信息都差不多，但是如果在生产环境中要查看这类信息的时候，尽量不要使用 `netstat` 命令，因为它的性能不好，在系统比较繁忙的情况下，如果频繁使用 `netstat` 命令则会对性能的开销雪上加霜，所以更推荐你使用性能更好的 `ss` 命令。

从下面这张图，你可以看到这两个命令的输出内容：

![img](xiaolin-os.assets/showsocket.png)

可以发现，输出的内容都差不多， 比如都包含了 socket 的状态（*State*）、接收队列（*Recv-Q*）、发送队列（*Send-Q*）、本地地址（*Local Address*）、远端地址（*Foreign Address*）、进程 PID 和进程名称（*PID/Program name*）等。

接收队列（*Recv-Q*）和发送队列（*Send-Q*）比较特殊，在不同的 socket 状态。它们表示的含义是不同的。

当 socket 状态处于 `Established`时：

- *Recv-Q* 表示 socket 缓冲区中还没有被应用程序读取的字节数；
- *Send-Q* 表示 socket 缓冲区中还没有被远端主机确认的字节数；

而当 socket 状态处于 `Listen` 时：

- *Recv-Q* 表示全连接队列的长度；
- *Send-Q* 表示全连接队列的最大长度；



![半连接队列与全连接队列](xiaolin-os.assets/3.jpg)



`ss` 命令输出的统计信息相比 `netsat` 比较少，`ss` 只显示已经连接（*estab*）、关闭（*closed*）、孤儿（*orphaned*） socket 等简要统计。

```bash
ubuntu@hk1:~$ ss -s
Total: 1128
TCP:   992 (estab 683, closed 249, orphaned 0, timewait 249)

Transport Total     IP        IPv6
RAW	  1         0         1        
UDP	  16        9         7        
TCP	  743       558       185      
INET	  760       567       193      
FRAG	  0         0         0 
```

