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

//sys_log
//系统后台 - 操作日记

// +gen *
type SysLog struct {
	LogId       int    `db:"log_id" json:"log_id"`           //
	AdminId     int    `db:"admin_id" json:"admin_id"`       // 用户ID
	AdminName   string `db:"admin_name" json:"admin_name"`   // 用户姓名
	LogType     int8   `db:"log_type" json:"log_type"`       // 1.后台登录  2.系统管理 3.红包管理 4.交易管理 5.鲤鱼夺宝
	Description string `db:"description" json:"description"` // 描述
	Ip          string `db:"ip" json:"ip"`                   // IP地址
	Status      int8   `db:"status" json:"status"`           // 1 成功 2 失败
	CTime       int64  `db:"c_time" json:"c_time"`           // 添加时间
}

type sysLogOp struct{}

var SysLogOp = &sysLogOp{}
var DefaultSysLog = &SysLog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *sysLogOp) Get(log_id int) (*SysLog, error) {
	obj := &SysLog{}
	sql := "select `log_id`,`admin_id`,`admin_name`,`log_type`,`description`,`ip`,`status`,`c_time` from sys_log where log_id=? "
	err := db.JYLogDB.Get(obj, sql,
		log_id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *sysLogOp) SelectAll() ([]*SysLog, error) {
	objList := []*SysLog{}
	sql := "select `log_id`,`admin_id`,`admin_name`,`log_type`,`description`,`ip`,`status`,`c_time` from sys_log"
	err := db.JYLogDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *sysLogOp) QueryByMap(m map[string]interface{}) ([]*SysLog, error) {
	result := []*SysLog{}
	var params []interface{}

	sql := "select `log_id`,`admin_id`,`admin_name`,`log_type`,`description`,`ip`,`status`,`c_time` from sys_log where 1=1 "
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

func (op *sysLogOp) QueryByMapComparison(m map[string]interface{}) ([]*SysLog, error) {
	result := []*SysLog{}
	var params []interface{}

	sql := "select `log_id`,`admin_id`,`admin_name`,`log_type`,`description`,`ip`,`status`,`c_time` from sys_log where 1=1 "
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

func (op *sysLogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*SysLog, error) {
	result := []*SysLog{}
	var params []interface{}

	sql := "select `log_id`,`admin_id`,`admin_name`,`log_type`,`description`,`ip`,`status`,`c_time` from sys_log where 1=1 "
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

func (op *sysLogOp) GetByMap(m map[string]interface{}) (*SysLog, error) {
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
func (op *sysLogOp) Insert(m *SysLog) (int64, error) {
	return op.InsertTx(db.JYLogDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *sysLogOp) InsertTx(ext sqlx.Ext, m *SysLog) (int64, error) {
	sql := "insert into sys_log(admin_id,admin_name,log_type,description,ip,status,c_time) values(?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.AdminId,
		m.AdminName,
		m.LogType,
		m.Description,
		m.Ip,
		m.Status,
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
func (i *SysLog) Update() {
    _,err := db.JYLogDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *sysLogOp) Update(m *SysLog) error {
	return op.UpdateTx(db.JYLogDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *sysLogOp) UpdateTx(ext sqlx.Ext, m *SysLog) error {
	sql := `update sys_log set admin_id=?,admin_name=?,log_type=?,description=?,ip=?,status=?,c_time=? where log_id=?`
	_, err := ext.Exec(sql,
		m.AdminId,
		m.AdminName,
		m.LogType,
		m.Description,
		m.Ip,
		m.Status,
		m.CTime,
		m.LogId,
	)

	if err != nil {
		game_error.RaiseError(err)
		return err
	}

	return nil
}

// 用主键做条件，更新map里包含的字段名
func (op *sysLogOp) UpdateWithMap(log_id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYLogDB, log_id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *sysLogOp) UpdateWithMapTx(ext sqlx.Ext, log_id int, m map[string]interface{}) error {

	sql := `update sys_log set %s where 1=1 and log_id=? ;`

	var params []interface{}
	var set_sql string
	for k, v := range m {
		set_sql += fmt.Sprintf(" %s=? ", k)
		params = append(params, v)
	}
	params = append(params, log_id)
	_, err := ext.Exec(fmt.Sprintf(sql, set_sql), params...)
	return err
}

/*
func (i *SysLog) Delete(){
    _,err := db.JYLogDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *sysLogOp) Delete(log_id int) error {
	return op.DeleteTx(db.JYLogDB, log_id)
}

// 根据主键删除相关记录,Tx
func (op *sysLogOp) DeleteTx(ext sqlx.Ext, log_id int) error {
	sql := `delete from sys_log where 1=1
        and log_id=?
        `
	_, err := ext.Exec(sql,
		log_id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *sysLogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from sys_log where 1=1 `
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

func (op *sysLogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYLogDB, m)
}

func (op *sysLogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from sys_log where 1=1 "
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
