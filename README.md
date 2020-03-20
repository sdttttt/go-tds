# Go-tds (Development stage)

![Go](https://github.com/sdttttt/go-tds/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/sdttttt/go-tds)](https://goreportcard.com/report/github.com/sdttttt/go-tds)
[![codebeat badge](https://codebeat.co/badges/9040bc68-655c-4d3e-be12-661554bacecf)](https://codebeat.co/projects/github-com-sdttttt-go-tds-master)
[![codecov](https://codecov.io/gh/sdttttt/go-tds/branch/master/graph/badge.svg)](https://codecov.io/gh/sdttttt/go-tds)

一个基于RPC的轻量级分布式微服务框架(库).
以及服务注册中心.

它是简单的, 用户主要专注于开发原生态的RPC服务,不需要关心其他事情.只需要少量的配置,它就能很快开始工作.

## Example

目前还处于开发阶段.不适用于生产环境.

## Why do need go-tds?

或许您和我一样, 在学习大型框架时需要花费大量的时间去理解他们.
`go-tds`就是为此诞生的, 我们希望任何人都可以快速上手微服务.

能帮助您快速构建属于您自己的分布式应用.

## Architecture

它的架构和`Dubbo`类似,简单易用，首先有3个名词你需要知道，
如果你学习过微服务，那就更简单了！

**Hub** 

它是独立出来的注册中心, 通信采用`gRPC`, 
每一个注册进入`Hub`的服务会在这里登记, 当有消费者来使用服务, 
Hub仅仅只是将服务的地址给它们, 不会做其他任何事情. 
在这中间可以使用负载均衡.来加强您的应用.

**Provider** 

您可以使用Go提供的原生RPC编写您的业务,
也可以使用其他任何基于socket通信的服务,
然后将服务的一些信息注册到Hub就行了！
我们喜欢简单的设计!

**Customer**

如果您的服务使用的是Go提供的原生RPC编写的, 那太好了!
我们提供了`trpc`库, 可以帮助你将您的服务,
无痛式的转移到`go-tds`中去.
不用担心，其他我们都帮您完成了!

## Future

- [x] 注册服务
- [x] 服务消费
- [ ] 负载均衡
- [ ] 服务治理
- [ ] 多中心化
