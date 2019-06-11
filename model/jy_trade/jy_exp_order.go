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

//jy_exp_order
//经验订单表

// +gen *
type JyExpOrder struct {
	Id         int64  `db:"id" json:"id"`                   //
	Uid        int    `db:"uid" json:"uid"`                 // 用户ID
	ActionType int8   `db:"action_type" json:"action_type"` // 1转入 2转出
	OrderType  int8   `db:"order_type" json:"order_type"`   // 经验类型 1通过平台认证，2自媒体相关资质 3资格认证 4发布付费文章 5首发付费文章
	Exp        int    `db:"exp" json:"exp"`                 // 经验值
	Status     int8   `db:"status" json:"status"`           // 状态1成功 2失败
	OtherId    int    `db:"other_id" json:"other_id"`       // 其他ID （文章id）
	Ip         string `db:"ip" json:"ip"`                   // 用户IP地址
	CTime      int64  `db:"c_time" json:"c_time"`           // 创建时间
	UTime      int64  `db:"u_time" json:"u_time"`           // 最后修改时间
}

type jyExpOrderOp struct{}

var JyExpOrderOp = &jyExpOrderOp{}
var DefaultJyExpOrder = &JyExpOrder{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyExpOrderOp) Get(id int64) (*JyExpOrder, error) {
	obj := &JyExpOrder{}
	sql := "select `id`,`uid`,`action_type`,`order_type`,`exp`,`status`,`other_id`,`ip`,`c_time`,`u_time` from jy_exp_order where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyExpOrderOp) SelectAll() ([]*JyExpOrder, error) {
	objList := []*JyExpOrder{}
	sql := "select `id`,`uid`,`action_type`,`order_type`,`exp`,`status`,`other_id`,`ip`,`c_time`,`u_time` from jy_exp_order"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyExpOrderOp) QueryByMap(m map[string]interface{}) ([]*JyExpOrder, error) {
	result := []*JyExpOrder{}
	var params []interface{}

	sql := "select `id`,`uid`,`action_type`,`order_type`,`exp`,`status`,`other_id`,`ip`,`c_time`,`u_time` from jy_exp_order where 1=1 "
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

func (op *jyExpOrderOp) QueryByMapComparison(m map[string]interface{}) ([]*JyExpOrder, error) {
	result := []*JyExpOrder{}
	var params []interface{}

	sql := "select `id`,`uid`,`action_type`,`order_type`,`exp`,`status`,`other_id`,`ip`,`c_time`,`u_time` from jy_exp_order where 1=1 "
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

func (op *jyExpOrderOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyExpOrder, error) {
	result := []*JyExpOrder{}
	var params []interface{}

	sql := "select `id`,`uid`,`action_type`,`order_type`,`exp`,`status`,`other_id`,`ip`,`c_time`,`u_time` from jy_exp_order where 1=1 "
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

func (op *jyExpOrderOp) GetByMap(m map[string]interface{}) (*JyExpOrder, error) {
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
func (op *jyExpOrderOp) Insert(m *JyExpOrder) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyExpOrderOp) InsertTx(ext sqlx.Ext, m *JyExpOrder) (int64, error) {
	sql := "insert into jy_exp_order(uid,action_type,order_type,exp,status,other_id,ip,c_time,u_time) values(?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.ActionType,
		m.OrderType,
		m.Exp,
		m.Status,
		m.OtherId,
		m.Ip,
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
func (i *JyExpOrder) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyExpOrderOp) Update(m *JyExpOrder) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyExpOrderOp) UpdateTx(ext sqlx.Ext, m *JyExpOrder) error {
	sql := `update jy_exp_order set uid=?,action_type=?,order_type=?,exp=?,status=?,other_id=?,ip=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.ActionType,
		m.OrderType,
		m.Exp,
		m.Status,
		m.OtherId,
		m.Ip,
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
func (op *jyExpOrderOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyExpOrderOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update jy_exp_order set %s where 1=1 and id=? ;`

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
func (i *JyExpOrder) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyExpOrderOp) Delete(id int64) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyExpOrderOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from jy_exp_order where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyExpOrderOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_exp_order where 1=1 `
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

func (op *jyExpOrderOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyExpOrderOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_exp_order where 1=1 "
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
