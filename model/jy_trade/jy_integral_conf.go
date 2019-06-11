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

//jy_integral_conf
//积分配置表

// +gen *
type JyIntegralConf struct {
	Id       int8   `db:"id" json:"id"`               // id
	GetType  int8   `db:"get_type" json:"get_type"`   // 积分类型（小于50为新手任务，大于50为每日任务）
	EventKey string `db:"event_key" json:"event_key"` // 对应位置key
	Name     string `db:"name" json:"name"`           // 积分任务标题
	Desc     string `db:"desc" json:"desc"`           // 描述
	Value    int    `db:"value" json:"value"`         // 积分值
	Type     int8   `db:"type" json:"type"`           // 任务类型（1:新手任务，2 每日任务）
	ImgUrl   string `db:"img_url" json:"img_url"`     // icon图标
	Status   int8   `db:"status" json:"status"`       // 状态 1禁用2启用
	Sort     int8   `db:"sort" json:"sort"`           // 排序
	CTime    int64  `db:"c_time" json:"c_time"`       // 创建时间
	UTime    int64  `db:"u_time" json:"u_time"`       // 改修时间
}

type jyIntegralConfOp struct{}

var JyIntegralConfOp = &jyIntegralConfOp{}
var DefaultJyIntegralConf = &JyIntegralConf{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyIntegralConfOp) Get(id int8) (*JyIntegralConf, error) {
	obj := &JyIntegralConf{}
	sql := "select `id`,`get_type`,`event_key`,`name`,`desc`,`value`,`type`,`img_url`,`status`,`sort`,`c_time`,`u_time` from jy_integral_conf where id=? "
	err := db.JYTradeDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyIntegralConfOp) SelectAll() ([]*JyIntegralConf, error) {
	objList := []*JyIntegralConf{}
	sql := "select `id`,`get_type`,`event_key`,`name`,`desc`,`value`,`type`,`img_url`,`status`,`sort`,`c_time`,`u_time` from jy_integral_conf"
	err := db.JYTradeDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyIntegralConfOp) QueryByMap(m map[string]interface{}) ([]*JyIntegralConf, error) {
	result := []*JyIntegralConf{}
	var params []interface{}

	sql := "select `id`,`get_type`,`event_key`,`name`,`desc`,`value`,`type`,`img_url`,`status`,`sort`,`c_time`,`u_time` from jy_integral_conf where 1=1 "
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

func (op *jyIntegralConfOp) QueryByMapComparison(m map[string]interface{}) ([]*JyIntegralConf, error) {
	result := []*JyIntegralConf{}
	var params []interface{}

	sql := "select `id`,`get_type`,`event_key`,`name`,`desc`,`value`,`type`,`img_url`,`status`,`sort`,`c_time`,`u_time` from jy_integral_conf where 1=1 "
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

func (op *jyIntegralConfOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyIntegralConf, error) {
	result := []*JyIntegralConf{}
	var params []interface{}

	sql := "select `id`,`get_type`,`event_key`,`name`,`desc`,`value`,`type`,`img_url`,`status`,`sort`,`c_time`,`u_time` from jy_integral_conf where 1=1 "
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

func (op *jyIntegralConfOp) GetByMap(m map[string]interface{}) (*JyIntegralConf, error) {
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
func (op *jyIntegralConfOp) Insert(m *JyIntegralConf) (int64, error) {
	return op.InsertTx(db.JYTradeDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyIntegralConfOp) InsertTx(ext sqlx.Ext, m *JyIntegralConf) (int64, error) {
	sql := "insert into jy_integral_conf(get_type,event_key,name,desc,value,type,img_url,status,sort,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.GetType,
		m.EventKey,
		m.Name,
		m.Desc,
		m.Value,
		m.Type,
		m.ImgUrl,
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
func (i *JyIntegralConf) Update() {
    _,err := db.JYTradeDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyIntegralConfOp) Update(m *JyIntegralConf) error {
	return op.UpdateTx(db.JYTradeDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyIntegralConfOp) UpdateTx(ext sqlx.Ext, m *JyIntegralConf) error {
	sql := `update jy_integral_conf set get_type=?,event_key=?,name=?,desc=?,value=?,type=?,img_url=?,status=?,sort=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.GetType,
		m.EventKey,
		m.Name,
		m.Desc,
		m.Value,
		m.Type,
		m.ImgUrl,
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
func (op *jyIntegralConfOp) UpdateWithMap(id int8, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYTradeDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyIntegralConfOp) UpdateWithMapTx(ext sqlx.Ext, id int8, m map[string]interface{}) error {

	sql := `update jy_integral_conf set %s where 1=1 and id=? ;`

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
func (i *JyIntegralConf) Delete(){
    _,err := db.JYTradeDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyIntegralConfOp) Delete(id int8) error {
	return op.DeleteTx(db.JYTradeDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyIntegralConfOp) DeleteTx(ext sqlx.Ext, id int8) error {
	sql := `delete from jy_integral_conf where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyIntegralConfOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_integral_conf where 1=1 `
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

func (op *jyIntegralConfOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYTradeDB, m)
}

func (op *jyIntegralConfOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_integral_conf where 1=1 "
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
