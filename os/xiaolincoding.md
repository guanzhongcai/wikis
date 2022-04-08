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
