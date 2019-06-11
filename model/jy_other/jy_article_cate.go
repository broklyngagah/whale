package JYOtherDB

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

//jy_article_cate
//话题分类表

// +gen *
type JyArticleCate struct {
	Id        int     `db:"id" json:"id"`                 //
	Pid       int     `db:"pid" json:"pid"`               // 父类ID(0为一级主分类)
	Title     string  `db:"title" json:"title"`           // 话题标题
	Desc      string  `db:"desc" json:"desc"`             // 话题描述
	Type      int8    `db:"type" json:"type"`             // 类型（1 无所属板块，2有所属板块）
	Thumb     string  `db:"thumb" json:"thumb"`           // 图片头像
	LikedNum  int     `db:"liked_num" json:"liked_num"`   // 关注数量
	RiskIndex float64 `db:"risk_index" json:"risk_index"` // 风险指数
	StockCode string  `db:"stock_code" json:"stock_code"` // 股票代码
	ReleCode  string  `db:"rele_code" json:"rele_code"`   // 相关股票代码（json字符串）
	Extend    string  `db:"extend" json:"extend"`         // 拓展字段，json格式 {type:1,time:1514244515,time_range:9-11月}  type1预估时间，type2确定时间，type3追踪时间
	Status    int8    `db:"status" json:"status"`         // 状态（1已发布，2未发布，9删除）
	OptTime   int64   `db:"opt_time" json:"opt_time"`     // 发布时间
	Sort      int     `db:"sort" json:"sort"`             // 排序
	CTime     int64   `db:"c_time" json:"c_time"`         // 添加时间
	UTime     int64   `db:"u_time" json:"u_time"`         // 更新时间
}

type jyArticleCateOp struct{}

var JyArticleCateOp = &jyArticleCateOp{}
var DefaultJyArticleCate = &JyArticleCate{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyArticleCateOp) Get(id int) (*JyArticleCate, error) {
	obj := &JyArticleCate{}
	sql := "select `id`,`pid`,`title`,`desc`,`type`,`thumb`,`liked_num`,`risk_index`,`stock_code`,`rele_code`,`extend`,`status`,`opt_time`,`sort`,`c_time`,`u_time` from jy_article_cate where id=? "
	err := db.JYOtherDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyArticleCateOp) SelectAll() ([]*JyArticleCate, error) {
	objList := []*JyArticleCate{}
	sql := "select `id`,`pid`,`title`,`desc`,`type`,`thumb`,`liked_num`,`risk_index`,`stock_code`,`rele_code`,`extend`,`status`,`opt_time`,`sort`,`c_time`,`u_time` from jy_article_cate"
	err := db.JYOtherDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyArticleCateOp) QueryByMap(m map[string]interface{}) ([]*JyArticleCate, error) {
	result := []*JyArticleCate{}
	var params []interface{}

	sql := "select `id`,`pid`,`title`,`desc`,`type`,`thumb`,`liked_num`,`risk_index`,`stock_code`,`rele_code`,`extend`,`status`,`opt_time`,`sort`,`c_time`,`u_time` from jy_article_cate where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s=? ", k)
		params = append(params, v)
	}
	err := db.JYOtherDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyArticleCateOp) QueryByMapComparison(m map[string]interface{}) ([]*JyArticleCate, error) {
	result := []*JyArticleCate{}
	var params []interface{}

	sql := "select `id`,`pid`,`title`,`desc`,`type`,`thumb`,`liked_num`,`risk_index`,`stock_code`,`rele_code`,`extend`,`status`,`opt_time`,`sort`,`c_time`,`u_time` from jy_article_cate where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s? ", k)
		params = append(params, v)
	}
	err := db.JYOtherDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyArticleCateOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyArticleCate, error) {
	result := []*JyArticleCate{}
	var params []interface{}

	sql := "select `id`,`pid`,`title`,`desc`,`type`,`thumb`,`liked_num`,`risk_index`,`stock_code`,`rele_code`,`extend`,`status`,`opt_time`,`sort`,`c_time`,`u_time` from jy_article_cate where 1=1 "
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

	err := db.JYOtherDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *jyArticleCateOp) GetByMap(m map[string]interface{}) (*JyArticleCate, error) {
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
func (op *jyArticleCateOp) Insert(m *JyArticleCate) (int64, error) {
	return op.InsertTx(db.JYOtherDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyArticleCateOp) InsertTx(ext sqlx.Ext, m *JyArticleCate) (int64, error) {
	sql := "insert into jy_article_cate(pid,title,desc,type,thumb,liked_num,risk_index,stock_code,rele_code,extend,status,opt_time,sort,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Pid,
		m.Title,
		m.Desc,
		m.Type,
		m.Thumb,
		m.LikedNum,
		m.RiskIndex,
		m.StockCode,
		m.ReleCode,
		m.Extend,
		m.Status,
		m.OptTime,
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
func (i *JyArticleCate) Update() {
    _,err := db.JYOtherDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleCateOp) Update(m *JyArticleCate) error {
	return op.UpdateTx(db.JYOtherDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleCateOp) UpdateTx(ext sqlx.Ext, m *JyArticleCate) error {
	sql := `update jy_article_cate set pid=?,title=?,desc=?,type=?,thumb=?,liked_num=?,risk_index=?,stock_code=?,rele_code=?,extend=?,status=?,opt_time=?,sort=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Pid,
		m.Title,
		m.Desc,
		m.Type,
		m.Thumb,
		m.LikedNum,
		m.RiskIndex,
		m.StockCode,
		m.ReleCode,
		m.Extend,
		m.Status,
		m.OptTime,
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
func (op *jyArticleCateOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYOtherDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyArticleCateOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update jy_article_cate set %s where 1=1 and id=? ;`

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
func (i *JyArticleCate) Delete(){
    _,err := db.JYOtherDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyArticleCateOp) Delete(id int) error {
	return op.DeleteTx(db.JYOtherDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyArticleCateOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from jy_article_cate where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyArticleCateOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_article_cate where 1=1 `
	for k, v := range m {
		sql += fmt.Sprintf(" and  %s=? ", k)
		params = append(params, v)
	}
	count := int64(-1)
	err := db.JYOtherDB.Get(&count, sql, params...)
	if err != nil {
		game_error.RaiseError(err)
	}
	return count
}

func (op *jyArticleCateOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYOtherDB, m)
}

func (op *jyArticleCateOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_article_cate where 1=1 "
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
