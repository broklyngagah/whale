package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_zone
//城市表

// +gen
type SysZone struct {
	ZoneId      int     `db:"zone_id" json:"zone_id"`             // 城市ID
	GeoLat      float64 `db:"geo_lat" json:"geo_lat"`             // 纬度
	GeoLng      float64 `db:"geo_lng" json:"geo_lng"`             // 经度
	Affiliation int     `db:"Affiliation" json:"Affiliation"`     // zone_id的子id
	ZoneName    string  `db:"zone_name" json:"zone_name"`         // 城市名字
	Rank        int8    `db:"Rank" json:"Rank"`                   // 0表示省级，1表示市级，2表是区，县
	CN          string  `db:"CN" json:"CN"`                       // 中文
	EN          string  `db:"EN" json:"EN"`                       // 英文
	Postcode    string  `db:"Postcode" json:"Postcode"`           // 邮编
	PhoneCode   string  `db:"Phone_Code" json:"Phone_Code"`       // 电话前缀
	WID         int8    `db:"WID" json:"WID"`                     //
	Top         int     `db:"Top" json:"Top"`                     //
	JingdianNo  int8    `db:"jingdian_no" json:"jingdian_no"`     //
	Orderid     int8    `db:"orderid" json:"orderid"`             //
	BaiduProNum int8    `db:"baidu_pro_num" json:"baidu_pro_num"` //
	CityRank    int8    `db:"city_rank" json:"city_rank"`         //
	Hits        int     `db:"hits" json:"hits"`                   //
	UpEn        string  `db:"up_en" json:"up_en"`                 // 英文简写
	Letter      string  `db:"letter" json:"letter"`               // 首字母
	Location    string  `db:"location" json:"location"`           // 地理位置
	Population  string  `db:"population" json:"population"`       // 人口
	SubArea     string  `db:"sub_area" json:"sub_area"`           // 几个市县
	ZoneArea    string  `db:"zone_area" json:"zone_area"`         // 面积
	Weather     string  `db:"weather" json:"weather"`             // 地区天气
	Shortname   string  `db:"shortname" json:"shortname"`         // 简称
	Capital     string  `db:"capital" json:"capital"`             // 首都
	IsUsed      int8    `db:"is_used" json:"is_used"`             // 已经开通城市 0-未开通 1-已开通
	Area        string  `db:"area" json:"area"`                   // 地区：华东，华南
	IsHot       int8    `db:"is_hot" json:"is_hot"`               // 热门城市 1为热门城市
	IsGsd       int8    `db:"is_gsd" json:"is_gsd"`               // 是否为归属地使用
	IsUsed400   int8    `db:"is_used_400" json:"is_used_400"`     // 400开通城市 1-未开通 2-已开通
	HeavenCode  string  `db:"heaven_code" json:"heaven_code"`     // 天府存管地区代码
}

var DefaultSysZone = SysZone{}

type sysZoneCache struct {
	objMap  map[int]*SysZone
	objList []*SysZone
}

var SysZoneCache = &sysZoneCache{}

func (c *sysZoneCache) LoadAll() {
	sql := "select `zone_id`,`geo_lat`,`geo_lng`,`Affiliation`,`zone_name`,`Rank`,`CN`,`EN`,`Postcode`,`Phone_Code`,`WID`,`Top`,`jingdian_no`,`orderid`,`baidu_pro_num`,`city_rank`,`hits`,`up_en`,`letter`,`location`,`population`,`sub_area`,`zone_area`,`weather`,`shortname`,`capital`,`is_used`,`area`,`is_hot`,`is_gsd`,`is_used_400`,`heaven_code` from sys_zone"
	c.objList = make([]*SysZone, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysZone)
	for _, v := range c.objList {
		c.objMap[v.ZoneId] = v
	}
}

func (c *sysZoneCache) All() []*SysZone {
	return c.objList
}

func (c *sysZoneCache) Count() int {
	return len(c.objList)
}

func (c *sysZoneCache) Get(zone_id int) (*SysZone, bool) {
	key := zone_id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysZoneCache) Update(v *SysZone) {
	key := v.ZoneId
	c.objMap[key] = v
}
