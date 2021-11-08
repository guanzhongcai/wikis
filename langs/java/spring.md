# spring



## IOC理论推导

dao: UserDao UserDaoImpl UserDaoMysqlImpl UserDaoOracleImpl
service: UserService UserServiceImpl

原代码：

```java


private UserDao userDao = new UserDaoImpl();
public void getUser() {
  public void getUser() {
    userDao.getUser();
  }
}
```

IOC代码：

```java
// 需要什么，自己set后再get
UserService userService = new UserServiceImpl();
((UserServiceImpl)userService).setUserDao(new UserDaoImp());
//((UserServiceImpl)userService).setUserDao(new UserDaoMysqlImp());
//((UserServiceImpl)userService).setUserDao(new UserDaoOrcalImp());

userDao.getUser();
```

用户的需求可能会影响我们原来的代码，IOC可以改很少，实现需求只做源码文件新增后的注入！！

使用了set注入后，程序变成了被动的接受对象！而非程序员主动创建对象！程序员不用再手动去创建对象了！