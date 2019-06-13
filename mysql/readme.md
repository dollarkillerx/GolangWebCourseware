Mysql主从同步
===
### 主节点配置文件
master.conf
``` 
[client]
default-character-set=utf8
[mysql]
default-character-set=utf8
[mysqld]
log_bin = log  #开启二进制日志，用于从节点的历史复制回放
collation-server = utf8_unicode_ci
init-connect='SET NAMES utf8'
character-set-server = utf8
server_id = 1  #需保证主库和从库的server_id不同， 假设主库设为1
replicate-do-db=fileserver  #需要复制的数据库名，需复制多个数据库的话则重复设置这个选项
```
### 从节点配置文件
slave.conf
``` 
[client]
default-character-set=utf8
[mysql]
default-character-set=utf8
[mysqld]
log_bin = log  #开启二进制日志，用于从节点的历史复制回放
collation-server = utf8_unicode_ci
init-connect='SET NAMES utf8'
character-set-server = utf8
server_id = 1  #需保证主库和从库的server_id不同， 假设主库设为1
replicate-do-db=fileserver  #需要复制的数据库名，需复制多个数据库的话则重复设置这个选项
```

### 操作指南
第一部在本地创建好配置文件
创建主容器持久化目录
创建从节点持久化目录
```
dollarkiller@worldlink:~/data/test/mysql$ ls
master.conf  slave.conf
dollarkiller@worldlink:~/data/test/mysql$ mkdir master
dollarkiller@worldlink:~/data/test/mysql$ mkdir slave
dollarkiller@worldlink:~/data/test/mysql$ pwd
/home/dollarkiller/data/test/mysql
```

启动主节点容器
```
docker run -d --name mysql-master -p 3307:3306 -v /home/dollarkiller/data/test/mysql/master.conf:/etc/mysql/mysql.conf.d/mysqld.cnf -v  /home/dollarkiller/data/test/mysql/maste:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7.26
```
启动从节点
```
docker run -d --name mysql-slave -p 3308:3306 -v /home/dollarkiller/data/test/mysql/master.conf:/etc/mysql/mysql.conf.d/mysqld.cnf -v  /home/dollarkiller/data/test/mysql/maste:/var/lib/slave -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7.26
```
CHANGE MASTER TO MASTER_HOST='172.17.0.2',MASTER_PORT=3306,MASTER_USER='slave',MASTER_PASSWORD='slave',MASTER_LOG_FILE='log.000004',MASTER_LOG_POS=154;