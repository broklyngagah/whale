package db

import (
	"fmt"
	"sync"

	"changit.cn/contra/bot/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type loader interface {
	LoadAll()
}

var (
	BaseDB         *sqlx.DB
	BaseDataCaches = make(map[string]loader)
	once           = &sync.Once{}
)

func InitDB() {
	BaseDB = initSqlxDB(config.GetDBDsn(), "[BASE_DB] -> ", 32, 2)
	if err := BaseDB.Ping(); err != nil {
		fmt.Println("ping db fail.")
	} else {
		fmt.Println("ping db suc.")
	}
}

func initSqlxDB(dbConfig, logHeader string, maxOpen, maxIdle int) *sqlx.DB {
	db := sqlx.MustConnect("mysql", dbConfig)
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)
	return db
}
