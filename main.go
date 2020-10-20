package main

import (
	"github.com/prometheus/common/log"
	"shuai/lock_holder/zk_locker"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	zk_locker.Init()
	wg = sync.WaitGroup{}
}

func main() {
	testZkLock()
}

func testZkLock() {
	num := 10
	wg.Add(num)
	for i := 0; i < num; i++ {
		go mock(zk_locker.Key, i)
	}
	wg.Wait()
}

func mock(uniqueKey string, id int) {
	defer wg.Done()
	if err := zk_locker.Lock(uniqueKey); err != nil {
		log.Infof("[%v] the lock was in use", id)
		return
	}

	log.Infof("[%v] get the lock.., doing his work", id)
	time.Sleep(time.Second * 2)
	log.Infof("[%v] was done", id)

	if err := zk_locker.UnLock(uniqueKey); err != nil {
		log.Infof("[%v] UnLock the lock err=%v", id, err)
	}
}
