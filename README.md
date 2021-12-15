# petty

## 部署etcd,mysql,redis

```
cd docker
docker-compose up -d
```

## 编译

```
cd xq
make
```

## 运行

## 创建数据库

```
mysql
mysql> create database petty
```

### server

```
cd xq
sh serve.sh
```

### client

```
cd xq
sh run-cli.sh
```

```
# 登录 (username/password)
-> login a a
# 注册 (username/password)
-> register a a
# 查看table列表
-> tables
```

### Todo

1. auto login
2. status
   服务端推流。广播消息，聊天消息，系统通知等。
3. 添加好友
4. 搜索，自动补全
5. recent logs, common logs
6. counters, statistics
7. distributed locks
