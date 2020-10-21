package main

import (
	"shuai/lock_holder/db_locker"
	"shuai/lock_holder/zk_locker"
)

func init() {
	zk_locker.Init()
	db_locker.Init()
}

func main() {
	//zk_locker.TestZkLock()
	db_locker.TestDBLock()
}
