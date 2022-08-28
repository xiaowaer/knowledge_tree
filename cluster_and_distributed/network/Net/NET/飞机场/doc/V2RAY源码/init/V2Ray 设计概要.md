# V2Ray 设计概要

## 目标
* V2Ray 自身提供基础的翻墙功能，单独使用可满足日常需求；
* V2Ray 可被用于二次开发，可为用户提供个性化的翻墙体验，从而也达到难以统一屏蔽之效果；
* V2Ray 为模块化设计，模块之间互相独立。每个模块可单独使用，也可和其它模块搭配使用。

## 架构

### 术语
* VPoint：一个 V2Ray 服务器称为 VPoint
* VSet：本机上的一组 VPoint
* VSuperSet：多机环境中的多个 VSet
* VSource：用户所使用的需要翻墙的软件，比如浏览器
* VEnd：用户需要访问的网站
* VUser：一个受到 VPoint 认证的帐号
* VID：全局唯一的 ID，类似于 UUID


### 工作流程

VPoint( V2Ray 服务器)  可提收来自 VSource (用户所使用的需要翻墙的软件，比如浏览器) 或其它 VPoint 的请求，并将请求转发至配置中的下一个 VPoint（或 VSet 或 VSuperSet） 或目标网站，然后将所得到的应答回复给请求来源。

VPoint( V2Ray 服务器) 采用白名单机制，只接受已认证帐号的请求。

### 通信协议
* VPoint 之间默认使用自有 VMess 协议，或第三方自定义协议。

* VPoint 和客户端之间可使用以下协议：
  * HTTP Proxy
  * SOCKS 5 Proxy
  * PPTP / L2TP / SSTP 等 VPN 隧道
  * 其它自定义协议
* VPoint 和目标网站之间使用以下协议：
  * HTTP / HTTPS
  * UDP (DNS)

#### VMess
VMess 为 V2Ray 的原生协议，设计用于两个 VPoint 之间的通信。[详细设计](https://github.com/V2Ray/v2ray-core/blob/master/spec/vmess.md)

### VUser
* 每个 VUser 有一个 VID

### VPoint
* 每个 VPoint 有一个 VID，运行时生成
* 每个 VPoint 可使用独立的配置文件，或从 VSet 继承
* 一个 VPoint 监听主机上的一个特定端口（可配置），用于接收和发送数据
* 一个 VPoint 运行于一个独立的进程，可设定其使用的系统帐户

### VSet
TODO

### VSuperSet
TODO

## VPoint 详细设计
一个 VPoint 包含五个部分：
* 配置文件处理：读取和解析配置文件

* 输入：负责与客户端建立连接（如 TCP），接收客户端的消息

* 控制中心：负责处理事件
  * 加密解密
  * VPoint 负载均衡

* VPoint 进程间通信

* 输出：负责向客户端发送消息

### 配置文件
配置文件使用 JSON / ProtoBuf 兼容格式，定义 TODO

### 加密
TODO

### 任务处理
TODO

### 控制中心
控制中心响应以下事件：

**INIT**
* 输入：用户 VID
* 输出：如果用户 VID 有效："OK"，否则关闭连接

**MSG**
* 输入：VMess 消息
* 输出：对应的响应消息

**END**
* 输入：用户 VID
* 输出：如果用户 VID 有效：关闭连接

## 编程语言
暂定为 golang。