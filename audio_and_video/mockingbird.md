### A3_mockingbird

#### <span style="color:SlateBlue ">1.为什么用mockingbird?</span>

```
  因为我们想通过视频格式来做知识分享,
  视频格式主要包含两个方面,视频和音频.

  音频方面:我们语速很慢,而且声音不好听,很不如语言转文字.
  解决痛点:配音.

```

#### <span style="color:SlateBlue ">怎么下载并开始使用?</span>

这里要感谢:
    up主:什么都懂一点的奶糖的视频教程.

```
    主要流程:

        下载软件依赖:
             1.Python 3.8 或更高 (ok)
             2.PyTorch 
             3.ffmpeg
             4.还有一些项目内置的依赖 

    问题:磁盘空间不足,安装虚拟机的时候不要默认,/tmp会很小,/ 挂载的也很小
    解决方案:
        
        1.新建一个盘.
        2.lsblk
        3. mkfs.ext4 /dev/sdc
        4. mount 

    2.扩容加磁盘组lvm,
        fdisk -l
        fdisk /dev/sdd n p 1 t 1 w
        5. pip3 install torch torchvision torchaudio -i https://mirrors.aliyun.com/pypi/simple/ -v



```


