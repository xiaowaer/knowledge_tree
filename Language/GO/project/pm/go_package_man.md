# A3_GO语言环境关于IDE的配置 

#### <span style="color:SlateBlue ">GO的包管理</span>

#### <span style="color:SlateBlue ">初体验</span>

```
    1. 其他语言的包管理.
        (平时都是IDE或者工具帮助我们过度了这个file 结构->project的过程 )
      对比java,java的包管理底层都是统一的jar格式,ruby是gems.
      java可以用gradle,maven来做包管理,ruby可以用Bundler来打包.
        这些包管理都是基于约定的.目录格式统一.
    
    2.go 的包管理是怎么样子的?
        

    3.仓库,模块,包,源代码之间的关系.

        模块:是包的集合,这个集合内的包将一起被发布,版本化和分发. 
            特点:
                模块是可以远程仓库拉取的单位.
                模块被模块路径标识.

    4.go.mod 文件  
        记录模块路径

    5.模块路径
    组成:通常由一个存储库root路径,directory within the repository,
         and a major version suffix (only for major version 2 or higher).

    2.简单安装 
        go install -v honnef.co/go/tools/cmd/staticcheck@latest
        go install -v honnef.co/go/tools/cmd/staticcheck@2020.2.1
   
   3.外壳是操作系统命令行

   4.https://staticcheck.io/docs/checks (检查内容)

```

#### <span style="color:SlateBlue ">2.staticcheck的优势?</span>

```
    1.它很负责,超过150次的检查确保您的代码处于最佳状态。
    2.Staticcheck can easily be integrated with your code review and CI systems, 
        preventing buggy code from ever getting committed.
        --可以集成到ci之中去,作为自动化code review的一部分.
    3.Staticcheck 是大多数主要编辑器的一部分，
        它已经与 gopls 集成在一起，发现 bug 并提供自动修复。
    --配合gopls,有bug提醒并修复的作用.
   
```

#### <span style="color:SlateBlue ">3.staticcheck命令参数?</span>

```
文档不多,说明和简单.
    Usage: staticcheck [flags] [packages]
        lags:
         -checks checks	Comma-separated list of checks to enable. (default "inherit")
         [要检查的内容]  --核心
         -explain check	Print description of check
         [告诉我们静态检查的结果] --核心
         -f format  Output format (valid choices are 'stylish', 'text' and 'json') (default text)
         [静态检查结果的输出文本格式]
         -fail checks	Comma-separated list of checks that can cause a non-zero exit status. (default "all")
        [失败结果检出] --核心
         -go version  Target Go version in the format '1.x', or the literal 'module' to use the module's Go version (default "module")

         -list-checks  List all available checks
        [检查列表]
         -matrix   Read a build config matrix from stdin

         -merge    Merge results of multiple Staticcheck runs

         -show-ignored Don't filter ign ored diagnostics
    
         -tags build tags  List of build tags

         -tests Include tests (default true)

         -version Print version and exit

# A3_GO语言套件 -- staticcheck

#### <span style="color:SlateBlue ">staticcheck</span>

#### <span style="color:SlateBlue ">1.staticcheck是什么?</span>

学习参考地址:https://staticcheck.io/

```
    1.简述
        staticcheck 是 Go 的一套静态分析工具。它专注于查找错误，代码简单，性能和编辑器集成。
        Staticcheck 是 Go 编程语言最先进的连接器。通过使用静态分析，
        它可以发现 bug 和性能问题，提供简化，并强制执行样式规则。


    2.简单安装 
        go install -v honnef.co/go/tools/cmd/staticcheck@latest
        go install -v honnef.co/go/tools/cmd/staticcheck@2020.2.1
   
   3.外壳是操作系统命令行

   4.https://staticcheck.io/docs/checks (检查内容)

```

#### <span style="color:SlateBlue ">2.staticcheck的优势?</span>

```
    1.它很负责,超过150次的检查确保您的代码处于最佳状态。
    2.Staticcheck can easily be integrated with your code review and CI systems, 
        preventing buggy code from ever getting committed.
        --可以集成到ci之中去,作为自动化code review的一部分.
    3.Staticcheck 是大多数主要编辑器的一部分，
        它已经与 gopls 集成在一起，发现 bug 并提供自动修复。
    --配合gopls,有bug提醒并修复的作用.
   
```

#### <span style="color:SlateBlue ">3.staticcheck命令参数?</span>

```
文档不多,说明和简单.
    Usage: staticcheck [flags] [packages]
        lags:
         -checks checks	Comma-separated list of checks to enable. (default "inherit")
         [要检查的内容]  --核心
         -explain check	Print description of check
         [告诉我们静态检查的结果] --核心
         -f format  Output format (valid choices are 'stylish', 'text' and 'json') (default text)
         [静态检查结果的输出文本格式]
         -fail checks	Comma-separated list of checks that can cause a non-zero exit status. (default "all")
        [失败结果检出] --核心
         -go version  Target Go version in the format '1.x', or the literal 'module' to use the module's Go version (default "module")

         -list-checks  List all available checks
        [检查列表]
         -matrix   Read a build config matrix from stdin

         -merge    Merge results of multiple Staticcheck runs

         -show-ignored Don't filter ign ored diagnostics
    
         -tags build tags  List of build tags

         -tests Include tests (default true)

         -version Print version and exit

```

#### <span style="color:SlateBlue ">4.staticcheck 检查内容</span>

```
   
```

#### <span style="color:SlateBlue ">4.staticcheck使用小demo?</span>

```
   
```


#### <span style="color:SlateBlue ">3.怎么配置好我们的staticcheck?</span>

```
   
```

#### <span style="color:SlateBlue ">2.V2RAY学习</span>```

#### <span style="color:SlateBlue ">4.staticcheck使用小demo?</span>

```
   
```



#### <span style="color:SlateBlue ">3.怎么配置好我们的staticcheck?</span>

```
   
```

#### <span style="color:SlateBlue ">2.V2RAY学习</span>