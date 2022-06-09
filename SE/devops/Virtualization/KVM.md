
### A3_kvm

#### <span style="color:SlateBlue ">KVM</span>

```
   这里我们可以理解成KVM就是内核的一个拓展包,而且现在默认也集成到内核里面了.

   基于内核的虚拟机 Kernel-based Virtual Machine（KVM）是一种内建于 Linux® 中的开源虚拟化技术。
   具体而言，KVM 可帮助您将 Linux 转变为虚拟机监控程序，
   使主机计算机能够运行多个隔离的虚拟环境，即虚拟客户机或虚拟机（VM）。

  基于KVM的虚拟机是在Linux操作系统上运行的虚拟机应用程序，虚拟机是Linux中的一个进程。
  KVM通常以kvm.ko + kvm_intel.ko/kvm_amd.ko的内核模块形式提供，虚拟机厂商基于kvm内核模块开发虚拟机应用程序。
  KVM提供了CPU虚拟化、部分中断虚拟化，虚拟机厂商自己提供UI、虚拟设备等功能以提供完整的虚拟机。
  基于KVM的虚拟机典型的有qemu-kvm，其使用qemu提供UI和虚拟设备功能，使用KVM提供CPU虚拟化、中断虚拟化功能。

    作者：ScratchLab
    链接：https://www.zhihu.com/question/24123210/answer/2442179667
    来源：知乎
    著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 忘记从哪里看到的，大约说 xen 这种 hypervisor 在内核更下层，
 实际上需要重新造很多内核的轮子还造的不如内核现有的机制好，
 比如你多个虚拟机执行时间片啥的也得有个调度器吧，
 io 虚拟化也得有个 io 调度器吧，kvm 这方面就直接重用 linux 内核里现成的，
 各种发挥内核的 economic of scope 的优势。

    作者：fleuria
    链接：https://www.zhihu.com/question/24123210/answer/732118076
    来源：知乎
    著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

```


#### <span style="color:SlateBlue ">1.kvm是怎么运行的?</span>

```
   KVM 将 Linux 转变为 1 类（裸机恢复）虚拟机监控程序。
   所有虚拟机监控程序都需要一些操作系统层面的组件才能运行虚拟机，
   如内存管理器、进程调度程序、输入/输出（I/O）堆栈、设备驱动程序、安全管理器以及网络堆栈等。
   由于 KVM 是 Linux 内核的一部分，因此所有这些组件它都有。
   每个虚拟机都像普通的 Linux 进程一样实施，
   由标准的 Linux 调度程序进行调度，并且使用专门的虚拟硬件，如网卡、图形适配器、CPU、内存和磁盘等。
        
```

#### <span style="color:SlateBlue ">2.kvm使用</span>

检查状态

```
       1. 安装 kvm 虚拟化相关软件包
        [root@kvm ~]# yum install -y kvm virt-* libvirt bridge-utils qemu-img

       2. 查看 kvm 模块是否加载到内核

        [root@kvm ~]# lsmod | grep kvm_intel
        kvm_intel              53484  0 
        kvm                   316506  1 kvm_intel
    
     如果没有加载，可以尝试执行命令：modprobe kvm_intel ，不行的话，试试重启宿主机。

     


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



