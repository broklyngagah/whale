package JYOtherDB

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

//jy_crontab
//定时任务表

// +gen *
type JyCrontab struct {
	Id      int   `db:"id" json:"id"`             // 自增ID
	Aid     int   `db:"aid" json:"aid"`           // 文章ID
	Type    int   `db:"type" json:"type"`         // 1定时提交审核2定时发布
	IsAdmin int8  `db:"is_admin" json:"is_admin"` // 是否管理员1是2否
	STime   int64 `db:"s_time" json:"s_time"`     // 触发时间
	Status  int8  `db:"status" json:"status"`     // 1已处理2未处理
	CTime   int64 `db:"c_time" json:"c_time"`     // 添加时间
	UTime   int64 `db:"u_time" json:"u_time"`     // 更新时间
}

type jyCrontabOp struct{}

var JyCrontabOp = &jyCrontabOp{}
var DefaultJyCrontab = &JyCrontab{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyCrontabOp) Get(id int) (*JyCrontab, error) {
	obj := &JyCrontab{}
	sql := "select `id`,`aid`,`type`,`is_admin`,`s_time`,`status`,`c_time`,`u_time` from jy_crontab where id=? "
	err := db.JYOtherDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyCrontabOp) SelectAll() ([]*JyCrontab, error) {
	objList := []*JyCrontab{}
	sql := "select `id`,`aid`,`type`,`is_admin`,`s_time`,`status`,`c_time`,`u_time` from jy_crontab"
	err := db.JYOtherDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyCrontabOp) QueryByMap(m map[string]interface{}) ([]*JyCrontab, error) {
	result := []*JyCrontab{}
	var params []interface{}

	sql := "select `id`,`aid`,`type`,`is_admin`,`s_time`,`status`,`c_time`,`u_time` from jy_crontab where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s=? ", k)
		params = append(params, v)
	}
	err := db.JYOtherDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyCrontabOp) QueryByMapComparison(m map[string]interface{}) ([]*JyCrontab, error) {
	result := []*JyCrontab{}
	var params []interface{}

	sql := "select `id`,`aid`,`type`,`is_admin`,`s_time`,`status`,`c_time`,`u_time` from jy_crontab where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s? ", k)
		params = append(params, v)
	}
	err := db.JYOtherDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyCrontabOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyCrontab, error) {
	result := []*JyCrontab{}
	var params []interface{}

	sql := "select `id`,`aid`,`type`,`is_admin`,`s_time`,`status`,`c_time`,`u_time` from jy_crontab where 1=1 "
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

	err := db.JYOtherDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyCrontabOp) GetByMap(m map[string]interface{}) (*JyCrontab, error) {
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
func (op *jyCrontabOp) Insert(m *JyCrontab) (int64, error) {
	return op.InsertTx(db.JYOtherDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyCrontabOp) InsertTx(ext sqlx.Ext, m *JyCrontab) (int64, error) {
	sql := "insert into jy_crontab(aid,type,is_admin,s_time,status,c_time,u_time) values(?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Aid,
		m.Type,
		m.IsAdmin,
		m.STime,
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
func (i *JyCrontab) Update() {
    _,err := db.JYOtherDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyCrontabOp) Update(m *JyCrontab) error {
	return op.UpdateTx(db.JYOtherDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyCrontabOp) UpdateTx(ext sqlx.Ext, m *JyCrontab) error {
	sql := `update jy_crontab set aid=?,type=?,is_admin=?,s_time=?,status=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Aid,
		m.Type,
		m.IsAdmin,
		m.STime,
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
func (op *jyCrontabOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYOtherDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyCrontabOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update jy_crontab set %s where 1=1 and id=? ;`

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
func (i *JyCrontab) Delete(){
    _,err := db.JYOtherDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyCrontabOp) Delete(id int) error {
	return op.DeleteTx(db.JYOtherDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyCrontabOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from jy_crontab where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyCrontabOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_crontab where 1=1 `
	for k, v := range m {
		sql += fmt.Sprintf(" and  %s=? ", k)
		params = append(params, v)
	}
	count := int64(-1)
	err := db.JYOtherDB.Get(&count, sql, params...)
	if err != nil {
		game_error.RaiseError(err)
	}
	return count
}

func (op *jyCrontabOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYOtherDB, m)
}

func (op *jyCrontabOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_crontab where 1=1 "
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
