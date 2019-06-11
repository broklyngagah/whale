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

//jy_user_account_detail
//用户资产明细表

// +gen *
type JyUserAccountDetail struct {
	Id          int64   `db:"id" json:"id"`                     //
	Uid         int     `db:"uid" json:"uid"`                   // 用户ID
	Oid         int64   `db:"oid" json:"oid"`                   // 订单id
	AccountType int8    `db:"account_type" json:"account_type"` // 账户类型 1普通用户账户，2大V收入账户
	Money       float64 `db:"money" json:"money"`               // 充值金额
	EnableMoney float64 `db:"enable_money" json:"enable_money"` // 剩余金额
	UseMoney    float64 `db:"use_money" json:"use_money"`       // 已使用金额
	CTime       int64   `db:"c_time" json:"c_time"`             // 创建时间
	UTime       int64   `db:"u_time" json:"u_time"`             // 最后修改时间
}

type jyUserAccountDetailOp struct{}

var JyUserAccountDetailOp = &jyUserAccountDetailOp{}
var DefaultJyUserAccountDetail = &JyUserAccountDetail{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyUserAccountDetailOp) Get(id int64) (*JyUserAccountDetail, error) {
	obj := &JyUserAccountDetail{}
	sql := "select `id`,`uid`,`oid`,`account_type`,`money`,`enable_money`,`use_money`,`c_time`,`u_time` from jy_user_account_detail where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyUserAccountDetailOp) SelectAll() ([]*JyUserAccountDetail, error) {
	objList := []*JyUserAccountDetail{}
	sql := "select `id`,`uid`,`oid`,`account_type`,`money`,`enable_money`,`use_money`,`c_time`,`u_time` from jy_user_account_detail"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyUserAccountDetailOp) QueryByMap(m map[string]interface{}) ([]*JyUserAccountDetail, error) {
	result := []*JyUserAccountDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`account_type`,`money`,`enable_money`,`use_money`,`c_time`,`u_time` from jy_user_account_detail where 1=1 "
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

func (op *jyUserAccountDetailOp) QueryByMapComparison(m map[string]interface{}) ([]*JyUserAccountDetail, error) {
	result := []*JyUserAccountDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`account_type`,`money`,`enable_money`,`use_money`,`c_time`,`u_time` from jy_user_account_detail where 1=1 "
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

func (op *jyUserAccountDetailOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyUserAccountDetail, error) {
	result := []*JyUserAccountDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`account_type`,`money`,`enable_money`,`use_money`,`c_time`,`u_time` from jy_user_account_detail where 1=1 "
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

func (op *jyUserAccountDetailOp) GetByMap(m map[string]interface{}) (*JyUserAccountDetail, error) {
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
func (op *jyUserAccountDetailOp) Insert(m *JyUserAccountDetail) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyUserAccountDetailOp) InsertTx(ext sqlx.Ext, m *JyUserAccountDetail) (int64, error) {
	sql := "insert into jy_user_account_detail(uid,oid,account_type,money,enable_money,use_money,c_time,u_time) values(?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Oid,
		m.AccountType,
		m.Money,
		m.EnableMoney,
		m.UseMoney,
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
func (i *JyUserAccountDetail) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyUserAccountDetailOp) Update(m *JyUserAccountDetail) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyUserAccountDetailOp) UpdateTx(ext sqlx.Ext, m *JyUserAccountDetail) error {
	sql := `update jy_user_account_detail set uid=?,oid=?,account_type=?,money=?,enable_money=?,use_money=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Oid,
		m.AccountType,
		m.Money,
		m.EnableMoney,
		m.UseMoney,
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
func (op *jyUserAccountDetailOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyUserAccountDetailOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update jy_user_account_detail set %s where 1=1 and id=? ;`

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
func (i *JyUserAccountDetail) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyUserAccountDetailOp) Delete(id int64) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyUserAccountDetailOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from jy_user_account_detail where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyUserAccountDetailOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_user_account_detail where 1=1 `
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

func (op *jyUserAccountDetailOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyUserAccountDetailOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_user_account_detail where 1=1 "
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
