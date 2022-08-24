### A3_GO_testing

#### <span style="color:SlateBlue ">1.为什么要测试?</span>

```
    为了让测试驱动开发.
    通过测试驱动(TDD)来写代码,什么是测试驱动?
        巨简单,就是先写一个单元测试,结合mock,通过测试反馈来持续修改代码,知道测试不报错.
        
```
#### <span style="color:SlateBlue ">1.编码链路</span>

```
    愿景:不用自己写强耦合于语言的单元测试代码,
         希望可以用一种更接近于人类语言风格的DSL来完成测试用例的编写,
         之后再帮我们转化成单元测试接口.

 统一DSL语言的测试编写->[(bdd 框架)]->单元测试代码->[TDD(方法论)]->业务代码
       ->[(可视化手段 swag)+(接口风格 restful)]-> restful接口    --后端落地 
```


#### <span style="color:SlateBlue ">TDD的闭环</span>

```
过了一段时间，我发现如果我更改代码并导致测试失败，我可以查看测试方法名称并识别代码的预期行为。

通常有三种情况发生:
    1.我引入了一个错误。解决方案: 修复这个错误
    2.预期的行为仍然是相关的，但已经移动到其他地方。解决方案: 移动测试，也许改变它
    3.解决方案:Delete the test. 删除测试
    总之:只需完成上面描述的过程: 要么引入一个 bug，要么改变行为，要么测试不再相关。
    
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



