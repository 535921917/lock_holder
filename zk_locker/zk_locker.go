package zk_locker

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/samuel/go-zookeeper/zk"
	"sort"
	"strings"
)

func Lock(uniqueKey string) error {
	//check 父目录是否已经存在
	if err := CheckRootPath(uniqueKey); err != nil {
		log.Errorf("check zk root path err:%v", err)
		return err
	}

	//创建临时顺序节点 return 形如 /zk_locker/uid0000000013
	myLockPath, err := conn.Create(FormatLockChildPath(Key), []byte(""), int32(TmpOrderNode), zk.WorldACL(zk.PermAll))
	if err != nil {
		return err
	}

	//获取父目录下最小节点
	children, _, err := conn.Children(FormatLockRootPath(Key))
	sort.Strings(children)
	if err != nil || len(children) == 0 {
		return err
	}

	//当前节点是最小节点 则加锁成功
	if strings.EqualFold(FormatLockRootPath(uniqueKey)+"/"+children[0], myLockPath) {
		return nil
	}

	//当前不是最小节点 ，寻找当前节点的上一个节点
	var watchNode string
	for i := len(children) - 1; i >= 0; i-- {
		childPath := FormatLockRootPath(uniqueKey) + "/" + children[i]
		if strings.Compare(childPath, myLockPath) < 0 {
			watchNode = childPath
			break
		}
	}

	//监听当前节点的前一个节点
	if watchNode != "" {
		_, _, ch, err := conn.ExistsW(watchNode)
		if err != nil {
			return err
		}
		event := <-ch
		if event.Type == zk.EventNodeDeleted {
			return nil
		}
	}
	return fmt.Errorf("无watchNode")
}

func UnLock(uniqueKey string) error {
	children, _, err := conn.Children(FormatLockRootPath(uniqueKey))
	sort.Strings(children)
	if err != nil || len(children) == 0 {
		return err
	}
	//删除最小节点
	if err := conn.Delete(FormatLockRootPath(uniqueKey)+"/"+children[0], -1); err != nil {
		return err
	}
	return nil
}

func CheckRootPath(uniqueKey string) error {
	exists, _, err := conn.Exists(FormatLockRootPath(uniqueKey))
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	if exists, _, _ := conn.Exists(ZkLockRootPath); !exists {
		_, err = conn.Create(ZkLockRootPath, []byte("zk_root_path"), int32(PermanentNode), zk.WorldACL(zk.PermAll))
	}
	if exists, _, _ := conn.Exists(FormatLockRootPath(uniqueKey)); !exists {
		_, err = conn.Create(FormatLockRootPath(uniqueKey), []byte(""), int32(PermanentNode), zk.WorldACL(zk.PermAll))
	}
	return err
}

func FormatLockChildPath(uniqueKey string) string {
	return fmt.Sprintf("%s/%s/lock", ZkLockRootPath, uniqueKey)
}
func FormatLockRootPath(uniqueKey string) string {
	return fmt.Sprintf("%s/%s", ZkLockRootPath, uniqueKey)
}
