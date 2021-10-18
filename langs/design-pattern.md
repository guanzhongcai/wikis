# 设计模式

分为三大类：

- 创建型设计模式：Creational，解决如何创建对象
- 结构型设计模式：Structional，解决如何实现类或对象的组合
- 行为型设计模式：Behavioral，解决类或对象怎样交互已经怎样分配职责



设计模式的基础是：**多态**



## 观察者模式

以英雄们打boss为例，此设计模式的关键要点：

- 被观察目标：boss

- 观察者：heroA、heroB、。。
- 通知观察者：boss.notify
- 接收通知：hero.update

```cpp
AbstractHero *heroA = new HeroA();
AbstractHero *heroB = new HeroB();
AbstractHero *heroC = new HeroC();

AbstratBoss* boss = new Boss();
boss->addHero(heroA);
boss->addHero(heroB);

boss->deleteHero(heroC);
boss->notify(); // forEach heros: heroX->update();

delete heroA, heroB, boss;
```



## 策略模式

以英雄所使用的武器为例，此设计模式的关键要点：

- 英雄character类：接收不同的武器，和使用武器的策略接口
- 每种武器都继承于抽象的武器策略类，且各自实现基类的UseWeapon策略实现接口

```cpp
Character* character = new Character();

WeaponStrategy* knife = new Knife();
character.setWeapon(knife);
character.useWeapon();
delete knife;

WeaponStrategy* ak47 = new AK47();
character.setWeapon(ak47);
character.useWeapon();
delete ak47;
```



## 命令模式

以客户端和服务器间交互命令为例，此设计模式的关键要点：

- 将各命令包装成一个command类
- 各command类继承于一个抽象command类
- server类有个command命令队列，用于命令的存储和消费
- server类还有个命令的add方法，用于添加命令
- server类还有个命令的handle方法，不间断的消费命令

```cpp
class HandleClientProtocol {
public:
  void AddMoney() {};
  void AddDiamond() {};
}

class AbstractCommand {
public:
  virtual void handle() = 0;
}

class AddMoneyCommand : public AbstractCommand {
public:
  AddMoneyCommand(HandleClientProtocol* protocol) {
    this->mProtocol = protocol;
  }
  virtual void handle() {
    this->mProtocol->AddMoney();
  }
private:
  HandleClientProtocol* mProtocol;
}

class Server {
private:
  queue<AbstractCommand*> mCommands;
}
HandleClientProtocol* protol = new HandleClientProtocol();

AbstractCommand* addMoney = new AddMoneyCommand(protocol);
AbstractCommand* addDiamond = new AddDiamondCommand(protocol);

Server* server = new Server();
server->addRequest(addMoney);
server->addRequest(addDiamond);

server->startHandle(); // consume commands
```



## 代理模式

在代理模式（Proxy Pattern）中，一个类代表另一个类的功能。这种类型的设计模式属于结构型模式。

在代理模式中，我们创建具有现有对象的对象，以便向外界提供功能接口。

**使用场景：**按职责来划分，通常有以下使用场景： 1、远程代理。 2、虚拟代理。 3、Copy-on-Write 代理。 4、保护（Protect or Access）代理。 5、Cache代理。 6、防火墙（Firewall）代理。 7、同步化（Synchronization）代理。 8、智能引用（Smart Reference）代理。

**注意事项：** 1、和适配器模式的区别：适配器模式主要改变所考虑对象的接口，而代理模式不能改变所代理类的接口。 2、和装饰器模式的区别：装饰器模式为了增强功能，而代理模式是为了加以控制。

栗子：加载图片类，代理图片类，前者从磁盘加载，后者只代理图片的访问接口，不care实现（怎么加载）