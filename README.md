# satori-sdk-go
go语言开发的[satori](https://satori.js.org/zh-CN/)协议的sdk客户端

[![](https://img.shields.io/github/license/dezhishen/satori-sdk-go.svg?style=for-the-badge&logo=github)](./LICENSE)
[![](https://img.shields.io/github/stars/dezhishen/satori-sdk-go.svg?style=for-the-badge&logo=github)](https://github.com/dezhishen/satori-sdk-go/stargazers)
[![](https://img.shields.io/github/forks/dezhishen/satori-sdk-go.svg?style=for-the-badge&logo=github)](https://github.com/dezhishen/satori-sdk-go/network/members)
[![](https://img.shields.io/github/contributors/dezhishen/satori-sdk-go.svg?style=for-the-badge&logo=github)](https://github.com/dezhishen/satori-sdk-go/graphs/contributors)

[![](https://img.shields.io/github/commit-activity/m/dezhishen/satori-sdk-go?logo=github&style=for-the-badge)](https://github.com/dezhishen/satori-sdk-go/graphs/commit-activity)
[![](https://img.shields.io/github/last-commit/dezhishen/satori-sdk-go.svg?style=for-the-badge&logo=github)](https://github.com/dezhishen/satori-sdk-go/commits)
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
    - [x] Model
    - [ ] [元素](https://satori.js.org/zh-CN/protocol/elements.html)构建器和解析器
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
