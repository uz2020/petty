package config

import (
	"github.com/spf13/viper"
)

type Conf struct {
	ListenAddr, Service, EtcdUrl               string
	MysqlDb, MysqlAddr, MysqlUser, MysqlPasswd string
	RedisAddr                                  string
}

func (conf *Conf) Init() {
	conf.ListenAddr = viper.GetString("LISTEN_ADDR")
	conf.Service = viper.GetString("SERVICE")
	conf.EtcdUrl = viper.GetString("ETCD_URL")

	conf.MysqlDb = viper.GetString("MYSQL_DB")
	conf.MysqlAddr = viper.GetString("MYSQL_ADDR")
	conf.MysqlUser = viper.GetString("MYSQL_USER")
	conf.MysqlPasswd = viper.GetString("MYSQL_PASSWD")

	conf.RedisAddr = viper.GetString("REDIS_ADDR")
}
