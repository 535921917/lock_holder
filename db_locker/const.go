package db_locker

const SqlTemplate = "insert into `global_locks` (lock_key,create_time) select '%s','%d' from dual where not exists  (select 1 from global_locks where lock_key = '%s');"
