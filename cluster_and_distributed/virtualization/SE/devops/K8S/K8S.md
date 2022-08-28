
### A3_K8S

加油!办法总比困难多哦!
#### <span style="color:SlateBlue ">在学习K8S中遇到的问题.</span>


####  <H6> <span style="color:SlateBlue ">最基本的问题?</span></H6>

```
    K8S是什么?
        软件,他自身是不提供虚拟化能力的,虚拟化能力是又想openstack,kvm,docker基础软件这些来提供.
    

```

####  <H6> <span style="color:SlateBlue ">资源问题:存储</span></H6>

```
    分布式存储和本地存储资源的在K8S上不同.

```

#### <span style="color:SlateBlue ">0.Kubernetes中的概念</span>

```
    资源对象:
        Node
        Pod
        Replication
        Controller
        Service

    资源对象控制命令行:kubectl
        执行对资源对象的增删改查并保存在etcd中持久化.

    k8s可以采用YAML或json格式声明(定义或创建)一个Kubernetes对象,对象格式是特定的.
        资源对象中的Alpha和Beta Annotations 可能会被取消.

    

```



#### <span style="color:SlateBlue ">1.K8S是什么?</span>

```
    K8S是一个给容器编排引擎,
        它的功能是自动部署,管理和伸缩容器.
    
    当你部署了Kubernetes,你就会得到一个集群,这个集群由表示控制平面的组件和一组称为节点的机器组成.
    最低配置: 只是有表示控制平面的组件和一个工作节点(放置容器服务).

    k8s 分成两侧一个是Master,主要是一组控制面板的进程,pod 调度,弹性伸缩这些功能依赖它实现.
        另一侧是pod节点侧,主要是业务开发.
    
    Master 的高可用建议是三台.
    Master组件:
        kube-apiserver:
            提供了Http Rest接口的关键服务进程,是API调用入口.
        
        kube-controller-manager:
            Kubernetes里所有资源对象的自动化控制中心,可以理解成资源对象"总管"

        kube-scheduler: 负责资源调度的进程.

        etcd服务: k8s中所有的资源对象的数据对被保存在etcd中.

        除了Master,Kubernetes集群中的其他机器被称为Node.

        node节点运行的过程:
            kubelet:负责Pod对应的容器的创建,启停等任务,同时与Master密切协作,实现集群管理.
        
            kube-proxy:实现Kubernetes Service的通信与负载均衡机制的重要组件.

            docker Engine: 负责本机的容器创建和管理工作

        node和master 之间的联系:
            kubelet 向master 注册自己,并定时向Master 汇报健康信息,资源情况.
            master 根据这些信息实现高效均衡的调度策略.

    pod:
        每个pod包含一个Pause 容器 和若干紧密相关的用户业务容器.

    为什么要引入Pod这种结构:
        pod 里面是一个完整的业务单元. 

    pod 网络:通常使用Flannel,Open vSwitch等来做.(容器间通信)

    pod的简单介绍:
        pod分为动态pod和静态pod
        静态pod,不存储在etcd中,而是存放在某个具体的Node上的一个具体文件中,并且只在这个Node上启动,运行.

        动态pod一旦被创建,就会被放在etcd中存储,随后Master调度到某个具体的Node上并进行绑定,
        随后该Pod被对应的Node上的kubelet进程实例化成一组相关的的哦仓库儿容器并启动.
        当Pod里某个容器停止时,K8S自动重启,失败就将这个Node上的所有Pod重新调度到其他节点上.

    Label(标签)
        可以用来标记不同版本的应用到不同的环境中;例如,部署不同版本的应用到不同的环境中;
        监控和分析应用(日志记录,监控,告警)等.

        
    Label Selector在Kubernetes中重要使用场景:
        kube-controller进程通过在资源对象RC上定义的Label Selector 来筛选要监控的Pod 副本数量.
        使Pod副本数量始终符合预期设定的全自动控制流程.

        kube-proxy进程通过Service 的 Label Selector来选择对应的Pod,自动建立每个Service到对应Pod的请求转发路由表
        实现智能负载均衡.

        通过某些Node定义特定的Label,可以实现基于Pod定向调度的特性.

    Replication Controller
        RC是K8S的核心概念,它其实定义了一个期望的场景,
            Pod 期待的副本数量.
            用于筛选Pod的Label Selector
            当Pod副本数量小于预期数量时,用于创建新Pod的Pod模板.

  
```
#### <span style="color:SlateBlue ">2.为什么用K8S?</span>

```
    1.容器云编排的统一规范就是K8S,龙头. CI/CD的最终制品就是容器.
    2.k8s降低了落地分布式系统的成本.
    
```

#### <span style="color:SlateBlue ">3.为什么你会对 k8s 有一种畏难情绪?</span>

```
    1.不了解集群的概念,

```

#### <span style="color:SlateBlue ">3. k8s 怎么用?</span>

```
    看见一个简单的kubectl 部署用RC,service配置文件起容器的感觉.
    kubectl 有种和 docker类似的亲切感.



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



