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

//jy_order
//交易订单表

// +gen *
type JyOrder struct {
	Id           int64   `db:"id" json:"id"`                       //
	Uid          int     `db:"uid" json:"uid"`                     // 用户ID
	OrderId      string  `db:"order_id" json:"order_id"`           // 订单编号
	TradeNo      string  `db:"trade_no" json:"trade_no"`           // 订单流水号
	AccountType  int8    `db:"account_type" json:"account_type"`   // 账户类型 1普通用户账户，2大V收入账户
	OrderType    int8    `db:"order_type" json:"order_type"`       // 转入【从这些类型 1现金充值到鱼币 2文章收入 3订阅收入 4现金购文自动转鱼币、转出到【51提现银行卡 52购买文章 53订阅用户】
	AccountMoney float64 `db:"account_money" json:"account_money"` // 账户发生金额money+discount_fee
	Money        float64 `db:"money" json:"money"`                 // 订单发生本金
	PayMoney     float64 `db:"pay_money" json:"pay_money"`         // 实际支付金额money+cash_fee
	DiscountFee  float64 `db:"discount_fee" json:"discount_fee"`   // 优惠金额
	CashFee      float64 `db:"cash_fee" json:"cash_fee"`           // 手续费或佣金（转入为正数，转出为负数）
	FeePercent   float64 `db:"fee_percent" json:"fee_percent"`     // 手续费或佣金比例
	PayType      int8    `db:"pay_type" json:"pay_type"`           // 支付方式 1微信 2支付宝
	PayTime      int64   `db:"pay_time" json:"pay_time"`           // 转入:支付时间 转出:提现或使用时间
	Status       int8    `db:"status" json:"status"`               // 1待付款 5已付款 9取消买入 ，11转出申请 15 转出成功 19取消转出
	OtherType    int8    `db:"other_type" json:"other_type"`       // 其他类型 1正常流程  2购买文章 3订阅用户
	OtherId      int     `db:"other_id" json:"other_id"`           // 其他ID  (other_type=1等于jy_trade_conf 配置id，other_type=2时等于aid，other_type=3时等于user_pay_conf的id)
	Platform     int8    `db:"platform" json:"platform"`           // 平台类型设备(1-ios;2-android;3-wap;4-PC
	Ip           string  `db:"ip" json:"ip"`                       // 用户IP地址
	CTime        int64   `db:"c_time" json:"c_time"`               // 创建时间
	UTime        int64   `db:"u_time" json:"u_time"`               // 订单最后修改时间
}

type jyOrderOp struct{}

var JyOrderOp = &jyOrderOp{}
var DefaultJyOrder = &JyOrder{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyOrderOp) Get(id int64) (*JyOrder, error) {
	obj := &JyOrder{}
	sql := "select `id`,`uid`,`order_id`,`trade_no`,`account_type`,`order_type`,`account_money`,`money`,`pay_money`,`discount_fee`,`cash_fee`,`fee_percent`,`pay_type`,`pay_time`,`status`,`other_type`,`other_id`,`platform`,`ip`,`c_time`,`u_time` from jy_order where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyOrderOp) SelectAll() ([]*JyOrder, error) {
	objList := []*JyOrder{}
	sql := "select `id`,`uid`,`order_id`,`trade_no`,`account_type`,`order_type`,`account_money`,`money`,`pay_money`,`discount_fee`,`cash_fee`,`fee_percent`,`pay_type`,`pay_time`,`status`,`other_type`,`other_id`,`platform`,`ip`,`c_time`,`u_time` from jy_order"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyOrderOp) QueryByMap(m map[string]interface{}) ([]*JyOrder, error) {
	result := []*JyOrder{}
	var params []interface{}

	sql := "select `id`,`uid`,`order_id`,`trade_no`,`account_type`,`order_type`,`account_money`,`money`,`pay_money`,`discount_fee`,`cash_fee`,`fee_percent`,`pay_type`,`pay_time`,`status`,`other_type`,`other_id`,`platform`,`ip`,`c_time`,`u_time` from jy_order where 1=1 "
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

func (op *jyOrderOp) QueryByMapComparison(m map[string]interface{}) ([]*JyOrder, error) {
	result := []*JyOrder{}
	var params []interface{}

	sql := "select `id`,`uid`,`order_id`,`trade_no`,`account_type`,`order_type`,`account_money`,`money`,`pay_money`,`discount_fee`,`cash_fee`,`fee_percent`,`pay_type`,`pay_time`,`status`,`other_type`,`other_id`,`platform`,`ip`,`c_time`,`u_time` from jy_order where 1=1 "
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

func (op *jyOrderOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyOrder, error) {
	result := []*JyOrder{}
	var params []interface{}

	sql := "select `id`,`uid`,`order_id`,`trade_no`,`account_type`,`order_type`,`account_money`,`money`,`pay_money`,`discount_fee`,`cash_fee`,`fee_percent`,`pay_type`,`pay_time`,`status`,`other_type`,`other_id`,`platform`,`ip`,`c_time`,`u_time` from jy_order where 1=1 "
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

func (op *jyOrderOp) GetByMap(m map[string]interface{}) (*JyOrder, error) {
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
func (op *jyOrderOp) Insert(m *JyOrder) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyOrderOp) InsertTx(ext sqlx.Ext, m *JyOrder) (int64, error) {
	sql := "insert into jy_order(uid,order_id,trade_no,account_type,order_type,account_money,money,pay_money,discount_fee,cash_fee,fee_percent,pay_type,pay_time,status,other_type,other_id,platform,ip,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.OrderId,
		m.TradeNo,
		m.AccountType,
		m.OrderType,
		m.AccountMoney,
		m.Money,
		m.PayMoney,
		m.DiscountFee,
		m.CashFee,
		m.FeePercent,
		m.PayType,
		m.PayTime,
		m.Status,
		m.OtherType,
		m.OtherId,
		m.Platform,
		m.Ip,
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
func (i *JyOrder) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyOrderOp) Update(m *JyOrder) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyOrderOp) UpdateTx(ext sqlx.Ext, m *JyOrder) error {
	sql := `update jy_order set uid=?,order_id=?,trade_no=?,account_type=?,order_type=?,account_money=?,money=?,pay_money=?,discount_fee=?,cash_fee=?,fee_percent=?,pay_type=?,pay_time=?,status=?,other_type=?,other_id=?,platform=?,ip=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.OrderId,
		m.TradeNo,
		m.AccountType,
		m.OrderType,
		m.AccountMoney,
		m.Money,
		m.PayMoney,
		m.DiscountFee,
		m.CashFee,
		m.FeePercent,
		m.PayType,
		m.PayTime,
		m.Status,
		m.OtherType,
		m.OtherId,
		m.Platform,
		m.Ip,
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
func (op *jyOrderOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyOrderOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update jy_order set %s where 1=1 and id=? ;`

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
func (i *JyOrder) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyOrderOp) Delete(id int64) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyOrderOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from jy_order where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyOrderOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_order where 1=1 `
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

func (op *jyOrderOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyOrderOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_order where 1=1 "
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
