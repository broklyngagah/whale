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

//user_integral_log
//用户领取积分日志

// +gen *
type UserIntegralLog struct {
	Id       int64   `db:"id" json:"id"`             //
	Uid      int     `db:"uid" json:"uid"`           // 用户ID
	Integral float64 `db:"integral" json:"integral"` // 积分值
	GetType  int8    `db:"get_type" json:"get_type"` // 积分类型（1:注册2首次加关注3关注大V 4给好友分享APP 5充值 6开户 7每日登录 8分享文章 9评论 10发帖）
	Type     int8    `db:"type" json:"type"`         // 任务类型（1:新手任务，2 每日任务）
	CTime    int64   `db:"c_time" json:"c_time"`     // 创建时间
	UTime    int64   `db:"u_time" json:"u_time"`     // 最后修改时间
}

type userIntegralLogOp struct{}

var UserIntegralLogOp = &userIntegralLogOp{}
var DefaultUserIntegralLog = &UserIntegralLog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userIntegralLogOp) Get(id int64) (*UserIntegralLog, error) {
	obj := &UserIntegralLog{}
	sql := "select `id`,`uid`,`integral`,`get_type`,`type`,`c_time`,`u_time` from user_integral_log where id=? "
	err := db.JYLogDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userIntegralLogOp) SelectAll() ([]*UserIntegralLog, error) {
	objList := []*UserIntegralLog{}
	sql := "select `id`,`uid`,`integral`,`get_type`,`type`,`c_time`,`u_time` from user_integral_log"
	err := db.JYLogDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userIntegralLogOp) QueryByMap(m map[string]interface{}) ([]*UserIntegralLog, error) {
	result := []*UserIntegralLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`integral`,`get_type`,`type`,`c_time`,`u_time` from user_integral_log where 1=1 "
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

func (op *userIntegralLogOp) QueryByMapComparison(m map[string]interface{}) ([]*UserIntegralLog, error) {
	result := []*UserIntegralLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`integral`,`get_type`,`type`,`c_time`,`u_time` from user_integral_log where 1=1 "
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

func (op *userIntegralLogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserIntegralLog, error) {
	result := []*UserIntegralLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`integral`,`get_type`,`type`,`c_time`,`u_time` from user_integral_log where 1=1 "
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

func (op *userIntegralLogOp) GetByMap(m map[string]interface{}) (*UserIntegralLog, error) {
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
func (op *userIntegralLogOp) Insert(m *UserIntegralLog) (int64, error) {
	return op.InsertTx(db.JYLogDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userIntegralLogOp) InsertTx(ext sqlx.Ext, m *UserIntegralLog) (int64, error) {
	sql := "insert into user_integral_log(uid,integral,get_type,type,c_time,u_time) values(?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Integral,
		m.GetType,
		m.Type,
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
func (i *UserIntegralLog) Update() {
    _,err := db.JYLogDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userIntegralLogOp) Update(m *UserIntegralLog) error {
	return op.UpdateTx(db.JYLogDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userIntegralLogOp) UpdateTx(ext sqlx.Ext, m *UserIntegralLog) error {
	sql := `update user_integral_log set uid=?,integral=?,get_type=?,type=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Integral,
		m.GetType,
		m.Type,
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
func (op *userIntegralLogOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYLogDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userIntegralLogOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update user_integral_log set %s where 1=1 and id=? ;`

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
func (i *UserIntegralLog) Delete(){
    _,err := db.JYLogDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userIntegralLogOp) Delete(id int64) error {
	return op.DeleteTx(db.JYLogDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userIntegralLogOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from user_integral_log where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userIntegralLogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_integral_log where 1=1 `
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

func (op *userIntegralLogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYLogDB, m)
}

func (op *userIntegralLogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_integral_log where 1=1 "
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
