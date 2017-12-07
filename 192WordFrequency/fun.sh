#!/bin/bash

#                           连续空格需要去掉                                                                         n参数：以数字排序
sed 's/\ /\n/g' words.txt | sed '/^$/d' | awk '{sum[$1]++;}END{for (word in sum) printf "%s %d\n",word,sum[word]}' | sort -k2nr
