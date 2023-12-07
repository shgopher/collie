<p align="center">
  <a href="https://github.com/googege/collie">
    <img width="20%" alt="github.com/googege/collie" src="./collie.png">
  </a>
</p>
<p align="center">
<a href="https://www.bilibili.com/video/BV1za4y1e7U9/">视频演示</a> | 
<a href="https://www.youtube.com/watch?v=3nDGifbqdug">Video presentation</a>
</p>
<p align="center">
    Picture processing tools like compressing pictures, adding watermarks, etc ,it support batch processing.
</p>

<p align="center">
    <a href="https://travis-ci.com/googege/collie.svg?branch=master">
    <img src="https://travis-ci.com/googege/collie.svg?branch=master"/>
    </a>
  <a href="https://goreportcard.com/report/github.com/googege/collie"><img alt="collie" src="https://goreportcard.com/badge/github.com/googege/collie"></a>
  <a href="https://golang.org"><img alt="golang" src="https://img.shields.io/badge/awesome-golang-blue.svg"></a>
  <a href="https://pkg.go.dev/search?q=googege" rel="nofollow"><img src="https://camo.githubusercontent.com/a9a286d43bdfff9fb41b88b25b35ea8edd2634fc/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f646572656b7061726b65722f64656c76653f7374617475732e737667" alt="GoDoc" data-canonical-src="https://godoc.org/github.com/derekparker/delve?status.svg" style="max-width:100%;"></a>
  <a href="https://raw.githubusercontent.com/googege/collie/master/LICENSE" rel="nofollow"><img src="https://img.shields.io/badge/license-BSD 3 Clause-blue.svg" alt="license" data-canonical-src="https://img.shields.io/badge/license-BDS3-blue.svg" style="max-width:100%;"></a>
</p> 

<p align="center">
    Learn more: <a href="https://github.com/shgopher" target="_blank">shgopher</a>
</p>


### 下载地址
[download binary file](https://github.com/shgopher/collie/releases)

### 如何使用

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
## HERE

|项目|介绍|
|:---:|:---:|
|便宜服务器推荐|[阿里云](https://www.aliyun.com/minisite/goods?userCode=ol87kpmz)，[梯子服务器](https://app.cloudcone.com/?ref=2525):支持支付宝|
|微信公众号|科科人神|
|我的社交平台|[b站](https://space.bilibili.com/478621088)|

## Stargazers over time

[![Stargazers over time](https://starchart.cc/googege/collie.svg)](https://starchart.cc/googege/collie)






