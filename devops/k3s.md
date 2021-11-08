# k3s

中文文档：

https://docs.rancher.cn/k3s/



有道笔记：

https://note.youdao.com/ynoteshare/index.html?id=3cce690cb42a4ff43c392bd346f07dcb&type=note&_time=1636242085879



从0到1基础入门：全面了解k3s QA文档

https://shimo.im/docs/HjYY8c836jPgPJK9/read



check env after installation:

```bash
kubectl get all -n kube-system
```

安装docker命令版而非默认containerd命令版的k3s:

```bash
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="--docker" sh -s -
```

k3s安装目录：

```bash
ls /var/lib/rancher/k3s
agent data server
```









