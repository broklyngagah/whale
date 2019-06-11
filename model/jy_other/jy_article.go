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

//jy_article
//文章表

// +gen *
type JyArticle struct {
	Id           int     `db:"id" json:"id"`                       //
	Tid          int     `db:"tid" json:"tid"`                     // 分类id
	Uid          int     `db:"uid" json:"uid"`                     // 发布者uid（0表示系统发布）
	Title        string  `db:"title" json:"title"`                 // 文章标题
	Keyword      string  `db:"keyword" json:"keyword"`             // 关键字
	Desc         string  `db:"desc" json:"desc"`                   // 文章短描述
	Thumb        string  `db:"thumb" json:"thumb"`                 // 封面缩略图(多张json)
	Type         int     `db:"type" json:"type"`                   // 文章类型（1：文章 2：帖子）
	Islink       int8    `db:"islink" json:"islink"`               // 1跳转0不跳转
	Url          string  `db:"url" json:"url"`                     // 跳转地址
	ChargeType   int8    `db:"charge_type" json:"charge_type"`     // 收费类型（1免费，2收费）
	ChargeStatus int8    `db:"charge_status" json:"charge_status"` // 付费状态 1开启 2关闭 （关闭付费功能，不再被购买）
	Price        float64 `db:"price" json:"price"`                 // 收费文章鱼币单价
	Integral     int     `db:"integral" json:"integral"`           // 收费文章积分单价
	Status       int8    `db:"status" json:"status"`               // 状态（1已发布  2草稿箱  3定时提交审核 4 审核中  5待发布 9删除）
	IsComment    int8    `db:"is_comment" json:"is_comment"`       // 状态（1开启评论，2关闭评论）
	IsShow       int8    `db:"is_show" json:"is_show"`             // 是否显示文字1显示2不显示
	IsPush       int8    `db:"is_push" json:"is_push"`             // 推送1是2否
	Sort         int     `db:"sort" json:"sort"`                   // 排序
	Adminid      int     `db:"adminid" json:"adminid"`             // 1管理员0非管理员
	OptTime      int64   `db:"opt_time" json:"opt_time"`           // 发布时间
	CTime        int64   `db:"c_time" json:"c_time"`               // 添加时间
	UTime        int64   `db:"u_time" json:"u_time"`               // 更新时间
}

type jyArticleOp struct{}

var JyArticleOp = &jyArticleOp{}
var DefaultJyArticle = &JyArticle{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyArticleOp) Get(id int) (*JyArticle, error) {
	obj := &JyArticle{}
	sql := "select `id`,`tid`,`uid`,`title`,`keyword`,`desc`,`thumb`,`type`,`islink`,`url`,`charge_type`,`charge_status`,`price`,`integral`,`status`,`is_comment`,`is_show`,`is_push`,`sort`,`adminid`,`opt_time`,`c_time`,`u_time` from jy_article where id=? "
	err := db.JYOtherDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyArticleOp) SelectAll() ([]*JyArticle, error) {
	objList := []*JyArticle{}
	sql := "select `id`,`tid`,`uid`,`title`,`keyword`,`desc`,`thumb`,`type`,`islink`,`url`,`charge_type`,`charge_status`,`price`,`integral`,`status`,`is_comment`,`is_show`,`is_push`,`sort`,`adminid`,`opt_time`,`c_time`,`u_time` from jy_article"
	err := db.JYOtherDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyArticleOp) QueryByMap(m map[string]interface{}) ([]*JyArticle, error) {
	result := []*JyArticle{}
	var params []interface{}

	sql := "select `id`,`tid`,`uid`,`title`,`keyword`,`desc`,`thumb`,`type`,`islink`,`url`,`charge_type`,`charge_status`,`price`,`integral`,`status`,`is_comment`,`is_show`,`is_push`,`sort`,`adminid`,`opt_time`,`c_time`,`u_time` from jy_article where 1=1 "
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

func (op *jyArticleOp) QueryByMapComparison(m map[string]interface{}) ([]*JyArticle, error) {
	result := []*JyArticle{}
	var params []interface{}

	sql := "select `id`,`tid`,`uid`,`title`,`keyword`,`desc`,`thumb`,`type`,`islink`,`url`,`charge_type`,`charge_status`,`price`,`integral`,`status`,`is_comment`,`is_show`,`is_push`,`sort`,`adminid`,`opt_time`,`c_time`,`u_time` from jy_article where 1=1 "
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

func (op *jyArticleOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyArticle, error) {
	result := []*JyArticle{}
	var params []interface{}

	sql := "select `id`,`tid`,`uid`,`title`,`keyword`,`desc`,`thumb`,`type`,`islink`,`url`,`charge_type`,`charge_status`,`price`,`integral`,`status`,`is_comment`,`is_show`,`is_push`,`sort`,`adminid`,`opt_time`,`c_time`,`u_time` from jy_article where 1=1 "
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

func (op *jyArticleOp) GetByMap(m map[string]interface{}) (*JyArticle, error) {
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
func (op *jyArticleOp) Insert(m *JyArticle) (int64, error) {
	return op.InsertTx(db.JYOtherDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyArticleOp) InsertTx(ext sqlx.Ext, m *JyArticle) (int64, error) {
	sql := "insert into jy_article(tid,uid,title,keyword,desc,thumb,type,islink,url,charge_type,charge_status,price,integral,status,is_comment,is_show,is_push,sort,adminid,opt_time,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Tid,
		m.Uid,
		m.Title,
		m.Keyword,
		m.Desc,
		m.Thumb,
		m.Type,
		m.Islink,
		m.Url,
		m.ChargeType,
		m.ChargeStatus,
		m.Price,
		m.Integral,
		m.Status,
		m.IsComment,
		m.IsShow,
		m.IsPush,
		m.Sort,
		m.Adminid,
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
func (i *JyArticle) Update() {
    _,err := db.JYOtherDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleOp) Update(m *JyArticle) error {
	return op.UpdateTx(db.JYOtherDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleOp) UpdateTx(ext sqlx.Ext, m *JyArticle) error {
	sql := `update jy_article set tid=?,uid=?,title=?,keyword=?,desc=?,thumb=?,type=?,islink=?,url=?,charge_type=?,charge_status=?,price=?,integral=?,status=?,is_comment=?,is_show=?,is_push=?,sort=?,adminid=?,opt_time=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Tid,
		m.Uid,
		m.Title,
		m.Keyword,
		m.Desc,
		m.Thumb,
		m.Type,
		m.Islink,
		m.Url,
		m.ChargeType,
		m.ChargeStatus,
		m.Price,
		m.Integral,
		m.Status,
		m.IsComment,
		m.IsShow,
		m.IsPush,
		m.Sort,
		m.Adminid,
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
func (op *jyArticleOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYOtherDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyArticleOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update jy_article set %s where 1=1 and id=? ;`

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
func (i *JyArticle) Delete(){
    _,err := db.JYOtherDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyArticleOp) Delete(id int) error {
	return op.DeleteTx(db.JYOtherDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyArticleOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from jy_article where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyArticleOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_article where 1=1 `
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

func (op *jyArticleOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYOtherDB, m)
}

func (op *jyArticleOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_article where 1=1 "
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
