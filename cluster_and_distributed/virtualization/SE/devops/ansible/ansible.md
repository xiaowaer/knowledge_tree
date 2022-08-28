
### A3_ansible


#### <span style="color:SlateBlue ">1.为什么要用ansible?</span>

```
    根本原因:
        看了一眼ansible-galaxy,并爱上了.
   表现:用docker安装gitea的时候rootless run命令的写法无从下手.
   核心疼点:
        为什么我每次下载一个应用的时候,都要定制化的写一些启动命令,安装配置?
    解决方案:
        我们去学ansible,ansible在github上有50k star,有很多人用,
        你想到的任务,别人基本都用ansible实现过,复用别人的成果,解放自己.

```
#### <span style="color:SlateBlue ">1.使用ansible的一些小条件?</span>

```
    1.使用ssh连接到远程服务器
    2.会处理bash命令行的输入输出( 管道和重定向)
    3.安装软件包
    4.使用sudo命令
    5.检查和设置文件权限
    6.启动和停止系统服务
    7.设置环境变量.
    8.编写脚本(语言不限)

    注意:只要你不需要编写自己的ansible模块就不用会写python.

```

#### <span style="color:SlateBlue ">2.什么是ansible?</span>

```
    ansible愿景: 
        让自动化配置和部署变得简单.
        --这就引入了一个新的概念: 配置管理.

    ansible是什么?
        是用python写的一个配置管理工具.
    
    配置管理是做什么的?
        它是一种获取那些始终保持最新的配置与信息的方法.
        你不再需要频繁地搜索正确的文档或者查找以前的笔记.
    
    ansible为什么强大?
        你可以在用ansible远程执行shell,并且ansible已经有了很多现成的任务模板.
    
    ansible的操作幂等性:
        模块操作是幂等的,有了就不再重复叠加操作,优点,更加安全.

    ansible的package抽象:
        某些模块替代yum,apt等包管理工具.
    
    ansible playbook:
        注意:不提供不同系统环境之间复用.

    ansible role 是ansible的精华,Ansible Galaxy,是一个在线的role仓库.

    ansible模板可以用yaml 和 jinja2.
    

```

#### <span style="color:SlateBlue ">2.怎么学ansible?</span>

```
    国内写ansible的书并不多,先找一本看上去入门--<< 奔跑吧,ansible>>
```

#### <span style="color:SlateBlue ">3.BDD的核心价值是什么?</span>

```
   精简的统一表达:
    通常公司内部使用的故事模板是这样的:
        As a[x]
        I want [y]
        so that [z]

    demo for atm machine:
        Title: Customer withdraws cash**

        As a customer,
        I want to withdraw cash from an ATM,
        so that I don't have to wait in line at the bank.

        Scenario(场景):Account is in credit
        Given the account is in credit
        And the card is valid
        And the dispenser contains cash
        When the customer requests cash
        Then ensure the account is debited
        And ensure cash is dispensed
        And ensure the card is returned

```

#### <span style="color:SlateBlue ">3.BDD与闭包</span>


```
   Closure.md
   (本人巨讨厌闭包,语法看着巨恶心.)

```
#### <span style="color:SlateBlue ">4.BDD怎么用? </span>


```
    愿景:不用自己写强耦合于语言的单元测试代码,
         希望可以用一种更接近于人类语言风格的DSL来完成测试用例的编写,
         之后再帮我们转化成单元测试接口.

 统一DSL语言的测试编写->[(bdd 框架)]->单元测试代码->[TDD(方法论)]->业务代码
       ->[(可视化手段 swag)+(接口风格 restful)]-> restful接口    --后端落地 
```

#### <span style="color:SlateBlue ">单测demo V2_DEMO</span>

```
   基于v2_0.1.121版本,即V2的第一个测试代码.
     {
         被测试类: socks.go
            功能说明:
         测试类: socks_test.go
         测试框架: testing

     }


```

#### <span style="color:SlateBlue ">linux是用什么支撑起GUI编程的?</span>

```
    Linux系统的图形架构:
----------------------------------------------
普通PC的Linux操作系统架构如下：
    Linux Kernel + XLib + GLib + GNOME GNOME/GTK
    Linux Kernel + XLib + Qt   + KDE    KDE/Qt
而在一般采用Qt嵌入式操作系统的架构如下：
    Linux Embedded Kernel + Framebuffer + Qte + Qtopia
-------------------------------------------------------
　内核 -->X服务器<-[通过X11协议交谈]-> 
(图形化实现不限,可以是QT,也可以是GTK)窗口管理器（综合桌面环境）--> X应用程序。

X server是xfree86/xorg驱动下的显示设备鼠标键盘统称，
X client通过X11协议和xfree86/xorg实现的X server通信，
比如，告诉它画一个左上角坐标为(x,y)，宽为w，高为h的窗口，
xfree86就让显示器把屏幕上的小灯（像素）打亮，然后你就看到了一个窗口。
为了方便开发人员编写X clients，就有了Xlib来封装协议；
Xlib不够方便，于是就有了qt和gtk，提供了很多窗口控件（widgets）。
为了方便用户，就出现了gnome和kde等桌面管理系统。
一般来说，linux用户看到的界面就是其中之一了。gnome用的是gtk库，kde用的是qt库。

```



