### A3_Ti6

#### <span style="color:SlateBlue ">1.什么是GUI?</span>

```
    接口:
        我们平时用编程叫计算机帮我们做事情就是通过写接口的方式实现.

    我们需要程序是有图形化的,比如有按钮,弹出窗口这些...
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



