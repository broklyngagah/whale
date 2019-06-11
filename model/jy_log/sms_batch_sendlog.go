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

//sms_batch_sendlog
//批量发送短信日志

// +gen *
type SmsBatchSendlog struct {
	Id           int    `db:"id" json:"id"`                       //
	TplId        int    `db:"tpl_id" json:"tpl_id"`               // 短信模板ID
	RCode        string `db:"r_code" json:"r_code"`               // 回调code
	RDesc        string `db:"r_desc" json:"r_desc"`               // 回调内容
	MsgId        string `db:"msg_id" json:"msg_id"`               // 第三方短信发送ID
	ChannelId    int    `db:"channel_id" json:"channel_id"`       // 短信通道标识
	ChannelTitle string `db:"channel_title" json:"channel_title"` // 通道名称
	Sender       string `db:"sender" json:"sender"`               // 发送者
	Status       int8   `db:"status" json:"status"`               // 状态  1成功，2失败
	STime        int64  `db:"s_time" json:"s_time"`               // 设置发送时间
	CTime        int64  `db:"c_time" json:"c_time"`               // 创建时间
}

type smsBatchSendlogOp struct{}

var SmsBatchSendlogOp = &smsBatchSendlogOp{}
var DefaultSmsBatchSendlog = &SmsBatchSendlog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *smsBatchSendlogOp) Get(id int) (*SmsBatchSendlog, error) {
	obj := &SmsBatchSendlog{}
	sql := "select `id`,`tpl_id`,`r_code`,`r_desc`,`msg_id`,`channel_id`,`channel_title`,`sender`,`status`,`s_time`,`c_time` from sms_batch_sendlog where id=? "
	err := db.JYLogDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *smsBatchSendlogOp) SelectAll() ([]*SmsBatchSendlog, error) {
	objList := []*SmsBatchSendlog{}
	sql := "select `id`,`tpl_id`,`r_code`,`r_desc`,`msg_id`,`channel_id`,`channel_title`,`sender`,`status`,`s_time`,`c_time` from sms_batch_sendlog"
	err := db.JYLogDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *smsBatchSendlogOp) QueryByMap(m map[string]interface{}) ([]*SmsBatchSendlog, error) {
	result := []*SmsBatchSendlog{}
	var params []interface{}

	sql := "select `id`,`tpl_id`,`r_code`,`r_desc`,`msg_id`,`channel_id`,`channel_title`,`sender`,`status`,`s_time`,`c_time` from sms_batch_sendlog where 1=1 "
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

func (op *smsBatchSendlogOp) QueryByMapComparison(m map[string]interface{}) ([]*SmsBatchSendlog, error) {
	result := []*SmsBatchSendlog{}
	var params []interface{}

	sql := "select `id`,`tpl_id`,`r_code`,`r_desc`,`msg_id`,`channel_id`,`channel_title`,`sender`,`status`,`s_time`,`c_time` from sms_batch_sendlog where 1=1 "
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

func (op *smsBatchSendlogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*SmsBatchSendlog, error) {
	result := []*SmsBatchSendlog{}
	var params []interface{}

	sql := "select `id`,`tpl_id`,`r_code`,`r_desc`,`msg_id`,`channel_id`,`channel_title`,`sender`,`status`,`s_time`,`c_time` from sms_batch_sendlog where 1=1 "
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

func (op *smsBatchSendlogOp) GetByMap(m map[string]interface{}) (*SmsBatchSendlog, error) {
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
func (op *smsBatchSendlogOp) Insert(m *SmsBatchSendlog) (int64, error) {
	return op.InsertTx(db.JYLogDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *smsBatchSendlogOp) InsertTx(ext sqlx.Ext, m *SmsBatchSendlog) (int64, error) {
	sql := "insert into sms_batch_sendlog(tpl_id,r_code,r_desc,msg_id,channel_id,channel_title,sender,status,s_time,c_time) values(?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.TplId,
		m.RCode,
		m.RDesc,
		m.MsgId,
		m.ChannelId,
		m.ChannelTitle,
		m.Sender,
		m.Status,
		m.STime,
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
func (i *SmsBatchSendlog) Update() {
    _,err := db.JYLogDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *smsBatchSendlogOp) Update(m *SmsBatchSendlog) error {
	return op.UpdateTx(db.JYLogDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *smsBatchSendlogOp) UpdateTx(ext sqlx.Ext, m *SmsBatchSendlog) error {
	sql := `update sms_batch_sendlog set tpl_id=?,r_code=?,r_desc=?,msg_id=?,channel_id=?,channel_title=?,sender=?,status=?,s_time=?,c_time=? where id=?`
	_, err := ext.Exec(sql,
		m.TplId,
		m.RCode,
		m.RDesc,
		m.MsgId,
		m.ChannelId,
		m.ChannelTitle,
		m.Sender,
		m.Status,
		m.STime,
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
func (op *smsBatchSendlogOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYLogDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *smsBatchSendlogOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update sms_batch_sendlog set %s where 1=1 and id=? ;`

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
func (i *SmsBatchSendlog) Delete(){
    _,err := db.JYLogDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *smsBatchSendlogOp) Delete(id int) error {
	return op.DeleteTx(db.JYLogDB, id)
}

// 根据主键删除相关记录,Tx
func (op *smsBatchSendlogOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from sms_batch_sendlog where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *smsBatchSendlogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from sms_batch_sendlog where 1=1 `
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

func (op *smsBatchSendlogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYLogDB, m)
}

func (op *smsBatchSendlogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from sms_batch_sendlog where 1=1 "
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
