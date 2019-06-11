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

//user_feedback
//用户反馈表

// +gen *
type UserFeedback struct {
	Id         int    `db:"id" json:"id"`                 //
	Uid        int    `db:"uid" json:"uid"`               // 用户UID
	Contact    string `db:"contact" json:"contact"`       // 联系方式
	Content    string `db:"content" json:"content"`       // 内容
	Picture    string `db:"picture" json:"picture"`       // 图片
	Appversion string `db:"appversion" json:"appversion"` // 客户端版本
	Remark     string `db:"remark" json:"remark"`         // 管理员备注
	Status     int8   `db:"status" json:"status"`         // 1已处理2未处理
	OptTime    int64  `db:"opt_time" json:"opt_time"`     // 处理时间
	CTime      int64  `db:"c_time" json:"c_time"`         // 添加时间
	UTime      int64  `db:"u_time" json:"u_time"`         // 更新时间
}

type userFeedbackOp struct{}

var UserFeedbackOp = &userFeedbackOp{}
var DefaultUserFeedback = &UserFeedback{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userFeedbackOp) Get(id int) (*UserFeedback, error) {
	obj := &UserFeedback{}
	sql := "select `id`,`uid`,`contact`,`content`,`picture`,`appversion`,`remark`,`status`,`opt_time`,`c_time`,`u_time` from user_feedback where id=? "
	err := db.JYMemberDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userFeedbackOp) SelectAll() ([]*UserFeedback, error) {
	objList := []*UserFeedback{}
	sql := "select `id`,`uid`,`contact`,`content`,`picture`,`appversion`,`remark`,`status`,`opt_time`,`c_time`,`u_time` from user_feedback"
	err := db.JYMemberDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userFeedbackOp) QueryByMap(m map[string]interface{}) ([]*UserFeedback, error) {
	result := []*UserFeedback{}
	var params []interface{}

	sql := "select `id`,`uid`,`contact`,`content`,`picture`,`appversion`,`remark`,`status`,`opt_time`,`c_time`,`u_time` from user_feedback where 1=1 "
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

func (op *userFeedbackOp) QueryByMapComparison(m map[string]interface{}) ([]*UserFeedback, error) {
	result := []*UserFeedback{}
	var params []interface{}

	sql := "select `id`,`uid`,`contact`,`content`,`picture`,`appversion`,`remark`,`status`,`opt_time`,`c_time`,`u_time` from user_feedback where 1=1 "
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

func (op *userFeedbackOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserFeedback, error) {
	result := []*UserFeedback{}
	var params []interface{}

	sql := "select `id`,`uid`,`contact`,`content`,`picture`,`appversion`,`remark`,`status`,`opt_time`,`c_time`,`u_time` from user_feedback where 1=1 "
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

func (op *userFeedbackOp) GetByMap(m map[string]interface{}) (*UserFeedback, error) {
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
func (op *userFeedbackOp) Insert(m *UserFeedback) (int64, error) {
	return op.InsertTx(db.JYMemberDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userFeedbackOp) InsertTx(ext sqlx.Ext, m *UserFeedback) (int64, error) {
	sql := "insert into user_feedback(uid,contact,content,picture,appversion,remark,status,opt_time,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Contact,
		m.Content,
		m.Picture,
		m.Appversion,
		m.Remark,
		m.Status,
		m.OptTime,
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
func (i *UserFeedback) Update() {
    _,err := db.JYMemberDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userFeedbackOp) Update(m *UserFeedback) error {
	return op.UpdateTx(db.JYMemberDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userFeedbackOp) UpdateTx(ext sqlx.Ext, m *UserFeedback) error {
	sql := `update user_feedback set uid=?,contact=?,content=?,picture=?,appversion=?,remark=?,status=?,opt_time=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Contact,
		m.Content,
		m.Picture,
		m.Appversion,
		m.Remark,
		m.Status,
		m.OptTime,
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
func (op *userFeedbackOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYMemberDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userFeedbackOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update user_feedback set %s where 1=1 and id=? ;`

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
func (i *UserFeedback) Delete(){
    _,err := db.JYMemberDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userFeedbackOp) Delete(id int) error {
	return op.DeleteTx(db.JYMemberDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userFeedbackOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from user_feedback where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userFeedbackOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_feedback where 1=1 `
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

func (op *userFeedbackOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYMemberDB, m)
}

func (op *userFeedbackOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_feedback where 1=1 "
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
