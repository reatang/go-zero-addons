# go-zero-addons
go-zero 附加功能

本库存放的是一些可以更好的使用 [go-zero](https://github.com/zeromicro/go-zero) 框架的额外功能

## rest 服务

### 前缀路由 ahttpx.NewPrefixPriorityRouter()

用法：
```go
package main

import (
	"github.com/zeromicro/go-zero/rest/router"
	"github.com/zeromicro/go-zero/rest"
	"github.com/reatang/go-zero-addons/ahttpx"
)


func main() {
    // ...
	
	server := rest.MustNewServer(c.RestConf, rest.WithRouter(ahttpx.NewPrefixPriorityRouter(router.NewRouter())))
	defer server.Stop()
	
	// ...
}

```

## grpc 服务

### grpc支持json格式

用法：
```go
package main

import (
    "github.com/reatang/go-zero-addons/agrpcx/codecx"
)


func main() {
	// 直接注册，客户端可通过 http2协议用json格式调用grpc接口
	// Context-Type: application/json
    codecx.RegisterCodecJson()
}
```
