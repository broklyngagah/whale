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

//user_login_log
//用户登录日志表

// +gen *
type UserLoginLog struct {
	Id         int    `db:"id" json:"id"`                   //
	Uid        int    `db:"uid" json:"uid"`                 // 用户ID user表ID
	Ip         string `db:"ip" json:"ip"`                   // 本次登录IP
	Platform   int8   `db:"platform" json:"platform"`       // 设备类型:（1：ios；2：android；3：wap；4：PC）
	AppVersion string `db:"app_version" json:"app_version"` // app版本号
	Source     int8   `db:"source" json:"source"`           // 平台（360,91,baidu，...)
	LoginTime  int64  `db:"login_time" json:"login_time"`   // 用户登录时间
	LoginUrl   string `db:"login_url" json:"login_url"`     // 登录入口
}

type userLoginLogOp struct{}

var UserLoginLogOp = &userLoginLogOp{}
var DefaultUserLoginLog = &UserLoginLog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userLoginLogOp) Get(id int) (*UserLoginLog, error) {
	obj := &UserLoginLog{}
	sql := "select `id`,`uid`,`ip`,`platform`,`app_version`,`source`,`login_time`,`login_url` from user_login_log where id=? "
	err := db.JYLogDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userLoginLogOp) SelectAll() ([]*UserLoginLog, error) {
	objList := []*UserLoginLog{}
	sql := "select `id`,`uid`,`ip`,`platform`,`app_version`,`source`,`login_time`,`login_url` from user_login_log"
	err := db.JYLogDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userLoginLogOp) QueryByMap(m map[string]interface{}) ([]*UserLoginLog, error) {
	result := []*UserLoginLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`ip`,`platform`,`app_version`,`source`,`login_time`,`login_url` from user_login_log where 1=1 "
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

func (op *userLoginLogOp) QueryByMapComparison(m map[string]interface{}) ([]*UserLoginLog, error) {
	result := []*UserLoginLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`ip`,`platform`,`app_version`,`source`,`login_time`,`login_url` from user_login_log where 1=1 "
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

func (op *userLoginLogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserLoginLog, error) {
	result := []*UserLoginLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`ip`,`platform`,`app_version`,`source`,`login_time`,`login_url` from user_login_log where 1=1 "
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

func (op *userLoginLogOp) GetByMap(m map[string]interface{}) (*UserLoginLog, error) {
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
func (op *userLoginLogOp) Insert(m *UserLoginLog) (int64, error) {
	return op.InsertTx(db.JYLogDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userLoginLogOp) InsertTx(ext sqlx.Ext, m *UserLoginLog) (int64, error) {
	sql := "insert into user_login_log(uid,ip,platform,app_version,source,login_time,login_url) values(?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Ip,
		m.Platform,
		m.AppVersion,
		m.Source,
		m.LoginTime,
		m.LoginUrl,
	)
	if err != nil {
		game_error.RaiseError(err)
		return -1, err
	}
	affected, _ := result.RowsAffected()
	return affected, nil
}

/*
func (i *UserLoginLog) Update() {
    _,err := db.JYLogDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userLoginLogOp) Update(m *UserLoginLog) error {
	return op.UpdateTx(db.JYLogDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userLoginLogOp) UpdateTx(ext sqlx.Ext, m *UserLoginLog) error {
	sql := `update user_login_log set uid=?,ip=?,platform=?,app_version=?,source=?,login_time=?,login_url=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Ip,
		m.Platform,
		m.AppVersion,
		m.Source,
		m.LoginTime,
		m.LoginUrl,
		m.Id,
	)

	if err != nil {
		game_error.RaiseError(err)
		return err
	}

	return nil
}

// 用主键做条件，更新map里包含的字段名
func (op *userLoginLogOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYLogDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userLoginLogOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update user_login_log set %s where 1=1 and id=? ;`

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
func (i *UserLoginLog) Delete(){
    _,err := db.JYLogDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userLoginLogOp) Delete(id int) error {
	return op.DeleteTx(db.JYLogDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userLoginLogOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from user_login_log where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userLoginLogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_login_log where 1=1 `
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

func (op *userLoginLogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYLogDB, m)
}

func (op *userLoginLogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_login_log where 1=1 "
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
