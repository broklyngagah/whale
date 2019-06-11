package db

import (
	"github.com/jmoiron/sqlx"
	"time"
	"github.com/gocql/gocql"
	"sync"
	"gopkg.in/redis.v4"
	"fmt"
	"carp.cn/whale/config"
	"carp.cn/whale/zaplogger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/bradfitz/gomemcache/memcache"
	"carp.cn/whale/pkg/mongo"
)

const driverName = "mysql"

type loader interface {
	LoadAll()
}

var (
	BaseDataCaches   = make(map[string]loader)
	StatsDBCassandra *gocql.Session
	JYBaseDB         *sqlx.DB
	JYLogDB          *sqlx.DB
	JYMemberDB       *sqlx.DB
	JYOtherDB        *sqlx.DB
	JYTradeDB        *sqlx.DB
	once             = &sync.Once{}
	Rds              *redis.Client
	MemCache         *memcache.Client
)

func InitDB() {
	once.Do(func() {
		cnf := config.Get()

		JYBaseDB = initSqlxDB(cnf.JYBaseDB.GetDSN(), "[JY_BASE_DB] -> ", cnf.JYBaseDB.GetMaxOpen(), cnf.JYBaseDB.GetMaxIdle())
		JYLogDB = initSqlxDB(cnf.JYLogDB.GetDSN(), "[JY_LOG_DB] -> ", cnf.JYLogDB.GetMaxOpen(), cnf.JYLogDB.GetMaxIdle())
		JYMemberDB = initSqlxDB(cnf.JYMemberDB.GetDSN(), "[JY_MEMBER_DB] -> ", cnf.JYMemberDB.GetMaxOpen(), cnf.JYMemberDB.GetMaxIdle())
		JYOtherDB = initSqlxDB(cnf.JYOtherDB.GetDSN(), "[JY_OTHER_DB] -> ", cnf.JYOtherDB.GetMaxOpen(), cnf.JYOtherDB.GetMaxIdle())
		JYTradeDB = initSqlxDB(cnf.JYTradeDB.GetDSN(), "[JY_TRADE_DB] -> ", cnf.JYTradeDB.GetMaxOpen(), cnf.JYTradeDB.GetMaxIdle())

		Rds = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cnf.GetRedisHost(), cnf.GetRedisPort()),
			Password: cnf.GetRedisPass(),
			DB:       0,
		})

		mc := memcache.New(cnf.MemCache.GetAddr())
		if mc == nil {
			zaplogger.Fatal("base memcache client failed.")
		} else {
			mc.MaxIdleConns = cnf.MemCache.MaxIdleConns
			mc.Timeout = cnf.MemCache.GetTimeout()
			MemCache = mc
		}

		mongo.Init(mongo.GenSimpleUrl(cnf.MongoDB.Host, cnf.MongoDB.Port))

		zaplogger.Info("Init DB success.")
	})

}

func initSqlxDB(dbConfig, logHeader string, maxOpen, maxIdle int) *sqlx.DB {
	db := sqlx.MustConnect(driverName, dbConfig)
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)
	// https://github.com/go-sql-driver/mysql/issues/446
	db.SetConnMaxLifetime(time.Second * 14400)
	return db
}

