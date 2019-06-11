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

//chat_message_log
//正在直播的文字信息

// +gen *
type ChatMessageLog struct {
	Id      int    `db:"id" json:"id"`             // ID
	Uid     int    `db:"uid" json:"uid"`           // 用户ID
	Secret  int8   `db:"secret" json:"secret"`     // 私密消息1是2否
	RoomId  int    `db:"room_id" json:"room_id"`   // 房间ID
	ReplyId int    `db:"reply_id" json:"reply_id"` // 被回复的那条记录ID
	Status  int8   `db:"status" json:"status"`     //
	Content string `db:"content" json:"content"`   // 只存文字
	CTime   int64  `db:"c_time" json:"c_time"`     // 创建时间
	UTime   int64  `db:"u_time" json:"u_time"`     // 更新时间
}

type chatMessageLogOp struct{}

var ChatMessageLogOp = &chatMessageLogOp{}
var DefaultChatMessageLog = &ChatMessageLog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *chatMessageLogOp) Get(id int) (*ChatMessageLog, error) {
	obj := &ChatMessageLog{}
	sql := "select `id`,`uid`,`secret`,`room_id`,`reply_id`,`status`,`content`,`c_time`,`u_time` from chat_message_log where id=? "
	err := db.JYLogDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *chatMessageLogOp) SelectAll() ([]*ChatMessageLog, error) {
	objList := []*ChatMessageLog{}
	sql := "select `id`,`uid`,`secret`,`room_id`,`reply_id`,`status`,`content`,`c_time`,`u_time` from chat_message_log"
	err := db.JYLogDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *chatMessageLogOp) QueryByMap(m map[string]interface{}) ([]*ChatMessageLog, error) {
	result := []*ChatMessageLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`secret`,`room_id`,`reply_id`,`status`,`content`,`c_time`,`u_time` from chat_message_log where 1=1 "
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

func (op *chatMessageLogOp) QueryByMapComparison(m map[string]interface{}) ([]*ChatMessageLog, error) {
	result := []*ChatMessageLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`secret`,`room_id`,`reply_id`,`status`,`content`,`c_time`,`u_time` from chat_message_log where 1=1 "
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

func (op *chatMessageLogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*ChatMessageLog, error) {
	result := []*ChatMessageLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`secret`,`room_id`,`reply_id`,`status`,`content`,`c_time`,`u_time` from chat_message_log where 1=1 "
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

func (op *chatMessageLogOp) GetByMap(m map[string]interface{}) (*ChatMessageLog, error) {
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
func (op *chatMessageLogOp) Insert(m *ChatMessageLog) (int64, error) {
	return op.InsertTx(db.JYLogDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *chatMessageLogOp) InsertTx(ext sqlx.Ext, m *ChatMessageLog) (int64, error) {
	sql := "insert into chat_message_log(uid,secret,room_id,reply_id,status,content,c_time,u_time) values(?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Secret,
		m.RoomId,
		m.ReplyId,
		m.Status,
		m.Content,
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
func (i *ChatMessageLog) Update() {
    _,err := db.JYLogDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *chatMessageLogOp) Update(m *ChatMessageLog) error {
	return op.UpdateTx(db.JYLogDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *chatMessageLogOp) UpdateTx(ext sqlx.Ext, m *ChatMessageLog) error {
	sql := `update chat_message_log set uid=?,secret=?,room_id=?,reply_id=?,status=?,content=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Secret,
		m.RoomId,
		m.ReplyId,
		m.Status,
		m.Content,
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
func (op *chatMessageLogOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYLogDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *chatMessageLogOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update chat_message_log set %s where 1=1 and id=? ;`

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
func (i *ChatMessageLog) Delete(){
    _,err := db.JYLogDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *chatMessageLogOp) Delete(id int) error {
	return op.DeleteTx(db.JYLogDB, id)
}

// 根据主键删除相关记录,Tx
func (op *chatMessageLogOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from chat_message_log where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *chatMessageLogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from chat_message_log where 1=1 `
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

func (op *chatMessageLogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYLogDB, m)
}

func (op *chatMessageLogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from chat_message_log where 1=1 "
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
