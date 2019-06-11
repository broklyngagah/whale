package JYTradeDB

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

//jy_pay_log
//第三方支付日志表

// +gen *
type JyPayLog struct {
	Id           int    `db:"id" json:"id"`                       //
	Uid          int    `db:"uid" json:"uid"`                     // 用户id
	OrderId      string `db:"order_id" json:"order_id"`           // 订单编号
	Msg          string `db:"msg" json:"msg"`                     // 描述
	OriginalText string `db:"original_text" json:"original_text"` // 原始报文
	EncryptText  string `db:"encrypt_text" json:"encrypt_text"`   // 加密报文
	PayType      int8   `db:"pay_type" json:"pay_type"`           // 通道 1微信 2 支付宝
	Type         int8   `db:"type" json:"type"`                   // 区分同一支付通道报文
	CTime        int64  `db:"c_time" json:"c_time"`               // 创建时间
}

type jyPayLogOp struct{}

var JyPayLogOp = &jyPayLogOp{}
var DefaultJyPayLog = &JyPayLog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyPayLogOp) Get(id int) (*JyPayLog, error) {
	obj := &JyPayLog{}
	sql := "select `id`,`uid`,`order_id`,`msg`,`original_text`,`encrypt_text`,`pay_type`,`type`,`c_time` from jy_pay_log where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyPayLogOp) SelectAll() ([]*JyPayLog, error) {
	objList := []*JyPayLog{}
	sql := "select `id`,`uid`,`order_id`,`msg`,`original_text`,`encrypt_text`,`pay_type`,`type`,`c_time` from jy_pay_log"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyPayLogOp) QueryByMap(m map[string]interface{}) ([]*JyPayLog, error) {
	result := []*JyPayLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`order_id`,`msg`,`original_text`,`encrypt_text`,`pay_type`,`type`,`c_time` from jy_pay_log where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s=? ", k)
		params = append(params, v)
	}
	err := db.JYTradeDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyPayLogOp) QueryByMapComparison(m map[string]interface{}) ([]*JyPayLog, error) {
	result := []*JyPayLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`order_id`,`msg`,`original_text`,`encrypt_text`,`pay_type`,`type`,`c_time` from jy_pay_log where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s? ", k)
		params = append(params, v)
	}
	err := db.JYTradeDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyPayLogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyPayLog, error) {
	result := []*JyPayLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`order_id`,`msg`,`original_text`,`encrypt_text`,`pay_type`,`type`,`c_time` from jy_pay_log where 1=1 "
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

	err := db.JYTradeDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyPayLogOp) GetByMap(m map[string]interface{}) (*JyPayLog, error) {
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
func (op *jyPayLogOp) Insert(m *JyPayLog) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyPayLogOp) InsertTx(ext sqlx.Ext, m *JyPayLog) (int64, error) {
	sql := "insert into jy_pay_log(uid,order_id,msg,original_text,encrypt_text,pay_type,type,c_time) values(?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.OrderId,
		m.Msg,
		m.OriginalText,
		m.EncryptText,
		m.PayType,
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
func (i *JyPayLog) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyPayLogOp) Update(m *JyPayLog) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyPayLogOp) UpdateTx(ext sqlx.Ext, m *JyPayLog) error {
	sql := `update jy_pay_log set uid=?,order_id=?,msg=?,original_text=?,encrypt_text=?,pay_type=?,type=?,c_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.OrderId,
		m.Msg,
		m.OriginalText,
		m.EncryptText,
		m.PayType,
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
func (op *jyPayLogOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyPayLogOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update jy_pay_log set %s where 1=1 and id=? ;`

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
func (i *JyPayLog) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyPayLogOp) Delete(id int) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyPayLogOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from jy_pay_log where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyPayLogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_pay_log where 1=1 `
	for k, v := range m {
		sql += fmt.Sprintf(" and  %s=? ", k)
		params = append(params, v)
	}
	count := int64(-1)
	err := db.JYTradeDB.Get(&count, sql, params...)
	if err != nil {
		game_error.RaiseError(err)
	}
	return count
}

func (op *jyPayLogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyPayLogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_pay_log where 1=1 "
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
