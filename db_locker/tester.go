package db_locker

import (
	"github.com/prometheus/common/log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func TestDBLock() {
	wg = sync.WaitGroup{}

	num := 10
	wg.Add(num)
	for i := 0; i < num; i++ {
		time.Sleep(time.Second)
		go Mock("user_id_111", i)
	}
	wg.Wait()
}

func Mock(uniqueKey string, goId int) {
	defer wg.Done()
	//goId := base.GoID()

	if err := TryLock(uniqueKey, 1500); err != nil {
		log.Infof("[%v] the lock was in use", goId)
		return
	}

	log.Infof("[%v] get the lock.., doing his work", goId)
	time.Sleep(time.Second * 4)
	log.Infof("[%v] was done", goId)

	if err := UnLock(uniqueKey); err != nil {
		log.Errorf("[%v] UnLock the lock err=%v", goId, err)
	}
}
