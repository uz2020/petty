# petty

## 部署etcd

```
cd etcd
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
-> tables
-> login a a
```
