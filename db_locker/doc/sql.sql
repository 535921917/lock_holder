CREATE TABLE `global_locks` (
                                `id` int(11) NOT NULL AUTO_INCREMENT,
                                `lock_key` varchar(60) NOT NULL COMMENT '锁名称',
                                `create_time` datetime NOT NULL COMMENT '创建时间',
                                PRIMARY KEY (`id`),
                                UNIQUE KEY `lockKey` (`lock_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='全局锁';
