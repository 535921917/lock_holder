package zk_locker

import (
	"github.com/prometheus/common/log"
	"shuai/lock_holder/base"
	"sync"
	"time"
)

var wg sync.WaitGroup

func TestZkLock() {
	wg = sync.WaitGroup{}

	num := 10
	wg.Add(num)
	for i := 0; i < num; i++ {
		go Mock(Key)
	}
	wg.Wait()
}

func Mock(uniqueKey string) {
	defer wg.Done()
	goId := base.GoID()

	if err := Lock(uniqueKey); err != nil {
		log.Infof("[%v] the lock was in use", goId)
		return
	}

	log.Infof("[%v] get the lock.., doing his work", goId)
	time.Sleep(time.Second * 2)
	log.Infof("[%v] was done", goId)

	if err := UnLock(uniqueKey); err != nil {
		log.Infof("[%v] UnLock the lock err=%v", goId, err)
	}
}
