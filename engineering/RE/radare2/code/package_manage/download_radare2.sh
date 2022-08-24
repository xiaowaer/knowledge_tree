#!/bin/bash
#set -eux
# 
# 重要经验:如果缺了什么pc包,可以通过debian官网查看!
# 看看那个镜像里面有,然后找个类似名字的apt install 
#  要sudo以上的权限
#      apt-get install libcapsone-dev; apt-get install  capstone-tool:
#

create_download_dir(){
if [ ! -d "/home/netty/devops/debug/radare2" ]; 
then
  mkdir /home/netty/devops/debug/radare2
  echo '已经创建安装目录'
fi
#进入安装目录
cd "/home/netty/devops/debug/radare2"
}

select_download_src(){
# 拉取源文件方式  
echo '现在开始拉取要安装的radare2文件'
echo "请选择拉取方式,
            1.(默认)从github拉取,
            2.网络有问题,要转本地文件安装"  
#读取从键盘的输入,模式不匹配直接,打印退出. 
read -n1 -p "Do you want to select git to pull repo [Y/N]?" answer


case $answer in
Y | y)
      echo "fine ,我们去拉取仓库去"
      cur_dir=$(pwd)
      echo "当前目录为 '$cur_dir'"
      git -v
      git clone https://github.com/radareorg/radare2 
       echo "仓库以下好"
      ;;
N | n)
      echo "知道了,你要使用本地文件,功能尚未完善";;
*)
     echo "error choice"
     exit 0
     ;;
esac

}
install(){
     radare2/sys/install.sh #要sudo以上权限,和 capstone
}
uninstall(){
      #To uninstall the current build of r2 run make uninstall
     # To uninstall ALL the system installations of r2 do: sudo make purge
}

#######################################主执行目录#################################
#切换目录,没有就创建
#查看路径是否存在
main(){

create_download_dir;
select_download_src;
install;

}

main
# echo 'hello_world!'
