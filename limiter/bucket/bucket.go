package bucket

import "github.com/gin-gonic/gin"

const (
	LOCKTYPE_CAS   = 1
	LOCKTYPE_MUTEX = 2
	LOCKTYPE_IP    = 3
)

type TokenBucket interface {
	GetToken(c *gin.Context) bool
	AddToken()
}
