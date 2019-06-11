package mongo

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"carp.cn/whale/zaplogger"
	"go.uber.org/zap"
	"fmt"
)

var (
	_URL     string
	_Session *mgo.Session
)

type Counter struct {
	ID  string `bson:"_id"`
	Seq int
}

func GenAuthUrl(user, pass, host string, port int, db string) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", user, pass, host, port, db)
}
func GenSimpleUrl(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func Init(url string) {

	_URL = url

	// mongodb的Session只在程序开启的时候建立连接，之后直接Copy共用现有连接
	// 如果没有copy, 各goroutine共用同一个Session会导致所有mongodb的访问变成线性的
	var err error
	for true {
		_Session, err = mgo.Dial(_URL)
		if err != nil {
			zaplogger.Error("failed to connect mongodb", zap.Error(err))
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	Factory = NewPool(128)
}

//func NewSession() *mgo.Session {
//	return _Session.Copy()
//}

func GetNextSequence(db *mgo.Database, name string) (int, error) {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	var counter Counter
	_, err := db.C("counters").FindId(name).Apply(change, &counter)
	return counter.Seq, err
}

