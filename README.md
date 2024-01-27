# go-gin

存放 gin 相关的增强工具

## 安装

```shell
go get -u github.com/dreamsxin/go-gin
```

## 缓存

### 使用本地缓存
```go
package main

import (
	"time"

	"github.com/dreamsxin/go-gin/cache"
	"github.com/dreamsxin/go-gin/cache/persist"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	app.GET("/hello",
		cache.CacheByRequestURI(memoryStore, 2*time.Second),
		func(c *gin.Context) {
			c.String(200, "hello world")
		},
	)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
```

### 使用redis作为缓存
```go
package main

import (
	"time"

	"github.com/dreamsxin/go-gin/cache"
	"github.com/dreamsxin/go-gin/cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	app := gin.New()

	redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	}))

	app.GET("/hello",
		cache.CacheByRequestURI(redisStore, 2*time.Second),
		func(c *gin.Context) {
			c.String(200, "hello world")
		},
	)
	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
```

# 压测
```
wrk -c 500 -d 1m -t 5 http://127.0.0.1:8080/hello
```
