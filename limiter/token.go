package limiter

import (
	"net/http"
	"time"

	"github.com/dreamsxin/go-gin/limiter/bucket"
	"github.com/gin-gonic/gin"
)

func TokenBucketLimiter(buckLockType int, limit, num int32, dur time.Duration) gin.HandlerFunc {

	var b bucket.TokenBucket
	switch buckLockType {
	case bucket.LOCKTYPE_CAS:
		b = bucket.GetBucketCAS(limit, num, time.NewTicker(dur))
	case bucket.LOCKTYPE_MUTEX:
		b = bucket.GetBucketMutex(limit, num, time.NewTicker(dur))
	}
	return func(context *gin.Context) {
		if ok := b.GetToken(); !ok {
			context.Status(http.StatusServiceUnavailable)
			context.Abort()
		} else {
			context.Next()
		}
	}
}
