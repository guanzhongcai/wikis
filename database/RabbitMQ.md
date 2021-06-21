<!-- TITLE: RabbitMQ -->

# RabbitMQ

- 默认管理后台：http://localhost:15672/

- 如果打不开，请[参考](https://blog.csdn.net/weixin_43641651/article/details/89848530)以下解决：

  ```powershell
  cd /d C:\Program Files\RabbitMQ Server\rabbitmq_server-3.7.8\sbin
  # 显示启动的服务 E和e分别表示显性和隐性启动
  rabbitmq-plugins.bat list
  
  # 启动插件
  rabbitmq-plugins.bat enable rabbitmq_management
  ```
  
  

## 基本概念

| Name        | Concept                                                      |
| ----------- | ------------------------------------------------------------ |
| Broker      | 简单来说就是消息队列服务器实体。                             |
| Exchange    | 消息交换机，它指定消息按什么规则，路由到哪个队列。           |
| Queue       | 消息队列载体，每个消息都会被投入到一个或多个队列。           |
| Binding     | 绑定，它的作用就是把exchange和queue按照路由规则绑定起来。    |
| Routing Key | 路由关键字，exchange根据这个关键字进行消息投递。             |
| vhost       | 虚拟主机，一个broker里可以开设多个vhost，用作不同用户的权限分离。 |
| producer    | 消息生产者，就是投递消息的程序。                             |
| consumer    | 消息消费者，就是接受消息的程序。                             |
| channel     | 消息通道，在客户端的每个连接里，可建立多个channel，每个channel代表一个会话任务。 |



![]()



消息队列的使用过程大概如下：

（1）客户端连接到消息队列服务器，打开一个channel。
（2）客户端声明一个exchange，并设置相关属性。
（3）客户端声明一个queue，并设置相关属性。
（4）客户端使用routing key，在exchange和queue之间建立好绑定关系。
（5）客户端投递消息到exchange。

![image-20210616134611571](uploads/RabbitMQ-flow.png)
