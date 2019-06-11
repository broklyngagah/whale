package JYTradeDB

import (
	"errors"
	"fmt"
	"strings"

	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
	"changit.cn/contra/server/game_error"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

//This file is generate by scripts,don't edit it

//jy_user_subscribe
//用户购买订阅表(记录整个交易链路)

// +gen *
type JyUserSubscribe struct {
	Id         int64  `db:"id" json:"id"`                   //
	Uid        int    `db:"uid" json:"uid"`                 // 订阅者ID
	GoodName   string `db:"good_name" json:"good_name"`     // 订阅商品名称
	InOid      int64  `db:"in_oid" json:"in_oid"`           // 买方充值的订单ID (非充值购买时为0)
	OutOid     int64  `db:"out_oid" json:"out_oid"`         // 买方转出的订单ID
	PassiveOid int64  `db:"passive_oid" json:"passive_oid"` // 卖方转入的订单ID
	PassiveUid int    `db:"passive_uid" json:"passive_uid"` // 被订阅者UID
	OtherId    int    `db:"other_id" json:"other_id"`       // 其他ID  (订阅用户时jy_trade_conf 配置id，购买文章时aid）
	OrderType  int8   `db:"order_type" json:"order_type"`   // 类型 1 现金购买 2 鱼币购买  3 积分购买）
	Type       int8   `db:"type" json:"type"`               // 订阅类型：1购买文章(永久)，2订阅用户(有e_time)
	STime      int64  `db:"s_time" json:"s_time"`           // 开始时间
	ETime      int64  `db:"e_time" json:"e_time"`           // 结束时间
	CTime      int64  `db:"c_time" json:"c_time"`           // 添加时间
	UTime      int64  `db:"u_time" json:"u_time"`           // 更行时间
}

type jyUserSubscribeOp struct{}

var JyUserSubscribeOp = &jyUserSubscribeOp{}
var DefaultJyUserSubscribe = &JyUserSubscribe{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyUserSubscribeOp) Get(id int64) (*JyUserSubscribe, error) {
	obj := &JyUserSubscribe{}
	sql := "select `id`,`uid`,`good_name`,`in_oid`,`out_oid`,`passive_oid`,`passive_uid`,`other_id`,`order_type`,`type`,`s_time`,`e_time`,`c_time`,`u_time` from jy_user_subscribe where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyUserSubscribeOp) SelectAll() ([]*JyUserSubscribe, error) {
	objList := []*JyUserSubscribe{}
	sql := "select `id`,`uid`,`good_name`,`in_oid`,`out_oid`,`passive_oid`,`passive_uid`,`other_id`,`order_type`,`type`,`s_time`,`e_time`,`c_time`,`u_time` from jy_user_subscribe"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyUserSubscribeOp) QueryByMap(m map[string]interface{}) ([]*JyUserSubscribe, error) {
	result := []*JyUserSubscribe{}
	var params []interface{}

	sql := "select `id`,`uid`,`good_name`,`in_oid`,`out_oid`,`passive_oid`,`passive_uid`,`other_id`,`order_type`,`type`,`s_time`,`e_time`,`c_time`,`u_time` from jy_user_subscribe where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s=? ", k)
		params = append(params, v)
	}
	err := db.JYTradeDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyUserSubscribeOp) QueryByMapComparison(m map[string]interface{}) ([]*JyUserSubscribe, error) {
	result := []*JyUserSubscribe{}
	var params []interface{}

	sql := "select `id`,`uid`,`good_name`,`in_oid`,`out_oid`,`passive_oid`,`passive_uid`,`other_id`,`order_type`,`type`,`s_time`,`e_time`,`c_time`,`u_time` from jy_user_subscribe where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s? ", k)
		params = append(params, v)
	}
	err := db.JYTradeDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyUserSubscribeOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyUserSubscribe, error) {
	result := []*JyUserSubscribe{}
	var params []interface{}

	sql := "select `id`,`uid`,`good_name`,`in_oid`,`out_oid`,`passive_oid`,`passive_uid`,`other_id`,`order_type`,`type`,`s_time`,`e_time`,`c_time`,`u_time` from jy_user_subscribe where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s? ", k)
		params = append(params, v)
	}

	if len(orderby) > 0 {
		for k, v := range orderby {
			if len(v) < 2 {
				continue
			}
			opr := v[:1]
			switch opr {
			case "-":
				orderby[k] = fmt.Sprintf("%s desc", v[1:len(v)])
			case "+":
				orderby[k] = fmt.Sprintf("%s", v[1:len(v)])
			default:
				orderby[k] = fmt.Sprintf("%s", v)
			}
		}
		sql += fmt.Sprintf(" order by %s", strings.Join(orderby, ", "))
	}

	if len(clause) > 0 {
		sql += fmt.Sprintf(" %s", strings.Join(clause, " "))
	}

	if limit > 0 {
		sql += fmt.Sprintf(" limit %d offset %d", limit, offset)
	}

	zaplogger.Info("[SQL]:"+sql, zap.Reflect("| params:", params))

	err := db.JYTradeDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyUserSubscribeOp) GetByMap(m map[string]interface{}) (*JyUserSubscribe, error) {
	lst, err := op.QueryByMap(m)
	if err != nil {
		return nil, err
	}
	if len(lst) == 1 {
		return lst[0], nil
	} else if len(lst) == 0 {
		return nil, nil
	}

	return nil, errors.New("Get multi rows.")
}

// 插入数据，自增长字段将被忽略
func (op *jyUserSubscribeOp) Insert(m *JyUserSubscribe) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyUserSubscribeOp) InsertTx(ext sqlx.Ext, m *JyUserSubscribe) (int64, error) {
	sql := "insert into jy_user_subscribe(uid,good_name,in_oid,out_oid,passive_oid,passive_uid,other_id,order_type,type,s_time,e_time,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.GoodName,
		m.InOid,
		m.OutOid,
		m.PassiveOid,
		m.PassiveUid,
		m.OtherId,
		m.OrderType,
		m.Type,
		m.STime,
		m.ETime,
		m.CTime,
		m.UTime,
	)
	if err != nil {
		game_error.RaiseError(err)
		return -1, err
	}
	affected, _ := result.RowsAffected()
	return affected, nil
}

/*
func (i *JyUserSubscribe) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyUserSubscribeOp) Update(m *JyUserSubscribe) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyUserSubscribeOp) UpdateTx(ext sqlx.Ext, m *JyUserSubscribe) error {
	sql := `update jy_user_subscribe set uid=?,good_name=?,in_oid=?,out_oid=?,passive_oid=?,passive_uid=?,other_id=?,order_type=?,type=?,s_time=?,e_time=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.GoodName,
		m.InOid,
		m.OutOid,
		m.PassiveOid,
		m.PassiveUid,
		m.OtherId,
		m.OrderType,
		m.Type,
		m.STime,
		m.ETime,
		m.CTime,
		m.UTime,
		m.Id,
	)

	if err != nil {
		game_error.RaiseError(err)
		return err
	}

	return nil
}

// 用主键做条件，更新map里包含的字段名
func (op *jyUserSubscribeOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyUserSubscribeOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update jy_user_subscribe set %s where 1=1 and id=? ;`

	var params []interface{}
	var set_sql string
	for k, v := range m {
		set_sql += fmt.Sprintf(" %s=? ", k)
		params = append(params, v)
	}
	params = append(params, id)
	_, err := ext.Exec(fmt.Sprintf(sql, set_sql), params...)
	return err
}

/*
func (i *JyUserSubscribe) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyUserSubscribeOp) Delete(id int64) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyUserSubscribeOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from jy_user_subscribe where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyUserSubscribeOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_user_subscribe where 1=1 `
	for k, v := range m {
		sql += fmt.Sprintf(" and  %s=? ", k)
		params = append(params, v)
	}
	count := int64(-1)
	err := db.JYTradeDB.Get(&count, sql, params...)
	if err != nil {
		game_error.RaiseError(err)
	}
	return count
}

func (op *jyUserSubscribeOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyUserSubscribeOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_user_subscribe where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s=? ", k)
		params = append(params, v)
	}
	result, err := ext.Exec(sql, params...)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
