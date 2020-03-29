# Go-tds (Development stage)

![Go](https://github.com/sdttttt/go-tds/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/sdttttt/go-tds)](https://goreportcard.com/report/github.com/sdttttt/go-tds)
[![codebeat badge](https://codebeat.co/badges/9040bc68-655c-4d3e-be12-661554bacecf)](https://codebeat.co/projects/github-com-sdttttt-go-tds-master)
[![codecov](https://codecov.io/gh/sdttttt/go-tds/branch/master/graph/badge.svg)](https://codecov.io/gh/sdttttt/go-tds) [![Join the chat at https://gitter.im/go-tds/community](https://badges.gitter.im/go-tds/community.svg)](https://gitter.im/go-tds/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

一个基于RPC的轻量级分布式微服务框架(库), 以及服务注册中心.
使用Go编写.

它是简单的, 用户可以专注于开发服务,不需要关心其他事情.
只需要少量的配置,它就能很快开始工作.

## Why do need go-tds?

或许您和我一样, 在学习大型框架时需要花费大量的时间去理解他们,
`go-tds`就是为此诞生的, 它非常简单, 容易上手, 
我们希望任何人都可以快速着手开发自己的分布式应用.

## Architecture

它的架构和`Dubbo`类似,简单易用，首先有3个名词你需要知道，
如果你学习过一些别的微服务框架（比如Spring Cloud Netflix），那就更简单了！

**Hub** 

它是独立出来的注册中心, 通信采用`gRPC`, 
每一个注册进入`Hub`的服务会在这里登记, 当有消费者来使用服务, 
Hub仅仅只是将服务的地址给它们, 不会做其他任何事情. 
这意味着在服务通信方面,您可以采用**任何网络协议**作为通信媒介.
您还可以完全自定义Hub的负载均衡算法.

**Provider** 

您可以使用任意网络协议作为您的通信方式,
然后将一些服务信息注册到Hub就可以开始工作了.
我们喜欢简单的设计!

**Customer**

如果您的服务使用的是Go提供的原生RPC编写的, 那太好了!
我们提供了`trpc`库, 可以在消费者端以很简单的方式,调用您的RPC服务,
不用担心，其他我们都帮您完成了!


# Example

> **Warning: 目前还处于开发阶段.不适用于生产环境.**

### Start Hub

首先您需要在启动hub.请`clone`本仓库的代码.

```shell
git clone https://github.com/sdttttt/go-tds.git
```

然后编译,运行它.

```
go build -v -o hub
./hub
```

当然, 您可以自定义你的Hub配置文件.修改目录下的`tclient.yaml`中的配置即可.

```yaml
hub:
  address: localhost
  port: 1234 // 默认在1234端口
  checkSurvivalTime: 120 // 检查服务生存间隔时间 (Unit: Seconds)
```

### Provider Service

您需要在您的目录下配置`tclient.yaml`文件.可以在主函数的当前目录,也可以在上一级.
如果您需要改变您的配置文件路径,可以使用`configuration.ChangeConfigFilePath`.

```yaml
hub:
  address: localhost
  port: 1234

self:
  address: localhost
  port: 5555
  survivalTime: 45 // 设置服务的心跳间隔, 这个值得小于Hub上的服务检查时间间隔
```

编写您的服务, 这里使用golang的RPC作为示范.

功能很简单,将第一次参数值作为返回值.

```go
package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/sdttttt/go-tds/trpc"
)

type API struct{}

func (a *API) Hello(in int, out *int) error {
	*out = 100
	return nil
}

func main() {
  // 注册填入您的服务名即可
	trpc.Register("API.Hello")
	api := new(API)
  
  // 在这里注册它
  rpc.Register(api)

	l, err := net.Listen("tcp", ":4321")
	if err != nil {
		log.Println(err)
		return
	}

	rpc.HandleHTTP()

	go http.Serve(l, nil)

	select {}
}

```

启动它!

### Customer

服务消费端需要知道Hub的地址.

```yaml
hub:
  address: localhost
  port: 1234
```


```go
package main

import (
	"log"

	"github.com/sdttttt/go-tds/trpc"
)

func main() {
	var one int = 1
	var two int
  
  // 服务端如果使用golang的RPC
  // go-tds提供了简单便利的库, trpc
	err := trpc.Call("API.Hello", one, &two)

	if err != nil {
		log.Fatalln(err)
	}

	println(two)
}

```

启动它! 您会在终端中看到

```shell
localhost : 4321
API.Hello (0x88c720,0xc000136030) (0x8770e0,0xc000136008)
1
```

## Future

未来要做的事情

- [x] 注册服务
- [x] 服务消费
- [x] 负载均衡
- [x] 服务治理
- [ ] 多中心化