#!/bin/bash

#考虑行数小于10的情况
# wc -l 命令会输出文件名
[ `wc -l file.txt | awk '{print $1}'` -lt 10 ] && echo "" && exit
head -n10 file.txt | tail -n1
