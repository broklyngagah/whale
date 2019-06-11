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

//user_pay_conf
//用户购买订阅配置表

// +gen *
type UserPayConf struct {
	Id       int     `db:"id" json:"id"`               // id
	Uid      int     `db:"uid" json:"uid"`             // 用户ID
	Type     int8    `db:"type" json:"type"`           // 类型 1 一个月，3一季。6半年，12一年
	Name     string  `db:"name" json:"name"`           // 配置金额（1个月，3个月，半年，一年等）
	Days     int     `db:"days" json:"days"`           // 有效天数
	PayMoney float64 `db:"pay_money" json:"pay_money"` // 支付金额
	Status   int8    `db:"status" json:"status"`       // 1未审核 2已审核 9删除
	Sort     int8    `db:"sort" json:"sort"`           // 排序
	CTime    int64   `db:"c_time" json:"c_time"`       // 创建时间
	UTime    int64   `db:"u_time" json:"u_time"`       // 改修时间
}

type userPayConfOp struct{}

var UserPayConfOp = &userPayConfOp{}
var DefaultUserPayConf = &UserPayConf{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userPayConfOp) Get(id int) (*UserPayConf, error) {
	obj := &UserPayConf{}
	sql := "select `id`,`uid`,`type`,`name`,`days`,`pay_money`,`status`,`sort`,`c_time`,`u_time` from user_pay_conf where id=? "
	err := db.JYMemberDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userPayConfOp) SelectAll() ([]*UserPayConf, error) {
	objList := []*UserPayConf{}
	sql := "select `id`,`uid`,`type`,`name`,`days`,`pay_money`,`status`,`sort`,`c_time`,`u_time` from user_pay_conf"
	err := db.JYMemberDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userPayConfOp) QueryByMap(m map[string]interface{}) ([]*UserPayConf, error) {
	result := []*UserPayConf{}
	var params []interface{}

	sql := "select `id`,`uid`,`type`,`name`,`days`,`pay_money`,`status`,`sort`,`c_time`,`u_time` from user_pay_conf where 1=1 "
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

func (op *userPayConfOp) QueryByMapComparison(m map[string]interface{}) ([]*UserPayConf, error) {
	result := []*UserPayConf{}
	var params []interface{}

	sql := "select `id`,`uid`,`type`,`name`,`days`,`pay_money`,`status`,`sort`,`c_time`,`u_time` from user_pay_conf where 1=1 "
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

func (op *userPayConfOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserPayConf, error) {
	result := []*UserPayConf{}
	var params []interface{}

	sql := "select `id`,`uid`,`type`,`name`,`days`,`pay_money`,`status`,`sort`,`c_time`,`u_time` from user_pay_conf where 1=1 "
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

func (op *userPayConfOp) GetByMap(m map[string]interface{}) (*UserPayConf, error) {
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
func (op *userPayConfOp) Insert(m *UserPayConf) (int64, error) {
	return op.InsertTx(db.JYMemberDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userPayConfOp) InsertTx(ext sqlx.Ext, m *UserPayConf) (int64, error) {
	sql := "insert into user_pay_conf(uid,type,name,days,pay_money,status,sort,c_time,u_time) values(?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Type,
		m.Name,
		m.Days,
		m.PayMoney,
		m.Status,
		m.Sort,
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
func (i *UserPayConf) Update() {
    _,err := db.JYMemberDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userPayConfOp) Update(m *UserPayConf) error {
	return op.UpdateTx(db.JYMemberDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userPayConfOp) UpdateTx(ext sqlx.Ext, m *UserPayConf) error {
	sql := `update user_pay_conf set uid=?,type=?,name=?,days=?,pay_money=?,status=?,sort=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Type,
		m.Name,
		m.Days,
		m.PayMoney,
		m.Status,
		m.Sort,
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
func (op *userPayConfOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYMemberDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userPayConfOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update user_pay_conf set %s where 1=1 and id=? ;`

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
func (i *UserPayConf) Delete(){
    _,err := db.JYMemberDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userPayConfOp) Delete(id int) error {
	return op.DeleteTx(db.JYMemberDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userPayConfOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from user_pay_conf where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userPayConfOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_pay_conf where 1=1 `
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

func (op *userPayConfOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYMemberDB, m)
}

func (op *userPayConfOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_pay_conf where 1=1 "
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
