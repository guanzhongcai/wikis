# 打包

```bash
mvn clean install -DskipTests=true
```



## docker

```dockerfile
FROM carsharing/alpine-oraclejdk8-bash
```

在容器外部提供一个映射路径，`-volumn webapps:/root/app/webapps`，在外部放置项目，就自动同步到内部了！



查看各容器占用的系统资源：

```bash
docker stats
```

```bash
CONTAINER ID   NAME        CPU %     MEM USAGE / LIMIT     MEM %     NET I/O           BLOCK I/O        PIDS
4f1eb7e71c39   portainer   0.00%     6.227MiB / 1.942GiB   0.31%     6.94kB / 1.73kB   19.4MB / 307kB   8
```



启动ES：

```bash
# 限制JVM内存最小64M最大512M
docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms=64m -Xmx512m" elasticsearch:7.6.2
```

