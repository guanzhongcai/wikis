# Docker



## docker和传统VM比较

| 比较项         | 传统VM                                                       | Docker                                                   |
| -------------- | ------------------------------------------------------------ | -------------------------------------------------------- |
| 解决的核心问题 | 资源配置                                                     | 应用的开发、测试、部署                                   |
| 实现方法       | 通过Hypervisor层对宿主机硬件资源进行虚拟化。                 | 直接使用宿主机操作系统调度硬件资源，资源利用率上远超VM。 |
| 创建速度       | 提前统一配置、统一管理                                       | 容器是利用宿主机的系统内核创建的，可以在几秒内大量创建   |
| 缺点           | 会有计算、I/O、网络性能损耗，因本质多了一层软件，运行一个完整的OS。 | 因共享内核，对安全和隔离问题做出了一定妥协。             |

由于容器不需要进行硬件虚拟以及运行完整操作系统等额外开销，`Docker` 对系统资源的利用率更高。无论是应用执行速度、内存损耗或者文件存储速度，都要比传统虚拟机技术更高效。



### docker技术实现要点

Go语言开发。基于Linux内核的CGroup + Namespace，以及AUFS类的Union FS技术，是对进程进行封装隔离的轻量级容器虚拟化。

```mermaid
graph LR;
硬件资源-->|落实资源管理|CGroup-->|封装|LXC-->|封装|Docker
```

LXC是Linux原生的容器工具，利用LXC容器能有效地将单个操作系统管理的资源划分到鼓励的组中，以更好地在孤立的组之间平衡有冲突的资源使用需求。

这样既不需要指令级模拟，也不需要即时编译。容器可以在核心CPU本地运行命令，而不需要任何专门的解释机制。



### tips

- Linux中都有一个进程号为1的init进程，系统服务的父进程都是init。但是**Docker容器中进程号为1 的进程号是bash**，而不是init。这得益于LXC功能。宿主机器中运行的Docker服务就是该容器中Ubuntu系统的init进程。
- 每个运行的容器仅仅是宿主机器中运行的一个进程而已，容器中运行的程序其实也是宿主机器中的一个进程。
- Docker通过**CGroup**将属于每个容器的进程分为一组进行资源（内存、CPU、网络、硬盘）控制，通过**Namespace**将属于同一个容器的进程划分为一组，使分属于同一个容器的进程拥有独立的进程名字和独立的进程号！

- 在Docker出现之前。很多技术方案就是**直接令应用调用CGroup隔离**来运行资源的，但是这种隔离是粗粒度、硬编码的，想同时隔离资源和进程组，Docker方案做的最好。
- docker的启动需要相当大的用户权限，所以实际上docker服务用户组的权限相当于root，并且该权限会通过fork传递下去，这个是安全隐患。

### docker进程模型

![img](pictures/docker-workflow.jpg)

```mermaid
graph LR;
dockerd-->|fork|docker-containerd-->|fork|docker-contatinerd-shim-->|run|镜像
```

- docker服务启动的第一个进程是 `/usr/bin/dockerd`，是这个docker服务端启动的入口，即Docker Daemon、Docker Engine

- dockerd的子进程docker-container的，是docker服务端的核心进程，负责与docker客户端、docker容器进行通信交互，例如执行 `docker run`命令，fork出docker容器进程

  ```bash
  # 启动参数：listen，打开一个sock描述符，实现所有docker容器和docker客户端之间的通信
  /../docker-containerd -l unix:///var/run/Docker/libcontainerd/Docker-containerd.sock
  ```

  

### unix domain socket

一种更高效的IPC机制，使用Socket API，将应用层数据从一个进程复制到另一个进程，不需要结果网络协议栈！

| 比较项         | Unix Domain socker                                           | 网络通信                                                   |
| -------------- | ------------------------------------------------------------ | ---------------------------------------------------------- |
| 可靠性         | IPC机制本质上是可靠的通信                                    | 网络协议是为不可靠的通信设计的                             |
| 原理           | 将应用层数据从一个进程复制到另一个进程                       | 通过网络协议栈，打包拆包、计算校验、维护序号、应答等做通信 |
| 地址           | 是一个socker类型的文件在文件系统中的路径。<br/>这个文件由bind()方法创建，若已存在，则返回错误 | 是IP地址加端口号                                           |
| address family | AF_UNIX                                                      | AF_INET                                                    |



### 容器中进程启动的两种模式

所有docker容器内启动的进程全部都是宿主机上的独立进程。该进程号是不是docker容器进程本身，要依据dockerfile的写法：

| 比较项    | shell方式执行进程                                            | exec方式执行进程                                          |
| --------- | ------------------------------------------------------------ | --------------------------------------------------------- |
| 命令格式  | /bin/sh -c "executable param1 param2"                        | CMD ["executable", "param1", "param2"]                    |
| redis示例 | ...<br />CMD "/usr/bin/redis-server"                         | ...<br />CMD ["/usr/bin/redis-server"]                    |
| ps -ef    | PID CMD<br />1 /bin/sh -c "/usr/bin/redis-server"<br />5 /usr/bin/redis-server *:6379<br />8 ps -ef | PID CMD<br />1 /usr/bin/redis-server *:6379<br />7 ps -ef |
| 释义      | 1号进程为shell                                               | 1号进程为redis-server                                     |
| 容器退出  | 需要对容器进程增加SIGTERM的处理逻辑，否则docker stop不能做到优雅退出，docker daemon默认10秒超时后退出 | docker stop能自动优雅退出                                 |

结论：

如果容器中包含多个进程，需要1号进程能够正确地传播SIGTERM信号来结束素有的子进程，之后再推出。

**令每个容器中只包含一个进程，同时都采用exec模式启动进程。**也是docker官方文档推荐做法。





## docker与微服务

| 微服务                    | docker                                                       |
| ------------------------- | ------------------------------------------------------------ |
| X轴水平克隆、水平扩展能力 | docker镜像快速部署，镜像即代码                               |
| Y轴功能分模块解耦         | docker镜像独立完整，用docker-compose等技术串联docker容器启动 |
| Z轴分区部署               | docker与数据服务结合，一键式扩展？                           |

在常见的传统部署模式中：

1. 用边缘节点Endpoint来做公网入口
2. 配合防火墙和三层交换机进行内外网隔离和网络安全区的划分
3. 边缘节点会通过nginx/haproxy或者lvs进行四层或七层上的分发和路由
4. 边缘节点的高可用性可以通过keepalived进行主备，通过冗余节点保证CAP定理中的AP（可用性和分区容错性）

```mermaid
graph TD;
公网入口-->边缘节点Endpoint-->|网络安全|防火墙-->|内外网隔离|三层交换机-->|分发路由|Nginx/HAproxy/lvs-->|高可用主备冗余|keepalived
```





## docker的三大组件

- 镜像
- 容器（运行时）
- 仓库



## 容器的本质

容器 = CGroups + Namespace + Rootfs

### Namespace

Namespac是Linux提供的内核级别的环境隔离方法。

基于实现了内部资源无法访问外部资源的简单隔离的chroot技术。

Namespace的实现基于三个系统方法：

1. clone()：实现线程的系统调用，用来创建一个新的进程
2. unshare()：把某个进程脱离某个Namespace
3. setns()：把某个进程加入某个Namespace

**namespace概念**

namespace 是 Linux 内核用来隔离内核资源的方式。通过 namespace 可以让一些进程只能看到与自己相关的一部分资源，而另外一些进程也只能看到与它们自己相关的资源，这两拨进程根本就感觉不到对方的存在。具体的实现方式是把一个或多个进程的相关资源指定在同一个 namespace 中。

Linux namespaces 是对全局系统资源的一种封装隔离，使得处于不同 namespace 的进程拥有独立的全局系统资源，改变一个 namespace 中的系统资源只会影响当前 namespace 里的进程，对其他 namespace 中的进程没有影响。

**namespace用途**

Linux 内核实现 namespace 的一个主要目的就是实现轻量级虚拟化(容器)服务。在同一个 namespace 下的进程可以感知彼此的变化，而对外界的进程一无所知。这样就可以让容器中的进程产生错觉，认为自己置身于一个独立的系统中，从而达到隔离的目的。也就是说 linux 内核提供的 namespace 技术为 docker 等容器技术的出现和发展提供了基础条件。


### Rootfs

Rootfs是docker容器在**启动时**其内部进程的文件系统，即docker容器的根目录。

该目录下有docker容器所需要的：系统文件、工具、容器文件等。和Linux系统内核启动时挂载的Rootfs目录的思想。

docker源码中通过下面的方法在进程中切换Rootfs：

```go
syscall.PivotRoot(rootfs, pivotDir)
```



### CGroups

Control Groups，CGroups 就是把进程放到一个组里面统一加以控制。具体提供了：

1. 资源限制：Resource LImitation，超过上限就发出OOM信息
2. 优先级分配：Prioritization，对不同进程分配CPU时间片数量及硬盘IO带宽，相当于控制了进程运行的优先级
3. 资源统计：Accounting，如CPU使用时长、内存用来，适用于计费
4. 进程控制：Control，对进程组执行挂起、恢复等操作

CGroups本质是内核附加在程序上的一系列的钩子（hook）。

```bash
# 给docker容器设置内存限制为128M
docker run -m 128m redis

# docker会在系统的hierarchy中为每个容器创建CGroups
cd /sys/fs/CGroup/memory/Docker/$container_id

# 查看CGroups的内存限制
cat memory.limit_in_bytes
134217728

# 查看CGroups中进程所使用的内存大小
cat memory.usage_in_bytes
430080
```



## Docker容器的运行时模型

### Linux中的特殊进程

- ID为0的是调度进程，该进程是内核的一部分，不执行如何磁盘上的程序。
- ID为1的是init进程，init通常读取与系统有关的初始化文件例如：/etc/rc*文件、/etc/inittab文件、/etc/init.d中的文件
- ID为2的页守护进程，负责支持虚拟存储器系统的分页操作



### Linux进程模型

Linux中的父进程用fork命令创建子进程，然后调用exec执行子进程函数，进程ID可以复用，但要使用延迟算法，防止将新进程误认为使用同一ID的某个已经终止的先前进程。

Linux在进行fork操作的时候，会首先调用 copy_process 函数，然后根据父进程传入的flag判断是否要新建Namespace，随后复制父进程的进程描述符 task_struct。

task_struct中包括当前进程的各种系统配置信息，包括：

- 网络描述
- PID描述
- UID描述
- MNT描述等



### Docker进程模型

Docker启动的时候，也是：

```mermaid
graph LR;
Docker-containerd-->|fork命令|子进程-->|exec方式|启动
```

容器进程被fork之后，便创建了Namespace，下面就执行一系列的初始化操作了，分三个阶段：

1. dockerinit负责初始化网络栈
2. ENTRYPOINT负责完成用户态配置
3. CMD负责启动入口

启动后的docker容器和docker daemon就是通过sock文件描述符进行通信的。



## Docker逻辑架构

分为：Client + Docker_Host + Registry

所有Docker Client的命令，docker build、docker pull、docker run都可以使用HTTPS、HTTP的restful api来通信：

```bash
vi /etc/sysconfig/docker
# 添加
DOCKER_OPTS="-H tcp://0.0.0.0:2375"
server docker restart
# test
curl -v -X GET localhost:2375/_ping
```

基于这些RESTful API，可以在自己的平台上实现从编译到发布的全流程自动化。

## Docker Registry选型: Habor

使用Habor。

- Habor以Docker公司开源的Registry为基础，提供了管理UI、基于角色的访问控制、AD/LDAP基础、日志审核等功能，还支持中文。是VMVare公司开源的企业级Docker Registry项目。

- Habor支持Docker Compose一键式安装。

- Habor支持镜像复制，可以在开发环境、测试环境、生产环境做镜像的复制



## Docker单机网络模型

这个Docker体系中最复杂、对生产上线最具影响力的就是网络模式。

### Bridge模式

默认的。使用etho0虚拟网桥进行通信。

执行 docker run -p 命令时，Docker实际是在 iptables 上遵循 DNAT 规则，实现了端口转发的功能。

Docker安装好后，Docker守护进程就会调用Linux内核，生成一个虚拟网桥。所有容器的地址段都是：

```ip
172.17.0.1/16
```

Linux网桥的本质是用一组代码模拟网络协议栈，类似软件交换机。



### Host模式

此模式容器将不会获得一个独立的Network Namespace，而是会和宿主机共用一个Network Namespace，直接使用宿主机的端口和IP地址。

```bash
docker run -d --net=host --privileged=true tomcat

docker exec -it 84ee1dce1de806e50bcf19a5042c81713c1a841eab72c8c14d66ab93d5d73c84 ip addr show
```

`--privileged=true` 命令使容器会被允许直接配置主机的网络堆栈。



### Container模式

将新建容器的进程让道一个已存在的容器的网络栈中，两者的进程直接通过lo回环接口进行通信：

```bash
docker run --net=container:$container_id -d tomcat
```



### None模式

令Docker新容器放到隔离的网络栈中，但是不进行网络配置。

```bash
docker run  --net=none -it tomcat ip addr show
```

![docker-network-models](docker-network-models.png)



## Docker底层原理

- docker是面向软件开发者的，没有虚拟机的完整的os概念和硬件资源的预划分
- docker依托自己的docker engine实现了硬件资源的调度，移除了VM的hypervisor的概念

![](./pictures/docker_vs_vm.png)



### Docker for Mac配置镜像源：

在任务栏点击 Docker Desktop 应用图标 -> Perferences，在左侧导航菜单选择 Docker Engine，在右侧输入栏编辑 json 文件。将

https://n8bn2y81.mirror.aliyuncs.com加到"registry-mirrors"的数组里，点击 Apply & Restart按钮，等待Docker重启并应用配置的镜像加速器。



## 镜像分层原理

- 通过history查看镜像分层

  ```shell
  docker history nginx
  ```

- 镜像由若干个层组成，每个有size的层都放置不同的存储目录中

- 这些层是由docker的overlay2驱动的

- 使用 `docker inspect containerID` 查看每层文件的存储信息

  ```
          "GraphDriver": {
              "Data": {
                  "LowerDir": "/var/lib/docker/overlay2/0f75953d648f7360dd8056f932950d12141d89591a0498d8050a47eed5be1ebe-init/diff:/var/lib/docker/overlay2/e89983d5fa44d196b9bcc7e244bc3242207e576a9c08e6eb38fe6416111f813e/diff:/var/lib/docker/overlay2/e4163f54be653626a97e00dd6d7bc28c1449f55ef3b766a4791e1759ffbcb2b9/diff:/var/lib/docker/overlay2/a3764ea52f54653232be22ddcd516866ec56e65f09c67c61ed9205a04c913bf9/diff:/var/lib/docker/overlay2/bfeff87f5084396e793b015d060eec517a6e2fb65435b7668450e274cf495c19/diff:/var/lib/docker/overlay2/d3a68dbf675141cf3d2575f965e3d0b2c4aed69fb9c2046d9687bc83af4428ba/diff:/var/lib/docker/overlay2/9a2625e05ef607bc8b507efd8f9d6ef37663ecee2e449927765640b594871580/diff",
                  "MergedDir": "/var/lib/docker/overlay2/0f75953d648f7360dd8056f932950d12141d89591a0498d8050a47eed5be1ebe/merged",
                  "UpperDir": "/var/lib/docker/overlay2/0f75953d648f7360dd8056f932950d12141d89591a0498d8050a47eed5be1ebe/diff",
                  "WorkDir": "/var/lib/docker/overlay2/0f75953d648f7360dd8056f932950d12141d89591a0498d8050a47eed5be1ebe/work"
              },
              "Name": "overlay2"
          },
  ```

- 镜像层是Readonly的，容器层是Read/Write的

- 静态的只读层是可以被别的镜像锁引用的，避免不必要的磁盘占用

- 镜像层都是映射到宿主机的/var/lib/docker/overlay2目录中，分为 *下级目录、上级目录、合并目录、工作目录*



## docker网络

![image-20210606162702465](../images/docker-network.png)



# Docker Machine

Docker Machine 是一种可以让您在虚拟主机上安装 Docker 的工具，并可以使用 docker-machine 命令来管理主机。

Docker Machine 也可以集中管理所有的 docker 主机，比如快速的给 100 台服务器安装上 docker。

```bash
# 1、列出可用的机器
docker-machine ls

# 2、创建机器
docker-machine create --driver virtualbox test

# 3、查看机器的 ip
docker-machine ip test

# 4、停止机器
docker-machine stop test

# 5、启动机器
docker-machine start test

# 6、进入机器
docker-machine ssh test
```



# Swarm 集群管理

Docker Swarm 是 Docker 的集群管理工具。它将 Docker 主机池转变为单个虚拟 Docker 主机。

不同主机间的容器互通，要组织成局域网，需要用到 `docker swarm`，主要用在小型环境中使用。大型用kubernetes。



swarm 集群由管理节点（manager）和工作节点（work node）构成。

- **swarm mananger**：负责整个集群的管理工作包括集群配置、服务管理等所有跟集群有关的工作。
- **work node**：即图中的 available node，主要负责运行相应的服务来执行任务（task）。

![img](../images/services-diagram.png)



```bash
# 创建 docker 机器：
docker-machine create -d virtualbox swarm-manager

docker-machine ls

# 初始化 swarm 集群，进行初始化的这台机器，就是集群的管理节点。
docker-machine ssh swarm-manager
docker swarm init --advertise-addr 192.168.99.107 #这里的 IP 为创建机器时分配的 ip。

# output
# docker swarm join --token SWMTKN-1-24chiircvefa75pyya86cq7uy3hrplrbteqcctpnmr221xtk0f-88xty7cs1qmwso3y9lshs4bk8 192.168.65.3:2377

# 增加工作节点node：
docker swarm join --token SWMTKN-1-4oogo9qziq768dma0uh3j0z0m5twlm10iynvz7ixza96k6jh9p-ajkb6w7qd06y1e33yrgko64sk 192.168.99.107:2377

# 查看节点
docker node ls

# 生成Manage join tokens
docker swarm join-token manager

# 输出
# docker swarm join --token SWMTKN-1-24chiircvefa75pyya86cq7uy3hrplrbteqcctpnmr221xtk0f-a7wl14eo9qwjbm7gbz55kw44d 192.168.65.3:2377
```

实际命令：

```bash
# 4节点的nginx服务，集群中任意nodeIP:81都可以访问到这个服务，和容器副本数无关
docker service create --name my_nginx -p 81:80 --replicas 4 nginx

# 查看container/node分布
docker service ps my_nginx

# 扩缩容
docker service update --replicas 2 my_nginx
```



# docker stack

```bash
docker stack deploy -c docker-compose.yml stack-demo
```

层级关系，由外到内：

1. docker node ls - 节点
2. docker stack ls - 堆栈
3. service ls - 服务
4. docker service ps - task



## dockerfile

![image-20210606220210762](../images/dockerfile.png)





![image-20210606220604049](../images/dockerfile-flow.png)

### WORKDIR

会一层层的迭路径

```dockerfile
WORKDIR /data
WORKDIR bb # 此处表示相对路径，即位于容器中的目录在：/data/bb
ADD https://mirros.../tomcat-8.5.tar.gz /data/bb

ENV BASE_DIR /data/bb
COPY . $BASE_DIR


```



### ADD

可以指定一个url，复制到容器中，甚至是个归档包，会做自动的解包



### ENTRYPOINT

类似于 CMD 指令，但其不会被 docker run 的命令行参数指定的指令所覆盖，而且这些命令行参数会被当作参数送给 ENTRYPOINT 指令指定的程序。

但是, 如果运行 docker run 时使用了 --entrypoint 选项，将覆盖 CMD 指令指定的程序。

**优点**：在执行 docker run 的时候可以指定 ENTRYPOINT 运行所需的参数。

**注意**：如果 Dockerfile 中如果存在多个 ENTRYPOINT 指令，仅最后一个生效。

格式：

```dockerfile
ENTRYPOINT ["<executeable>","<param1>","<param2>",...]
```

可以搭配 CMD 命令使用：**一般是变参才会使用 CMD ，这里的 CMD 等于是在给 ENTRYPOINT 传参**，以下示例会提到。

示例：

假设已通过 Dockerfile 构建了 nginx:test 镜像：

```dockerfile
FROM nginx

ENTRYPOINT ["nginx", "-c"] # 定参
CMD ["/etc/nginx/nginx.conf"] # 变参 
```

1、不传参运行

```bash
$ docker run  nginx:test
```

容器内会默认运行以下命令，启动主进程。

```bash
nginx -c /etc/nginx/nginx.conf
```

2、传参运行

```
$ docker run  nginx:test -c /etc/nginx/new.conf
```

容器内会默认运行以下命令，启动主进程(/etc/nginx/new.conf:假设容器内已有此文件)

```bash
nginx -c /etc/nginx/new.conf
```

Example：

```dockerfile
...
ENTRYPOINT ["ls", "/data"]
```

构建镜像时附加一个参数，会和ENTRYPOINT一并执行，即打印出两个目录下的ls：

```bash
docker run imageName /data/bb
```



### LABEL

```dockerfile
FROM alpine
LABEL k="v" k1="v1"
```

`LABEL`通常紧随`FROM`之后的下一行，在`docker inspect`的时候，会看到`LABEL`信息，用以打标签等目的



### ONBUILD

```dockerfile
FROM alpine
LABEL k="v" k1="v1"
ONBUILD ENV C=100
CMD echo $C
```

在自己的镜像里产生的容器不会输入任何东西，当如果以此为基础镜像而构建出来的容器，就会输出100！

即：`ONBUILD`只在被当做基础镜像时，命令才有效。



### ARG

只在`build`时有效的变量声明，在`docker`命令中可以修改（囿于成见）：

```bash
docker build --build-args A=3 .
```



### ENV

在`build`和`runtime`是都有效的变量声明

```dockerfile
ARG A=2
ENV B $A
```



### MATAINER

镜像的作者信息，在 `docker inspect` 中可以看到



### 对续行的参数按照字母表排序

以防止重复导入包

```dockerfile
RUN apt-get 
```



### 多阶段构建

多阶段构建可以让我们大幅度减小最终的镜像大小，而不需要去想办法减少中间层和文件的数量。因为镜像是在生成过程的最后阶段生成的，所以可以利用生成缓存来最小化镜像层。



例如，如果你的构建包含多个层，则可以将他们从变化频率较低（以确保生成缓存可重用）到变化频率较高的顺序排序：

- 安装构建应用程序所需的依赖工具
- 安装或更新依赖项
- 构建你的应用

比如我们构建一个Go应用程序的Dockerfile可能类似于这样：

```dockerfile
FROM golang:1.11-alpine AS build

# 安装项目需要的工具
# 运行 `docker build --no-cache .` 来更新依赖

RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

# 通过 Gopkg.toml 和 Gopkg.lock 获取项目的依赖
# 仅在更新 Gopkg 文件时才重新构建这些层（COPY/ADD时会自动检查文件是否变化）
COPY Gopkg.lock Gopkg.toml /go/src/project/
WORKDIR /go/src/project/
# 安装依赖库
RUN dep ensure -vendor-only

# 拷贝整个项目进行构建
# 当项目下面有文件变化的时候该层才会重新构建

COPY . /go/src/project/
RUN go build -o /bin/project

# 将打包后的二进制文件拷贝到 scratch 镜像下面，将镜像大小降到最低
FROM scratch 
COPY --from=build /bin/project /bin/project
ENTRYPOINT ["/bin/project"]
CMD ["--help"]
```



### 尽量使用管道

使用管道操作，所以没有中间文件需要删除

```dockerfile
RUN mkdir -p /user/src/things \
		&& curl -SL http://example.com/big.tar.xz \
		| tar -xJC /usr/src/things \
		&& make -C /usr/src/things all
```

带管道的命令，其返回值是最有一台命令的返回值。所以如果管道前的命令出错而管道后的命令执行正常，则docker不会认为这条指令有问题。如果需要所有的管道命令都正常执行，可以增加 `set -o pipefail`:

```dockerfile
RUN ["bin/bash", "-c" "set -o pipefail && wget -O https://some.site | wc -l > /number"]
```

部分shell（比如默认的sh）不支持 `set -o pipefail`，但`bash`支持。



### extra_hosts

添加主机名映射。类似 `docker client --add-host`。

```yaml
extra_hosts:
 - "somehost:162.242.195.82"
 - "otherhost:50.31.209.229"
```

以上会在此服务的内部容器中 /etc/hosts 创建一个具有 ip 地址和主机名的映射关系：

```
162.242.195.82  somehost
50.31.209.229   otherhost
```



### secrets

存储敏感数据，例如密码：

```yaml
version: "3.1"
services:

mysql:
  image: mysql
  environment:
    MYSQL_ROOT_PASSWORD_FILE: /run/secrets/my_secret
  secrets:
    - my_secret

secrets:
  my_secret:
    file: ./my_secret.txt
```



### stop_grace_period

指定在容器无法处理 SIGTERM (或者任何 stop_signal 的信号)，等待多久后发送 SIGKILL 信号关闭容器。

```yaml
stop_grace_period: 1s # 等待 1 秒
stop_grace_period: 1m30s # 等待 1 分 30 秒 
```

默认的等待时间是 10 秒。



### stop_signal

设置停止容器的替代信号。默认情况下使用 SIGTERM 。

以下示例，使用 SIGUSR1 替代信号 SIGTERM 来停止容器。

```yaml
stop_signal: SIGUSR1
```



### save & load

```bash
docker save -o imagexxx.tar imagexxx:0.25.0
docker load -i imagexxx.tar
```



## docker-compose

- `-p, --project-name NAME` 指定项目名称，默认使用所在目录名称作为项目名
- `--x-network-driver DRIVER`  指定网络后端的驱动，默认为 `bridge`

- 建议把 `docker-compose`命名给起个alias别名：`dc`
- 容器服务自动注册IP到consul中？



### up

- 默认情况下，若服务容器已经存在，`docker-compose up` 将会尝试停止容器，然后重新创建（保持使用 `volumes-from` 挂在的卷），以保证新启动的服务匹配 `docker-compose.ym`文件的最新内容

```bash
dc up
dc stop
dc start
dc down
dc rm
```



### volume

```bash
# 删除unused映射卷
docker volume prune

# 显示正在使用的映射卷
docker volume ls

# 手动创建此卷名
docker volume create tomcatwebapps

# 容器内只读，宿主机读写，-v 宿主机路径:容器内路径
docker run -d -p 8090:8080 --name tomcat90 -v /root/apps:/usr/local/tomcat/webapps:ro tomcat:8.0-jre8
```

### network

```bash
# 手动创建网络
docker network create -d bridge hello

# 子网192.168.0.0/16 表示前面16位固定为192.168，后面的随意，即可选IP范围为：65535-2
# 子网192.168.0.0/24 表示前面24位固定为192.168.0，后面的随意，即可选IP范围为：255-1
# 255.255.0.0 为广播地址
docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 mynet
```

### 网络互连

```bash
# tomcat01将会出现两个IP！
# 两个网络直接互通是麻烦的，但可以把容器加入到了另一个网络中！
docker network connect mynet tomcat01

# 如此这般，就ping的通了
docker exec -it tomcat01 ping tomcat-net-01 
```



### docker

```bash
# 删除停止了的容器
docker rm -f $(docker ps -qa)

# 删除不用的image: 删除所有tag的mycentos7镜像
docker rmi -f $(docker images mycentos7 -qa)
```



### docker-compose.yml示例

文件夹hello目录中的 `docker-compose.yml`:（目录名hello也即project名）

```yaml
version "3.2"

services:
	tomcat01:
		container_name: tomcat01 # 指定容器名称 默认是项目名_服务名
		image: tomcat:8.0-jre8
		ports:
			- "8080:8080"
		volumes: # 完成宿主机和容器中目录数据卷共享
			#- /root/apps:/user/local/tomcat/webapps # 使用自定义路径映射 前者是容器中path
			- tomcatwebapps01:/user/local/tomcat/webapps
		networks: # 代表当前服务使用哪个网络桥
  		- hello # 表示和tomcat02在一个网络中
    links:
    	- tomcat02 # 要访问它，会在/etc/hosts里生成一个域名解析
  	depends_on:
  		- tomcat02
  		- redis
  		- mysql
  	healthcheck:
    	test: ["CMD", "curl", "-f", "http://localhost"]
    	interval: 1m30s
    	timeout: 10s
    	retries: 3

	tomcat02:
		image: tomcat:8.0-jre8
		ports:
			- "8080:8080"
		volumes: # 完成宿主机和容器中目录数据卷共享
			- tomcatwebapps02:/user/local/tomcat/webapps
		networks:
  		- hello
	
	mysql:
		image: mysql:5.7.32
		container_name: mysql
		ports:
			- "3307:3306"
		volumes:
    	- mysqldata:/var/lib/mysql
    	- mysqlconf:/etc/mysql
    environment:
    	- MYSQL_ROOT_PASSWORD=root
    env_file:
    	- ./mysql.env
    network:
    	- hello
	
	redis:
		image: redis:5.0.10
		container_name: redis
		ports:
			- "6379:6379"
		volumes:
    	- redisdata:/data
    networks:
    	- hello
    command: "redis-server --appendonly yes" # 覆盖容器内的原默认命令	
	
volumes:	# 声明上面服务所使用的自动创建的卷名
	tomcatwebapps01: # 声明制定的卷名 compose会自动创建该卷名，在前面加入项目名（即dc.yml所在的目录名）
		external:
			false
	tomcatwebapps02:
	mysqldata: 	# inspect会发现目录在：/var/lib/docker/volumes/hello_mysqldata/_data
	mysqlconf:

networks: # 定义服务用到的桥
	hello:	# 定义上面用到的网桥名称 默认创建就是bridge类型
		external:
			true	# 使用外部指定的网桥 需要已存在
```

`env_file`: 替换yml中的environment。从文件中获取环境变量。env文件中支持`#`注释，且都是`key=value`的写法

`mysql.env`：

```env
# MySQL的环境变量
MYSQL_ROOT_PASSWORD=root
```

useful commands:

```bash
# 使用服务id进入容器
docker-compose exec tomcat01 bash

# 列出当前运行的
docker-compose ps

# -f强制 -v数据卷
docker-compose rm [SERVICE] -f -v

# 重启 -t超时 默认10秒
docker-compose restart [SERVICE] -t 

# 查看容器内运行的进程
docker-compose top [SERVICE]

# 暂停和继续
docker-compose pause [SERVICE]
docker-compose unpause [SERVICE] 
```

- docker面向的是容器
- docker-compose面向的是服务

- 网桥？？



# docker容器日志

每个容器的日志默认都会以json-file的格式存储于 `/var/lib/docker/containers/containerID`下，生产中，对此目录做 `filebeat`的日志收集到ELK中。



# portainer

docker最专业的可视化工具。生产环境用的非常多！！

```bash
docker pull portainer/portainer
docker volume create portainer_data
docker run -d -P --name=portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer

# docker run -d -p 8000:8000 -p 9000:9000 --name=portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer

# potainer-agent
docker run -d -p 9001:9001 --name portainer_agent --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v /var/lib/docker/volumes:/var/lib/docker/volumes portainer/agent
```

登录网页后端：`http://localhost:9000`

- 操作另一台电脑，需要把 `/etc/docker/daemon.conf`中添加`http://remoteIP:2357`，然后从起docker，portainer这边就可管理了。



# ASP .net core 

方法1：

```dockerfile
# syntax=docker/dockerfile:1
FROM mcr.microsoft.com/dotnet/sdk:5.0 AS build-env
WORKDIR /app

# Copy csproj and restore as distinct layers
COPY *.csproj ./
RUN dotnet restore

# Copy everything else and build
COPY ../engine/examples ./
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:3.1
WORKDIR /app
COPY --from=build-env /app/out .
ENTRYPOINT ["dotnet", "aspnetapp.dll"]
```

方法2：

```dockerfile
  # syntax=docker/dockerfile:1
  FROM mcr.microsoft.com/dotnet/aspnet:5.0
  COPY bin/Release/netcoreapp3.1/publish/ App/
  WORKDIR /App
  ENTRYPOINT ["dotnet", "aspnetapp.dll"]
```

此方法假定您的项目已经构建，并且它从发布文件夹复制构建工件。请参阅有关[容器化 .Net Core 应用程序](https://docs.microsoft.com/en-us/dotnet/core/docker/build-container?tabs=windows#create-the-dockerfile)的 Microsoft 文档。

`docker build`这里的步骤将比方法 1 快得多，因为所有工件（artifacts）都在`docker build`步骤之外构建，并且与构建基础镜像相比，基础镜像的大小要小得多。

这种方法是 Jenkins、Azure DevOps、GitLab CI 等 CI 工具的首选，因为如果 Docker 不是唯一使用的部署模型，您可以在多个部署模型中使用相同的工件。此外，您将能够运行单元测试并发布代码覆盖率报告，或在 CI 构建的工件上使用自定义插件。



# redis-cluster

3个主节点，3个从节点

```bash
# redis容器中执行:创建集群
redis-cli --cluster create 172.38.0.11:6379 172.38.0.12:6379 172.38.0.13:6379 172.38.0.14:6379 172.38.0.15:6379 172.38.0.16:6379 --cluster-replicas 1

# 连接集群
redis-cli -c
127.0.0.1:6379> cluster info
127.0.0.1:6379> cluster nodes
```



# vagrant

类似docker/visualBox，启动虚拟机，Hashicorp的另一个流行产品！[官网](https://www.vagrantup.com/)

```bash
$ vagrant init hashicorp/bionic64
$ vagrant up
  Bringing machine 'default' up with 'virtualbox' provider...
  ==> default: Importing base box 'hashicorp/bionic64'...
  ==> default: Forwarding ports...
  default: 22 (guest)
  => 2222 (host) (adapter 1)
  ==> default: Waiting for machine to boot...

$ vagrant ssh
  vagrant@bionic64:~$ _
```

useful commands:

```bash
# 显示ssh配置
vagrant ssh-config

# 异地访问
ssh vagrant@host-ip -i ".vagrant/machines/default/hyperv/private_key"
# 密码默认也是vagrant

# 修改password可以登录的选项
# PasswordAuthentication yes
vim /etc/ssh/sshd_config
```



# vscode远程开发

安装 `remote-ssh`插件，左侧会多出一个按钮：`Remote Explorer`

然后添加一个新的target，可以分不同平台例如WSL或SSH，SSH的会自动发现 `.ssh/config` 的中配置，自动发现远程主机

![image-20210605073856081](../images/vagrant-vscode.png)



# container vs host kernal

我在此视频中解释了哪些容器：[https](https://www.youtube.com/watch?v=98ZNU1KqJfc) ://www.youtube.com/watch?v=98ZNU1KqJfc 。本质上，容器实际上并不存在。这是一个在内核上正常运行的进程，但有一个进入分层文件系统的 chroot 并受 cgroups 和命名空间的限制。cgroup 限制了它可以使用的内容（CPU/内存/磁盘），而命名空间限制了它可以看到的内容。但是没有涉及其他层。这是一个在内核上运行的普通 Linux 进程。



这就引出了第二点：Ubuntu 和 Oracle Linux 以及所有其他 Linux 都运行相同的 Linux 内核。您知道二进制文件和库是兼容的，因为 Linux 内核保证向前和向后兼容性。这就是为什么我们可以在一台机器上同时拥有 RHCK 和 UEK 的原因。它仍然被认为是 Ubuntu，因为发行版之间唯一真正不同的是他们做出的用户空间选择，而不是他们使用的内核。发行版对内核的唯一区别是它打开或关闭了哪些功能，但底层的二进制兼容性仍然存在。



对于容器来说，这也不是一个真正的问题，因为它们很少运行任何直接与内核交互的东西。这就是 glibc 的用途。



# 参考资料

- [docker核心基础](https://www.bilibili.com/video/BV1Vs411E7AR?p=11)

- [runnob-docker](https://www.runoob.com/docker/docker-dockerfile.html)

- [Docker —— 从入门到实践](https://yeasy.gitbook.io/docker_practice/)

- [docker中文网](https://docker_practice.gitee.io/zh-cn/)

- [Dockerfile最佳实践-video](https://www.bilibili.com/video/BV1kz4y1y7aC/?spm_id_from=333.788.recommend_more_video.4)

- [Dockerfile最佳实践-text](https://k8s.coding3min.com/docker-jing-xiang/best-dockerfile)

- [了解 Docker“容器主机”与 Linux 和 Windows 容器的“容器操作系统”](http://www.floydhilton.com/docker/2017/03/31/Docker-ContainerHost-vs-ContainerOS-Linux-Windows.html)
