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

//user_liked
//用户点赞表

// +gen *
type UserLiked struct {
	Id         int   `db:"id" json:"id"`                   // id
	Uid        int   `db:"uid" json:"uid"`                 // 点赞者ID
	PassiveUid int   `db:"passive_uid" json:"passive_uid"` // 被点赞者ID
	Aid        int   `db:"aid" json:"aid"`                 // 文章帖子ID
	Cmid       int   `db:"cmid" json:"cmid"`               // 评论id
	Status     int8  `db:"status" json:"status"`           // 状态（1点赞中，9取消点赞）
	Type       int   `db:"type" json:"type"`               // 文章类型（1：文章2帖子 3评论）
	CTime      int64 `db:"c_time" json:"c_time"`           // 创建时间
}

type userLikedOp struct{}

var UserLikedOp = &userLikedOp{}
var DefaultUserLiked = &UserLiked{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userLikedOp) Get(id int) (*UserLiked, error) {
	obj := &UserLiked{}
	sql := "select `id`,`uid`,`passive_uid`,`aid`,`cmid`,`status`,`type`,`c_time` from user_liked where id=? "
	err := db.JYMemberDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userLikedOp) SelectAll() ([]*UserLiked, error) {
	objList := []*UserLiked{}
	sql := "select `id`,`uid`,`passive_uid`,`aid`,`cmid`,`status`,`type`,`c_time` from user_liked"
	err := db.JYMemberDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userLikedOp) QueryByMap(m map[string]interface{}) ([]*UserLiked, error) {
	result := []*UserLiked{}
	var params []interface{}

	sql := "select `id`,`uid`,`passive_uid`,`aid`,`cmid`,`status`,`type`,`c_time` from user_liked where 1=1 "
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

func (op *userLikedOp) QueryByMapComparison(m map[string]interface{}) ([]*UserLiked, error) {
	result := []*UserLiked{}
	var params []interface{}

	sql := "select `id`,`uid`,`passive_uid`,`aid`,`cmid`,`status`,`type`,`c_time` from user_liked where 1=1 "
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

func (op *userLikedOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserLiked, error) {
	result := []*UserLiked{}
	var params []interface{}

	sql := "select `id`,`uid`,`passive_uid`,`aid`,`cmid`,`status`,`type`,`c_time` from user_liked where 1=1 "
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

func (op *userLikedOp) GetByMap(m map[string]interface{}) (*UserLiked, error) {
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
func (op *userLikedOp) Insert(m *UserLiked) (int64, error) {
	return op.InsertTx(db.JYMemberDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userLikedOp) InsertTx(ext sqlx.Ext, m *UserLiked) (int64, error) {
	sql := "insert into user_liked(uid,passive_uid,aid,cmid,status,type,c_time) values(?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.PassiveUid,
		m.Aid,
		m.Cmid,
		m.Status,
		m.Type,
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
func (i *UserLiked) Update() {
    _,err := db.JYMemberDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userLikedOp) Update(m *UserLiked) error {
	return op.UpdateTx(db.JYMemberDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userLikedOp) UpdateTx(ext sqlx.Ext, m *UserLiked) error {
	sql := `update user_liked set uid=?,passive_uid=?,aid=?,cmid=?,status=?,type=?,c_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.PassiveUid,
		m.Aid,
		m.Cmid,
		m.Status,
		m.Type,
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
func (op *userLikedOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYMemberDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userLikedOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update user_liked set %s where 1=1 and id=? ;`

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
func (i *UserLiked) Delete(){
    _,err := db.JYMemberDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userLikedOp) Delete(id int) error {
	return op.DeleteTx(db.JYMemberDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userLikedOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from user_liked where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userLikedOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_liked where 1=1 `
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

func (op *userLikedOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYMemberDB, m)
}

func (op *userLikedOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_liked where 1=1 "
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
