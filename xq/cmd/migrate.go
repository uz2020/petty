package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uz2020/petty/xq/config"
	"github.com/uz2020/petty/xq/db"
)

func initTable() {
	conf := &config.Conf{}
	conf.Init()
	newDb, err := db.InitDb(conf.MysqlUser, conf.MysqlPasswd, conf.MysqlAddr, conf.MysqlDb)
	if err != nil {
		panic(err)
	}

	d := newDb.Migrator()
	sqlDb, _ := newDb.DB()
	defer sqlDb.Close()

	if !d.HasTable(&db.TbUser{}) {
		err := d.CreateTable(&db.TbUser{})
		if err != nil {
			panic(err)
		}
	}
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Create table struct",
	Run: func(cmd *cobra.Command, args []string) {
		initTable()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
