package main

import (
	"fmt"
	"time"

	"github.com/dreamsxin/go-gin/limiter"
	"github.com/dreamsxin/go-gin/limiter/bucket"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//r.Use(gin.Recovery(), gin.Logger())
	//r.Use(limiter.TokenBucketLimiter(bucket.LOCKTYPE_CAS, 100, 100, time.Millisecond))
	//r.Use(limiter.TokenBucketLimiter(bucket.LOCKTYPE_MUTEX, 100, 1, time.Millisecond))
	r.Use(limiter.TokenBucketLimiter(bucket.LOCKTYPE_IP, 1, 1, time.Millisecond))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run("localhost:8080")
	if err != nil {
		fmt.Printf("run err %v\n", err)
		return
	}
}
