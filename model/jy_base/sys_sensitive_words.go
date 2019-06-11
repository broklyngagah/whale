package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_sensitive_words
//敏感词库

// +gen
type SysSensitiveWords struct {
	Id   int    `db:"id" json:"id"`     //
	Name string `db:"name" json:"name"` // 敏感词
}

var DefaultSysSensitiveWords = SysSensitiveWords{}

type sysSensitiveWordsCache struct {
	objMap  map[int]*SysSensitiveWords
	objList []*SysSensitiveWords
}

var SysSensitiveWordsCache = &sysSensitiveWordsCache{}

func (c *sysSensitiveWordsCache) LoadAll() {
	sql := "select `id`,`name` from sys_sensitive_words"
	c.objList = make([]*SysSensitiveWords, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysSensitiveWords)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysSensitiveWordsCache) All() []*SysSensitiveWords {
	return c.objList
}

func (c *sysSensitiveWordsCache) Count() int {
	return len(c.objList)
}

func (c *sysSensitiveWordsCache) Get(id int) (*SysSensitiveWords, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysSensitiveWordsCache) Update(v *SysSensitiveWords) {
	key := v.Id
	c.objMap[key] = v
}
