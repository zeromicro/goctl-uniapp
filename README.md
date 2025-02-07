# goctl-uniapp

![go-zero](https://img.shields.io/badge/Github-go--zero-brightgreen?link=https://github.com/zeromicro/go-zero&logo=github)
![License](https://img.shields.io/badge/License-MIT-blue?link=https://github.com/zeromicro/goctl-android/blob/main/LICENSE)
![Go](https://github.com/zeromicro/goctl-android/workflows/Go/badge.svg)

goctl-uniapp 是一款基于 goctl 的插件，用于生成 uniapp 使用的 http 客户端。

```bash
# 安装 goctl-uniapp 插件
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/goctl-uniapp
```

```bash
# 生成代码
goctl api plugin -plugin goctl-uniapp="uniapp" -api demo.api -dir .
```