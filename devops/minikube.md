## 创建 Minikube 集群

启动本地 kubernetes 集群

```bash
minikube start

minikube start --vm-driver=docker --image-mirror-country="cn" --registry-mirror=https://f1z25q5p.mirror.aliyuncs.com
```

在浏览器中打开 Kubernetes 仪表板（Dashboard）

```bash
minikube dashboard
```

如果不想打开 Web 浏览器，请使用 url 标志运行显示板命令以得到 URL：

```shell
minikube dashboard --url
```



## 创建 Deployment

Kubernetes [*Pod*](https://kubernetes.io/zh/docs/concepts/workloads/pods/) 是由一个或多个 为了管理和联网而绑定在一起的容器构成的组。 本教程中的 Pod 只有一个容器。 Kubernetes [*Deployment*](https://kubernetes.io/zh/docs/concepts/workloads/controllers/deployment/) 检查 Pod 的健康状况，并在 Pod 中的容器终止的情况下重新启动新的容器。 Deployment 是管理 Pod 创建和扩展的推荐方法。

使用 `kubectl create` 命令创建管理 Pod 的 Deployment。该 Pod 根据提供的 Docker 镜像运行 Container。

```shell
kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
```

查看 Deployment：

```shell
kubectl get deployments
```

输出结果类似于这样：

```
NAME         READY   UP-TO-DATE   AVAILABLE   AGE
hello-node   1/1     1            1           1m
```

查看 Pod：

```shell
kubectl get pods
```

输出结果类似于这样：

```
NAME                          READY     STATUS    RESTARTS   AGE
hello-node-5f76cf6ccf-br9b5   1/1       Running   0          1m
```

查看集群事件：

```shell
kubectl get events
```

查看 `kubectl` 配置：

```shell
kubectl config view
```



## 创建 Service[ ](https://kubernetes.io/zh/docs/tutorials/hello-minikube/#创建-service)

默认情况下，Pod 只能通过 Kubernetes 集群中的内部 IP 地址访问。 要使得 `hello-node` 容器可以从 Kubernetes 虚拟网络的外部访问，你必须将 Pod 暴露为 Kubernetes [*Service*](https://kubernetes.io/zh/docs/concepts/services-networking/service/)。

1. 使用 `kubectl expose` 命令将 Pod 暴露给公网：

   ```shell
   kubectl expose deployment hello-node --type=LoadBalancer --port=8080
   ```

   这里的 `--type=LoadBalancer` 参数表明你希望将你的 Service 暴露到集群外部。

   镜像 `k8s.gcr.io/echoserver` 中的应用程序代码仅监听 TCP 8080 端口。 

1. 查看你创建的 Service：

   ```shell
   kubectl get services
   ```

   输出结果类似于这样:

   ```
   NAME         TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
   hello-node   LoadBalancer   10.108.144.78   <pending>     8080:30369/TCP   21s
   kubernetes   ClusterIP      10.96.0.1       <none>        443/TCP          23m
   ```

   对于支持负载均衡器的云服务平台而言，平台将提供一个外部 IP 来访问该服务。 在 Minikube 上，`LoadBalancer` 使得服务可以通过命令 `minikube service` 访问。

1. 运行下面的命令：

   ```shell
   minikube service hello-node
   ```



## 启用插件

Minikube 有一组内置的 [插件](https://kubernetes.io/zh/docs/concepts/cluster-administration/addons/)， 可以在本地 Kubernetes 环境中启用、禁用和打开。

列出当前支持的插件：

```shell
minikube addons list
```

输出结果类似于这样：

```
addon-manager: enabled
dashboard: enabled
default-storageclass: enabled
efk: disabled
freshpod: disabled
gvisor: disabled
helm-tiller: disabled
ingress: disabled
ingress-dns: disabled
logviewer: disabled
metrics-server: disabled
nvidia-driver-installer: disabled
nvidia-gpu-device-plugin: disabled
registry: disabled
registry-creds: disabled
storage-provisioner: enabled
storage-provisioner-gluster: disabled
```

启用插件，例如 `metrics-server`：

```shell
minikube addons enable metrics-server
```

输出结果类似于这样：

```
metrics-server was successfully enabled
```

查看创建的 Pod 和 Service：

```shell
kubectl get pod,svc -n kube-system
```

输出结果类似于这样：

```
NAME                                        READY     STATUS    RESTARTS   AGE
pod/coredns-5644d7b6d9-mh9ll                1/1       Running   0          34m
pod/coredns-5644d7b6d9-pqd2t                1/1       Running   0          34m
pod/metrics-server-67fb648c5                1/1       Running   0          26s
pod/etcd-minikube                           1/1       Running   0          34m
pod/influxdb-grafana-b29w8                  2/2       Running   0          26s
pod/kube-addon-manager-minikube             1/1       Running   0          34m
pod/kube-apiserver-minikube                 1/1       Running   0          34m
pod/kube-controller-manager-minikube        1/1       Running   0          34m
pod/kube-proxy-rnlps                        1/1       Running   0          34m
pod/kube-scheduler-minikube                 1/1       Running   0          34m
pod/storage-provisioner                     1/1       Running   0          34m

NAME                           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)             AGE
service/metrics-server         ClusterIP   10.96.241.45    <none>        80/TCP              26s
service/kube-dns               ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP       34m
service/monitoring-grafana     NodePort    10.99.24.54     <none>        80:30002/TCP        26s
service/monitoring-influxdb    ClusterIP   10.111.169.94   <none>        8083/TCP,8086/TCP   26s
```

1. 禁用 `metrics-server`：

   ```shell
   minikube addons disable metrics-server
   ```

   输出结果类似于这样：

   ```
   metrics-server was successfully disabled
   ```



## 清理

现在可以清理你在集群中创建的资源：

```shell
kubectl delete service hello-node
kubectl delete deployment hello-node
```

可选地，停止 Minikube 虚拟机（VM）：

```shell
minikube stop
```

可选地，删除 Minikube 虚拟机（VM）：

```shell
minikube delete
```



# 使用 Minikube 创建集群

### Kubernetes 集群

- **Kubernetes 协调一个高可用计算机集群，每个计算机作为独立单元互相连接工作**
-  **Kubernetes 以更高效的方式跨集群自动分发和调度应用容器。**

一个 Kubernetes 集群包含两种类型的资源:

- **Master** 调度整个集群
- **Nodes** 负责运行应用

*Kubernetes 是一个生产级别的开源平台，可协调在计算机集群内和跨计算机集群的应用容器的部署（调度）和执行.*

![img](../images/module_01_cluster.svg)

*Master 管理集群，Node 用于托管正在运行的应用。*

**Node 使用 Master 暴露的 Kubernetes API 与 Master 通信。**

```bash
minikube version

minikube start

kubectl version

kubectl cluster-info

kubectl get nodes
```



# 使用 kubectl 创建 Deployment

- *Deployment 负责创建和更新应用程序的实例*

- Kubernetes Deployment 控制器会持续监视实例。 如果托管实例的节点关闭或被删除，则 Deployment 控制器会将该实例替换为群集中另一个节点上的实例。 **这提供了一种自我修复机制来解决机器故障维护问题。**
- *应用程序需要打包成一种受支持的容器格式，以便部署在 Kubernetes 上*（dockerize）

![img](../images/module_02_first_app.svg)

```bash
kubectl create deployment kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1

kubectl get deployments

# another terminal window
kubectl proxy

# test
curl http://localhost:8001/version
```



# 查看 pod 和工作节点

- *Pod 是一组一个或多个应用程序容器（例如 Docker），包括共享存储（卷), IP 地址和有关如何运行它们的信息。*

一个Pod中的容器共享资源:

- 共享存储，当作卷
- 网络，作为唯一的集群 IP 地址
- 有关每个容器如何运行的信息，例如容器映像版本或要使用的特定端口。

![img](../images/module_03_pods.svg)

- *如果它们紧耦合并且需要共享磁盘等资源，这些容器应在一个 Pod 中编排。*

- *工作节点是 Kubernetes 中的负责计算的机器，可能是VM或物理计算机，具体取决于群集。多个 Pod 可以在一个工作节点上运行。*

![img](../images/module_03_nodes.svg)

## 使用 kubectl 进行故障排除

- **kubectl get** - 列出资源
- **kubectl describe** - 显示有关资源的详细信息
- **kubectl logs** - 打印 pod 和其中容器的日志
- **kubectl exec** - 在 pod 中的容器上执行命令





## 参考资料

- https://kubernetes.io/zh/docs/tutorials/hello-minikube/

- [Katacoda搭建和使用教程](https://developer.aliyun.com/article/752183)

- [katacoda官方提供的各技术栈教程（docker、consul、k8s...)](https://www.katacoda.com/courses)

- [kubernetes-basics](https://kubernetes.io/zh/docs/tutorials/kubernetes-basics/deploy-app/deploy-intro/)

- [使用kubernetes 官网工具kubeadm部署kubernetes(使用阿里云镜像)](https://www.cnblogs.com/tylerzhou/p/10971336.html)
