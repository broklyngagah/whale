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

//jy_settlement_detail
//结算明细表

// +gen *
type JySettlementDetail struct {
	Id          int64   `db:"id" json:"id"`                     //
	Uid         int     `db:"uid" json:"uid"`                   // 用户ID
	Money       float64 `db:"money" json:"money"`               // 订单发生本金
	EnableMoney float64 `db:"enable_money" json:"enable_money"` // 可结算金额
	UseMoney    float64 `db:"use_money" json:"use_money"`       // 已结算金额
	EnableTime  int64   `db:"enable_time" json:"enable_time"`   // 可结算时间
	PayTime     int64   `db:"pay_time" json:"pay_time"`         // 结算时间
	Status      int8    `db:"status" json:"status"`             // 1未结算 2已结算
	ConfType    int8    `db:"conf_type" json:"conf_type"`       // 包月类型：0单篇，1一月，3一季，6半年，12一年
	Oid         string  `db:"oid" json:"oid"`                   // order表订单id
	Soid        string  `db:"soid" json:"soid"`                 // 结算订单id
	CTime       int64   `db:"c_time" json:"c_time"`             // 创建时间
	UTime       int64   `db:"u_time" json:"u_time"`             // 订单最后修改时间
}

type jySettlementDetailOp struct{}

var JySettlementDetailOp = &jySettlementDetailOp{}
var DefaultJySettlementDetail = &JySettlementDetail{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jySettlementDetailOp) Get(id int64) (*JySettlementDetail, error) {
	obj := &JySettlementDetail{}
	sql := "select `id`,`uid`,`money`,`enable_money`,`use_money`,`enable_time`,`pay_time`,`status`,`conf_type`,`oid`,`soid`,`c_time`,`u_time` from jy_settlement_detail where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jySettlementDetailOp) SelectAll() ([]*JySettlementDetail, error) {
	objList := []*JySettlementDetail{}
	sql := "select `id`,`uid`,`money`,`enable_money`,`use_money`,`enable_time`,`pay_time`,`status`,`conf_type`,`oid`,`soid`,`c_time`,`u_time` from jy_settlement_detail"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jySettlementDetailOp) QueryByMap(m map[string]interface{}) ([]*JySettlementDetail, error) {
	result := []*JySettlementDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`money`,`enable_money`,`use_money`,`enable_time`,`pay_time`,`status`,`conf_type`,`oid`,`soid`,`c_time`,`u_time` from jy_settlement_detail where 1=1 "
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

func (op *jySettlementDetailOp) QueryByMapComparison(m map[string]interface{}) ([]*JySettlementDetail, error) {
	result := []*JySettlementDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`money`,`enable_money`,`use_money`,`enable_time`,`pay_time`,`status`,`conf_type`,`oid`,`soid`,`c_time`,`u_time` from jy_settlement_detail where 1=1 "
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

func (op *jySettlementDetailOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JySettlementDetail, error) {
	result := []*JySettlementDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`money`,`enable_money`,`use_money`,`enable_time`,`pay_time`,`status`,`conf_type`,`oid`,`soid`,`c_time`,`u_time` from jy_settlement_detail where 1=1 "
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

func (op *jySettlementDetailOp) GetByMap(m map[string]interface{}) (*JySettlementDetail, error) {
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
func (op *jySettlementDetailOp) Insert(m *JySettlementDetail) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jySettlementDetailOp) InsertTx(ext sqlx.Ext, m *JySettlementDetail) (int64, error) {
	sql := "insert into jy_settlement_detail(uid,money,enable_money,use_money,enable_time,pay_time,status,conf_type,oid,soid,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Money,
		m.EnableMoney,
		m.UseMoney,
		m.EnableTime,
		m.PayTime,
		m.Status,
		m.ConfType,
		m.Oid,
		m.Soid,
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
func (i *JySettlementDetail) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jySettlementDetailOp) Update(m *JySettlementDetail) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jySettlementDetailOp) UpdateTx(ext sqlx.Ext, m *JySettlementDetail) error {
	sql := `update jy_settlement_detail set uid=?,money=?,enable_money=?,use_money=?,enable_time=?,pay_time=?,status=?,conf_type=?,oid=?,soid=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Money,
		m.EnableMoney,
		m.UseMoney,
		m.EnableTime,
		m.PayTime,
		m.Status,
		m.ConfType,
		m.Oid,
		m.Soid,
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
func (op *jySettlementDetailOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jySettlementDetailOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update jy_settlement_detail set %s where 1=1 and id=? ;`

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
func (i *JySettlementDetail) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jySettlementDetailOp) Delete(id int64) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jySettlementDetailOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from jy_settlement_detail where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jySettlementDetailOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_settlement_detail where 1=1 `
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

func (op *jySettlementDetailOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jySettlementDetailOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_settlement_detail where 1=1 "
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
