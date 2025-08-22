# Go-tds

![Go](https://github.com/sdttttt/go-tds/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/sdttttt/go-tds)](https://goreportcard.com/report/github.com/sdttttt/go-tds)
[![codebeat badge](https://codebeat.co/badges/9040bc68-655c-4d3e-be12-661554bacecf)](https://codebeat.co/projects/github-com-sdttttt-go-tds-master)
[![codecov](https://codecov.io/gh/sdttttt/go-tds/branch/master/graph/badge.svg)](https://codecov.io/gh/sdttttt/go-tds) [![Join the chat at https://gitter.im/go-tds/community](https://badges.gitter.im/go-tds/community.svg)](https://gitter.im/go-tds/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

一个微服务验证框架，使用go编写. 
实现最简单的服务注册，调度，调用等功能。

# Example


### Hub



```shell
git clone https://github.com/sdttttt/go-tds.git
```



```
go build -v -o hub
./hub
```

```yaml
hub:
  address: localhost
  port: 1234 // 默认在1234端口
  checkSurvivalTime: 120 // 检查服务生存间隔时间 (Unit: Seconds)
```

### Provider Service


```yaml
hub:
  address: localhost
  port: 1234

self:
  address: localhost
  port: 5555
  survivalTime: 45 // 设置服务的心跳间隔, 这个值得小于Hub上的服务检查时间间隔
```

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

### Customer


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


```shell
localhost : 4321
API.Hello (0x88c720,0xc000136030) (0x8770e0,0xc000136008)
1
```
