package bucket

const (
	LOCKTYPE_CAS   = 1
	LOCKTYPE_MUTEX = 2
)

type TokenBucket interface {
	GetToken() bool
	AddToken()
}
