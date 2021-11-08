# springboot

### 修改端口

修改`resources/application.properties`

server.port = 8081



### 修改banner

新建：`resources/banner.txt`

从网址：www.bootschool/ascii 的在线工具中获取艺术字/图



### 自动配置：

pom.xml

- spring-boot-dependencies: 核心依赖在父工程中



### 自定义properties配置文件

qinjiang.properties:

```properties
name=qinjiang
```

应用：

```java
@PropertySource(value = "classpath:qinjiang.properties")
public class Person {
  // SPEL表达式取出配置文件的值
  @Value("${name}")
  private String name;
}
```

注意properties文件需要配置项目setting的此文件属性为utf8存储，否则默认的会乱码。



### yaml配置文件值

application.yaml

```yml
person:
  name: qingjiang${random.uuid}
  age: ${random.int}
  dog:
    name: ${person.hello:hello}_旺财 # person.hello如果不存在，就用后面的hello做默认值
```



spring.io官方文档：

https://docs.spring.io/spring-boot/docs/2.1.6.RELEASE/reference/htmlsingle/



### springboot的多环境配置

可以选择激活哪个环境下配置文件(dev/test/pro..)

application.properties中选择application-dev.properties、application-test.properties..：

```properties
# 运行环境选择
spring.profiles.active=test

# 上传的文件大小
spring.servlet.multipart.max-file-size=10MB
```

application.yaml可用一个文件包括三个环境:

```yaml
---
server:
  port: 8081
spring:
  profiles:
    active: dev
---
server:
  port: 8082
spring:
  profiles: dev
---
server:
  port: 8083
spring:
  profiles:test




```



官方文档：

https://docs.spring.io/spring-boot/docs/current/reference/html/using.html#using.build-systems.starters



SpringBoot2核心技术与响应式编程:

https://www.yuque.com/atguigu/springboot

教程源码：https://gitee.com/leifengyang/springboot2



Java8的接口默认实现，即实现类对interface的所有方法不用都实现，可以只实现几个，其他Java8会默认实现。



```java
public static void main(String[] args) {
  CnfigurableApplicationContext run = SpringApplication.run(MainApplication.class, args);

  // 查看容器里面的组件
  String[] names = run.getBeanDefinitionNames();
  for (String name : names) {
    println(name);
  }
}
```



bean

```java
@Configuration(proxyBeanMethods = true) // 告诉springboot这是一个配置类 == 配置文件
public class MyConfig {
  @Bean
  public User user01() {
    return new User("zhangsan", 18);
  }
  
  @Bean("tom")
  public Pet tomcatPet() {
    return new Pet("tomcat");
  }
}
```



## 参考资料

- [java model类_java 实体类entity，model模型，javabean的理解以及使用场景](https://blog.csdn.net/weixin_34646919/article/details/114200617)

- [lombok注解来简化java代码](https://blog.csdn.net/xue632777974/article/details/80437452)



# dubbo

https://dubbo.apache.org/zh/

Dubbo采用全Spring配置方式，透明化接入应用，对应用没有任何api侵入，只需用Spring加载dubbo的配置即可。dubbo基于Spring的schema扩展进行加载。



