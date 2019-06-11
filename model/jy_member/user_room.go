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

//user_room
//直播房间表

// +gen *
type UserRoom struct {
	Id            int    `db:"id" json:"id"`                         // ID
	Uid           int    `db:"uid" json:"uid"`                       // 用户ID
	LoginLastip   string `db:"login_lastip" json:"login_lastip"`     // 最后登录IP
	LoginLasttime int64  `db:"login_lasttime" json:"login_lasttime"` // 最后登录时间
	Status        int8   `db:"status" json:"status"`                 // 是否开启1是2否
	CTime         int64  `db:"c_time" json:"c_time"`                 // 创建时间
	UTime         int64  `db:"u_time" json:"u_time"`                 // 更新时间
}

type userRoomOp struct{}

var UserRoomOp = &userRoomOp{}
var DefaultUserRoom = &UserRoom{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userRoomOp) Get(id int) (*UserRoom, error) {
	obj := &UserRoom{}
	sql := "select `id`,`uid`,`login_lastip`,`login_lasttime`,`status`,`c_time`,`u_time` from user_room where id=? "
	err := db.JYMemberDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userRoomOp) SelectAll() ([]*UserRoom, error) {
	objList := []*UserRoom{}
	sql := "select `id`,`uid`,`login_lastip`,`login_lasttime`,`status`,`c_time`,`u_time` from user_room"
	err := db.JYMemberDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userRoomOp) QueryByMap(m map[string]interface{}) ([]*UserRoom, error) {
	result := []*UserRoom{}
	var params []interface{}

	sql := "select `id`,`uid`,`login_lastip`,`login_lasttime`,`status`,`c_time`,`u_time` from user_room where 1=1 "
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

func (op *userRoomOp) QueryByMapComparison(m map[string]interface{}) ([]*UserRoom, error) {
	result := []*UserRoom{}
	var params []interface{}

	sql := "select `id`,`uid`,`login_lastip`,`login_lasttime`,`status`,`c_time`,`u_time` from user_room where 1=1 "
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

func (op *userRoomOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserRoom, error) {
	result := []*UserRoom{}
	var params []interface{}

	sql := "select `id`,`uid`,`login_lastip`,`login_lasttime`,`status`,`c_time`,`u_time` from user_room where 1=1 "
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

func (op *userRoomOp) GetByMap(m map[string]interface{}) (*UserRoom, error) {
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
func (op *userRoomOp) Insert(m *UserRoom) (int64, error) {
	return op.InsertTx(db.JYMemberDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userRoomOp) InsertTx(ext sqlx.Ext, m *UserRoom) (int64, error) {
	sql := "insert into user_room(uid,login_lastip,login_lasttime,status,c_time,u_time) values(?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.LoginLastip,
		m.LoginLasttime,
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
func (i *UserRoom) Update() {
    _,err := db.JYMemberDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userRoomOp) Update(m *UserRoom) error {
	return op.UpdateTx(db.JYMemberDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userRoomOp) UpdateTx(ext sqlx.Ext, m *UserRoom) error {
	sql := `update user_room set uid=?,login_lastip=?,login_lasttime=?,status=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.LoginLastip,
		m.LoginLasttime,
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
func (op *userRoomOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYMemberDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userRoomOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update user_room set %s where 1=1 and id=? ;`

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
func (i *UserRoom) Delete(){
    _,err := db.JYMemberDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userRoomOp) Delete(id int) error {
	return op.DeleteTx(db.JYMemberDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userRoomOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from user_room where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userRoomOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_room where 1=1 `
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

func (op *userRoomOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYMemberDB, m)
}

func (op *userRoomOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_room where 1=1 "
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
