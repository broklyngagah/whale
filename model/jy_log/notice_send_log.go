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

//notice_send_log
//推送日志表

// +gen *
type NoticeSendLog struct {
	Id          int    `db:"id" json:"id"`                     //
	Uid         int    `db:"uid" json:"uid"`                   // 用户id
	Tel         string `db:"tel" json:"tel"`                   // 手机号码，只有发短信才会用到
	OtherId     int    `db:"other_id" json:"other_id"`         // 文章ID、公告ID
	ContentType int8   `db:"content_type" json:"content_type"` // 1重要消息(全推) 2优质文章(全推) 3订阅更新(个推，订阅者)
	PushType    int8   `db:"push_type" json:"push_type"`       // 1极光推送 2短信推送
	PushTime    int64  `db:"push_time" json:"push_time"`       // 可推送时间
	Status      int8   `db:"status" json:"status"`             // 1未推送 2已推送
	PushKey     string `db:"push_key" json:"push_key"`         // md5(uid+content_type+day_time+aid或announce_id)唯一索引，day_time当天时间
	CTime       int64  `db:"c_time" json:"c_time"`             // 创建时间
	UTime       int64  `db:"u_time" json:"u_time"`             // 更新时间
}

type noticeSendLogOp struct{}

var NoticeSendLogOp = &noticeSendLogOp{}
var DefaultNoticeSendLog = &NoticeSendLog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *noticeSendLogOp) Get(id int) (*NoticeSendLog, error) {
	obj := &NoticeSendLog{}
	sql := "select `id`,`uid`,`tel`,`other_id`,`content_type`,`push_type`,`push_time`,`status`,`push_key`,`c_time`,`u_time` from notice_send_log where id=? "
	err := db.JYLogDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *noticeSendLogOp) SelectAll() ([]*NoticeSendLog, error) {
	objList := []*NoticeSendLog{}
	sql := "select `id`,`uid`,`tel`,`other_id`,`content_type`,`push_type`,`push_time`,`status`,`push_key`,`c_time`,`u_time` from notice_send_log"
	err := db.JYLogDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *noticeSendLogOp) QueryByMap(m map[string]interface{}) ([]*NoticeSendLog, error) {
	result := []*NoticeSendLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`tel`,`other_id`,`content_type`,`push_type`,`push_time`,`status`,`push_key`,`c_time`,`u_time` from notice_send_log where 1=1 "
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

func (op *noticeSendLogOp) QueryByMapComparison(m map[string]interface{}) ([]*NoticeSendLog, error) {
	result := []*NoticeSendLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`tel`,`other_id`,`content_type`,`push_type`,`push_time`,`status`,`push_key`,`c_time`,`u_time` from notice_send_log where 1=1 "
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

func (op *noticeSendLogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*NoticeSendLog, error) {
	result := []*NoticeSendLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`tel`,`other_id`,`content_type`,`push_type`,`push_time`,`status`,`push_key`,`c_time`,`u_time` from notice_send_log where 1=1 "
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

func (op *noticeSendLogOp) GetByMap(m map[string]interface{}) (*NoticeSendLog, error) {
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
func (op *noticeSendLogOp) Insert(m *NoticeSendLog) (int64, error) {
	return op.InsertTx(db.JYLogDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *noticeSendLogOp) InsertTx(ext sqlx.Ext, m *NoticeSendLog) (int64, error) {
	sql := "insert into notice_send_log(uid,tel,other_id,content_type,push_type,push_time,status,push_key,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Tel,
		m.OtherId,
		m.ContentType,
		m.PushType,
		m.PushTime,
		m.Status,
		m.PushKey,
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
func (i *NoticeSendLog) Update() {
    _,err := db.JYLogDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *noticeSendLogOp) Update(m *NoticeSendLog) error {
	return op.UpdateTx(db.JYLogDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *noticeSendLogOp) UpdateTx(ext sqlx.Ext, m *NoticeSendLog) error {
	sql := `update notice_send_log set uid=?,tel=?,other_id=?,content_type=?,push_type=?,push_time=?,status=?,push_key=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Tel,
		m.OtherId,
		m.ContentType,
		m.PushType,
		m.PushTime,
		m.Status,
		m.PushKey,
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
func (op *noticeSendLogOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYLogDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *noticeSendLogOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update notice_send_log set %s where 1=1 and id=? ;`

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
func (i *NoticeSendLog) Delete(){
    _,err := db.JYLogDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *noticeSendLogOp) Delete(id int) error {
	return op.DeleteTx(db.JYLogDB, id)
}

// 根据主键删除相关记录,Tx
func (op *noticeSendLogOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from notice_send_log where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *noticeSendLogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from notice_send_log where 1=1 `
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

func (op *noticeSendLogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYLogDB, m)
}

func (op *noticeSendLogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from notice_send_log where 1=1 "
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
