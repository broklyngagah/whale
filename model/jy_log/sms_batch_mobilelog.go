package JYLogDB

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

//sms_batch_mobilelog
//批量发送短信手机号记录

// +gen *
type SmsBatchMobilelog struct {
	Id      int    `db:"id" json:"id"`             //
	BatchId int    `db:"batch_id" json:"batch_id"` // 短信批量发送ID（sms_batch_sendlog表）
	Tel     string `db:"tel" json:"tel"`           // 手机号码
	Content string `db:"content" json:"content"`   // 短信内容
	CTime   int64  `db:"c_time" json:"c_time"`     // 创建时间
}

type smsBatchMobilelogOp struct{}

var SmsBatchMobilelogOp = &smsBatchMobilelogOp{}
var DefaultSmsBatchMobilelog = &SmsBatchMobilelog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *smsBatchMobilelogOp) Get(id int) (*SmsBatchMobilelog, error) {
	obj := &SmsBatchMobilelog{}
	sql := "select `id`,`batch_id`,`tel`,`content`,`c_time` from sms_batch_mobilelog where id=? "
	err := db.JYLogDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *smsBatchMobilelogOp) SelectAll() ([]*SmsBatchMobilelog, error) {
	objList := []*SmsBatchMobilelog{}
	sql := "select `id`,`batch_id`,`tel`,`content`,`c_time` from sms_batch_mobilelog"
	err := db.JYLogDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *smsBatchMobilelogOp) QueryByMap(m map[string]interface{}) ([]*SmsBatchMobilelog, error) {
	result := []*SmsBatchMobilelog{}
	var params []interface{}

	sql := "select `id`,`batch_id`,`tel`,`content`,`c_time` from sms_batch_mobilelog where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s=? ", k)
		params = append(params, v)
	}
	err := db.JYLogDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *smsBatchMobilelogOp) QueryByMapComparison(m map[string]interface{}) ([]*SmsBatchMobilelog, error) {
	result := []*SmsBatchMobilelog{}
	var params []interface{}

	sql := "select `id`,`batch_id`,`tel`,`content`,`c_time` from sms_batch_mobilelog where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s? ", k)
		params = append(params, v)
	}
	err := db.JYLogDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *smsBatchMobilelogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*SmsBatchMobilelog, error) {
	result := []*SmsBatchMobilelog{}
	var params []interface{}

	sql := "select `id`,`batch_id`,`tel`,`content`,`c_time` from sms_batch_mobilelog where 1=1 "
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

	err := db.JYLogDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *smsBatchMobilelogOp) GetByMap(m map[string]interface{}) (*SmsBatchMobilelog, error) {
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
func (op *smsBatchMobilelogOp) Insert(m *SmsBatchMobilelog) (int64, error) {
	return op.InsertTx(db.JYLogDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *smsBatchMobilelogOp) InsertTx(ext sqlx.Ext, m *SmsBatchMobilelog) (int64, error) {
	sql := "insert into sms_batch_mobilelog(batch_id,tel,content,c_time) values(?,?,?,?)"
	result, err := ext.Exec(sql,
		m.BatchId,
		m.Tel,
		m.Content,
		m.CTime,
	)
	if err != nil {
		game_error.RaiseError(err)
		return -1, err
	}
	affected, _ := result.RowsAffected()
	return affected, nil
}

/*
func (i *SmsBatchMobilelog) Update() {
    _,err := db.JYLogDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *smsBatchMobilelogOp) Update(m *SmsBatchMobilelog) error {
	return op.UpdateTx(db.JYLogDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *smsBatchMobilelogOp) UpdateTx(ext sqlx.Ext, m *SmsBatchMobilelog) error {
	sql := `update sms_batch_mobilelog set batch_id=?,tel=?,content=?,c_time=? where id=?`
	_, err := ext.Exec(sql,
		m.BatchId,
		m.Tel,
		m.Content,
		m.CTime,
		m.Id,
	)

	if err != nil {
		game_error.RaiseError(err)
		return err
	}

	return nil
}

// 用主键做条件，更新map里包含的字段名
func (op *smsBatchMobilelogOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYLogDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *smsBatchMobilelogOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update sms_batch_mobilelog set %s where 1=1 and id=? ;`

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
func (i *SmsBatchMobilelog) Delete(){
    _,err := db.JYLogDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *smsBatchMobilelogOp) Delete(id int) error {
	return op.DeleteTx(db.JYLogDB, id)
}

// 根据主键删除相关记录,Tx
func (op *smsBatchMobilelogOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from sms_batch_mobilelog where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *smsBatchMobilelogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from sms_batch_mobilelog where 1=1 `
	for k, v := range m {
		sql += fmt.Sprintf(" and  %s=? ", k)
		params = append(params, v)
	}
	count := int64(-1)
	err := db.JYLogDB.Get(&count, sql, params...)
	if err != nil {
		game_error.RaiseError(err)
	}
	return count
}

func (op *smsBatchMobilelogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYLogDB, m)
}

func (op *smsBatchMobilelogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from sms_batch_mobilelog where 1=1 "
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
