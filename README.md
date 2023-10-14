# satori-sdk-go
go语言开发的satori协议的sdk客户端

[![](https://img.shields.io/github/license/dezhishen/satori-sdk-go.svg?style=for-the-badge&logo=github)](./LICENSE)
[![](https://img.shields.io/github/stars/dezhishen/satori-sdk-go.svg?style=for-the-badge&logo=github)](https://github.com/dezhishen/satori-sdk-go/stargazers)
[![](https://img.shields.io/github/forks/dezhishen/satori-sdk-go.svg?style=for-the-badge&logo=github)](https://github.com/dezhishen/satori-sdk-go/network/members)
[![](https://img.shields.io/github/contributors/dezhishen/satori-sdk-go.svg?style=for-the-badge&logo=github)](https://github.com/dezhishen/satori-sdk-go/graphs/contributors)

[![](https://img.shields.io/github/commit-activity/m/dezhishen/satori-sdk-go?logo=github&style=for-the-badge)](https://github.com/dezhishen/satori-sdk-go/graphs/commit-activity)
[![](https://img.shields.io/github/last-commit/dezhishen/satori-sdk-go.svg?style=for-the-badge&logo=github)](https://github.com/dezhishen/satori-sdk-go/commits)
## todo
- [ ] 示例
- [ ] 配置
- [ ] 客户端创建
  - [x] http
  - [x] websocket
  - [x] Webhook
- [ ] 资源
  - [ ] Channel
    - [ ] API
    - [ ] Event
    - [x] Model
  - [ ] Guild
    - [ ] API
    - [ ] Event
    - [x] Model
  - [ ] GuildMember
    - [ ] API
    - [ ] Event
    - [ ] Model
  - [ ] GuildRole
    - [ ] API
    - [ ] Event
    - [ ] Model
  - [ ] Login
    - [ ] API
    - [ ] Event
    - [ ] Model
  - [ ] Message
    - [ ] API
    - [ ] Event
    - [ ] Model
  - [ ] Reaction
    - [ ] API
    - [ ] Event
    - [ ] Model
  - [ ] User
    - [ ] API
    - [ ] Event
    - [ ] Model
- [ ] API
- [ ] EVENT

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