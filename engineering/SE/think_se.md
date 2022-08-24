### A3_Ti3


#### <span style="color:SlateBlue ">1.什么是软件工程?</span>

```
    愿景:让软件开发变得多快好省.

    是什么? 
        ACM与IEEE Computer Society联合修定的SWEBOK[14]（Software Engineering Body of Knowledge）提到，软件工程领域中的核心知识包括：
            软件需求（Software requirements）
            软件设计（Software design）
            软件建构（Software construction）
            软件测试（Software test）
            软件维护与更新（Software maintenance）
            软件构型管理（Software Configuration Management, SCM）
            软件工程管理（Software Engineering Management）
            软件开发过程（Software Development Process）
            软件工程工具与方法（Software Engineering Tools and methods）
            软件质量（Software Quality）

```

#### <span style="color:SlateBlue ">2.为什么要引入软件工程?</span>

```
    为了解决软件危机.
```

#### <span style="color:SlateBlue ">3.什么是软件危机?</span>

```

    软件开发不符合最新的市场需求:
        人话:
            1.你做的垃圾比别人贵.
            2.你做的垃圾,没人用.
            3.你做的垃圾,不适合协同开发,改不了.


    第一次软件机遇:
        
        期待: 加快速度,简化编程.
        现状: 
            IBM研发了一款701计算机,在当时快的一批,但是没人能很好的利用它.
        问题:
            汇编语言的编程,成为了限制了我们高效的使用这台机器.
            1.汇编人才太少
            2.汇编写的也慢.
            --所以没什么人能用上这样的机器.
        解决方案: FORTRAN语言诞生,第一门高级语言,从此软件开发的大军迈向了高级语言的使用.

        =>过渡历史: IBM704绑定FORTRAN,程序员开始用fortran,1960-1970Lisp,Algol,COBOL,Snobol等语言陆续诞生.

    第二次软件机遇:
        发现: 
        原来开发软件系统比预想的难,编程成为职业,软件开始被看成生意. (你能做别人不能做的事情,你就可以挣钱.)
        {
            :FORTRAN => 主要用于工程和科学研究
            :COBOL => 主要用于一般商业数据处理  => 解决了如何在编程语言中处理业务数据.
            :SQL => 主要用于数据库.
            :Lisp => 主要用于研究人工智能.   
            
            在没有C和UINX的时代是巨型机的天下.
        }
        发展:pc机出现,BASIC语言,intel芯片开始流行,普通个体开始广泛接触计算机,VB语言流行.
        期待: 加快速度,简化编程.

        现状: 
            肯汤普森讨厌ibm 360系列产品=>后面就有PDP装上UNIX并且开发了C语言的故事.
            ibm 360主机不适合分时共享系统,  IBM大型机是批处理任务系统.
            =>Unix就是这种基于分布式系统概念的交互系统

        问题:接触编程的人从商用转移到民用  
        解决方案: unix和c开始普及

    第三次软件机遇:
            
            问题:C,C++代码长了之后代码结构越来越难维护,网络时代降临,款平台的问题出现.
            本质:对了一个类的数据结构,并且根据这个数据结构,要特定的方式组织代码.
            解决方案:面向对象的高级语言诞生,java,ruby等 
            java优点: 语法比C++简单,规范严格,面向对象,具有垃圾收集机制.

    是不是有了java就很完美,明显不是的!

         第四次软件机遇:
            问题:
                以前很多软件架构都是单核架构的,现在cpu都多核心了,并行编程是必须的.
                因为现在互联网用户巨多,软件系统有三个追求：高性能、高并发、高可用,也就是软件要支持三高场景.
            

         第五次软件机遇:   
            问题:
                用户多了,就有傻逼用户,所以要面对经常变更的需求,这个事情就要求软件开发方法足够敏捷.
            
            解决方案:

```

#### <span style="color:SlateBlue ">2.开发模型选型?</span>

```
    我们选敏捷开发.
```

#### <span style="color:SlateBlue ">3.方法论和架构</span>

```
    我们选devops,看板,插件式开发,

```
