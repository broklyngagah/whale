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

//jy_trade_conf
//系统充值鱼币配置表

// +gen *
type JyTradeConf struct {
	Id           int     `db:"id" json:"id"`                       // id
	Name         string  `db:"name" json:"name"`                   // 牛币名称
	PayMoney     float64 `db:"pay_money" json:"pay_money"`         // 实际支付金额
	AccountMoney float64 `db:"account_money" json:"account_money"` // 到账金额
	Status       int8    `db:"status" json:"status"`               // 1未审核 2已审核 9删除
	Sort         int8    `db:"sort" json:"sort"`                   // 排序
	CTime        int64   `db:"c_time" json:"c_time"`               // 创建时间
	UTime        int64   `db:"u_time" json:"u_time"`               // 改修时间
}

type jyTradeConfOp struct{}

var JyTradeConfOp = &jyTradeConfOp{}
var DefaultJyTradeConf = &JyTradeConf{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyTradeConfOp) Get(id int) (*JyTradeConf, error) {
	obj := &JyTradeConf{}
	sql := "select `id`,`name`,`pay_money`,`account_money`,`status`,`sort`,`c_time`,`u_time` from jy_trade_conf where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyTradeConfOp) SelectAll() ([]*JyTradeConf, error) {
	objList := []*JyTradeConf{}
	sql := "select `id`,`name`,`pay_money`,`account_money`,`status`,`sort`,`c_time`,`u_time` from jy_trade_conf"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyTradeConfOp) QueryByMap(m map[string]interface{}) ([]*JyTradeConf, error) {
	result := []*JyTradeConf{}
	var params []interface{}

	sql := "select `id`,`name`,`pay_money`,`account_money`,`status`,`sort`,`c_time`,`u_time` from jy_trade_conf where 1=1 "
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

func (op *jyTradeConfOp) QueryByMapComparison(m map[string]interface{}) ([]*JyTradeConf, error) {
	result := []*JyTradeConf{}
	var params []interface{}

	sql := "select `id`,`name`,`pay_money`,`account_money`,`status`,`sort`,`c_time`,`u_time` from jy_trade_conf where 1=1 "
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

func (op *jyTradeConfOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyTradeConf, error) {
	result := []*JyTradeConf{}
	var params []interface{}

	sql := "select `id`,`name`,`pay_money`,`account_money`,`status`,`sort`,`c_time`,`u_time` from jy_trade_conf where 1=1 "
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

func (op *jyTradeConfOp) GetByMap(m map[string]interface{}) (*JyTradeConf, error) {
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
func (op *jyTradeConfOp) Insert(m *JyTradeConf) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyTradeConfOp) InsertTx(ext sqlx.Ext, m *JyTradeConf) (int64, error) {
	sql := "insert into jy_trade_conf(name,pay_money,account_money,status,sort,c_time,u_time) values(?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Name,
		m.PayMoney,
		m.AccountMoney,
		m.Status,
		m.Sort,
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
func (i *JyTradeConf) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyTradeConfOp) Update(m *JyTradeConf) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyTradeConfOp) UpdateTx(ext sqlx.Ext, m *JyTradeConf) error {
	sql := `update jy_trade_conf set name=?,pay_money=?,account_money=?,status=?,sort=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Name,
		m.PayMoney,
		m.AccountMoney,
		m.Status,
		m.Sort,
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
func (op *jyTradeConfOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyTradeConfOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update jy_trade_conf set %s where 1=1 and id=? ;`

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
func (i *JyTradeConf) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyTradeConfOp) Delete(id int) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyTradeConfOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from jy_trade_conf where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyTradeConfOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_trade_conf where 1=1 `
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

func (op *jyTradeConfOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyTradeConfOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_trade_conf where 1=1 "
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
