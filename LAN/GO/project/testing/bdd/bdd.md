
### A3_bdd

#### <span style="color:SlateBlue ">0.为什么用BDD?</span>

```
    希望可以先清楚的表达自己要做的事情是什么,提高表达能力.
```
#### <span style="color:SlateBlue ">1.什么是BDD?</span>

```
   1.它具有可读性更好(接近英语)的DSL,而非特定的编程语言。
   2.减少跨语言交流带来的信息丢失.
   3.给于用户故事
   4.希望 => DSL 自动转码到 特定语言单测. 
        
```

#### <span style="color:SlateBlue ">2.BDD发展历史</span>

```
        1.Dan North创造了首个BDD框架：JBehave[1]；
        2.之后是Ruby语言的基于故事的RBehave[4]，后来被纳入了RSpec项目[5]。
        3.RSpec中第一个基于故事的框架，后来被主要由Aslak Hellesøy开发的Cucumber 取代。
        4.提出了特性注入(Feature Injection)[6]的想法，使BDD可以覆盖分析空间并提供从初期的展望到编码和发布的整个软件生命周期过程的改造。


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



