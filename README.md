# Go-tds

![Go](https://github.com/sdttttt/go-tds/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/sdttttt/go-tds)](https://goreportcard.com/report/github.com/sdttttt/go-tds)
[![codebeat badge](https://codebeat.co/badges/9040bc68-655c-4d3e-be12-661554bacecf)](https://codebeat.co/projects/github-com-sdttttt-go-tds-master)
[![codecov](https://codecov.io/gh/sdttttt/go-tds/branch/master/graph/badge.svg)](https://codecov.io/gh/sdttttt/go-tds) [![Join the chat at https://gitter.im/go-tds/community](https://badges.gitter.im/go-tds/community.svg)](https://gitter.im/go-tds/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

一个基于RPC的最小微服务库, 包含服务注册中心.
使用Go编写. 总代码不超过2000行. 它只实现了微服务最核心的功能。

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
