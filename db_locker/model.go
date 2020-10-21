package db_locker

type UserInfo struct {
	ID       int64  `gorm:"column:id"`
	UserID   int64  `gorm:"column:user_id"`
	UserName string `gorm:"column:user_name"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

type GlobalLock struct {
	ID         int64  `gorm:"column:id"`
	LockKey    string `gorm:"column:lock_key"`
	CreateTime int64  `gorm:"column:create_time"`
}

func (GlobalLock) TableName() string {
	return "global_locks"
}
