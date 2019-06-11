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

//user_follow_relation
//用户关注关联表

// +gen *
type UserFollowRelation struct {
	Id        int64 `db:"id" json:"id"`                 //
	Uid       int   `db:"uid" json:"uid"`               // 用户ID
	FollowNum int   `db:"follow_num" json:"follow_num"` // 关注数量
	FansNum   int   `db:"fans_num" json:"fans_num"`     // 粉丝数量
	CTime     int64 `db:"c_time" json:"c_time"`         // 创建时间
	UTime     int64 `db:"u_time" json:"u_time"`         // 最后修改时间
}

type userFollowRelationOp struct{}

var UserFollowRelationOp = &userFollowRelationOp{}
var DefaultUserFollowRelation = &UserFollowRelation{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userFollowRelationOp) Get(id int64) (*UserFollowRelation, error) {
	obj := &UserFollowRelation{}
	sql := "select `id`,`uid`,`follow_num`,`fans_num`,`c_time`,`u_time` from user_follow_relation where id=? "
	err := db.JYMemberDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userFollowRelationOp) SelectAll() ([]*UserFollowRelation, error) {
	objList := []*UserFollowRelation{}
	sql := "select `id`,`uid`,`follow_num`,`fans_num`,`c_time`,`u_time` from user_follow_relation"
	err := db.JYMemberDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userFollowRelationOp) QueryByMap(m map[string]interface{}) ([]*UserFollowRelation, error) {
	result := []*UserFollowRelation{}
	var params []interface{}

	sql := "select `id`,`uid`,`follow_num`,`fans_num`,`c_time`,`u_time` from user_follow_relation where 1=1 "
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

func (op *userFollowRelationOp) QueryByMapComparison(m map[string]interface{}) ([]*UserFollowRelation, error) {
	result := []*UserFollowRelation{}
	var params []interface{}

	sql := "select `id`,`uid`,`follow_num`,`fans_num`,`c_time`,`u_time` from user_follow_relation where 1=1 "
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

func (op *userFollowRelationOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserFollowRelation, error) {
	result := []*UserFollowRelation{}
	var params []interface{}

	sql := "select `id`,`uid`,`follow_num`,`fans_num`,`c_time`,`u_time` from user_follow_relation where 1=1 "
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

func (op *userFollowRelationOp) GetByMap(m map[string]interface{}) (*UserFollowRelation, error) {
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
func (op *userFollowRelationOp) Insert(m *UserFollowRelation) (int64, error) {
	return op.InsertTx(db.JYMemberDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userFollowRelationOp) InsertTx(ext sqlx.Ext, m *UserFollowRelation) (int64, error) {
	sql := "insert into user_follow_relation(uid,follow_num,fans_num,c_time,u_time) values(?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.FollowNum,
		m.FansNum,
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
func (i *UserFollowRelation) Update() {
    _,err := db.JYMemberDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userFollowRelationOp) Update(m *UserFollowRelation) error {
	return op.UpdateTx(db.JYMemberDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userFollowRelationOp) UpdateTx(ext sqlx.Ext, m *UserFollowRelation) error {
	sql := `update user_follow_relation set uid=?,follow_num=?,fans_num=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.FollowNum,
		m.FansNum,
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
func (op *userFollowRelationOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYMemberDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userFollowRelationOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update user_follow_relation set %s where 1=1 and id=? ;`

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
func (i *UserFollowRelation) Delete(){
    _,err := db.JYMemberDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userFollowRelationOp) Delete(id int64) error {
	return op.DeleteTx(db.JYMemberDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userFollowRelationOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from user_follow_relation where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userFollowRelationOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_follow_relation where 1=1 `
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

func (op *userFollowRelationOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYMemberDB, m)
}

func (op *userFollowRelationOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_follow_relation where 1=1 "
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
