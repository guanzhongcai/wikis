# architect



<img src="../net/nginx/reverse-proxy.png" style="zoom:50%;" />

访问8088两次，8099一次！

```nginx
{
	server localhost:8088 weight 2;
	server localhost:8099 weight 1;
}
```

