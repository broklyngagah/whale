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

//jy_article_comment
//文章评论表

// +gen *
type JyArticleComment struct {
	Id       int    `db:"id" json:"id"`               // id
	Aid      int    `db:"aid" json:"aid"`             // 文章ID
	Uid      int    `db:"uid" json:"uid"`             // 评论者ID
	Mid      int    `db:"mid" json:"mid"`             // 根评论id(0表示一级评论)
	Pid      int    `db:"pid" json:"pid"`             // 父级评论
	Ruid     int    `db:"ruid" json:"ruid"`           // 被回复者的UID
	Content  string `db:"content" json:"content"`     // 评论内容
	LikedNum int8   `db:"liked_num" json:"liked_num"` // 点赞数量
	Pathinfo string `db:"pathinfo" json:"pathinfo"`   // 评论id路径
	Status   int8   `db:"status" json:"status"`       // 状态（1评论中，4待审核，9删除评论）
	CTime    int64  `db:"c_time" json:"c_time"`       // 评论时间
	UTime    int64  `db:"u_time" json:"u_time"`       // 更新时间
}

type jyArticleCommentOp struct{}

var JyArticleCommentOp = &jyArticleCommentOp{}
var DefaultJyArticleComment = &JyArticleComment{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyArticleCommentOp) Get(id int) (*JyArticleComment, error) {
	obj := &JyArticleComment{}
	sql := "select `id`,`aid`,`uid`,`mid`,`pid`,`ruid`,`content`,`liked_num`,`pathinfo`,`status`,`c_time`,`u_time` from jy_article_comment where id=? "
	err := db.JYOtherDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyArticleCommentOp) SelectAll() ([]*JyArticleComment, error) {
	objList := []*JyArticleComment{}
	sql := "select `id`,`aid`,`uid`,`mid`,`pid`,`ruid`,`content`,`liked_num`,`pathinfo`,`status`,`c_time`,`u_time` from jy_article_comment"
	err := db.JYOtherDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyArticleCommentOp) QueryByMap(m map[string]interface{}) ([]*JyArticleComment, error) {
	result := []*JyArticleComment{}
	var params []interface{}

	sql := "select `id`,`aid`,`uid`,`mid`,`pid`,`ruid`,`content`,`liked_num`,`pathinfo`,`status`,`c_time`,`u_time` from jy_article_comment where 1=1 "
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

func (op *jyArticleCommentOp) QueryByMapComparison(m map[string]interface{}) ([]*JyArticleComment, error) {
	result := []*JyArticleComment{}
	var params []interface{}

	sql := "select `id`,`aid`,`uid`,`mid`,`pid`,`ruid`,`content`,`liked_num`,`pathinfo`,`status`,`c_time`,`u_time` from jy_article_comment where 1=1 "
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

func (op *jyArticleCommentOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyArticleComment, error) {
	result := []*JyArticleComment{}
	var params []interface{}

	sql := "select `id`,`aid`,`uid`,`mid`,`pid`,`ruid`,`content`,`liked_num`,`pathinfo`,`status`,`c_time`,`u_time` from jy_article_comment where 1=1 "
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

func (op *jyArticleCommentOp) GetByMap(m map[string]interface{}) (*JyArticleComment, error) {
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
func (op *jyArticleCommentOp) Insert(m *JyArticleComment) (int64, error) {
	return op.InsertTx(db.JYOtherDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyArticleCommentOp) InsertTx(ext sqlx.Ext, m *JyArticleComment) (int64, error) {
	sql := "insert into jy_article_comment(aid,uid,mid,pid,ruid,content,liked_num,pathinfo,status,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Aid,
		m.Uid,
		m.Mid,
		m.Pid,
		m.Ruid,
		m.Content,
		m.LikedNum,
		m.Pathinfo,
		m.Status,
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
func (i *JyArticleComment) Update() {
    _,err := db.JYOtherDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleCommentOp) Update(m *JyArticleComment) error {
	return op.UpdateTx(db.JYOtherDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleCommentOp) UpdateTx(ext sqlx.Ext, m *JyArticleComment) error {
	sql := `update jy_article_comment set aid=?,uid=?,mid=?,pid=?,ruid=?,content=?,liked_num=?,pathinfo=?,status=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Aid,
		m.Uid,
		m.Mid,
		m.Pid,
		m.Ruid,
		m.Content,
		m.LikedNum,
		m.Pathinfo,
		m.Status,
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
func (op *jyArticleCommentOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYOtherDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyArticleCommentOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update jy_article_comment set %s where 1=1 and id=? ;`

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
func (i *JyArticleComment) Delete(){
    _,err := db.JYOtherDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyArticleCommentOp) Delete(id int) error {
	return op.DeleteTx(db.JYOtherDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyArticleCommentOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from jy_article_comment where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyArticleCommentOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_article_comment where 1=1 `
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

func (op *jyArticleCommentOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYOtherDB, m)
}

func (op *jyArticleCommentOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_article_comment where 1=1 "
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
