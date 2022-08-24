### A3_GUI_MY_DESKTOP_POLYBAR

#### <span style="color:SlateBlue ">POLYBAR</span>

#### <H6><span style="color:SlateBlue ">1.POLYBAR是什么?</span></H6>

先让我们深吸一口气,虽然使用I3有点麻烦,但是值得.

```
    本质:
        基于X11(不用关心上层是KDE还是GTK)编写的一套平铺窗口管理器.

    什么是 window manager?
        核心功能:
            提供窗口边框的X客户端,他控制图形程序的外观和行为方式:
                控制边框,标题栏,窗口大小
        额外功能:
            提供应用程序面板,Fluxbox提供窗口标签功能，此外还有启动程序的菜单、窗口管理器配置菜单等。
            窗口管理器一般不提供额外的组件，比如图标之类的，它们一般由桌面环境提供。
            因此，窗口管理器通常不怎么耗费系统资源。

    分类:
        堆叠式（或悬浮式）窗口管理器，顾名思义，不同窗口可以相互重叠，就像桌子上随意摆放的白纸一样。

        平铺式（或直译瓦片式）窗口管理器，其中的窗口不能够重叠，而是像瓦片一样挨个摆放。
            这种窗口管理器一般比较依赖键盘操作，较少使用鼠标。
            此类窗口管理器一般也是高度可定制的。

```

#### <H6><span style="color:SlateBlue ">为什么用I3?</span></H6>

```
    1.提升开发效率
    2.好看

```

#### <H6><span style="color:SlateBlue ">怎么用I3?</span></H6>

```
    1.安装: apt install i3 
    2.安装桌面
        apt intall feh
        feh --bg-scale xxx
    
    3. 终端透明 
        
```