package bucket

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Client struct {
	IP       string
	LastTime int64
	Num      int32
	Mtx      sync.RWMutex
}

type TokenBucketIP struct {
	limit    int32 // 总 ip 数
	num      int32 // 单个 ip 连接数
	LastTime int64
	clientIP map[string]*Client
	ticker   *time.Ticker
}

var (
	bucketIP     *TokenBucketIP
	singleLockIP sync.Mutex
)

func GetBucketIP(limit, num int32, ticker *time.Ticker) *TokenBucketIP {
	if bucketIP != nil {
		return bucketIP
	}
	singleLockIP.Lock()
	defer singleLockIP.Unlock()
	if bucketIP == nil {
		bucketIP = &TokenBucketIP{
			limit:    limit,
			num:      num,
			ticker:   ticker,
			clientIP: make(map[string]*Client),
		}
		go func() {
			for {
				select {
				case <-bucketIP.ticker.C:
					bucketIP.AddToken()
				}
			}
		}()
	}
	return bucketIP
}

func (bucketIP *TokenBucketIP) AddToken() {
	bucketIP.LastTime = time.Now().Unix()
	bucketIP.clientIP = make(map[string]*Client)
}

func (bucketIP *TokenBucketIP) GetToken(c *gin.Context) bool {

	if len(bucketIP.clientIP) >= int(bucketIP.limit) {
		return false
	}
	singleLockIP.Lock()
	ip := c.ClientIP()
	client, ok := bucketIP.clientIP[ip]
	if !ok {
		client = &Client{IP: ip, LastTime: bucketIP.LastTime}
		bucketIP.clientIP[ip] = client
	}
	singleLockIP.Unlock()

	client.Mtx.Lock()
	defer client.Mtx.Unlock()
	if client.Num >= bucketIP.num {
		return false
	}
	client.Num++
	return true
}
