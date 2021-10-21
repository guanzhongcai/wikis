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
spring.profiles.active=test
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

