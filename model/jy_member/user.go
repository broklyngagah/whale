package JYMemberDB

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

//user
//用户表

// +gen *
type User struct {
	Id       int    `db:"id" json:"id"`             // 用户ID
	Tel      string `db:"tel" json:"tel"`           // 手机号码
	OpenId   string `db:"open_id" json:"open_id"`   // 用户对外的标识
	Nickname string `db:"nickname" json:"nickname"` // 昵称
	Headimg  string `db:"headimg" json:"headimg"`   // 头像
	IsLock   int8   `db:"is_lock" json:"is_lock"`   // 是否锁定 1正常 2禁用 3账号异常 4禁言 5未激活(web端注册) 9销户
	Source   int8   `db:"source" json:"source"`     // 渠道:(1-web,2-360,3-91,4-baidu)
	Platform int8   `db:"platform" json:"platform"` // 平台类型设备:（1：ios；2：android；3：wap；4：PC
	Level    int8   `db:"level" json:"level"`       // 用户等级（1普通用户 10新手实习 15正式认证）
	CTime    int64  `db:"c_time" json:"c_time"`     // 创建时间
	UTime    int64  `db:"u_time" json:"u_time"`     // 修改时间
}

type userOp struct{}

var UserOp = &userOp{}
var DefaultUser = &User{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userOp) Get(id int) (*User, error) {
	obj := &User{}
	sql := "select `id`,`tel`,`open_id`,`nickname`,`headimg`,`is_lock`,`source`,`platform`,`level`,`c_time`,`u_time` from user where id=? "
	err := db.JYMemberDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userOp) SelectAll() ([]*User, error) {
	objList := []*User{}
	sql := "select `id`,`tel`,`open_id`,`nickname`,`headimg`,`is_lock`,`source`,`platform`,`level`,`c_time`,`u_time` from user"
	err := db.JYMemberDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userOp) QueryByMap(m map[string]interface{}) ([]*User, error) {
	result := []*User{}
	var params []interface{}

	sql := "select `id`,`tel`,`open_id`,`nickname`,`headimg`,`is_lock`,`source`,`platform`,`level`,`c_time`,`u_time` from user where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s=? ", k)
		params = append(params, v)
	}
	err := db.JYMemberDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *userOp) QueryByMapComparison(m map[string]interface{}) ([]*User, error) {
	result := []*User{}
	var params []interface{}

	sql := "select `id`,`tel`,`open_id`,`nickname`,`headimg`,`is_lock`,`source`,`platform`,`level`,`c_time`,`u_time` from user where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s? ", k)
		params = append(params, v)
	}
	err := db.JYMemberDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *userOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*User, error) {
	result := []*User{}
	var params []interface{}

	sql := "select `id`,`tel`,`open_id`,`nickname`,`headimg`,`is_lock`,`source`,`platform`,`level`,`c_time`,`u_time` from user where 1=1 "
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

	err := db.JYMemberDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *userOp) GetByMap(m map[string]interface{}) (*User, error) {
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
func (op *userOp) Insert(m *User) (int64, error) {
	return op.InsertTx(db.JYMemberDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userOp) InsertTx(ext sqlx.Ext, m *User) (int64, error) {
	sql := "insert into user(tel,open_id,nickname,headimg,is_lock,source,platform,level,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Tel,
		m.OpenId,
		m.Nickname,
		m.Headimg,
		m.IsLock,
		m.Source,
		m.Platform,
		m.Level,
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
func (i *User) Update() {
    _,err := db.JYMemberDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userOp) Update(m *User) error {
	return op.UpdateTx(db.JYMemberDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userOp) UpdateTx(ext sqlx.Ext, m *User) error {
	sql := `update user set tel=?,open_id=?,nickname=?,headimg=?,is_lock=?,source=?,platform=?,level=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Tel,
		m.OpenId,
		m.Nickname,
		m.Headimg,
		m.IsLock,
		m.Source,
		m.Platform,
		m.Level,
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
func (op *userOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYMemberDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update user set %s where 1=1 and id=? ;`

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
func (i *User) Delete(){
    _,err := db.JYMemberDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userOp) Delete(id int) error {
	return op.DeleteTx(db.JYMemberDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from user where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user where 1=1 `
	for k, v := range m {
		sql += fmt.Sprintf(" and  %s=? ", k)
		params = append(params, v)
	}
	count := int64(-1)
	err := db.JYMemberDB.Get(&count, sql, params...)
	if err != nil {
		game_error.RaiseError(err)
	}
	return count
}

func (op *userOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYMemberDB, m)
}

func (op *userOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user where 1=1 "
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
