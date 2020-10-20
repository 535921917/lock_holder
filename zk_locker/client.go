package zk_locker

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var (
	conn *zk.Conn
)

func Init() {
	hosts := []string{"10.227.27.232:2181"}
	var err error
	conn, _, err = zk.Connect(hosts, time.Second*5)
	if err != nil {
		log.Errorf("init zk conn err=%v", err)
		panic("init zk conn err")
	}
}

func Create() {
	path, err := conn.Create("/zk/holder", []byte("test_value"), int32(PermanentOrderNode), zk.WorldACL(zk.PermAll))
	fmt.Printf("path=%v,err=%v\n", path, err)
}
func Query() {
	data, stat, err := conn.Get("/zk/holder0000000016")
	fmt.Printf("data=%v,stat=%+v,err=%v\n", string(data), stat, err)
}

func Exists() {
	exists, stat, err := conn.Exists("/zk/holder")
	fmt.Printf("exists=%v,stat=%+v,err=%v\n", exists, stat, err)
}

func Delete() {
	_, stat, _ := conn.Get("/zk/holder0000000016")
	_ = conn.Delete("/zk/holder0000000016", stat.Version)
}

func Watch() {
	_, _, ch, _ := conn.ExistsW("/zk/holder0000000017")
	event := <-ch
	if event.Type == zk.EventNodeDeleted {
		fmt.Printf("delete u haha")
	}
}
func Children() {
	children, stat, err := conn.Children("/zk")
	fmt.Printf("children=%v,stat=%+v,err=%v\n", children, stat, err)
}
