package zk_locker

type ZkNodeType int32

//flags 的四种取值
const (
	PermanentNode ZkNodeType = iota
	TmpNode
	PermanentOrderNode
	TmpOrderNode
)

const ZkLockRootPath = "/zk_locker"

const Key = "user_id"