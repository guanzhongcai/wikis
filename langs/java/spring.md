# Spring



2. 属性如何注入

```java
@Componet
public class User {
  public String name;
  
  //相当于 <property name="name" value="kuangshen"/>
  @Value("kuangshen2")
  public void setName(String name) {
    this.name = name;
  }
}
```

3. 衍生的注解

   @Component有几个衍生注解，譬如Spring MVC的三层架构分层：

   - dao [@Repository]
   - service [@Service]
   - controller [@Controller]

   这四个注解功能都是一样的，都是代表将某个类注册到Spring中，装配Bean



## 代理模式

静态代理：

给被代理类增加一些方法，且代理类实现被代理类的主要方法，需要时直接调用，实现了业务的分工，例如spring aop。

动态代理：



JDK中有个Proxy类、InvocationHandler实现代理，在 `java.lang.reflect`包，使用反射实现的。

```java
public static Object newProxyInstance(ClassLoader loader,
                                     class<?>[] interfaces,
                                     InvocationHandler h)
  										throws IllegalArgumentException
```



## 参考资料

- https://www.bilibili.com/video/BV1WE411d7Dv?p=14