#!/bin/sh

export SERVER_SERVICE="pet/games/xq"
export SERVER_ETCD_URL="http://localhost:2379"

# mysql
export SERVER_MYSQL_DB="petty"
export SERVER_MYSQL_ADDR="127.0.0.1:3306"
export SERVER_MYSQL_USER="root"
export SERVER_MYSQL_PASSWD="1"

# redis
export SERVER_REDIS_ADDR="127.0.0.1:6379"
