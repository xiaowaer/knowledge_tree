### A3_Ti2

#### <span style="color:SlateBlue ">0.问题点:gdb并不能直观的展示出二进制文件的运行.</span>:

```
    当用GDB调试我们的代码的时候存在的几个痛点:
        1.UI不能看见寄存器,不能看见内存和堆栈的使用情况.
        2.GDB 内嵌的命令太多,不够聚焦.
        3.GDB的文档垃圾,看见就不想看.
        4.老问题,国内的教程很少.
    
    说白了缺UI,缺文档,缺插件.
```

+ <h5 style="color: SlateBlue ">1.Radare2框架简介</h5>

学习参考:https://book.rada.re/first_steps/overview.html

```
    是什么?
        二进制分析文件分析框架:
    为什么用?
        1.我们还处于观望状态,目前还不知道能不能用上啊.
    我们的愿景:
        希望通过它能让我能对二进制文件的执行有更深入的理解.
    框架简介:
        通过命令集合的方式来使用,每个命令是独立的,也可以协同作用.
    命令集合简介:
        radare2:
           核心命令:允许我们打开大量输入/输出源,并像普通文件一样展示他们.
                    包括磁盘、网络连接、内核驱动程序、调试中的进程等等.
           提供功能:
                移动文件,分析数据,反汇编,二进制补丁,数据比较,搜索,替换和可视化.
                拓展,支持Python,ruby,javaScript,Lua和Perl
        
        rabin2:
            作用:从二进制文件中提取信息,例如ELF,PE,JavaClass,Mach-O,以及其他.
            信息包括:symbols,imports,file information,cross references(xrefs),library dependencied和sections.

        rasm2:
            简述:一个支持多种架构的汇编及反汇编程序.
            不支持:暂时未知.

```

+ <h5 style="color: SlateBlue ">怎么用Radare2框架?</h5>

```
    参考资料:
        官方文档:https://book.rada.re/
    
    


```