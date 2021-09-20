# Postgresql



## PITR恢复

`PITR（Point In Time Recovery）`过程与常规恢复不同的是，从配置参数 `archive_command`中甚至的归档目录中读取`WAL`段，从 `backup_label` 文件读取检查点位置。



## WAL段切换

当发生下面任意一种情况是，WAL段会发生切换：

1. WAL段已经被填满
2. 调用函数 pg_switch_wal()
3. 启用了 archive_mode 且已经超过 archive_timeout配置的时间（？？）

当WAL段文件发生切换时，自动将其复制到归档区域，archive_command参数配置归档区域。



## 参考资料

- [【2020PG亚洲大会】禹晓：详谈WAL文件](https://www.bilibili.com/video/BV1Wv411b7iQ?from=search&seid=16961779459008666835&spm_id_from=333.337.0.0)