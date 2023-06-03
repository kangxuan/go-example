package distributedLock

import "context"

type DistributedLock interface {
	Lock(ctx context.Context, uuid string) error
	UnLock(ctx context.Context, uuid string) error
}

type DistributedLockFactory struct {
}

func (f *DistributedLockFactory) NewDistributedLock(lockType string) DistributedLock {
	switch lockType {
	case "redis":
		return &RedisDistributedLock{}
	default:
		panic("还未实现")
	}
}
