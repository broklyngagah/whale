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

//sys_trade_log
//交易错误日志

// +gen *
type SysTradeLog struct {
	Id          int    `db:"id" json:"id"`                     //
	Uid         int    `db:"uid" json:"uid"`                   // 用户ID
	Oid         int64  `db:"oid" json:"oid"`                   // 订单ID
	AccountType int8   `db:"account_type" json:"account_type"` // 1普通用户账户，2大V用户账户
	OrderType   int8   `db:"order_type" json:"order_type"`     // 转入【从这些类型 1现金充值到鱼币 2文章收入 3订阅收入 4现金购文自动转鱼币、转出到【51提现银行卡 52购买文章 53订阅用户】
	Platform    int8   `db:"platform" json:"platform"`         // 平台类型设备(1-ios;2-android;3-wap;4-PC;5微信游戏;6-ios回馈版)
	OtherType   int8   `db:"other_type" json:"other_type"`     // 其他类型 1正常流程  2购买文章 3订阅用户
	OtherId     int    `db:"other_id" json:"other_id"`         // 其他ID  (other_type=1等于jy_trade_conf 配置id，other_type=2时等于aid，other_type=3时等于user_pay_conf的id)
	Content     string `db:"content" json:"content"`           // 提交的参数
	ErrContent  string `db:"err_content" json:"err_content"`   // 错误信息
	ActionUrl   string `db:"action_url" json:"action_url"`     // 操作入口
	CTime       int64  `db:"c_time" json:"c_time"`             // 创建时间
	Status      int8   `db:"status" json:"status"`             // 1未处理  9已处理
}

type sysTradeLogOp struct{}

var SysTradeLogOp = &sysTradeLogOp{}
var DefaultSysTradeLog = &SysTradeLog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *sysTradeLogOp) Get(id int) (*SysTradeLog, error) {
	obj := &SysTradeLog{}
	sql := "select `id`,`uid`,`oid`,`account_type`,`order_type`,`platform`,`other_type`,`other_id`,`content`,`err_content`,`action_url`,`c_time`,`status` from sys_trade_log where id=? "
	err := db.JYLogDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *sysTradeLogOp) SelectAll() ([]*SysTradeLog, error) {
	objList := []*SysTradeLog{}
	sql := "select `id`,`uid`,`oid`,`account_type`,`order_type`,`platform`,`other_type`,`other_id`,`content`,`err_content`,`action_url`,`c_time`,`status` from sys_trade_log"
	err := db.JYLogDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *sysTradeLogOp) QueryByMap(m map[string]interface{}) ([]*SysTradeLog, error) {
	result := []*SysTradeLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`account_type`,`order_type`,`platform`,`other_type`,`other_id`,`content`,`err_content`,`action_url`,`c_time`,`status` from sys_trade_log where 1=1 "
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

func (op *sysTradeLogOp) QueryByMapComparison(m map[string]interface{}) ([]*SysTradeLog, error) {
	result := []*SysTradeLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`account_type`,`order_type`,`platform`,`other_type`,`other_id`,`content`,`err_content`,`action_url`,`c_time`,`status` from sys_trade_log where 1=1 "
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

func (op *sysTradeLogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*SysTradeLog, error) {
	result := []*SysTradeLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`oid`,`account_type`,`order_type`,`platform`,`other_type`,`other_id`,`content`,`err_content`,`action_url`,`c_time`,`status` from sys_trade_log where 1=1 "
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

func (op *sysTradeLogOp) GetByMap(m map[string]interface{}) (*SysTradeLog, error) {
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
func (op *sysTradeLogOp) Insert(m *SysTradeLog) (int64, error) {
	return op.InsertTx(db.JYLogDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *sysTradeLogOp) InsertTx(ext sqlx.Ext, m *SysTradeLog) (int64, error) {
	sql := "insert into sys_trade_log(uid,oid,account_type,order_type,platform,other_type,other_id,content,err_content,action_url,c_time,status) values(?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Oid,
		m.AccountType,
		m.OrderType,
		m.Platform,
		m.OtherType,
		m.OtherId,
		m.Content,
		m.ErrContent,
		m.ActionUrl,
		m.CTime,
		m.Status,
	)
	if err != nil {
		game_error.RaiseError(err)
		return -1, err
	}
	affected, _ := result.RowsAffected()
	return affected, nil
}

/*
func (i *SysTradeLog) Update() {
    _,err := db.JYLogDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *sysTradeLogOp) Update(m *SysTradeLog) error {
	return op.UpdateTx(db.JYLogDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *sysTradeLogOp) UpdateTx(ext sqlx.Ext, m *SysTradeLog) error {
	sql := `update sys_trade_log set uid=?,oid=?,account_type=?,order_type=?,platform=?,other_type=?,other_id=?,content=?,err_content=?,action_url=?,c_time=?,status=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Oid,
		m.AccountType,
		m.OrderType,
		m.Platform,
		m.OtherType,
		m.OtherId,
		m.Content,
		m.ErrContent,
		m.ActionUrl,
		m.CTime,
		m.Status,
		m.Id,
	)

	if err != nil {
		game_error.RaiseError(err)
		return err
	}

	return nil
}

// 用主键做条件，更新map里包含的字段名
func (op *sysTradeLogOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYLogDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *sysTradeLogOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update sys_trade_log set %s where 1=1 and id=? ;`

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
func (i *SysTradeLog) Delete(){
    _,err := db.JYLogDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *sysTradeLogOp) Delete(id int) error {
	return op.DeleteTx(db.JYLogDB, id)
}

// 根据主键删除相关记录,Tx
func (op *sysTradeLogOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from sys_trade_log where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *sysTradeLogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from sys_trade_log where 1=1 `
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

func (op *sysTradeLogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYLogDB, m)
}

func (op *sysTradeLogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from sys_trade_log where 1=1 "
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
