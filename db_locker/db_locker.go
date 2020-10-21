package db_locker

import (
	"fmt"
	"github.com/prometheus/common/log"
	"shuai/lock_holder/base"
	"time"
)

func Lock(lockKey string) error {
	rows := InsertGlobalLockWithTest(lockKey)
	if rows == 1 {
		log.Debugf("locking success,lockKey =%v", lockKey)
		return nil
	}
	return fmt.Errorf("加锁失败")
}

//带有超时时间的Lock
func TryLock(lockKey string, expireTime int64) error {
	rows := InsertGlobalLockWithTest(lockKey)
	if rows == 1 {
		log.Debugf("locking success,lockKey =%v", lockKey)
		return nil
	}
	oldLock := SelectGlobalLockByLockKey(lockKey)
	if base.GetMilliTimeStamp(time.Now())-oldLock.CreateTime < expireTime {
		return fmt.Errorf("加锁失败")
	}
	DeleteGlobalLockById(oldLock)
	return Lock(lockKey)
}

func UnLock(lockKey string) error {
	return DeleteGlobalLockByLockKey(lockKey)
}
