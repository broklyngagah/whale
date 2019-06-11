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

//jy_order_log
//订单付款日志表

// +gen *
type JyOrderLog struct {
	Id       int     `db:"id" json:"id"`               //
	Oid      int     `db:"oid" json:"oid"`             // 订单ID
	OrderId  string  `db:"order_id" json:"order_id"`   // 订单编号
	Uid      int     `db:"uid" json:"uid"`             // 用户ID
	PayMoney float64 `db:"pay_money" json:"pay_money"` // 交易金额
	PayType  int8    `db:"pay_type" json:"pay_type"`   // 支付方式 1微信 2 支付宝
	Status   int8    `db:"status" json:"status"`       // 1付款成功 2异常支付 9作废
	ErrorMsg string  `db:"error_msg" json:"error_msg"` // 支付失败原因
	CTime    int64   `db:"c_time" json:"c_time"`       // 付款时间
	UTime    int64   `db:"u_time" json:"u_time"`       // 修改时间
}

type jyOrderLogOp struct{}

var JyOrderLogOp = &jyOrderLogOp{}
var DefaultJyOrderLog = &JyOrderLog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyOrderLogOp) Get(id int) (*JyOrderLog, error) {
	obj := &JyOrderLog{}
	sql := "select `id`,`oid`,`order_id`,`uid`,`pay_money`,`pay_type`,`status`,`error_msg`,`c_time`,`u_time` from jy_order_log where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyOrderLogOp) SelectAll() ([]*JyOrderLog, error) {
	objList := []*JyOrderLog{}
	sql := "select `id`,`oid`,`order_id`,`uid`,`pay_money`,`pay_type`,`status`,`error_msg`,`c_time`,`u_time` from jy_order_log"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyOrderLogOp) QueryByMap(m map[string]interface{}) ([]*JyOrderLog, error) {
	result := []*JyOrderLog{}
	var params []interface{}

	sql := "select `id`,`oid`,`order_id`,`uid`,`pay_money`,`pay_type`,`status`,`error_msg`,`c_time`,`u_time` from jy_order_log where 1=1 "
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

func (op *jyOrderLogOp) QueryByMapComparison(m map[string]interface{}) ([]*JyOrderLog, error) {
	result := []*JyOrderLog{}
	var params []interface{}

	sql := "select `id`,`oid`,`order_id`,`uid`,`pay_money`,`pay_type`,`status`,`error_msg`,`c_time`,`u_time` from jy_order_log where 1=1 "
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

func (op *jyOrderLogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyOrderLog, error) {
	result := []*JyOrderLog{}
	var params []interface{}

	sql := "select `id`,`oid`,`order_id`,`uid`,`pay_money`,`pay_type`,`status`,`error_msg`,`c_time`,`u_time` from jy_order_log where 1=1 "
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

func (op *jyOrderLogOp) GetByMap(m map[string]interface{}) (*JyOrderLog, error) {
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
func (op *jyOrderLogOp) Insert(m *JyOrderLog) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyOrderLogOp) InsertTx(ext sqlx.Ext, m *JyOrderLog) (int64, error) {
	sql := "insert into jy_order_log(oid,order_id,uid,pay_money,pay_type,status,error_msg,c_time,u_time) values(?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Oid,
		m.OrderId,
		m.Uid,
		m.PayMoney,
		m.PayType,
		m.Status,
		m.ErrorMsg,
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
func (i *JyOrderLog) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyOrderLogOp) Update(m *JyOrderLog) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyOrderLogOp) UpdateTx(ext sqlx.Ext, m *JyOrderLog) error {
	sql := `update jy_order_log set oid=?,order_id=?,uid=?,pay_money=?,pay_type=?,status=?,error_msg=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Oid,
		m.OrderId,
		m.Uid,
		m.PayMoney,
		m.PayType,
		m.Status,
		m.ErrorMsg,
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
func (op *jyOrderLogOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyOrderLogOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update jy_order_log set %s where 1=1 and id=? ;`

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
func (i *JyOrderLog) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyOrderLogOp) Delete(id int) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyOrderLogOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from jy_order_log where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyOrderLogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_order_log where 1=1 `
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

func (op *jyOrderLogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyOrderLogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_order_log where 1=1 "
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
