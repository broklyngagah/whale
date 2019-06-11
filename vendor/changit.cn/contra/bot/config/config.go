package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	AuthUrl        string `toml:"auth_url"`
	RedisAddr      string `toml:"redis_addr"`
	RedisPwd       string `toml:"redis_pwd"`
	Concurrency    int    `toml:"concurrency"`
	CountPerCon    int    `toml:"count_per_con"`
	ServerId       int    `toml:"server_id"`
	RobotInterval  int    `toml:"robot_interval"`
	UsernamePrefix string `toml:"username_prefix"`
	DBDsn          string `toml:"db_dsn"`
	IsMatch        bool   `toml:"is_match"`
}

var cfg Config

func InitConfig(filename string) {
	_, err := toml.DecodeFile(filename, &cfg)
	if err != nil {
		log.Fatalf("toml.DecodeFile error(%v).\n", err)
	}
}

func GetAuthUrl() string {
	return cfg.AuthUrl
}

func GetRedisAddr() string {
	return cfg.RedisAddr
}

func GetRedisPwd() string {
	return cfg.RedisPwd
}

func GetConcurrency() int {
	return cfg.Concurrency
}

func GetServerId() int {
	return cfg.ServerId
}

func GetRobotInterval() int {
	return cfg.RobotInterval
}

func GetUsernamePrefix() string {
	return cfg.UsernamePrefix
}

func GetDBDsn() string {
	return cfg.DBDsn
}

func IsMatch() bool {
	return cfg.IsMatch
}
