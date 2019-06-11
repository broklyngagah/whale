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

//jy_settlement_log
//交易结算日志表

// +gen *
type JySettlementLog struct {
	Id             int64   `db:"id" json:"id"`                           //
	Uid            int     `db:"uid" json:"uid"`                         // 用户ID
	BatchNo        string  `db:"batch_no" json:"batch_no"`               // 批次号
	OrderId        string  `db:"order_id" json:"order_id"`               // 订单编号
	TradeNo        string  `db:"trade_no" json:"trade_no"`               // 订单流水号(结算单号)
	PayMoney       float64 `db:"pay_money" json:"pay_money"`             //
	BalanceMoney   float64 `db:"balance_money" json:"balance_money"`     // 打款金额
	AccountingTime int64   `db:"accounting_time" json:"accounting_time"` // 记账操作时间
	AdminId        int     `db:"admin_id" json:"admin_id"`               // 记账人uid
	OptTime        int64   `db:"opt_time" json:"opt_time"`               // 当月时间YYYYMM，例如201712
	CTime          int64   `db:"c_time" json:"c_time"`                   // 创建时间
	UTime          int64   `db:"u_time" json:"u_time"`                   // 订单最后修改时间
}

type jySettlementLogOp struct{}

var JySettlementLogOp = &jySettlementLogOp{}
var DefaultJySettlementLog = &JySettlementLog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jySettlementLogOp) Get(id int64) (*JySettlementLog, error) {
	obj := &JySettlementLog{}
	sql := "select `id`,`uid`,`batch_no`,`order_id`,`trade_no`,`pay_money`,`balance_money`,`accounting_time`,`admin_id`,`opt_time`,`c_time`,`u_time` from jy_settlement_log where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jySettlementLogOp) SelectAll() ([]*JySettlementLog, error) {
	objList := []*JySettlementLog{}
	sql := "select `id`,`uid`,`batch_no`,`order_id`,`trade_no`,`pay_money`,`balance_money`,`accounting_time`,`admin_id`,`opt_time`,`c_time`,`u_time` from jy_settlement_log"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jySettlementLogOp) QueryByMap(m map[string]interface{}) ([]*JySettlementLog, error) {
	result := []*JySettlementLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`batch_no`,`order_id`,`trade_no`,`pay_money`,`balance_money`,`accounting_time`,`admin_id`,`opt_time`,`c_time`,`u_time` from jy_settlement_log where 1=1 "
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

func (op *jySettlementLogOp) QueryByMapComparison(m map[string]interface{}) ([]*JySettlementLog, error) {
	result := []*JySettlementLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`batch_no`,`order_id`,`trade_no`,`pay_money`,`balance_money`,`accounting_time`,`admin_id`,`opt_time`,`c_time`,`u_time` from jy_settlement_log where 1=1 "
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

func (op *jySettlementLogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JySettlementLog, error) {
	result := []*JySettlementLog{}
	var params []interface{}

	sql := "select `id`,`uid`,`batch_no`,`order_id`,`trade_no`,`pay_money`,`balance_money`,`accounting_time`,`admin_id`,`opt_time`,`c_time`,`u_time` from jy_settlement_log where 1=1 "
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

func (op *jySettlementLogOp) GetByMap(m map[string]interface{}) (*JySettlementLog, error) {
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
func (op *jySettlementLogOp) Insert(m *JySettlementLog) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jySettlementLogOp) InsertTx(ext sqlx.Ext, m *JySettlementLog) (int64, error) {
	sql := "insert into jy_settlement_log(uid,batch_no,order_id,trade_no,pay_money,balance_money,accounting_time,admin_id,opt_time,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.BatchNo,
		m.OrderId,
		m.TradeNo,
		m.PayMoney,
		m.BalanceMoney,
		m.AccountingTime,
		m.AdminId,
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
func (i *JySettlementLog) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jySettlementLogOp) Update(m *JySettlementLog) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jySettlementLogOp) UpdateTx(ext sqlx.Ext, m *JySettlementLog) error {
	sql := `update jy_settlement_log set uid=?,batch_no=?,order_id=?,trade_no=?,pay_money=?,balance_money=?,accounting_time=?,admin_id=?,opt_time=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.BatchNo,
		m.OrderId,
		m.TradeNo,
		m.PayMoney,
		m.BalanceMoney,
		m.AccountingTime,
		m.AdminId,
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
func (op *jySettlementLogOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jySettlementLogOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update jy_settlement_log set %s where 1=1 and id=? ;`

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
func (i *JySettlementLog) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jySettlementLogOp) Delete(id int64) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jySettlementLogOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from jy_settlement_log where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jySettlementLogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_settlement_log where 1=1 `
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

func (op *jySettlementLogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jySettlementLogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_settlement_log where 1=1 "
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
