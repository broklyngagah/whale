package config

import (
	"io/ioutil"
	"encoding/json"
	"github.com/lunny/log"
)

var config Config

func Get() *Config {
	return &config
}

type Config struct {
	Log        *Log
	JYBaseDB   *DBRoute
	JYLogDB    *DBRoute
	JYMemberDB *DBRoute
	JYOtherDB  *DBRoute
	JYTradeDB  *DBRoute
	Redis      *MemRoute
	MemCache   *CacheRoute
	Amqp       *AccountRoute
	MongoDB    *AccountRoute
}

func (c *Config) GetRedisHost() string {
	return c.Redis.Host
}

func (c *Config) GetRedisPass() string {
	return c.Redis.Pass
}

func (c *Config) GetRedisPort() int {
	return c.Redis.Port
}

func LoadFromFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("config:", config)
}
