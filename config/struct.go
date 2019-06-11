package config

import (
	"fmt"
	"time"
)

type DBEngine string

const (
	MYSQL     DBEngine = "mysql"
	CASSANDRA DBEngine = "cassandra"
)

const (
	default_db_max_open      = 32
	default_db_max_idle      = 2
	default_redis_max_open   = 1
	default_redis_max_idle   = 1
	default_stat_log_workers = 64
)

type Log struct {
	Name  string
	Dir   string
	Level string
}

//--------------------------------------------------------------
type Route struct {
	Host string
	Port int
}

func (r Route) GetAddr() string {
	if r.Host != "" {
		return fmt.Sprintf("%s:%d", r.Host, r.Port)
	}
	return fmt.Sprintf("%s:%d", "127.0.0.1", r.Port)
}

//--------------------------------------------------------------
type AuthRoute struct {
	Host string
	Port int
	Pass string
}

func (r *AuthRoute) GetAddr() string {
	if r.Host != "" {
		return fmt.Sprintf("%s:%d", r.Host, r.Port)
	}
	return fmt.Sprintf("%s:%d", "127.0.0.1", r.Port)
}

//--------------------------------------------------------------
type DBRoute struct {
	Engine  DBEngine
	Host    string
	Port    int
	Name    string
	User    string
	Pass    string
	Params  string
	MaxOpen int
	MaxIdle int
	Workers int
}

func (r *DBRoute) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		r.User,
		r.Pass,
		r.Host,
		r.Port,
		r.Name,
		r.Params,
	)
}

func (r *DBRoute) GetMaxOpen() int {
	if r.MaxOpen <= 0 {
		return default_db_max_open
	}
	return r.MaxOpen
}

func (r *DBRoute) GetMaxIdle() int {
	if r.MaxIdle <= 0 {
		return default_db_max_idle
	}
	return r.MaxIdle
}

func (r *DBRoute) GetWorkers() int {
	if r.Workers <= 0 {
		return default_stat_log_workers
	}
	return r.Workers
}

//--------------------------------------------------------------
type MemRoute struct {
	Host      string
	Port      int
	Pass      string
	DB        int
	MaxIdle   int
	MaxActive int
}

func (r *MemRoute) GetMaxiIdle() int {
	if r.MaxIdle <= 0 {
		return default_redis_max_idle
	}
	return r.MaxIdle
}

func (r *MemRoute) GetMaxActive() int {
	if r.MaxActive <= 0 {
		return default_redis_max_open
	}
	return r.MaxActive
}

//--------------------------------------------------------------
type CacheRoute struct {
	Host         string
	Port         int
	Pass         string // password
	Timeout      int    // second
	MaxIdleConns int
}

func (c *CacheRoute) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *CacheRoute) GetTimeout() time.Duration {
	return time.Duration(c.Timeout)
}

//--------------------------------------------------------------
type Url struct {
	Host string
	Port int
	Path string
}

func (u *Url) GetUrl() string {
	var buf string
	if u.Port != 0 {
		buf = fmt.Sprintf("http://%s:%d", u.Host, u.Port)
	} else {
		buf = fmt.Sprintf("http://%s", u.Host)
	}
	if u.Path != "" {
		buf += u.Path
	}
	return buf
}

//--------------------------------------------------------------
type CassandraRoute struct {
	Host             string
	Port             int
	Name             string
	User             string
	Pass             string
	ConsistencyLevel string
	Timeout          int
}

//--------------------------------------------------------------
type ServerNodeRoute struct {
	ServerId int
	NodeId   int
	Consul   ConsulRoute
	PublicIP string
}

type ConsulRoute struct {
	Route
	DC    string
	Token string
}

type AccountRoute struct {
	Host string
	Port int
	User string
	Pass string
}
