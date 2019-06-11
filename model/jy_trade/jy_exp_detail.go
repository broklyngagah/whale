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

//jy_exp_detail
//用户经验明细表

// +gen *
type JyExpDetail struct {
	Id        int64 `db:"id" json:"id"`                 //
	Uid       int   `db:"uid" json:"uid"`               // 用户ID
	Oid       int64 `db:"oid" json:"oid"`               // 订单id
	Exp       int   `db:"exp" json:"exp"`               // 经验值
	EnableExp int   `db:"enable_exp" json:"enable_exp"` // 剩余可使用经验值
	UseExp    int   `db:"use_exp" json:"use_exp"`       // 已使用经验值
	CTime     int64 `db:"c_time" json:"c_time"`         // 创建时间
	UTime     int64 `db:"u_time" json:"u_time"`         // 最后修改时间
}

type jyExpDetailOp struct{}

var JyExpDetailOp = &jyExpDetailOp{}
var DefaultJyExpDetail = &JyExpDetail{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyExpDetailOp) Get(id int64) (*JyExpDetail, error) {
	obj := &JyExpDetail{}
	sql := "select `id`,`uid`,`oid`,`exp`,`enable_exp`,`use_exp`,`c_time`,`u_time` from jy_exp_detail where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyExpDetailOp) SelectAll() ([]*JyExpDetail, error) {
	objList := []*JyExpDetail{}
	sql := "select `id`,`uid`,`oid`,`exp`,`enable_exp`,`use_exp`,`c_time`,`u_time` from jy_exp_detail"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyExpDetailOp) QueryByMap(m map[string]interface{}) ([]*JyExpDetail, error) {
	result := []*JyExpDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`exp`,`enable_exp`,`use_exp`,`c_time`,`u_time` from jy_exp_detail where 1=1 "
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

func (op *jyExpDetailOp) QueryByMapComparison(m map[string]interface{}) ([]*JyExpDetail, error) {
	result := []*JyExpDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`exp`,`enable_exp`,`use_exp`,`c_time`,`u_time` from jy_exp_detail where 1=1 "
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

func (op *jyExpDetailOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyExpDetail, error) {
	result := []*JyExpDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`exp`,`enable_exp`,`use_exp`,`c_time`,`u_time` from jy_exp_detail where 1=1 "
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

func (op *jyExpDetailOp) GetByMap(m map[string]interface{}) (*JyExpDetail, error) {
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
func (op *jyExpDetailOp) Insert(m *JyExpDetail) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyExpDetailOp) InsertTx(ext sqlx.Ext, m *JyExpDetail) (int64, error) {
	sql := "insert into jy_exp_detail(uid,oid,exp,enable_exp,use_exp,c_time,u_time) values(?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Oid,
		m.Exp,
		m.EnableExp,
		m.UseExp,
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
func (i *JyExpDetail) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyExpDetailOp) Update(m *JyExpDetail) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyExpDetailOp) UpdateTx(ext sqlx.Ext, m *JyExpDetail) error {
	sql := `update jy_exp_detail set uid=?,oid=?,exp=?,enable_exp=?,use_exp=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Oid,
		m.Exp,
		m.EnableExp,
		m.UseExp,
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
func (op *jyExpDetailOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyExpDetailOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update jy_exp_detail set %s where 1=1 and id=? ;`

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
func (i *JyExpDetail) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyExpDetailOp) Delete(id int64) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyExpDetailOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from jy_exp_detail where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyExpDetailOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_exp_detail where 1=1 `
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

func (op *jyExpDetailOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyExpDetailOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_exp_detail where 1=1 "
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
