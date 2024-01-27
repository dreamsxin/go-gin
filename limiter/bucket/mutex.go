package bucket

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type TokenBucketMutex struct {
	//token []int
	limit  int32
	num    int32
	ticker *time.Ticker
	mu     *sync.Mutex
}

var (
	bucket     *TokenBucketMutex
	singleLock sync.Mutex
)

func GetBucketMutex(limit, num int32, ticker *time.Ticker) *TokenBucketMutex {
	if bucket != nil {
		return bucket
	}
	singleLock.Lock()
	defer singleLock.Unlock()
	if bucket == nil {
		if limit < num {
			limit = num
		}
		bucket = &TokenBucketMutex{
			limit:  limit,
			num:    num,
			ticker: ticker,
			mu:     &sync.Mutex{},
		}

		// 开始定时向令牌桶中添加token
		go func() {
			for {
				select {
				case <-bucket.ticker.C:
					bucket.AddToken()
				}
			}
		}()
	}
	return bucket
}

func (bucket *TokenBucketMutex) GetToken(c *gin.Context) bool {
	if bucket.num > 0 {
		bucket.mu.Lock()
		defer bucket.mu.Unlock()
		if bucket.num > 0 {
			bucket.num--
			return true
		}
	}
	return false
}

func (bucket *TokenBucketMutex) AddToken() {
	if bucket.num < bucket.limit {
		bucket.mu.Lock()
		defer bucket.mu.Unlock()
		if bucket.num < bucket.limit {
			bucket.num++
			return
		}
	}
}
