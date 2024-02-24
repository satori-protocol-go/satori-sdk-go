# satori-sdk-go
go语言开发的[satori](https://satori.chat/zh-CN)协议的sdk客户端

[![](https://img.shields.io/github/license/satori-protocol-go/satori-sdk-go.svg?logo=github)](./LICENSE)
[![](https://img.shields.io/github/stars/satori-protocol-go/satori-sdk-go.svg?logo=github)](https://github.com/satori-protocol-go/satori-sdk-go/stargazers)
[![](https://img.shields.io/github/forks/satori-protocol-go/satori-sdk-go.svg?logo=github)](https://github.com/satori-protocol-go/satori-sdk-go/network/members)
[![](https://img.shields.io/github/contributors/satori-protocol-go/satori-sdk-go.svg?logo=github)](https://github.com/satori-protocol-go/satori-sdk-go/graphs/contributors)

[![](https://img.shields.io/github/commit-activity/m/satori-protocol-go/satori-sdk-go?logo=github)](https://github.com/satori-protocol-go/satori-sdk-go/graphs/commit-activity)
[![](https://img.shields.io/github/last-commit/satori-protocol-go/satori-sdk-go.svg?logo=github)](https://github.com/satori-protocol-go/satori-sdk-go/commits)
[![wakatime](https://wakatime.com/badge/user/a2c981ca-317d-4b34-8ed9-4264fbfdb775/project/018b2817-27cb-454d-b957-5a4686855dcd.svg)](https://wakatime.com/badge/user/a2c981ca-317d-4b34-8ed9-4264fbfdb775/project/018b2817-27cb-454d-b957-5a4686855dcd)
## 架构图
![](doc/images/架构图.png)

## todo
- [x] 示例
- [x] 配置
- [ ] 日志标准化
- [x] 客户端创建
  - [x] API
    - [x] http
  - [x] EVENT
    - [x] websocket
    - [x] webhook
- 资源
  - [x] Channel
    - [x] API
    - [x] Event
    - [x] Model
  - [x] Guild
    - [x] API
    - [x] Event
    - [x] Model
  - [x] GuildMember
    - [x] API
    - [x] Event
    - [x] Model
  - [x] GuildRole
    - [x] API
    - [x] Event
    - [x] Model
  - [x] Login
    - [x] API
    - [x] Event
    - [x] Model
  - [x] Message
    - [x] API
    - [x] Event
    - [x] ~~Model~~ 由[satori-protocol-go/satori-model-go](https://github.com/satori-protocol-go/satori-model-go)提供
    - [x] ~~[元素]构建器和解析器~~ 由[satori-protocol-go/satori-model-go](https://github.com/satori-protocol-go/satori-model-go)提供
  - [x] Reaction
    - [x] API
    - [x] Event
    - [x] Model
  - [x] User
    - [x] API
    - [x] Event
    - [x] Model
## 示例
[examples/main.go](./examples/main.go)
## 目录
```
- examples # 使用示例
- pkg # sdk暴露给外部的能力
- - api # api
- - client # 客户端,包含api和event的多种方式客户端
- - event # 事件
- - resource # 资源
- - - channel
- - - guild
- - - ...
```
