# Collie 高性能图片压缩软件
## 赞助作者
- [cloudcone 服务器](https://app.cloudcone.com/?ref=2525) 主打一个性价比高，每月 3T 流量，带宽 1 GB/s，服务器地址美国洛杉矶，价格 5 刀/月。可支付宝付款！
- [just-my-socks 机场](https://justmysocks.net/members/aff.php?aff=29885
) 不怕 ip 被 ban，每月 500GB 流量，2.5 Gbps 带宽，支持 5 台设备，价格 5 刀/月。可支付宝付款！

## 下载地址
[download binary file](https://github.com/shgopher/collie/releases)

## 如何使用

-n int
是否输出跟源文件相同的名称，0：是，1：不是，如果不是，系统会给出一个随机代码，默认是输出相同的名字

-o string
输出的路径 (default ".")

-q int
输出的照片质量，范围是从1 - 100 (default 75)

-r string
指定的输入路径，路径是指的图片所处的文件夹，文件夹中还有文件夹不影响，系统会找到你指定文件夹中的所以照片，包括文件夹中的文件夹里的图片 (default "./test")

-w int
输出的照片尺寸，0是跟之前一样大，单位是px

> 如果权限不足，记得改权限 `chmod +x ./collie_macos`
- 用 windows 打开windows的命令行这样使用： ` collie.exe -r [your resource path] -o [your out put path] -q [the quality]`

- Mac和Linux： `collie_[linux/macos] -r resource path -o output path`

支持 png jpg jpeg and gif ,所有格式都按照jpeg来输出，所以GIF最好别用。因为输出的是第一张图

## Stargazers over time

[![Stargazers over time](https://starchart.cc/googege/collie.svg)](https://starchart.cc/googege/collie)






