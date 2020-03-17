# Go-tds (Development stage)

![Go](https://github.com/sdttttt/go-tds/workflows/Go/badge.svg)
[![codebeat badge](https://codebeat.co/badges/9040bc68-655c-4d3e-be12-661554bacecf)](https://codebeat.co/projects/github-com-sdttttt-go-tds-master)
[![codecov](https://codecov.io/gh/sdttttt/go-tds/branch/master/graph/badge.svg)](https://codecov.io/gh/sdttttt/go-tds)

一个基于RPC的分布式微服务框架,以及服务注册中心。

它是简单的, 用户主要专注于开发RPC服务,不需要关心其他事情.只需要少量的配置,它就能很快开始工作.

## Example

目前还处于开发阶段.不适用于生产环境.

## Why do need go-tds?

也许这是我个人的想法，我非常讨厌在学习一门（语言/框架）时需要花费巨额的时间成本。
他们真的太重了！

> Java学习的框架过程可以说是在套娃。`... -> Spring -> spring MVC -> ...`

对于新手来说，这个过程无疑是痛苦的。

一切都是为了尽快！
尽快开始这一切！为此才出现了`go-tds`,
它能很快的帮你构建完成您的分布式应用！
`go-tds`的架构非常简单！你能很快的分析他的源码！
并把它变成您自己的东西。技术一定要用自己最熟悉的。

我衷心的希望编程是一种享受，不是一个痛苦的过程。`By SDTTTTT`

## Architecture

它的架构和`Dubbo`类似,简单易用，首先有3个名词你需要知道，
如果你学习过微服务，那就更简单了！


**Hub:** 
它的工作比后两个稍微复杂一些，当`Provider`调用`Register`时，
会RPC远程调用`Hub`的`JoinServiceHub`方法，`Hub`只收集这些服务的地址以及服务名，
不关心具体细节。`Customer`可以在`Hub`中找到现有的服务名，并返回他们的地址信息，
`Hub`可以有多个同样的服务，便于做负载均衡。

**Provider:** 
提供商很简单，使用Golang原生的RPC，编写您的业务，
然后调用`Register`注册到Hub就行了！

**Customer**
消费方只需要和以前一样！调用您需要的RPC服务.
不用担心，其他我们都帮您完成了。

## Future

- [x] 注册服务
- [ ] 服务消费
- [ ] 负载均衡
