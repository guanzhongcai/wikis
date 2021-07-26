# k8s-aliyun



![](k8s/aliyun/framework.png)



1. 安装 VirtualBox：https://www.virtualbox.org
2. 安装阿里云版`MiniKube`: https://developer.aliyun.com/article/221687

```bash
# 启动
minikube start --vm-driver virtualbox
```



**容器的本质是？**

- 一个试图被隔离、资源受限的进程

- 容器里PID=1的进程就是应用本身
- 管理虚拟机=管理基础设施
- 管理容器=直接管理应用本身



**kubernetes是？**

- kubernetes就是云时代的操作系统！
- 以此类推，容器镜像其实就是：这个操作系统的软件安装包



## 容器

- 资源视图隔离 - Namespace
- 控制资源使用率 - CGroup
- 独立的文件系统 - chroot：在*nix下，chroot这个系统调用，可以把一个子目录变为根目录，在视图级别进行隔离。