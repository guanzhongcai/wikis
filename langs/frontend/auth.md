# Oauth2.0

1、官方定义：

OAuth 2.0是行业标准的**授权协议**。OAuth 2.0取代了2006年创建的原始OAuth协议所做的工作.OAuth 2.0专注于客户端开发人员的简单性，同时为Web应用程序，桌面应用程序，移动电话和客厅设备提供特定的**授权流程**。

2、应用场景：

比如淘宝商家平台，如果淘宝卖家想把自己的订单交给别的ERP厂商处理，这里卖家有两种选择：



**OAuth2.0的四个重要角色**

Resource Owner：资源拥有者，上例中的淘宝卖家；

Resource Server：资源服务器，上例中的淘宝卖家平台；

Client：第三方应用客户端，上例中的ERP厂商；

Authorization Server：授权服务器，管理Resource owner、Resource Server和Client的关系



## 协议流

```
     +--------+                               +---------------+
     |        |--(A)- Authorization Request ->|   Resource    |
     |        |                               |     Owner     |
     |        |<-(B)-- Authorization Grant ---|   淘宝卖家     |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(C)-- Authorization Grant -->| Authorization |
     | Client |                               |     Server    |
     |        |<-(D)----- Access Token -------|   授权服务器		|
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(E)----- Access Token ------>|    Resource   |
     |        |                               |     Server    |
     |        |<-(F)--- Protected Resource ---|  淘宝卖家平台   |
     +--------+                               +---------------+

                     Figure 1: Abstract Protocol Flow
```



## 刷新token

```
  +--------+                                           +---------------+
  |        |--(A)------- Authorization Grant --------->|               |
  |        |                                           |               |
  |        |<-(B)----------- Access Token -------------|               |
  |        |               & Refresh Token             |               |
  |        |                                           |               |
  |        |                            +----------+   |               |
  |        |--(C)---- Access Token ---->|          |   |               |
  |        |                            |          |   |               |
  |        |<-(D)- Protected Resource --| Resource |   | Authorization |
  | Client |                            |  Server  |   |     Server    |
  |        |--(E)---- Access Token ---->|          |   |               |
  |        |                            |          |   |               |
  |        |<-(F)- Invalid Token Error -|          |   |               |
  |        |                            +----------+   |               |
  |        |                                           |               |
  |        |--(G)----------- Refresh Token ----------->|               |
  |        |                                           |               |
  |        |<-(H)----------- Access Token -------------|               |
  +--------+           & Optional Refresh Token        +---------------+

               Figure 2: Refreshing an Expired Access Token
```



# OIDC

OIDC是OpenId Connect的简称

OpenId Connect的定义：

OpenID Connect 1.0是OAuth 2.0协议之上的简单身份层。它允许客户端根据授权服务器执行的身份验证来验证最终用户的身份，以及以可互操作和类似REST的方式获取有关最终用户的基本配置文件信息。

**OIDC=(Identity,Authentication)+OAuth2.0**

Oauth2.0本身是一个授权协议，并不提供身份认证功能。

OIDC使用OAuth2.0的授权服务器来为第三方用户提供身份认证，并将对于的身份认证信息传递给客户端，且可以适用于各种类型的客户端（比如服务端应用，移动APP，JS应用），且完全兼容OAuth2。



## OIDC 主要术语

主要的术语以及概念介绍（完整术语参见http://openid.net/specs/openid-connect-core-1_0.html#Terminology）：

1. EU：End User：一个人类用户。
2. RP：Relying Party ,用来代指OAuth2中的受信任的客户端，身份认证和授权信息的消费方；
3. OP：OpenID Provider，有能力提供EU认证的服务（比如OAuth2中的授权服务），用来为RP提供EU的身份认证信息；
4. ID Token：JWT格式的数据，包含EU身份认证的信息。
5. UserInfo Endpoint：用户信息接口（受OAuth2保护），当RP使用Access Token访问时，返回授权用户的信息，此接口必须使用HTTPS。



## 参考资料

- [Oauth2.0与OIDC区别](https://www.cnblogs.com/tlj2018/articles/10736788.html)

- [[认证 & 授权\] 4. OIDC（OpenId Connect）身份认证（核心部分）](https://www.cnblogs.com/linianhui/p/openid-connect-core.html)

- [The OAuth 2.0 Authorization Framework](https://datatracker.ietf.org/doc/html/rfc6749)

- [OAuth 2.0 Playground](https://www.oauth.com/playground/)
