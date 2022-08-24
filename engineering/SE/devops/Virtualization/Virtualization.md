
### A3_bdd

#### <span style="color:SlateBlue ">什么是虚拟化?</span>

```
    本质就是抽象并仿真

    是从不可以到我可以,而不是变成我很吊
        简单点说6核就是6核,就是你虚拟了128核,底层还是6核.
    (算力不以为虚拟化而飞跃式提升.)
    
    全虚拟机:模拟硬件.为了解决CPU性能过剩的问题。

    进程虚拟机:指令级虚拟化+跨平台.

    容器虚拟化: 操作系统虚拟,只有一个内核,但通过特定技术虚拟出一个隔离环境.
    
```

#### <span style="color:SlateBlue ">几种虚拟化技术对比</span>

```
作者：知乎用户G0K17q
链接：https://www.zhihu.com/question/24123210/answer/100874195
作者：知乎用户G0K17q
链接：https://www.zhihu.com/question/24123210/answer/100874195
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```

```
    eqmu : 
        如安卓模拟器,指令集是解释型的,慢的拉稀.
        好处是不调硬件,什么都能模拟.

    支持特定架构的虚拟化技术,比如我们用X64只想虚拟化X64
        我们可以变得更快,因为可用不用翻译指令,通过其他手段,来切换执行.

    Hypervisor/XEN: 
        IOS启动Hypervisor，Hypervisor首先启动第一个OS，那个就是Host，
        之后Host上的Emulator启动Guest，Emulator向Hypervisor请求资源，
        Hypervisor模拟一个环境给新的Guest，
        但如果Guest向Hypervisor请求IO或者其他特殊能力，
        这些请求就会被Hypervisor调度到Host上，让Host执行。

    KVM:
         而KVM的Hypervisor直接就是内核的一部分，
         这个Hypervisor的代码直接就在Linux的内核中，
         当Host启动的时候，它们一起加载，一同初始化，
         只是Hypervisor的代码工作在虚拟机调度器的状态，
         而其他代码工作在普通内核状态而已。

```