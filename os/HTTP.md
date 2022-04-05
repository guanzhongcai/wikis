# HTTP规范

## MIME-Type

- Multipupose Internet Mail Extensions

- **MIME-type** （现在称为“媒体类型(media type)”，但有时也是“内容类型(content type)”）是指示文件类型的字符串，与文件一起发送（例如，一个声音文件可能被标记为 `audio/ogg` ，一个图像文件可能是 `image/png` ）。它与传统Windows上的文件扩展名有相同目的。

  ```json
  Content-Type：text/HTML
  ```

  

- Email 附件的类型也是通过 MIME Type 指定的



## 客户端请求消息

客户端发送一个HTTP请求到服务器的请求报文消息包括四部分：

- 请求行（request line）
- 请求头部（header）
- 空行
- 请求数据

![img](../images/os/http-request-format.png)



```http
GET /hello.txt HTTP/1.1
User-Agent: curl/7.16.3 libcurl/7.16.3 OpenSSL/0.9.7l zlib/1.2.3
Host: www.example.com
Accept-Language: en, mi
```



## 服务器响应消息

HTTP响应也由四个部分组成，分别是：

- 状态行
- 消息报头
- 空行
- 响应正文

![img](../images/httpmessage.jpg)

```http
HTTP/1.1 200 OK
Date: Mon, 27 Jul 2009 12:28:53 GMT
Server: Apache
Last-Modified: Wed, 22 Jul 2009 19:15:56 GMT
ETag: "34aa387-d-1568eb00"
Accept-Ranges: bytes
Content-Length: 51
Vary: Accept-Encoding
Content-Type: text/plain

Hello World! My payload includes a trailing CRLF.
```



## HTTP响应头信息

| 应答头           | 说明                                                         |
| ---------------- | ------------------------------------------------------------ |
| Allow            | 服务器支持哪些请求方法（如GET、POST等）。                    |
| Content-Encoding | 文档的编码（Encode）方法。只有在解码之后才可以得到Content-Type头指定的内容类型。利用gzip压缩文档能够显著地减少HTML文档的下载时间。Java的GZIPOutputStream可以很方便地进行gzip压缩，但只有Unix上的Netscape和Windows上的IE 4、IE 5才支持它。因此，Servlet应该通过查看Accept-Encoding头（即request.getHeader("Accept-Encoding")）检查浏览器是否支持gzip，为支持gzip的浏览器返回经gzip压缩的HTML页面，为其他浏览器返回普通页面。 |
| Content-Length   | 表示内容长度。只有当浏览器使用持久HTTP连接时才需要这个数据。如果你想要利用持久连接的优势，可以把输出文档写入 ByteArrayOutputStream，完成后查看其大小，然后把该值放入Content-Length头，最后通过byteArrayStream.writeTo(response.getOutputStream()发送内容。 |
| Content-Type     | 表示后面的文档属于什么MIME类型。Servlet默认为text/plain，但通常需要显式地指定为text/html。由于经常要设置Content-Type，因此HttpServletResponse提供了一个专用的方法setContentType。 |
| Date             | 当前的GMT时间。你可以用setDateHeader来设置这个头以避免转换时间格式的麻烦。 |
| Expires          | 应该在什么时候认为文档已经过期，从而不再缓存它？             |
| Last-Modified    | 文档的最后改动时间。客户可以通过If-Modified-Since请求头提供一个日期，该请求将被视为一个条件GET，只有改动时间迟于指定时间的文档才会返回，否则返回一个304（Not Modified）状态。Last-Modified也可用setDateHeader方法来设置。 |
| Location         | 表示客户应当到哪里去提取文档。Location通常不是直接设置的，而是通过HttpServletResponse的sendRedirect方法，该方法同时设置状态代码为302。 |
| Server           | 服务器名字。是由Web服务器自己设置。                          |
| Set-Cookie       | 设置和页面关联的Cookie。                                     |
| WWW-Authenticate | 客户应该在Authorization头中提供什么类型的授权信息？在包含401（Unauthorized）状态行的应答中这个头是必需的。例如，response.setHeader("WWW-Authenticate", "BASIC realm=＼"executives＼"")。 |



## HTTP状态码分类

| 分类 | 分类描述                                       |
| ---- | ---------------------------------------------- |
| 1**  | 信息，服务器收到请求，需要请求者继续执行操作   |
| 2**  | 成功，操作被成功接收并处理                     |
| 3**  | 重定向，需要进一步的操作以完成请求             |
| 4**  | 客户端错误，请求包含语法错误或无法完成请求     |
| 5**  | 服务器错误，服务器在处理请求的过程中发生了错误 |



# HTTP 0.9

1991年发布

- 不支持请求头，只支持 `GET` 方法



# HTTP 1.0

1996年发布

- HTTP1.0 定义了三种请求方法： GET, POST 和 HEAD方法。
- 在请求中加入了HTTP版本号，如：`GET /coolshell/index.html HTTP/1.0`
- HTTP 开始有 header了，不管是request还是response 都有header了。
- 增加了HTTP Status Code 标识相关的状态码。
- 还有 `Content-Type` 可以传输其它的文件了。
- 每次请求连接不能复用
- 是串行请求



# HTTP 1.1

1997年发布

- HTTP1.1 新增了六种请求方法：OPTIONS、PUT、PATCH、DELETE、TRACE 和 CONNECT 方法。

- connection默认是keep-alive

  > 这是所谓的“**HTTP 长链接**” 或是 “**请求响应式的HTTP 持久链接**”。英文叫 HTTP Persistent connection.

- 支持pipeline网络传输，只要第一个请求发出去了，不必等其回来，就可以发第二个请求出去，可以减少整体的响应时间。（注：非幂等的POST 方法或是有依赖的请求是不能被pipeline化的）

- 支持 Chunked Responses ，也就是说，在Response的时候，不必说明 `Content-Length` 这样，客户端就不能断连接，直到收到服务端的EOF标识。这种技术又叫 “**服务端Push模型**”，或是 “**服务端Push式的HTTP 持久链接**”

- 增加了 cache control 机制。

- 协议头注增加了 Language, Encoding, Type 等等头，让客户端可以跟服务器端进行更多的协商。

- 还正式加入了一个很重要的头—— `HOST`这样的话，服务器就知道你要请求哪个网站了。因为可以有多个域名解析到同一个IP上，要区分用户是请求的哪个域名，就需要在HTTP的协议中加入域名的信息，而不是被DNS转换过的IP信息。

- 正式加入了 `OPTIONS` 方法，其主要用于 [跨源资源共享（CORS）](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/CORS) 应用。

- 因为并行的请求可能会导致浏览器的负载过重，所以比如Chrome会默认每个域名同时只有6个并发的请求

- 明文传输，没有压缩

- header太长

- 传输慢

| 方法    | 描述                                                         |
| ------- | ------------------------------------------------------------ |
| GET     | 请求指定的页面信息，并返回实体主体。                         |
| HEAD    | 类似于 GET 请求，只不过返回的响应中没有具体的内容（`body`），用于`获取报头` |
| POST    | 向指定资源提交数据进行处理请求（例如提交表单或者上传文件）。数据被包含在请求体中。POST 请求可能会导致新的资源的建立和/或已有资源的修改。 |
| PUT     | 从客户端向服务器传送的数据取代指定的文档的内容。             |
| DELETE  | 请求服务器删除指定的页面。                                   |
| CONNECT | HTTP/1.1 协议中预留给能够将连接改为管道方式的代理服务器。    |
| OPTIONS | 允许客户端`查看服务器的性能`。                               |
| TRACE   | 回显服务器收到的请求，主要用于`测试`或`诊断`。               |
| PATCH   | 是对 PUT 方法的补充，用来对已知资源进行`局部更新` 。         |



在HTTP/1.1 下，HTTP已经支持四种网络协议：

- 传统的短链接。
- 可重用TCP的的长链接模型。
- 服务端push的模型。
- WebSocket模型。



序号方法描述1GET请求指定的页面信息，并返回实体主体。2HEAD类似于 GET 请求，只不过返回的响应中没有具体的内容，用于获取报头3POST向指定资源提交数据进行处理请求（例如提交表单或者上传文件）。数据被包含在请求体中。POST 请求可能会导致新的资源的建立和/或已有资源的修改。4PUT从客户端向服务器传送的数据取代指定的文档的内容。5DELETE请求服务器删除指定的页面。6CONNECTHTTP/1.1 协议中预留给能够将连接改为管道方式的代理服务器。7OPTIONS允许客户端查看服务器的性能。8TRACE回显服务器收到的请求，主要用于测试或诊断。9PATCH是对 PUT 方法的补充，用来对已知资源进行局部更新 。

# HTTPS

使用TLS协议的HTTP。

生成HTTPS私有证书：

> 私有证书在网页中访问会显示此站点是不受信任的，需要网页中点击【继续前往】

```bash
# under ubuntu os
openssl genrsa -out ./server.key 2048
openssl req -new -x509 -key ./server.key -out ./server.pem -days 365
```

gin中开启https：

> gin中开启https会默认使用http2，如果不支持，就再切回http1.1

```go
r := gin.Default()
...
r.RunTLS(":9999", "./server.pem", "./server.key")
```



# HTTP 2.0

2015年发布。

- 可实现了文件分块后的乱序传输，最后组装好的BT下载思想
- 因为实现了server push，对客户端的单次请求，可以实现多次返回（js/css/html/imags/..)

- HTTP/2是一个二进制协议，增加了数据传输的效率。
- HTTP/2是可以在一个TCP链接中并发请求多个HTTP请求，移除了HTTP/1.1中的串行请求。
- HTTP/2会压缩头，如果你**同时发出多个请求，他们的头是一样的或是相似的**，那么，协议会帮你消除重复的部分。这就是所谓的HPACK算法（参看[RFC 7541](https://tools.ietf.org/html/rfc7541) 附录A）
- HTTP/2允许服务端在客户端放cache，又叫服务端push，也就是说，你没有请求的东西，我服务端可以先送给你放在你的本地缓存中。比如，你请求X，我服务端知道X依赖于Y，虽然你没有的请求Y，但我把把Y跟着X的请求一起返回客户端。
- 内部维护了一个“**优先级树**”来用于来做一些资源和请求的调度和控制。

- 截止2019年10月1日 ， 在全世界范围内已经有41%的网站开启了HTTP/2。



## 核心概念

| 概念       | 名称   | 含义                                                   |
| ---------- | ------ | ------------------------------------------------------ |
| connection | 连接   | 1 个 TCP 连接，包含一个或者多个 Stream。               |
| stream     | 数据流 | 一个双向通讯数据流，包含 1 条或者多条 Message。        |
| message    | 消息   | 对应 HTTP/1 中的请求或者响应，包含一条或者多条 Frame。 |
| frame      | 数据帧 | 最小单位，以二进制压缩格式存放 HTTP/1 中的内容。       |

帧（frame）、流（stream）、消息（message）示意图：

<img src="../images/http2-concepts.png" alt="在这里插入图片描述" style="zoom:50%;" />

> Tips：图中 `Stream` 表示多个数据流，它们可以源源不断地并发传送，同一个 `Stream` 流中的 `frame` 数据是串行发送的。

## Stream ID实现多路复用

- 接收端的实现可据此并发组装消息
- 同一 `Stream` 内的 `frame` 必须是有序的（同一 `Stream` 内的 `frame` 是串行的）
- `SETTINGS_MAX_CONCURRENT_STREAMS` 控制着并发 `Stream` 数



## Stream流特性

- 由客户端建立的 `Stream ID` 必须是奇数。
- 由服务端建立的 `Stream ID` 必须是偶数（如服务端主动向客户端推送消息）。
- 新建立的 Stream ID 必须大于曾经建立过的状态为 opened 或 reserved 的 Stream ID。
- 在新建立的流上发送帧时，意味着将更小 ID 且为 idle 状态的 Stream 设置为 Closed 状态。
- Stream ID 不能复用，长连接耗尽 ID 应创建新连接。
- `Stream ID` 为 `0` 的流仅用于传输控制帧。

- 心跳往返都是 `PING`，`WebSocket` 中往返心跳对应 `PING`、`PONG`。



## 帧

帧头部（3+1+1+4=9字节）示意图:

<img src="../images/http2-frame.png" alt="在这里插入图片描述" style="zoom:50%;" />

| 帧类型        | 类型编码 | 含义                                  |
| ------------- | -------- | ------------------------------------- |
| DATA          | 0x0      | 传递 HTTP 包体                        |
| HEADERS       | 0x1      | 传递 HTTP 头部                        |
| PRIORITY      | 0x2      | 指定 Stream 流的优先级                |
| RST_STREAM    | 0x3      | 终止 Stream 流                        |
| SETTINGS      | 0x4      | 修改连接或者 Stream 流的配置          |
| PUSH_PROMISE  | 0x5      | 服务端推送资源时描述请求的帧          |
| PING          | 0x6      | 心跳检测，兼具计算 RTT 往返时间的功能 |
| GOAWAY        | 0x7      | 优雅的终止连接或者通知错误            |
| WINDOW UPDATE | 0x8      | 实现流量控制                          |
| CONTINUATION  | 0x9      | 传递较大 HTTP 头部时的持续帧          |

> Tips：心跳往返都是 `PING`，`WebSocket` 中往返心跳对应 `PING`、`PONG`。



## SETTINGS 帧类型

| SETTINGS 帧类型                      | 含义                                                         |
| ------------------------------------ | ------------------------------------------------------------ |
| SETTINGS_HEADERS_TABLE_SIZE          | 通知对端索引表的最大尺寸（单位字节，初始 4096 字节）         |
| SETTINGS_ENABLE_PUSH(0x2)            | Value 设置为 0 时可禁用服务器推送功能， 1 表示启用           |
| SETTINGS_MAX_CONCURRENT_STREAMS(0x3) | 告诉接收端允许的最大并发 Stream 数量                         |
| SETTINGS_INITIAL_WINDOW_SIZE(0x4)    | 声明发送端的窗口大小，用于 Stream 级别流控，初始值 2^16-1，即 65535字节 |
| SETTINGS_MAX_FRAME_SIZE(0x5)         | 设置帧的最大大小，初始值 2^14，即 16384 字节                 |
| SETTINGS_MAX_HEADER_LIST_SIZE(0x6)   | 知会对端头部索引表的最大尺寸，单位字节，基于未压缩前的头部   |



## Wireshark抓包分析

> Tips：从图中可以看出客户端 `Settings - Max concurrent streams : 1000` 表示客户端允许的最大并发 `Stream` 数量是 `1000`。
>
> Tips：从图中可以看出服务端的 Settings - Max concurrent streams : 128 表示服务端允许的最大并发 Stream 数是 128，Settings - Max frame size : 16777215表示服务端允许的最大帧大小 16777215 字节。



# HTTP 3.0

2018年发布，已被 Chrome，Firefox，和Cloudflare支持

> HTTP/2主要的问题是：若干个HTTP的请求在复用一个TCP的连接，底层的TCP协议是不知道上层有多少个HTTP的请求的，所以，一旦发生丢包，造成的问题就是所有的HTTP请求都必需等待这个丢了的包被重传回来，哪怕丢的那个包不是我这个HTTP请求的。因为TCP底层是没有这个知识了。

- 从原来的基于TCP改为基于UDP

- 本质就是HTTP/2 + QUIC协议，QUIC：Quick UDP Internet Connections，QUIC需要把TCP的算法再实现一个Quick版的

- 一个HTTPS的连接，先是TCP的三次握手，然后是TLS的三次握手，要整出六次网络交互，一个链接才建好

- 如果QUIC成熟了，TCP是不是会有可能成为历史？！

-  HTTP/2的头压缩算法 HPACK，此时需要实现一版QPACK的算法

  > HPACK需要维护一个动态的字典表来分析请求的头中哪些是重复的，HPACK的这个数据结构需要在encoder和decoder端同步这个东西



# keep-alive

> https://www.cnblogs.com/caoweixiong/p/14720254.html

keep-alive 是客户端和服务端的一个约定，如果开启 keep-alive，则服务端在返回 response 后不关闭 TCP 连接；

**连接**能够在**短时间**内得到复用

- HTTP 1.0 中默认是关闭的，需要在http头加入"Connection: Keep-Alive"，才能启用Keep-Alive；
- HTTP 1.1 中默认启用Keep-Alive，如果加入"Connection: close "，才关闭。



## 使用场景

- **如果客户端和服务端的确需要进行多次通信，则开启 keep-alive 是更好的选择**，例如在微服务架构中，通常微服务的使用方和提供方会长期有交流。

- **在一些 TPS/QPS 很高的 REST 服务中，如果使用的是短连接（即没有开启keep-alive），则很可能发生`客户端端口`被占满的情形**。

  这是由于短时间内会创建大量TCP 连接，而在 TCP 四次挥手结束后，客户端的端口会处于 TIME_WAIT一段时间(2*MSL)，

  这期间端口不会被释放，从而导致端口被占满。这种情况下最好使用长连接。



## **如何处理keep-alive**

对于**客户端**来说，不论是浏览器，还是手机App，或者我们直接在Java代码中使用HttpUrlConnection，**只是负责在请求头中设置Keep-Alive**。

而**具体的连接复用时间的长短，通常是由web服务器控制的**。

**在HTTP协议中，Keep-Alive属性保持连接的时间长短是由服务端决定的，通常配置都是在几十秒左右。**

例如，在tomcat中，我们可以server.xml中配置以下属性：

![img](HTTP.assets/1577453-20210429234231235-636607618.png)

说明如下：

- **maxKeepAliveRequests：**一个连接上，最多可以发起多少次请求，默认100，超过这个次数后会关闭。
- **keepAliveTimeout：**底层socket连接最多保持多长时间，默认60秒，超过这个时间连接会被关闭。 

当然，这不是所有内容，在一些异常情况下，keepalive也会失效。tomcat会根据http响应的状态码，判断是否需要丢弃连接



## **客户端如何开启？**

现在我们用到的几乎所有工具都是默认开启长连接的：

- 对于**浏览器**而言，几乎你现在用的浏览器（包括 IE6）都默认使用 keep-alive 了。
- **Java8** 中的 **`HttpURLConnection`** 默认开启长连接，但是默认连接池中只保留 5 个长连接，如果同时超过 5 个在使用，则会创建新的连接，结束后多于 5 个的部分会被客户端主动关闭。
- **Apache `HttpClient`** 默认为每个地址保留 2 个长连接，连接池中最多共保留 20 个连接。
- **Python requests** 如果使用 session 则会默认开启长连接。



## **HTTP与TCP中keep-alive的区别**

- **HTTP协议(七层)**的Keep-Alive意图在于**连接复用**，希望可以短时间内在同一个连接上进行多次请求/响应。**核心在于：时间要短，速度要快。**

- **TCP协议(四层)**的KeepAlive机制意图在于**保活、心跳，检测连接错误**。**核心在于：虽然频率低，但是持久。**

　　  当一个TCP连接两端长时间没有数据传输时(通常默认配置是2小时)，发送keepalive探针，探测链接是否存活。

 

# 参考资料：

- [HTTP简介](https://www.runoob.com/http/http-tutorial.html)

- [HTTP/2 协议（帧、消息、流简单的抓包分析）](https://blog.csdn.net/qq_38937634/article/details/111352895?utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-1.control&dist_request_id=1331978.8272.16186134743356913&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-1.control)

- [HTTP/2 协议-Stream 的状态变迁](https://blog.csdn.net/qq_38937634/article/details/111420205)

- [HTTP的前世今生-coolshell](https://coolshell.cn/articles/19840.html)

