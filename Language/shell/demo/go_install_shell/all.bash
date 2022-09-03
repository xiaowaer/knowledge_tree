#!/usr/bin/env bash
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e
#如果任何语句的执行结果不是true则应该退出
#替代用法 set -o errexit

if [ ! -f make.bash ]; then
    # 如果找不到 make.sh
	echo 'all.bash must be run from $GOROOT/src' 1>&2
    # 说明目录不对 !找不到 make.bash
	exit 1
    # 程序退出 
    #
fi

OLDPATH="$PATH"
#定义变量 OLDPATH
. ./make.bash "$@" --no-banner
#$@ 是传给脚本的所有参数的列表 

#运行 make.bash`
bash run.bash --no-rebuild
PATH="$OLDPATH"
$GOTOOLDIR/dist banner  # print build info
