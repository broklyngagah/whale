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

//user_related
//与我相关中间表

// +gen *
type UserRelated struct {
	Id         int   `db:"id" json:"id"`                   //
	Uid        int   `db:"uid" json:"uid"`                 // 用户ID
	PassiveUid int   `db:"passive_uid" json:"passive_uid"` // 被关注者id
	Aid        int   `db:"aid" json:"aid"`                 // 文章ID
	Type       int8  `db:"type" json:"type"`               // 1文章，2帖子
	IsShow     int8  `db:"is_show" json:"is_show"`         // 1已阅读，0未阅读
	CTime      int64 `db:"c_time" json:"c_time"`           // 创建时间
}

type userRelatedOp struct{}

var UserRelatedOp = &userRelatedOp{}
var DefaultUserRelated = &UserRelated{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userRelatedOp) Get(id int) (*UserRelated, error) {
	obj := &UserRelated{}
	sql := "select `id`,`uid`,`passive_uid`,`aid`,`type`,`is_show`,`c_time` from user_related where id=? "
	err := db.JYMemberDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userRelatedOp) SelectAll() ([]*UserRelated, error) {
	objList := []*UserRelated{}
	sql := "select `id`,`uid`,`passive_uid`,`aid`,`type`,`is_show`,`c_time` from user_related"
	err := db.JYMemberDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userRelatedOp) QueryByMap(m map[string]interface{}) ([]*UserRelated, error) {
	result := []*UserRelated{}
	var params []interface{}

	sql := "select `id`,`uid`,`passive_uid`,`aid`,`type`,`is_show`,`c_time` from user_related where 1=1 "
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

func (op *userRelatedOp) QueryByMapComparison(m map[string]interface{}) ([]*UserRelated, error) {
	result := []*UserRelated{}
	var params []interface{}

	sql := "select `id`,`uid`,`passive_uid`,`aid`,`type`,`is_show`,`c_time` from user_related where 1=1 "
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

func (op *userRelatedOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserRelated, error) {
	result := []*UserRelated{}
	var params []interface{}

	sql := "select `id`,`uid`,`passive_uid`,`aid`,`type`,`is_show`,`c_time` from user_related where 1=1 "
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

func (op *userRelatedOp) GetByMap(m map[string]interface{}) (*UserRelated, error) {
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
func (op *userRelatedOp) Insert(m *UserRelated) (int64, error) {
	return op.InsertTx(db.JYMemberDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userRelatedOp) InsertTx(ext sqlx.Ext, m *UserRelated) (int64, error) {
	sql := "insert into user_related(uid,passive_uid,aid,type,is_show,c_time) values(?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.PassiveUid,
		m.Aid,
		m.Type,
		m.IsShow,
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
func (i *UserRelated) Update() {
    _,err := db.JYMemberDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userRelatedOp) Update(m *UserRelated) error {
	return op.UpdateTx(db.JYMemberDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userRelatedOp) UpdateTx(ext sqlx.Ext, m *UserRelated) error {
	sql := `update user_related set uid=?,passive_uid=?,aid=?,type=?,is_show=?,c_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.PassiveUid,
		m.Aid,
		m.Type,
		m.IsShow,
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
func (op *userRelatedOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYMemberDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userRelatedOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update user_related set %s where 1=1 and id=? ;`

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
func (i *UserRelated) Delete(){
    _,err := db.JYMemberDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userRelatedOp) Delete(id int) error {
	return op.DeleteTx(db.JYMemberDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userRelatedOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from user_related where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userRelatedOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_related where 1=1 `
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

func (op *userRelatedOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYMemberDB, m)
}

func (op *userRelatedOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_related where 1=1 "
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
