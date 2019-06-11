package kit

import (
	"github.com/bradfitz/gomemcache/memcache"
	"gopkg.in/redis.v4"
	"encoding/json"
	"sync"
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
	"go.uber.org/zap"
	"time"
)

const (
	SessionID_Expire_Time = 3600 * 24 * 7
)

var (
	MemCacheHelper *memcacheHelper
	RdsCacheHelper *rdscacheHelper
	once           = &sync.Once{}
)

// -----------------------------------------------------------
type memcacheHelper struct {
	cli *memcache.Client
}

func NewMemcacheHelper(cli *memcache.Client) *memcacheHelper {
	return &memcacheHelper{
		cli: cli,
	}
}

func (m *memcacheHelper) Set(key string, value []byte) error {

	err := m.cli.Set(&memcache.Item{Key: key, Value: value})
	if err != nil {
		return err
	}
	err = m.cli.Touch(key, SessionID_Expire_Time)
	zaplogger.Error("memcache set key expire time error", zap.Error(err), zap.String("key:", key))
	return nil
}

func (m *memcacheHelper) Get(key string) ([]byte, error) {
	it, err := m.cli.Get(key)
	if err != nil {
		return nil, err
	}
	return it.Value, nil
}

func (m *memcacheHelper) Touch(key string) error {
	return m.cli.Touch(key, SessionID_Expire_Time)
}

func (m *memcacheHelper) SetObject(key string, value interface{}) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return m.Set(key, v)
}
func (m *memcacheHelper) GetObject(key string, value interface{}) error {
	buf, err := m.Get(key)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, value)
}

// -----------------------------------------------------------
type rdscacheHelper struct {
	cli *redis.Client
}

func NewRdscacheHelper(cli *redis.Client) *rdscacheHelper {
	return &rdscacheHelper{cli: cli}
}

func (r *rdscacheHelper) Set(key, value string, duration time.Duration) error {
	return r.cli.Set(key, value, time.Duration(time.Now().Unix())+duration).Err()
}

func (r *rdscacheHelper) Get(key string) (string, error) {
	res := r.cli.Get(key)
	if res.Err() != nil {
		return "", res.Err()
	}
	return res.Result()
}

func (r *rdscacheHelper) Del(key string) error {
	res := r.cli.Del(key)
	if res.Err() != nil {
		return res.Err()
	}
	_, err := res.Result()
	if err != nil {
		return err
	}
	return nil
}

// ZADD：向有序列表中添加元素
func (r *rdscacheHelper) ZAdd(key string, members ...redis.Z) (int64, error) {
	res := r.cli.ZAdd(key, members...)
	if res.Err() != nil {
		return 0, res.Err()
	}
	return res.Result()
}

// ZCARD：计算有序列表长度
func (r *rdscacheHelper) ZCard(key string) (int64, error) {
	res := r.cli.ZCard(key)
	if res.Err() != nil {
		return 0, res.Err()
	}
	return res.Result()
}

// ZREVRANGE： score 值递减(从大到小)来排列 获取最新n条记录
func (r *rdscacheHelper) ZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	res := r.cli.ZRevRangeWithScores(key, start, stop)
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res.Result()
}

// ZREMRANGEBYRANK 移除有序集 key 中，指定排名(rank)区间内的所有成员
func (r *rdscacheHelper) ZRemRangeByRank(key string, start, stop int64) error {
	res := r.cli.ZRemRangeByRank(key, start, stop)
	if res.Err() != nil {
		return res.Err()
	}
	_, err := res.Result()
	return err
}

// -----------------------------------------------------------
func Init() {
	once.Do(func() {
		MemCacheHelper = NewMemcacheHelper(db.MemCache)
		RdsCacheHelper = NewRdscacheHelper(db.Rds)
	})
}
