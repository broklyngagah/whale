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

//jy_integral_account_verification
//用户积分核销明细表(主要用于转出对应的哪些明细里扣)

// +gen *
type JyIntegralAccountVerification struct {
	Id           int64 `db:"id" json:"id"`                       //
	Uid          int   `db:"uid" json:"uid"`                     // 用户ID
	Oid          int64 `db:"oid" json:"oid"`                     // 订单id
	DetailId     int64 `db:"detail_id" json:"detail_id"`         // 积分明细表的ID
	CashIntegral int   `db:"cash_integral" json:"cash_integral"` // 从detail_id这条转出金额
	Status       int8  `db:"status" json:"status"`               // 1正常9取消
	CTime        int64 `db:"c_time" json:"c_time"`               // 创建时间
	UTime        int64 `db:"u_time" json:"u_time"`               // 最后修改时间
}

type jyIntegralAccountVerificationOp struct{}

var JyIntegralAccountVerificationOp = &jyIntegralAccountVerificationOp{}
var DefaultJyIntegralAccountVerification = &JyIntegralAccountVerification{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyIntegralAccountVerificationOp) Get(id int64) (*JyIntegralAccountVerification, error) {
	obj := &JyIntegralAccountVerification{}
	sql := "select `id`,`uid`,`oid`,`detail_id`,`cash_integral`,`status`,`c_time`,`u_time` from jy_integral_account_verification where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyIntegralAccountVerificationOp) SelectAll() ([]*JyIntegralAccountVerification, error) {
	objList := []*JyIntegralAccountVerification{}
	sql := "select `id`,`uid`,`oid`,`detail_id`,`cash_integral`,`status`,`c_time`,`u_time` from jy_integral_account_verification"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyIntegralAccountVerificationOp) QueryByMap(m map[string]interface{}) ([]*JyIntegralAccountVerification, error) {
	result := []*JyIntegralAccountVerification{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`detail_id`,`cash_integral`,`status`,`c_time`,`u_time` from jy_integral_account_verification where 1=1 "
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

func (op *jyIntegralAccountVerificationOp) QueryByMapComparison(m map[string]interface{}) ([]*JyIntegralAccountVerification, error) {
	result := []*JyIntegralAccountVerification{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`detail_id`,`cash_integral`,`status`,`c_time`,`u_time` from jy_integral_account_verification where 1=1 "
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

func (op *jyIntegralAccountVerificationOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyIntegralAccountVerification, error) {
	result := []*JyIntegralAccountVerification{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`detail_id`,`cash_integral`,`status`,`c_time`,`u_time` from jy_integral_account_verification where 1=1 "
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

func (op *jyIntegralAccountVerificationOp) GetByMap(m map[string]interface{}) (*JyIntegralAccountVerification, error) {
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
func (op *jyIntegralAccountVerificationOp) Insert(m *JyIntegralAccountVerification) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyIntegralAccountVerificationOp) InsertTx(ext sqlx.Ext, m *JyIntegralAccountVerification) (int64, error) {
	sql := "insert into jy_integral_account_verification(uid,oid,detail_id,cash_integral,status,c_time,u_time) values(?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Oid,
		m.DetailId,
		m.CashIntegral,
		m.Status,
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
func (i *JyIntegralAccountVerification) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyIntegralAccountVerificationOp) Update(m *JyIntegralAccountVerification) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyIntegralAccountVerificationOp) UpdateTx(ext sqlx.Ext, m *JyIntegralAccountVerification) error {
	sql := `update jy_integral_account_verification set uid=?,oid=?,detail_id=?,cash_integral=?,status=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Oid,
		m.DetailId,
		m.CashIntegral,
		m.Status,
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
func (op *jyIntegralAccountVerificationOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyIntegralAccountVerificationOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update jy_integral_account_verification set %s where 1=1 and id=? ;`

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
func (i *JyIntegralAccountVerification) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyIntegralAccountVerificationOp) Delete(id int64) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyIntegralAccountVerificationOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from jy_integral_account_verification where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyIntegralAccountVerificationOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_integral_account_verification where 1=1 `
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

func (op *jyIntegralAccountVerificationOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyIntegralAccountVerificationOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_integral_account_verification where 1=1 "
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
