# Linux



## tty

TTY是电传打字机`Teletypewriter`的缩写，TTY是最流行的终端设备。

终端（terminal）=tty=文本的输入输出环境



## Linux中rz和sz命令用法详解

`rz，sz`是Linux/Unix同Windows进行ZModem文件传输的命令行工具。优点就是不用再开一个sftp工具登录上去上传下载文件。
`sz`：将选定的文件发送（send）到本地机器
`rz`：运行该命令会弹出一个文件选择窗口，从本地选择文件上传到Linux服务器
安装命令：

```bash
yum install lrzsz
```

从服务端发送文件到客户端：

```bash
sz filename
```

从客户端上传文件到服务端：

```
rz
```



## 需熟练掌握的命令

- 文件系统结构和基本操作 ls/chmod/chown/rm/find/ln/cat/mount/mkdir/tar/gzip …
- 学会使用一些文本操作命令 sed/awk/grep/tail/less/more …
- 学会使用一些管理命令 ps/top/lsof/netstat/kill/tcpdump/iptables/dd…



### sed



### awk



### dd



## lsof

lsof = list open files

```bash
# 使用-i:port来显示与指定端口相关的网络信息
lsof  -i :22

# 使用@host来显示指定到指定主机的连接
lsof  -i@172.16.12.5

# 使用@host:port显示基于主机与端口的连接
lsof  -i@172.16.12.5:22

# 找出正等候连接的端口。
lsof  -i -sTCP:LISTEN

# 找出已建立的连接
lsof  -i -sTCP:ESTABLISHED

# 消灭指定用户运行的所有程序
kill  -9  `lsof -t -u daniel`

# 使用-c查看指定的命令正在使用的文件和网络连接
lsof  -c syslog-ng

# 使用-p查看指定进程ID已打开的文件
lsof  -p 10075

# 显示开启文件abc.txt的进程
lsof abc.txt 

# 显示某个端口范围的打开的连接
lsof  -i @fw.google.com:2150=2180

# 同时使用-t和-c选项以给进程发送 HUP 信号
kill  -HUP `lsof -t -c sshd`
```

参考资料：

- [Linux 命令神器：lsof](https://www.jianshu.com/p/a3aa6b01b2e1)



## tcpdump

用简单的话来定义tcpdump，就是：dump the traffic on a network，根据使用者的定义对网络上的数据包进行截获的包分析工具。 tcpdump可以将网络中传送的数据包的“头”完全截获下来提供分析。它支持针对网络层、协议、主机、网络或端口的过滤，并提供and、or、not等逻辑语句来帮助你去掉无用的信息。

```bash
# 监视第一个网络接口上所有流过的数据包。
tcpdump

# 截获主机hostname发送的所有数据
tcpdump -i eth0 src host hostname

# 监视所有送到主机hostname的数据包
tcpdump -i eth0 dst host hostname

# 监视指定主机和端口的数据包
# 获取主机210.27.48.1接收或发出的telnet包
tcpdump tcp port 23 and host 210.27.48.1

# 使用tcpdump抓取HTTP包
# 0x4745 为"GET"前两个字母"GE",0x4854 为"HTTP"前两个字母"HT"。
tcpdump  -XvvennSs 0 -i eth0 tcp[20:2]=0x4745 or tcp[20:2]=0x4854

```

tcpdump 对截获的数据并没有进行彻底解码，数据包内的大部分内容是使用十六进制的形式直接打印输出的。显然这不利于分析网络故障，通常的解决办法是先使用带-w参数的tcpdump 截获数据并保存到文件中，然后再使用其他程序(如Wireshark)进行解码分析。当然也应该定义过滤规则，以避免捕获的数据包填满整个硬盘。



参考资料：

- [Linux tcpdump命令详解](https://www.cnblogs.com/ggjucheng/archive/2012/01/14/2322659.html)



## inode

[inode](https://baike.baidu.com/item/inode/16194#viewPageContent) 用来存放档案及目录的基本信息，包含时间、档名、使用者及群组等。

*inode* 是 UNIX 操作系统中的一种数据结构，其本质是[结构体](https://baike.baidu.com/item/结构体)，它包含了与文件系统中各个文件相关的一些重要信息。在 UNIX 中创建文件系统时，同时将会创建大量的 inode 。通常，文件系统磁盘空间中大约百分之一空间分配给了 inode 表。

inode 表包含一份清单，其中列出了对应文件系统的所有 inode 编号。当用户搜索或者访问一个文件时，UNIX 系统通过 inode 表查找正确的 inode 编号。在找到 inode 编号之后，相关的命令才可以访问该 inode ，并对其进行适当的更改。

例如，使用 vi 来编辑一个文件。当您键入 vi 时，在 inode 表中找到 inode 编号之后，才允许您打开该 inode 。在 vi 的编辑会话期间，更改了该 inode 中的某些属性，当您完成操作并键入 :wq 时，将关闭并释放该 inode 。



inode 中所包含的、UNIX 用户经常使用的一些重要信息：

● inode 编号

● 用来识别文件类型，以及用于 stat C 函数的模式信息

● 文件的链接数目

● 属主的ID (UID)

● 属主的组 ID (GID)

● 文件的大小

● 文件所使用的磁盘块的实际数目

● 最近一次修改的时间

● 最近一次访问的时间

● 最近一次更改的时间

从根本上讲， inode 中包含有关文件的所有信息（除了文件的实际名称以及实际数据内容之外）。



## iptables

- iptables是Linux [内核](https://baike.baidu.com/item/内核/108410)集成的 IP 信息包过滤系统，有利于在 Linux 系统上更好地控制 IP 信息包过滤和[防火墙](https://baike.baidu.com/item/防火墙/52767)配置。

- iptables是基于包过滤的防火墙工具、可以对流入、流出及流经服务器的数据包进行精细的控制。

- iptables工作在OSI七层的二、三层、四层。

- 防火墙在做[数据包过滤](https://baike.baidu.com/item/数据包过滤/7747768)决定时，有一套遵循和组成的规则，这些规则存储在专用的数据包过滤表中，而这些表集成在 Linux 内核中。在数据包过滤表中，规则被分组放在我们所谓的链（chain）中。而netfilter/iptables IP 数据包过滤系统是一款功能强大的工具，可用于添加、编辑和移除规则。

- 虽然 netfilter/iptables IP 信息包过滤系统被称为单个实体，但它实际上由两个组件netfilter 和 iptables 组成。

- **netfilter 组件也称为内核空间（kernelspace）**，是内核的一部分，由一些信息包过滤表组成，这些表包含内核用来控制信息包过滤处理的规则集。

- **iptables 组件是一种工具，也称为用户空间（userspace）**，它使插入、修改和除去信息包过滤表中的规则变得容易。

  ![img](../images/iptables-input-output.png)



![img](../images/iptables-workflow.png)



> 开放式系统互联通信参考模型（英语：Open System Interconnection Reference Model，缩写为 *OSI*），简称为*OSI*模型（*OSI* model），一种概念模型，由国际标准化组织提出，一个试图使各种计算机在世界范围内互连为网络的标准框架。

基本术语：

| Netfileter/iptables | 表（tables） | 链（chains） |  规则（policy）  |
| ------------------- | ------------ | ------------ | :--------------: |
| 小区的一栋楼        | 楼里的房子   | 房子里的柜子 | 增加、摆放的规则 |



启动并查看iptables

```bash
# 启动
/etc/init.d/iptables start

# 查看帮助
iptables -h

# 查看当前规则
iptables -L -v -x -n -t filter/nat
# -L ：列出一个或所有链的规则
# -v：显示详细信息、包括每条规则匹配包数量和匹配字节数
# -x：在v的基础上、进制自动单位换算（K,M）
# -n: 只显示IP地址和端口号码。不显示域名和服务名称
# -t : 接表名、如果不加-t，默认就是 –t filter
```

栗子：

```bash
# 例如：为了防止DOS太多连接进来，那么可以允许最多15个初始连接，超过的丢弃.
iptable -A INPUT -s 192.186.1.0/24 -p tcp --sync -m connlimit --connlimit-above 15 -j DROP

# ip范围应用
iptables -A FORWARD -m iprange --src-range 192.168.1.5-192.168.1.124 -j ACCEPT

# Syn-flood protection 洪水攻击保护:
iptables -A FORWARD -p tcp --syn -m limit --limit 1/s -j ACCEPT

# Furtive port scanner 端口扫描:
iptables -A FORWARD -p tcp --tcp-flags SYN,ACK,FIN,RST RST -m limit --limit 1/s -j ACCEPT

# Ping of death 禁止PING:
iptables -A FORWARD -p icmp --icmp-type echo-request -m limit --limit 1/s -j ACCEPT

# 禁止ssh默认的22端口 iptables 默认用的就是filter表 
# --jump  -j target 
# target 的常见的处理方法有ACCEPT（接受），DROP（丢弃），REJECT（拒接）其中、一般不使用REJECT行为、REJECT会带来安全隐患
iptables -A INPUT -p tcp -dport 22 -j DROP
iptables -t filter -A INPUT -p tcp --dport 22 -j DROP

# 源地址不是192.168.132.201 的禁止连接
iptables -A INPUT -I eth1 -s ! 192.168.132.201 -j DROP

# db 仅允许内部合法ip段访问mysql数据库
iptables –A INPUT –s 192.168.10.0/24 –p tcp –dport 3306 –j ACCEPT

对外提供HTTP服务的业务，要允许http服务通过、并且不限制IP

#http

iptables –A INPUT –p tcp –dport 80 –j ACCEPT

对内部提供http服务的业务、一般用特殊端口、并且限制合法IP连接或者VPN连接

iptables –A INPUT –s 192.168.132.0 –p tcp –m multiport –dport 8080,8081,8082,8888 –j ACCEPT
```



# bash 和 sh 的区别

shell的中文意思就是贝壳，其实比较类似于我们内核的壳。简而言之就是只要能够操作应用程序的接口都能够称为SHELL。狭义的shell指的是命令行方面的软件，广义的SHELL则包括图形界面。

1. 发现区别
   同样的 shell 脚本，使用 sh xxx.sh 和 bash xxx.sh 调用执行时结果不同，使用 sh 时会输出许多匪夷所思的结果，而使用 bash 时就完全按照预期。

2. 探究区别
   在一般的 Linux 系统中（例如 Ubuntu ）中，使用 sh 调用执行 shell 脚本相当于打开了 bash 的 POSIX 标准模式，这种模式在一定程度上保证了脚本的跨系统性（跨 UNIX 系统），即 /bin/sh 相当于 /bin/bash --posix，所以二者的一个区别就是有没有开启 POSIX 标准模式。

　　因为bash是sh的增强版本，在我们平常实地操作的时候如果sh这个命令不灵了我们应当使用bash。



## bash-completion

[Mac OS X 安裝 bash-completion](https://www.xiexianbin.cn/mac/2019-01-03-mac-git-bash-completion/index.html)

```
brew install bash-completion
```

Add bash-completion to your ~/.bash_profile:

```
[ -f /usr/local/etc/bash_completion ] && . /usr/local/etc/bash_completion
```



# POSIX解决什么问题

1. [POSIX解决什么问题](https://www.jianshu.com/p/7a17b34e05ee)

> 一般情况下，应用程序通过应用编程接口(API)而不是直接通过系统调用来编程(即并不需要和内核提供的系统调用来编程)。一个API定义了一组应用程序使用的编程接口。它们可以实现成调用一个系统，也可以通过调用多个系统来实现，而完全不使用任何系统调用也不存在问题。实际上，**API可以在各种不同的操作系统上实现给应用程序提供完全相同的接口**，而它们本身在这些系统上的实现却可能迥异。如下图，当应用程序调用printf()函数时，printf函数会调用C库中的printf，继而调用C库中的write，C库最后调用内核的write()。

![img](linux/app-syscall-kernal.png)

应用程序、C库和内核之间的关系

从程序员的角度看，系统调用无关紧要，只需要跟API打交道。相反，内核只跟系统调用打交道，库函数及应用程序是怎么系统调用不是内核所关心的。

> **完成同一功能，不同内核提供的系统调用（一个函数）是不同的**，例如创建进程，linux下是fork函数，windows下是creatprocess函数。好，我现在在linux下写一个程序，用到fork函数，那么这个程序该怎么往windows上移植？我需要把源代码里的fork通通改成creatprocess，然后重新编译...

主流的操作系统有两种，一种是Windows系统，另一种是Linux系统。由于操作系统的不同，API又分为Windows API和Linux API。在Windows平台开发出来的软件在Linux上无法运行，在Linux上开发的软件在Windows上又无法运行，这就导致了软件移植困难，POSIX(Protabl Operation System 可移植操作系统规范)应运而生。

**posix标准的出现就是为了解决这个问题。**linux和windows都要实现基本的posix标准，linux把fork函数封装成posix_fork（随便说的），windows把creatprocess函数也封装成posix_fork，都声明在unistd.h里。这样，程序员编写普通应用时候，只用包含unistd.h，调用posix_fork函数，程序就在源代码级别可移植了。

2. posix 是什么？

可移植操作系统接口Portable Operating System Interface of UNIX，POSIX标准定义了操作系统应该为应用程序提供的接口标准，是IEEE为要在各种UNIX操作系统上运行的软件而定义的一系列API标准的总称。



# 命名空间

Docker核心解决的问题是利用LXC来实现类似VM的功能，从而利用更加节省的硬件资源提供给用户更多的计算资源。而 LXC所实现的隔离性主要是来自内核的命名空间, 其中pid、net、ipc、mnt、uts 等命名空间将容器的进程、网络、消息、文件系统和hostname 隔离开。

[Linux中的clone()函数](https://www.cnblogs.com/xianzhedeyu/archive/2013/06/11/3132146.html)

```c
int clone(int (*fn)(void *), void *child_stack, int flags, void *arg);
```

这里fn是函数指针，我们知道进程的4要素，这个就是指向程序的指针，就是所谓的“剧本", child_stack明显是为子进程分配系统堆栈空间（在linux下系统堆栈空间是2页面，就是8K的内存，其中在这块内存中，低地址上放入了值，这个值就是进程控制块task_struct的值）,flags就是标志用来描述你需要从父进程继承那些资源， arg就是传给子进程的参数

flags取值：

 `CLONE_PARENT` 创建的子进程的父进程是调用者的父进程，新进程与创建它的进程成了“兄弟”而不是“父子”

 `CLONE_FS`     子进程与父进程共享相同的文件系统，包括root、当前目录、umask

 `CLONE_FILES`   子进程与父进程共享相同的文件描述符（file descriptor）表

 `CLONE_NEWNS` 在新的namespace启动子进程，namespace描述了进程的文件hierarchy

 `CLONE_SIGHAND` 子进程与父进程共享相同的信号处理（signal handler）表

 `CLONE_PTRACE` 若父进程被trace，子进程也被trace

 `CLONE_VFORK`  父进程被挂起，直至子进程释放虚拟内存资源

 `CLONE_VM`     子进程与父进程运行于相同的内存空间

 `CLONE_PID`     子进程在创建时PID与父进程一致

 `CLONE_THREAD`  Linux 2.4中增加以支持POSIX线程标准，子进程与父进程共享相同的线程群



# OS

## rootfs

**rootfs根文件系统首先是内核启动时所mount的第一个文件系统，内核代码映像文件保存在根文件系统中，而系统引导启动程序会在根文件系统挂载之后从中把一些基本的初始化脚本和服务等加载到内存中去运行。（**[百度百科](https://baike.baidu.com/item/%E6%A0%B9%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F/8029044?fr=aladdin)）

**尽管内核是 Linux 的核心，但`文件`却是用户与操作系统交互所采用的主要工具。**

Linux根文件系统中一般有如下的几个目录：

- /bin目录

  **该目录下的命令可以被root与一般账号所使用**，如命令：cat、chgrp、chmod、cp、ls、sh、kill、mount、umount、mkdir、[、test

- /sbin目录

  **存放只有系统管理员（俗称最高权限的root）能够使用的命令**，/sbin目录中存放的是基本的系统命令，它们用于启动系统和修复系统等，如命令：shutdown、reboot、fdisk、fsck、init等

- /dev目录

  **该目录下存放的是设备与设备接口的文件**，Linux系统下，以文件的方式访问各种设备，即通过读写某个设备文件操作某个具体硬件。比较重要的文件有/dev/null, /dev/zero, /dev/tty, /dev/lp*等。

- /etc目录

  **该目录下存放着系统主要的配置文件**，例如人员的账号密码文件、各种服务的配置文件等。

- /lib目录

  **该目录下存放共享库和可加载（[驱动程序](https://baike.baidu.com/item/驱动程序)），共享库用于启动系统。**运行根文件系统中的可执行程序，比如：/bin /sbin 目录下的程序。

- /home目录

  **系统默认的用户文件夹**，它是可选的，对于每个普通用户，在/home目录下都有一个以用户名命名的子目录，**里面存放用户相关的配置文件。**

- /root目录：

  **系统管理员（root）的主文件夹，即是根用户的目录**，与此对应，普通用户的目录是/home下的某个子目录。

- /usr目录

  /usr目录的内容可以存在另一个分区中，在系统启动后再挂接到根文件系统中的/usr目录下。里面存放的是共享、只读的程序和数据，这表明/usr目录下的内容可以在多个主机间共享，这些主要也符合[FHS](https://baike.baidu.com/item/FHS)标准的。**/usr中的文件应该是只读的，其他主机相关的，可变的文件应该保存在其他目录下，比如/var。**

- /var目录

  与/usr目录相反，**/var目录中存放可变的数据**，比如spool目录（mail,news），[log文件](https://baike.baidu.com/item/log文件/6646002)，临时文件。

- /proc目录

  这是一个空目录，常作为proc文件系统的挂接点，proc文件系统是个虚拟的文件系统，它没有实际的存储设备，里面的目录，**文件都是由内核临时生成的，用来表示系统的运行状态**，也可以操作其中的文件控制系统。

- /mnt目录

  用于临时挂载某个文件系统的挂接点，通常是空目录，也可以在里面创建一些空的子目录，比如/mnt/cdram /mnt/hda1 。**用来临时挂载光盘、移动存储设备等。**

- /tmp目录：

  用于存放临时文件，通常是空目录，一些需要生成临时文件的程序用到的/tmp目录下，**所以/tmp目录必须存在并可以访问**。



## chroot

CHROOT就是Change Root，也就是改变程序执行时所参考的根目录位置。CHROOT可以增进系统的安全性，限制使用者能做的事。

**作用**

1.限制被CHROOT的使用者所能执行的程式，如SetUid的程式，或是会造成 Load 的Compiler等等。

2.防止使用者存取某些特定档案，如/etc/passwd。

3.防止入侵者/bin/rm -rf /。

4.提供Guest服务以及处罚不乖的使用者。

5.增进系统的安全。



## Cgroup

**cgroups**，其名称源自**控制组群**（control groups）的简写，是[Linux内核](https://baike.baidu.com/item/Linux内核/10142820)的一个功能，用来限制、控制与分离一个[进程组](https://baike.baidu.com/item/进程组/1910809)的[资源](https://baike.baidu.com/item/资源)（如[CPU](https://baike.baidu.com/item/CPU/120556)、[内存](https://baike.baidu.com/item/内存/103614)、[磁盘](https://baike.baidu.com/item/磁盘/2842227)输入输出等）。



cgroups的一个设计目标是为不同的应用情况提供统一的接口，从控制单一进程（像[nice](https://baike.baidu.com/item/nice)）到[操作系统层虚拟化](https://baike.baidu.com/item/操作系统层虚拟化)（像[OpenVZ](https://baike.baidu.com/item/OpenVZ)，Linux-VServer，[LXC](https://baike.baidu.com/item/LXC)）。cgroups提供：

- **资源限制：**组可以被设置不超过设定的[内存](https://baike.baidu.com/item/内存)限制；这也包括[虚拟内存](https://baike.baidu.com/item/虚拟内存)。
- **优先级：**一些组可能会得到大量的CPU或磁盘IO吞吐量。
- **结算：**用来衡量系统确实把多少资源用到适合的目的上。
- **控制：**冻结组或检查点和重启动。



## Linux内核

Linux内核的主要模块（或组件）分以下几个部分：存储管理、CPU和[进程管理](https://baike.baidu.com/item/进程管理)、文件系统、设备管理和驱动、网络通信，以及系统的初始化（引导）、系统调用等。

主要的子系统如下：



### 系统调用接口

SCI 层提供了某些机制执行从[用户空间](https://baike.baidu.com/item/用户空间)到内核的[函数调用](https://baike.baidu.com/item/函数调用)。正如前面讨论的一样，这个接口依赖于[体系结构](https://baike.baidu.com/item/体系结构)，甚至在相同的处理器家族内也是如此。SCI 实际上是一个非常有用的函数调用多路复用和多路分解服务。在 ./[linux](https://baike.baidu.com/item/linux)/kernel 中您可以找到 SCI 的实现，并在 ./linux/arch 中找到依赖于体系结构的部分。



### 进程管理

[进程管理](https://baike.baidu.com/item/进程管理)还包括处理活动进程之间共享 CPU 的需求。内核实现了一种新型的[调度算法](https://baike.baidu.com/item/调度算法)，不管有多少个线程在竞争 CPU，这种算法都可以在固定时间内进行操作。这种算法就称为 O⑴ 调度程序，这个名字就表示它调度多个线程所使用的时间和调度一个线程所使用的时间是相同的。O⑴ 调度程序也可以支持多处理器（称为对称多处理器或 SMP）。您可以在 ./linux/kernel 中找到[进程管理](https://baike.baidu.com/item/进程管理)的源代码，在 ./linux/arch 中可以找到依赖于[体系结构](https://baike.baidu.com/item/体系结构)的源代码。



### 内存管理

内核所管理的另外一个重要资源是内存。为了提高效率，内存是按照所谓的内存页 方式进行管理的（对于大部分[体系结构](https://baike.baidu.com/item/体系结构)来说都是 4KB）。Linux 包括了管理可用内存的方式，以及物理和虚拟映射所使用的硬件机制。

不过[内存管理](https://baike.baidu.com/item/内存管理)要管理的可不止 4KB[缓冲区](https://baike.baidu.com/item/缓冲区)。Linux 提供了对 4KB缓冲区的抽象，例如 slab 分配器。这种内存管理模式使用 4KB缓冲区为基数，然后从中分配结构，并跟踪内存页使用情况，比如哪些内存页是满的，哪些页面没有完全使用，哪些页面为空。这样就允许该模式根据系统需要来动态调整内存使用。

为了支持多个用户使用内存，有时会出现可用内存被消耗光的情况。由于这个原因，页面可以移出内存并放入磁盘中。这个过程称为交换，因为页面会被从内存交换到硬盘上。[内存管理](https://baike.baidu.com/item/内存管理)的源代码可以在 ./[linux](https://baike.baidu.com/item/linux)/mm 中找到。



### 虚拟文件系统

虚拟文件系统（VFS）是 Linux 内核中非常有用的一个方面，因为它为文件系统提供了一个通用的接口抽象。VFS 在 SCI 和内核所支持的文件系统之间提供了一个交换层。

VFS 在用户和文件系统之间提供了一个交换层

在 VFS 上面，是对诸如 open、close、read 和 write 之类的函数的一个通用 API 抽象。在 VFS 下面是文件系统抽象，它定义了上层函数的实现方式。它们是给定文件系统（超过 50 个）的插件。文件系统的源代码可以在 ./[linux](https://baike.baidu.com/item/linux)/fs 中找到。

文件系统层之下是[缓冲区](https://baike.baidu.com/item/缓冲区)缓存，它为文件系统层提供了一个通用函数集（与具体文件系统无关）。这个缓存层通过将数据保留一段时间（或者随即预先读取数据以便在需要是就可用）优化了对[物理设备](https://baike.baidu.com/item/物理设备)的访问。缓冲区缓存之下是[设备驱动程序](https://baike.baidu.com/item/设备驱动程序)，它实现了特定物理设备的接口。



blkio -- 这个子系统为块设备设定输入/输出限制，比如物理设备（磁盘，固态硬盘，USB 等等）。

cpu -- 这个子系统使用调度程序提供对 CPU 的 cgroup 任务访问。

cpuacct -- 这个子系统自动生成 cgroup 中任务所使用的 CPU 报告。

cpuset -- 这个子系统为 cgroup 中的任务分配独立 CPU（在多核系统）和内存节点。

devices -- 这个子系统可允许或者拒绝 cgroup 中的任务访问设备。

freezer -- 这个子系统挂起或者恢复 cgroup 中的任务。

memory -- 这个子系统设定 cgroup 中任务使用的内存限制，并自动生成由那些任务使用的内存资源报告。

net_cls -- 这个子系统使用等级识别符（classid）标记网络数据包，可允许 Linux 流量控制程序（tc）识别从具体 cgroup 中生成的数据包。

ns -- 名称空间子系统。



# Linux Container

[Linux Containers](https://linuxcontainers.org/) (LXC) 是一种操作系统级别的虚拟化方法，用于在单个控制主机（LXC 主机）上运行多个隔离的 Linux 系统（容器）。它不提供虚拟机，而是提供一个拥有自己的CPU、内存、块I/O、网络等空间和资源控制机制的虚拟环境。这是由LXC 主机上的 Linux 内核中的[命名空间](https://en.wikipedia.org/wiki/Linux_namespaces)和[cgroups](https://wiki.archlinux.org/title/Cgroups)功能提供的。它类似于 chroot，但提供了更多的隔离。

使用容器的替代方案是[systemd-nspawn](https://wiki.archlinux.org/title/Systemd-nspawn)或[Docker](https://wiki.archlinux.org/title/Docker)。

摘自：

- https://wiki.archlinux.org/title/Linux_Containers

- https://linuxcontainers.org/





# rsync

rsync是[linux系统](https://baike.baidu.com/item/linux系统/1732935)下的数据镜像备份工具。使用快速增量备份工具Remote Sync可以远程同步，支持本地复制，或者与其他SSH、rsync主机同步。

目前，已支持跨平台，可以在Windows与Linux间进行数据同步。

https://baike.baidu.com/item/rsync/8086338?fr=aladdin

快速：第一次同步时 rsync 会复制全部内容，但在下一次只传输修改过的文件。rsync 在传输数据的过程中可以实行压缩及[解压缩](https://baike.baidu.com/item/解压缩)操作，因此可以使用更少的带宽。



**默认规则是文件大小或修改时间有变动**



## 服务器端启动

```bash
usr/bin/rsync --daemon --config=/etc/rsyncd/rsyncd.conf
```

也可以在/etc/rc.d/rc.local里加入让系统自动启动等.



## 客户端同步

```bash
# rsync -参数 用户名@同步服务器的IP::rsyncd.conf中那个方括号里的内容 本地存放路径
rsync -avzP nemo@192.168.10.1::nemo /backup
```

说明：

-a 参数，相当于-rlptgoD，-r 是递归 -l 是链接文件，意思是拷贝链接文件；-p 表示保持文件原有权限；-t 保持文件原有时间；-g 保持文件原有用户组；-o 保持文件原有属主；-D 相当于[块设备](https://baike.baidu.com/item/块设备)文件；

-z 传输时压缩；

-P 传输进度；

-v 传输时的进度等信息；



## 安装

Ubuntu安装: 

```bash
sudo apt-get install rsync
```

RedHat: 

```bash
yum install rsync
```



## rsync 协议

除了使用 SSH，如果另一台服务器安装并运行了 **rsync** 守护程序，则也可以用`rsync://`协议（默认端口873）进行传输。具体写法是服务器与目标目录之间使用双冒号分隔`::`。

> ```bash
> $ rsync -av source/ 192.168.122.32::module/destination
> ```

注意，上面地址中的`module`并不是实际路径名，而是 **rsync** 守护程序指定的一个资源名，由管理员分配。

如果想知道 **rsync** 守护程序分配的所有 module 列表，可以执行下面命令。

> ```bash
> $ rsync rsync://192.168.122.32
> ```

**rsync** 协议除了使用双冒号，也可以直接用`rsync://`协议指定地址。

> ```bash
> $ rsync -av source/ rsync://192.168.122.32/module/destination
> ```





# EOF

EOF，它是end of file的缩写，表示"文字流"（stream）的结尾。这里的"文字流"，可以是文件（file），也可以是标准输入（stdin）。

```c
#define EOF (-1)
```

**在Linux系统之中，EOF根本不是一个字符，而是当系统读取到文件结尾，所返回的一个信号值（也就是-1）。**至于系统怎么知道文件的结尾，资料上说是通过比较文件的长度。



Linux中，在新的一行的开头，按下Ctrl-D，就代表EOF（如果在一行的中间按下Ctrl-D，则表示输出"标准输入"的缓存区，所以这时必须按两次Ctrl-D）；Windows中，Ctrl-Z表示EOF。（顺便提一句，Linux中按下Ctrl-Z，表示将该进程中断，在后台挂起，用fg命令可以重新切回到前台；按下Ctrl-C表示终止该进程。）



## 参考资料

- [iptables入门指南 --- iptables详解 ---iptbales 防火墙](https://www.cnblogs.com/liang2580/articles/8400140.html)

- [EOF是什么？- 阮一峰](https://www.ruanyifeng.com/blog/2011/11/eof.html)

