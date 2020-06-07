<p align="center">
  <a href="https://github.com/googege/collie">
    <img width="20%" alt="github.com/googege/collie" src="./logo.jpg">
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

### Features
- 🚄 **Fast**: leveraging go's high concurrency capabilities, this tool has unparalleled speed.

- 🚲 **Easy**: you just need install it to the terminal and you can get started quickly.
### How to use
#### use go
```go
collie -r -o - w -q
```
- `-r` the resource photos dir path
- `-o` the output dir path
- `-w` the result photos' width ,or you can set 0,you will get the same width with the old one.
- `-q` the quality of the output photos from 20 to 100.

like:

```go
// add "gopath/bin" to PATH. 
go get -u github.com/googege/collie 
//
collie -r /Users/googege/Desktop/test -o /Users/googege/Desktop/app -w 0 -q 30
```
#### use binary file
[download binary file](https://github.com/shgopher/collie/releases)
- when you use windows you should open windows command and use ` collie.exe -r [your resource path] -o [your out put path] -q [the quality]`

- in mac or  linux open your bash or zsh command, use  `collie_[linux/macos] -r resource path -o output path`

suport png jpg jpeg and gif ,and all of the output is jpeg
## HERE

|项目|介绍|
|:---:|:---:|
|对我的赞助|![p](https://raw.githubusercontent.com/basicExploration/Demos/master/donate.png)|
|便宜服务器推荐|[阿里云](https://www.aliyun.com/minisite/goods?userCode=ol87kpmz)，[梯子服务器](https://app.cloudcone.com/?ref=2525):支持支付宝|
|微信公众号|![p](https://raw.githubusercontent.com/basicExploration/Demos/master/pluspro.png)|
|知识讨论微信群|![p](https://raw.githubusercontent.com/basicExploration/Demos/master/joinMyGroup.png)|
|我的社交平台|[b站](https://space.bilibili.com/23170151)，[YouTube](https://www.youtube.com/channel/UCM_-pFgD_HZDGD0yxfzguRQ?view_as=subscriber)，[微博](https://weibo.com/imgoogege)，抖音：googege|

## Stargazers over time

[![Stargazers over time](https://starchart.cc/googege/collie.svg)](https://starchart.cc/googege/collie)






