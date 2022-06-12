### A3_diskManager

#### <span style="color:SlateBlue ">痛点:</span>

```
    分区怎么做扩容:

        1.备份(防止数据丢失)--因为我们是虚拟机打个快照就好.
        
        2.df -h (查看分区挂载信息)


        
```

#### <span style="color:SlateBlue ">1.磁盘</span>

```
    1.磁盘是一种物理设备(看得见,摸得到),可以完成对数据的储存,修改.
    2.分区

```
#### <h6> <span style="color:SlateBlue ">1.1磁盘命名</span> </h6>

```
  在Linux下对 SCSI 和 SATA 设备是以 sd 命名的，第一个 scsi 设备是 sda，第二个是 sdb，依此类推。
  一般主板上有两个SCSI接口，因此一共可以安装四个SCSI设备。
  主 SCSI 上的两个设备分别对应 sda 和 sdb，第二个 SCSI 口上的两个设备对应 sdc 和 sdd。
  (IDE接口设备是用 hd 命名的，第一个设备是hda，第二个是hdb，依此类推。)
```

#### <span style="color:SlateBlue ">2.分区</span>

#### <h6> <span style="color:SlateBlue ">2.2分区是什么?</span> </h6>

```

    就是将磁盘分成一个个区域来处理.(因为磁盘本来自己也是按照扇区来划分的.)
    如果这个分区用满了是要添加新的空间的,分走了的空间就已经被占用了.(和分家产一个道理)

```

#### <h6> <span style="color:SlateBlue ">2.2分区格式化是什么?</span> </h6>

```

    不同文件系统底层有不同的磁盘分区使用方法,
    格式化是说明这块磁盘要做成什么文件系统.

```

#### <h6> <span style="color:SlateBlue ">2.2挂载是什么?</span> </h6>

```
    说明从某个目录下使用那个分区,并且赋予用户什么权限.

```

#### <span style="color:SlateBlue ">3.文件系统</span>

```
```

#### <span style="color:SlateBlue ">4.常用命令</span>

```
```