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

//jy_article_other
//文章其他表

// +gen *
type JyArticleOther struct {
	Id         int64 `db:"id" json:"id"`                   //
	Uid        int   `db:"uid" json:"uid"`                 // 用户ID
	Aid        int   `db:"aid" json:"aid"`                 // 文章id或分类id
	Tid        int   `db:"tid" json:"tid"`                 // 一级分类id
	ChargeType int8  `db:"charge_type" json:"charge_type"` // 收费类型（1免费，2收费）
	ShareNum   int8  `db:"share_num" json:"share_num"`     // 分享数量
	CommentNum int8  `db:"comment_num" json:"comment_num"` // 评论数量
	LikedNum   int   `db:"liked_num" json:"liked_num"`     // 点赞数量 或关注数量
	DayLiked   int   `db:"day_liked" json:"day_liked"`     // 每天的点赞数（0晨清空）
	CollectNum int   `db:"collect_num" json:"collect_num"` // 文章收藏数量
	Type       int8  `db:"type" json:"type"`               // 类型 1文章 2分类
	IsShow     int8  `db:"is_show" json:"is_show"`         // 1显示2不显示
	Status     int8  `db:"status" json:"status"`           // 状态（1:已发布，5未发布 ，9删除）
	DayTime    int64 `db:"day_time" json:"day_time"`       // 操作时间(当天0点时间)
	CTime      int64 `db:"c_time" json:"c_time"`           // 创建时间
	UTime      int64 `db:"u_time" json:"u_time"`           // 最后修改时间
	IsDav      int8  `db:"isDav" json:"isDav"`             // 是否是大V发布的1是2否
}

type jyArticleOtherOp struct{}

var JyArticleOtherOp = &jyArticleOtherOp{}
var DefaultJyArticleOther = &JyArticleOther{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyArticleOtherOp) Get(id int64) (*JyArticleOther, error) {
	obj := &JyArticleOther{}
	sql := "select `id`,`uid`,`aid`,`tid`,`charge_type`,`share_num`,`comment_num`,`liked_num`,`day_liked`,`collect_num`,`type`,`is_show`,`status`,`day_time`,`c_time`,`u_time`,`isDav` from jy_article_other where id=? "
	err := db.JYOtherDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyArticleOtherOp) SelectAll() ([]*JyArticleOther, error) {
	objList := []*JyArticleOther{}
	sql := "select `id`,`uid`,`aid`,`tid`,`charge_type`,`share_num`,`comment_num`,`liked_num`,`day_liked`,`collect_num`,`type`,`is_show`,`status`,`day_time`,`c_time`,`u_time`,`isDav` from jy_article_other"
	err := db.JYOtherDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyArticleOtherOp) QueryByMap(m map[string]interface{}) ([]*JyArticleOther, error) {
	result := []*JyArticleOther{}
	var params []interface{}

	sql := "select `id`,`uid`,`aid`,`tid`,`charge_type`,`share_num`,`comment_num`,`liked_num`,`day_liked`,`collect_num`,`type`,`is_show`,`status`,`day_time`,`c_time`,`u_time`,`isDav` from jy_article_other where 1=1 "
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

func (op *jyArticleOtherOp) QueryByMapComparison(m map[string]interface{}) ([]*JyArticleOther, error) {
	result := []*JyArticleOther{}
	var params []interface{}

	sql := "select `id`,`uid`,`aid`,`tid`,`charge_type`,`share_num`,`comment_num`,`liked_num`,`day_liked`,`collect_num`,`type`,`is_show`,`status`,`day_time`,`c_time`,`u_time`,`isDav` from jy_article_other where 1=1 "
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

func (op *jyArticleOtherOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyArticleOther, error) {
	result := []*JyArticleOther{}
	var params []interface{}

	sql := "select `id`,`uid`,`aid`,`tid`,`charge_type`,`share_num`,`comment_num`,`liked_num`,`day_liked`,`collect_num`,`type`,`is_show`,`status`,`day_time`,`c_time`,`u_time`,`isDav` from jy_article_other where 1=1 "
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

func (op *jyArticleOtherOp) GetByMap(m map[string]interface{}) (*JyArticleOther, error) {
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
func (op *jyArticleOtherOp) Insert(m *JyArticleOther) (int64, error) {
	return op.InsertTx(db.JYOtherDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyArticleOtherOp) InsertTx(ext sqlx.Ext, m *JyArticleOther) (int64, error) {
	sql := "insert into jy_article_other(uid,aid,tid,charge_type,share_num,comment_num,liked_num,day_liked,collect_num,type,is_show,status,day_time,c_time,u_time,isDav) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Aid,
		m.Tid,
		m.ChargeType,
		m.ShareNum,
		m.CommentNum,
		m.LikedNum,
		m.DayLiked,
		m.CollectNum,
		m.Type,
		m.IsShow,
		m.Status,
		m.DayTime,
		m.CTime,
		m.UTime,
		m.IsDav,
	)
	if err != nil {
		game_error.RaiseError(err)
		return -1, err
	}
	affected, _ := result.RowsAffected()
	return affected, nil
}

/*
func (i *JyArticleOther) Update() {
    _,err := db.JYOtherDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleOtherOp) Update(m *JyArticleOther) error {
	return op.UpdateTx(db.JYOtherDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyArticleOtherOp) UpdateTx(ext sqlx.Ext, m *JyArticleOther) error {
	sql := `update jy_article_other set uid=?,aid=?,tid=?,charge_type=?,share_num=?,comment_num=?,liked_num=?,day_liked=?,collect_num=?,type=?,is_show=?,status=?,day_time=?,c_time=?,u_time=?,isDav=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Aid,
		m.Tid,
		m.ChargeType,
		m.ShareNum,
		m.CommentNum,
		m.LikedNum,
		m.DayLiked,
		m.CollectNum,
		m.Type,
		m.IsShow,
		m.Status,
		m.DayTime,
		m.CTime,
		m.UTime,
		m.IsDav,
		m.Id,
	)

	if err != nil {
		game_error.RaiseError(err)
		return err
	}

	return nil
}

// 用主键做条件，更新map里包含的字段名
func (op *jyArticleOtherOp) UpdateWithMap(id int64, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYOtherDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyArticleOtherOp) UpdateWithMapTx(ext sqlx.Ext, id int64, m map[string]interface{}) error {

	sql := `update jy_article_other set %s where 1=1 and id=? ;`

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
func (i *JyArticleOther) Delete(){
    _,err := db.JYOtherDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyArticleOtherOp) Delete(id int64) error {
	return op.DeleteTx(db.JYOtherDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyArticleOtherOp) DeleteTx(ext sqlx.Ext, id int64) error {
	sql := `delete from jy_article_other where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyArticleOtherOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_article_other where 1=1 `
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

func (op *jyArticleOtherOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYOtherDB, m)
}

func (op *jyArticleOtherOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_article_other where 1=1 "
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
