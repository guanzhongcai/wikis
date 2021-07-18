# minikube



## 创建 Minikube 集群

启动本地 kubernetes 集群

```bash
minikube start

# 第一次start会默认下载kubernetes组件，使用阿里云镜像加速
minikube start --vm-driver=docker --image-mirror-country="cn" --registry-mirror=https://n8bn2y81.mirror.aliyuncs.com
```

在浏览器中打开 Kubernetes 仪表板（Dashboard）

```bash
minikube dashboard
```

如果不想打开 Web 浏览器，请使用 url 标志运行显示板命令以得到 URL：

```shell
minikube dashboard --url
```



## hello nginx

使用minikube配置nginx实战（参考[hello-minikube项目](https://kubernetes.io/zh/docs/tutorials/hello-minikube/)）：

```bash
kubectl create deployment hello-nginx --image=nginx --replicas=3

kubectl get deploy

kubectl get pod

kubectl expose deployment hello-nginx --type=LoadBalancer --port=80

minikube service hello-nginx
```

k8s的scheduler会自动在node间调度，来分配要启动的pod副本数



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
kubectl delete deployment hello-node # pod属于deploy，所以，删除deploy，就删除了pod
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



`kubernetes.yaml`示例：

```y
apiVersion: apps/v1
kind: Deployment
metadata:
  name: system-deployment
  labels:
    app: system
spec:
  selector:
    matchLabels:
      app: system
  template:
    metadata:
      labels:
        app: system
    spec:
      containers:
      - name: system-container
        image: system:1.0-SNAPSHOT
        ports:
        - containerPort: 9080
        readinessProbe:
          httpGet:
            path: /health/ready
            port: 9080
          initialDelaySeconds: 30
          periodSeconds: 10
          timeoutSeconds: 3
          failureThreshold: 1
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: inventory-deployment
  labels:
    app: inventory
spec:
  selector:
    matchLabels:
      app: inventory
  template:
    metadata:
      labels:
        app: inventory
    spec:
      containers:
      - name: inventory-container
        image: inventory:1.0-SNAPSHOT
        ports:
        - containerPort: 9080
        readinessProbe:
          httpGet:
            path: /health/ready
            port: 9080
          initialDelaySeconds: 30
          periodSeconds: 10
          timeoutSeconds: 3
          failureThreshold: 1
---
apiVersion: v1
kind: Service
metadata:
  name: system-service
spec:
  type: NodePort
  selector:
    app: system
  ports:
  - protocol: TCP
    port: 9080
    targetPort: 9080
    nodePort: 31000
---
apiVersion: v1
kind: Service
metadata:
  name: inventory-service
spec:
  type: NodePort
  selector:
    app: inventory
  ports:
  - protocol: TCP
    port: 9080
    targetPort: 9080
    nodePort: 32000
```

应用配置：

```bash
kubectl apply -f kubernetes.yaml
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

```bash
[root@k8smaster ~]# kubectl exec nginx-6799fc88d8-cq7df -- bash
[root@k8smaster ~]# kubectl exec -it nginx-6799fc88d8-cq7df -- bash
```



# 使用 Service 暴露您的应用

尽管每个 Pod 都有一个唯一的 IP 地址，但是如果没有 Service ，这些 IP 不会暴露在群集外部。Service 允许您的应用程序接收流量。

*Kubernetes 的 Service 是一个抽象层，它定义了一组 Pod 的逻辑集，并为这些 Pod 支持外部流量暴露、负载平衡和服务发现。*

- 将 Pod 暴露给外部通信
- 跨多个 Pod 的负载均衡
- 使用标签(Label)



Service 也可以用在 ServiceSpec 标记`type`的方式暴露

- *`ClusterIP`* (默认) - 在集群的内部 IP 上公开 Service 。这种类型使得 Service 只能从集群内访问。

- *`NodePort`* - 使用 NAT 在集群中每个选定 Node 的相同端口上公开 Service 。使用`<NodeIP>:<NodePort>` 从集群外部访问 Service。是 ClusterIP 的超集。

- *`LoadBalancer`* - 在当前云中创建一个外部负载均衡器(如果支持的话)，并为 Service 分配一个固定的外部IP。是 NodePort 的超集。

- *`ExternalName`* - 通过返回带有该名称的 `CNAME` 记录，使用任意名称(由 spec 中的`externalName`指定)公开 Service。不使用代理。这种类型需要`kube-dns`的v1.7或更高版本。

  > CNAME (canonical name) — an alias to an existing host.
  >
  > CNAME(规范名称)——现有主机的别名。



Service 匹配一组 Pod 是使用 [标签(Label)和选择器(Selector)](https://kubernetes.io/zh/docs/concepts/overview/working-with-objects/labels), 它们是允许对 Kubernetes 中的对象进行逻辑操作的一种分组原语。标签(Label)是附加在对象上的键/值对，可以以多种方式使用:

- **指定用于开发，测试和生产的对象**
- **嵌入版本标签**
- **使用 Label 将对象进行分类**

*你也可以在创建 Deployment 的同时用 `--expose`创建一个 Service 。*

![img](../images/module_04_labels.svg)

#### Step 1 Create a new service

```bash
# create a new service and expose it to external traffic
kubectl expose deployment/kubernetes-bootcamp --type="NodePort" --port 8080

kubectl get services

kubectl describe services/kubernetes-bootcamp
```

#### Step 2: Using labels

```bash
kubectl describe deployment

# -l: label
kubectl get pods -l app=kubernetes-bootcamp

kubectl get services -l app=kubernetes-bootcamp

# Get the name of the Pod and store it in the POD_NAME environment variable:
export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')
echo Name of the Pod: $POD_NAME

# apply a new label: use the label command followed by the object type, object name and the new label
kubectl label pods $POD_NAME version=v1

kubectl describe pods $POD_NAME

# query the list of pods using the new label:
kubectl get pods -l version=v1
```



#### Step 3 Deleting a service

```bash
kubectl delete service -l app=kubernetes-bootcamp
```



# 运行应用程序的多个实例

- 扩缩一个 Deployment

- 用 kubectl 扩缩应用程序
- *在运行 kubectl run 命令时，你可以通过设置 --replicas 参数来设置 Deployment 的副本数。*

- *扩缩是通过改变 Deployment 中的副本数量来实现的。*

- 一旦有了多个应用实例，就可以没有宕机地滚动更新。

> 服务 (Service)有一种负载均衡器类型，可以将网络流量均衡分配到外部可访问的 Pods 上。服务将会一直通过端点来监视 Pods 的运行，保证流量只分配到可用的 Pods 上。

#### Step 1: Scaling a deployment

```bash
# see the ReplicaSet created by the Deployment:
kubectl get rs

# scale the Deployment to 4 replicas
kubectl scale deployments/kubernetes-bootcamp --replicas=4

# list your Deployments 
kubectl get deployments

# check the number of Pods
kubectl get pods -o wide

# check the Deployment events log
kubectl describe deployments/kubernetes-bootcamp
```



#### Step 2: Load Balancing

```bash
# check the Service is load-balancing the traffic and find out the exposed IP and Port
kubectl describe services/kubernetes-bootcamp
```



#### Step 3: Scale Down

```bash
kubectl scale deployments/kubernetes-bootcamp --replicas=2
```



# 执行滚动更新

- 使用 kubectl 执行滚动更新。
- 在 Kubernetes 中，更新是经过版本控制的，任何 Deployment 更新都可以恢复到以前的（稳定）版本。
- *滚动更新允许通过使用新的实例逐步更新 Pod 实例从而实现 Deployments 更新，停机时间为零。*

#### Step 1: Update the version of the app

```bash
# update the image of the application to version 2, use the set image command, followed by the deployment name and the new image version:
kubectl set image deployments/kubernetes-bootcamp kubernetes-bootcamp=jocatalin/kubernetes-bootcamp:v2
```



#### Step 2: Verify an update

```bash
# confirm the update by running the rollout status command:
kubectl rollout status deployments/kubernetes-bootcamp
```



#### Step 3: Rollback an update

```bash
# 实验：deploy an image tagged with v10 
kubectl set image deployments/kubernetes-bootcamp kubernetes-bootcamp=gcr.io/google-samples/kubernetes-bootcamp:v10

# roll back the deployment to your last working version, use the rollout undo command:
kubectl rollout undo deployments/kubernetes-bootcamp
```



# Web 界面 (Dashboard)

## 部署 Dashboard UI

默认情况下不会部署 Dashboard。可以通过以下命令部署：

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.2.0/aio/deploy/recommended.yaml
```



```bash
# 查看当前运行的服务端口和名称
netstat -nlpt
```



# 参考资料

- https://kubernetes.io/zh/docs/tutorials/hello-minikube/
- [Katacoda搭建和使用教程](https://developer.aliyun.com/article/752183)
- [katacoda官方提供的各技术栈教程（docker、consul、k8s...)](https://www.katacoda.com/courses)
- [kubernetes-basics](https://kubernetes.io/zh/docs/tutorials/kubernetes-basics/deploy-app/deploy-intro/)
- [使用kubernetes 官网工具kubeadm部署kubernetes(使用阿里云镜像)](https://www.cnblogs.com/tylerzhou/p/10971336.html)
- [kubectl：启用 shell 自动补全功能](https://kubernetes.io/zh/docs/tasks/tools/install-kubectl-linux/#install-using-native-package-management)

