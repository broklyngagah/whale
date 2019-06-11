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

//jy_article_relation_code
//文章内容相关股票代码

// +gen *
type JyArticleRelationCode struct {
	Id     int    `db:"id" json:"id"`         //
	Aid    int    `db:"aid" json:"aid"`       // 文章ID
	Type   int    `db:"type" json:"type"`     // 文章类型（1：设置相关股 2：内容相关股）
	Symbol string `db:"symbol" json:"symbol"` // 相关股票代码(对应股票代码带后缀)
	CTime  int64  `db:"c_time" json:"c_time"` // 添加时间
}

type jyArticleRelationCodeOp struct{}

var JyArticleRelationCodeOp = &jyArticleRelationCodeOp{}
var DefaultJyArticleRelationCode = &JyArticleRelationCode{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyArticleRelationCodeOp) Get(id int) (*JyArticleRelationCode, error) {
	obj := &JyArticleRelationCode{}
	sql := "select `id`,`aid`,`type`,`symbol`,`c_time` from jy_article_relation_code where id=? "
	err := db.JYOtherDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyArticleRelationCodeOp) SelectAll() ([]*JyArticleRelationCode, error) {
	objList := []*JyArticleRelationCode{}
	sql := "select `id`,`aid`,`type`,`symbol`,`c_time` from jy_article_relation_code"
	err := db.JYOtherDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyArticleRelationCodeOp) QueryByMap(m map[string]interface{}) ([]*JyArticleRelationCode, error) {
	result := []*JyArticleRelationCode{}
	var params []interface{}

	sql := "select `id`,`aid`,`type`,`symbol`,`c_time` from jy_article_relation_code where 1=1 "
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

func (op *jyArticleRelationCodeOp) QueryByMapComparison(m map[string]interface{}) ([]*JyArticleRelationCode, error) {
	result := []*JyArticleRelationCode{}
	var params []interface{}

	sql := "select `id`,`aid`,`type`,`symbol`,`c_time` from jy_article_relation_code where 1=1 "
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

func (op *jyArticleRelationCodeOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyArticleRelationCode, error) {
	result := []*JyArticleRelationCode{}
	var params []interface{}

	sql := "select `id`,`aid`,`type`,`symbol`,`c_time` from jy_article_relation_code where 1=1 "
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

func (op *jyArticleRelationCodeOp) GetByMap(m map[string]interface{}) (*JyArticleRelationCode, error) {
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
func (op *jyArticleRelationCodeOp) Insert(m *JyArticleRelationCode) (int64, error) {
	return op.InsertTx(db.JYOtherDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyArticleRelationCodeOp) InsertTx(ext sqlx.Ext, m *JyArticleRelationCode) (int64, error) {
	sql := "insert into jy_article_relation_code(aid,type,symbol,c_time) values(?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Aid,
		m.Type,
		m.Symbol,
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
func (i *JyArticleRelationCode) Update() {
    _,err := db.JYOtherDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleRelationCodeOp) Update(m *JyArticleRelationCode) error {
	return op.UpdateTx(db.JYOtherDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleRelationCodeOp) UpdateTx(ext sqlx.Ext, m *JyArticleRelationCode) error {
	sql := `update jy_article_relation_code set aid=?,type=?,symbol=?,c_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Aid,
		m.Type,
		m.Symbol,
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
func (op *jyArticleRelationCodeOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYOtherDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyArticleRelationCodeOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update jy_article_relation_code set %s where 1=1 and id=? ;`

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
func (i *JyArticleRelationCode) Delete(){
    _,err := db.JYOtherDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyArticleRelationCodeOp) Delete(id int) error {
	return op.DeleteTx(db.JYOtherDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyArticleRelationCodeOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from jy_article_relation_code where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyArticleRelationCodeOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_article_relation_code where 1=1 `
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

func (op *jyArticleRelationCodeOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYOtherDB, m)
}

func (op *jyArticleRelationCodeOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_article_relation_code where 1=1 "
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
