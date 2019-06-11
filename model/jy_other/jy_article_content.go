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

//jy_article_content
//文章内容表

// +gen *
type JyArticleContent struct {
	Id            int    `db:"id" json:"id"`                         //
	Aid           int    `db:"aid" json:"aid"`                       // 文章ID
	WxTitle       string `db:"wx_title" json:"wx_title"`             // 微信分享标题
	WxDesc        string `db:"wx_desc" json:"wx_desc"`               // 微信分享描述
	WxUrl         string `db:"wx_url" json:"wx_url"`                 // 微信分享地址
	FreeContent   string `db:"free_content" json:"free_content"`     //
	Content       string `db:"content" json:"content"`               // 文章内容
	AppendContent string `db:"append_content" json:"append_content"` // 追加内容（多次时累加）
	CTime         int64  `db:"c_time" json:"c_time"`                 // 添加时间
	UTime         int64  `db:"u_time" json:"u_time"`                 // 更新时间
}

type jyArticleContentOp struct{}

var JyArticleContentOp = &jyArticleContentOp{}
var DefaultJyArticleContent = &JyArticleContent{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyArticleContentOp) Get(id int) (*JyArticleContent, error) {
	obj := &JyArticleContent{}
	sql := "select `id`,`aid`,`wx_title`,`wx_desc`,`wx_url`,`free_content`,`content`,`append_content`,`c_time`,`u_time` from jy_article_content where id=? "
	err := db.JYOtherDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyArticleContentOp) SelectAll() ([]*JyArticleContent, error) {
	objList := []*JyArticleContent{}
	sql := "select `id`,`aid`,`wx_title`,`wx_desc`,`wx_url`,`free_content`,`content`,`append_content`,`c_time`,`u_time` from jy_article_content"
	err := db.JYOtherDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyArticleContentOp) QueryByMap(m map[string]interface{}) ([]*JyArticleContent, error) {
	result := []*JyArticleContent{}
	var params []interface{}

	sql := "select `id`,`aid`,`wx_title`,`wx_desc`,`wx_url`,`free_content`,`content`,`append_content`,`c_time`,`u_time` from jy_article_content where 1=1 "
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

func (op *jyArticleContentOp) QueryByMapComparison(m map[string]interface{}) ([]*JyArticleContent, error) {
	result := []*JyArticleContent{}
	var params []interface{}

	sql := "select `id`,`aid`,`wx_title`,`wx_desc`,`wx_url`,`free_content`,`content`,`append_content`,`c_time`,`u_time` from jy_article_content where 1=1 "
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

func (op *jyArticleContentOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyArticleContent, error) {
	result := []*JyArticleContent{}
	var params []interface{}

	sql := "select `id`,`aid`,`wx_title`,`wx_desc`,`wx_url`,`free_content`,`content`,`append_content`,`c_time`,`u_time` from jy_article_content where 1=1 "
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

func (op *jyArticleContentOp) GetByMap(m map[string]interface{}) (*JyArticleContent, error) {
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
func (op *jyArticleContentOp) Insert(m *JyArticleContent) (int64, error) {
	return op.InsertTx(db.JYOtherDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyArticleContentOp) InsertTx(ext sqlx.Ext, m *JyArticleContent) (int64, error) {
	sql := "insert into jy_article_content(aid,wx_title,wx_desc,wx_url,free_content,content,append_content,c_time,u_time) values(?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Aid,
		m.WxTitle,
		m.WxDesc,
		m.WxUrl,
		m.FreeContent,
		m.Content,
		m.AppendContent,
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
func (i *JyArticleContent) Update() {
    _,err := db.JYOtherDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleContentOp) Update(m *JyArticleContent) error {
	return op.UpdateTx(db.JYOtherDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleContentOp) UpdateTx(ext sqlx.Ext, m *JyArticleContent) error {
	sql := `update jy_article_content set aid=?,wx_title=?,wx_desc=?,wx_url=?,free_content=?,content=?,append_content=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Aid,
		m.WxTitle,
		m.WxDesc,
		m.WxUrl,
		m.FreeContent,
		m.Content,
		m.AppendContent,
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
func (op *jyArticleContentOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYOtherDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyArticleContentOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update jy_article_content set %s where 1=1 and id=? ;`

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
func (i *JyArticleContent) Delete(){
    _,err := db.JYOtherDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyArticleContentOp) Delete(id int) error {
	return op.DeleteTx(db.JYOtherDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyArticleContentOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from jy_article_content where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyArticleContentOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_article_content where 1=1 `
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

func (op *jyArticleContentOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYOtherDB, m)
}

func (op *jyArticleContentOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_article_content where 1=1 "
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
