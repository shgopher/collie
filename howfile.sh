#!/bin/bash
###
 # @Author: shgopher shgopher@gmail.com
 # @Date: 2023-12-07 16:09:55
 # @LastEditors: shgopher shgopher@gmail.com
 # @LastEditTime: 2023-12-07 16:09:57
 # @FilePath: /collie/howfile.sh
 # @Description: 
 #   查看本计算机中到底有多少文件句柄
 # Copyright (c) 2023 by shgopher, All Rights Reserved. 
### 

while true; do
    
    count=$(lsof | wc -l)
    
    date
    echo "Open files count: $count"
    
    sleep 1
    
done