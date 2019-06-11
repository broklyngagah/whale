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

//jy_integral_detail
//用户积分明细表

// +gen *
type JyIntegralDetail struct {
	Id             int64 `db:"id" json:"id"`                           //
	Uid            int   `db:"uid" json:"uid"`                         // 用户ID
	Oid            int64 `db:"oid" json:"oid"`                         // 订单id
	Integral       int   `db:"integral" json:"integral"`               // 积分值
	EnableIntegral int   `db:"enable_integral" json:"enable_integral"` // 剩余可使用积分
	UseIntegral    int   `db:"use_integral" json:"use_integral"`       // 已使用积分
	CTime          int64 `db:"c_time" json:"c_time"`                   // 创建时间
	UTime          int64 `db:"u_time" json:"u_time"`                   // 最后修改时间
}

type jyIntegralDetailOp struct{}

var JyIntegralDetailOp = &jyIntegralDetailOp{}
var DefaultJyIntegralDetail = &JyIntegralDetail{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyIntegralDetailOp) Get(id int64) (*JyIntegralDetail, error) {
	obj := &JyIntegralDetail{}
	sql := "select `id`,`uid`,`oid`,`integral`,`enable_integral`,`use_integral`,`c_time`,`u_time` from jy_integral_detail where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyIntegralDetailOp) SelectAll() ([]*JyIntegralDetail, error) {
	objList := []*JyIntegralDetail{}
	sql := "select `id`,`uid`,`oid`,`integral`,`enable_integral`,`use_integral`,`c_time`,`u_time` from jy_integral_detail"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyIntegralDetailOp) QueryByMap(m map[string]interface{}) ([]*JyIntegralDetail, error) {
	result := []*JyIntegralDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`integral`,`enable_integral`,`use_integral`,`c_time`,`u_time` from jy_integral_detail where 1=1 "
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

func (op *jyIntegralDetailOp) QueryByMapComparison(m map[string]interface{}) ([]*JyIntegralDetail, error) {
	result := []*JyIntegralDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`integral`,`enable_integral`,`use_integral`,`c_time`,`u_time` from jy_integral_detail where 1=1 "
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

func (op *jyIntegralDetailOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyIntegralDetail, error) {
	result := []*JyIntegralDetail{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`integral`,`enable_integral`,`use_integral`,`c_time`,`u_time` from jy_integral_detail where 1=1 "
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

func (op *jyIntegralDetailOp) GetByMap(m map[string]interface{}) (*JyIntegralDetail, error) {
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
func (op *jyIntegralDetailOp) Insert(m *JyIntegralDetail) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyIntegralDetailOp) InsertTx(ext sqlx.Ext, m *JyIntegralDetail) (int64, error) {
	sql := "insert into jy_integral_detail(uid,oid,integral,enable_integral,use_integral,c_time,u_time) values(?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Oid,
		m.Integral,
		m.EnableIntegral,
		m.UseIntegral,
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
func (i *JyIntegralDetail) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyIntegralDetailOp) Update(m *JyIntegralDetail) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyIntegralDetailOp) UpdateTx(ext sqlx.Ext, m *JyIntegralDetail) error {
	sql := `update jy_integral_detail set uid=?,oid=?,integral=?,enable_integral=?,use_integral=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Oid,
		m.Integral,
		m.EnableIntegral,
		m.UseIntegral,
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
func (op *jyIntegralDetailOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyIntegralDetailOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update jy_integral_detail set %s where 1=1 and id=? ;`

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
func (i *JyIntegralDetail) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyIntegralDetailOp) Delete(id int64) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyIntegralDetailOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from jy_integral_detail where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyIntegralDetailOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_integral_detail where 1=1 `
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

func (op *jyIntegralDetailOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyIntegralDetailOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_integral_detail where 1=1 "
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
