# 小文件存储

由于在元数据管理、访问性能、存储效率等方面面临巨大的挑战性，因此海量小文件（LOSF，lots of small files）问题成为了工业界和学术界公认的难题。

通常我们认为大小在1MB以内的文件称为小文件，百万级数量及以上称为海量。



# SeaweedFS 

SeaweedFS 是一个快速的分布式存储系统，适用于数十亿个文件的 blob、对象、文件和数据湖！Blob 存储具有 O（1） 磁盘搜索、云分层。Filer支持Cloud Drive，跨DC主动 - 主动复制，Kubernetes，POSIX FUSE挂载，S3 API，S3网关，Hadoop，WebDAV，加密，擦除编码。



```bash
docker run --name seaweedfs -d -p 8333:8333 chrislusf/seaweedfs server -s3
```



# 参考资料

- [LOSF 海量小文件问题综述_xiaofei0859的博客-CSDN博客](https://blog.csdn.net/xiaofei0859/article/details/52817956)
- [chrislusf/seaweedfs：SeaweedFS是一个快速分布式存储系统，用于blob，对象，文件和数据湖，用于数十亿个文件！Blob 存储具有 O（1） 磁盘搜索、云分层。Filer支持Cloud Drive，跨DC主动 - 主动复制，Kubernetes，POSIX FUSE挂载，S3 API，S3网关，Hadoop，WebDAV，加密，擦除编码。 (github.com)](https://github.com/chrislusf/seaweedfs#quick-start-with-single-binary)

- [LOSF(Lots of small files)存储问题_刘爱贵的博客-CSDN博客](https://liuag.blog.csdn.net/article/details/5566410)

- [SeaweedFS基本介绍 - ROCKG - 博客园 (cnblogs.com)](https://www.cnblogs.com/rockg/p/11111337.html)
- 