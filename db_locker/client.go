package db_locker

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"shuai/lock_holder/base"
	"time"
)

var (
	db *gorm.DB
)

func Init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@/offline_data")
	if err != nil {
		panic(err)
	}
}

func GetConn() *gorm.DB {
	db := db.New()
	//db.LogMode(true)
	return db
}

func InsertGlobalLockWithTest(lockKey string) int64 {
	conn := GetConn()
	sql := fmt.Sprintf(SqlTemplate, lockKey, base.GetMilliTimeStamp(time.Now()), lockKey)
	res := conn.Exec(sql)
	if res.Error != nil {
		log.Errorf("InsertGlobalLockWithTest,res.err = %v", res.Error)
	}
	return res.RowsAffected
}

func DeleteGlobalLockById(globalLock *GlobalLock) {
	conn := GetConn()
	err := conn.Delete(globalLock).Error
	if err != nil {
		log.Errorf("DeleteGlobalLockById delete err =%v", err)
	}
}
func DeleteGlobalLockByLockKey(lockKey string) error {
	conn := GetConn()
	err := conn.Model(&GlobalLock{}).Delete(&GlobalLock{
		LockKey: lockKey,
	}).Error
	if err != nil {
		log.Errorf("DeleteGlobalLockByLockKey delete err =%v", err)
	}
	return err
}

func SelectGlobalLockByLockKey(lockKey string) *GlobalLock {
	conn := GetConn()
	lock := &GlobalLock{}
	err := conn.Model(&GlobalLock{}).Where("lock_key = ?", lockKey).First(lock).Error
	if err != nil {
		log.Errorf("SelectGlobalLockByLockKey query err =%v", err)
	}
	return lock
}
