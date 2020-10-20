#### 原理
1. 我们将锁抽象成目录，多个线程在此目录下创建瞬时的顺序节点，因为Zk会为我们保证节点的顺序性，所以可以利用节点的顺序进行锁的判断。
2. 首先创建顺序节点，然后获取当前目录下最小的节点，判断最小节点是不是当前节点，如果是那么获取锁成功，如果不是那么获取锁失败。
3. 获取锁失败的节点获取当前节点上一个顺序节点，对此节点注册监听，当节点删除的时候通知当前节点。
4. 当unlock的时候删除节点之后会通知下一个节点


#### 节点类型
- 0:永久，除非手动删除
- 1:短暂，session断开则改节点也被删除
- 2:会自动在节点后面添加序号
- 3:即，短暂且自动添加序号


#### Watch事件类型
1. EventNodeCreated：节点创建事件，需要watch一个不存在的节点，当节点被创建时触发，此watch通过conn.ExistsW(path string)设置
2. EventNodeDeleted：节点删除事件，需要watch一个已存在的节点，当节点被移除时触发，此watch通过conn.ExistsW(path string)设置
3. EventNodeDataChanged：节点数据变化事件，此watch通过conn.GetW(path string) 以及 conn.ExistsW(path string) 设置，
4. EventNodeChildrenChanged：子节点改变事件（数量改变），此watch通过conn.ChildrenW(path string)设置， 当path 下面增删子节点时触发（修改path下的子节点的内容时，不会触发通知）。
5. EventNoWatching：watch移除事件，服务端出于某些原因不再为客户端watch节点时触发