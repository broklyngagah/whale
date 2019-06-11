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

//jy_shenhe
//审核表

// +gen *
type JyShenhe struct {
	Id            int    `db:"id" json:"id"`                         //
	Tid           int    `db:"tid" json:"tid"`                       // 文章分类ID
	Aid           int    `db:"aid" json:"aid"`                       // 文章ID
	Adminid       int    `db:"adminid" json:"adminid"`               // 审核人UID
	Uid           int    `db:"uid" json:"uid"`                       // 发布人uid
	Tel           string `db:"tel" json:"tel"`                       // 手机号码
	Title         string `db:"title" json:"title"`                   // 文章标题
	Type          int8   `db:"type" json:"type"`                     // 1文章2帖子
	Level         int8   `db:"level" json:"level"`                   //
	Status        int8   `db:"status" json:"status"`                 // 1已通过2草稿8不通过3等待提交审核4待审核9删除
	Remark        string `db:"remark" json:"remark"`                 // 备注
	AppendContent string `db:"append_content" json:"append_content"` // 追加的内容
	PublicTime    int64  `db:"public_time" json:"public_time"`       // 文章发布时间
	CTime         int64  `db:"c_time" json:"c_time"`                 // 添加时间
	UTime         int64  `db:"u_time" json:"u_time"`                 // 更新时间
}

type jyShenheOp struct{}

var JyShenheOp = &jyShenheOp{}
var DefaultJyShenhe = &JyShenhe{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *jyShenheOp) Get(id int) (*JyShenhe, error) {
	obj := &JyShenhe{}
	sql := "select `id`,`tid`,`aid`,`adminid`,`uid`,`tel`,`title`,`type`,`level`,`status`,`remark`,`append_content`,`public_time`,`c_time`,`u_time` from jy_shenhe where id=? "
	err := db.JYOtherDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *jyShenheOp) SelectAll() ([]*JyShenhe, error) {
	objList := []*JyShenhe{}
	sql := "select `id`,`tid`,`aid`,`adminid`,`uid`,`tel`,`title`,`type`,`level`,`status`,`remark`,`append_content`,`public_time`,`c_time`,`u_time` from jy_shenhe"
	err := db.JYOtherDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *jyShenheOp) QueryByMap(m map[string]interface{}) ([]*JyShenhe, error) {
	result := []*JyShenhe{}
	var params []interface{}

	sql := "select `id`,`tid`,`aid`,`adminid`,`uid`,`tel`,`title`,`type`,`level`,`status`,`remark`,`append_content`,`public_time`,`c_time`,`u_time` from jy_shenhe where 1=1 "
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

func (op *jyShenheOp) QueryByMapComparison(m map[string]interface{}) ([]*JyShenhe, error) {
	result := []*JyShenhe{}
	var params []interface{}

	sql := "select `id`,`tid`,`aid`,`adminid`,`uid`,`tel`,`title`,`type`,`level`,`status`,`remark`,`append_content`,`public_time`,`c_time`,`u_time` from jy_shenhe where 1=1 "
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

func (op *jyShenheOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*JyShenhe, error) {
	result := []*JyShenhe{}
	var params []interface{}

	sql := "select `id`,`tid`,`aid`,`adminid`,`uid`,`tel`,`title`,`type`,`level`,`status`,`remark`,`append_content`,`public_time`,`c_time`,`u_time` from jy_shenhe where 1=1 "
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

func (op *jyShenheOp) GetByMap(m map[string]interface{}) (*JyShenhe, error) {
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
func (op *jyShenheOp) Insert(m *JyShenhe) (int64, error) {
	return op.InsertTx(db.JYOtherDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *jyShenheOp) InsertTx(ext sqlx.Ext, m *JyShenhe) (int64, error) {
	sql := "insert into jy_shenhe(tid,aid,adminid,uid,tel,title,type,level,status,remark,append_content,public_time,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Tid,
		m.Aid,
		m.Adminid,
		m.Uid,
		m.Tel,
		m.Title,
		m.Type,
		m.Level,
		m.Status,
		m.Remark,
		m.AppendContent,
		m.PublicTime,
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
func (i *JyShenhe) Update() {
    _,err := db.JYOtherDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyShenheOp) Update(m *JyShenhe) error {
	return op.UpdateTx(db.JYOtherDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *jyShenheOp) UpdateTx(ext sqlx.Ext, m *JyShenhe) error {
	sql := `update jy_shenhe set tid=?,aid=?,adminid=?,uid=?,tel=?,title=?,type=?,level=?,status=?,remark=?,append_content=?,public_time=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Tid,
		m.Aid,
		m.Adminid,
		m.Uid,
		m.Tel,
		m.Title,
		m.Type,
		m.Level,
		m.Status,
		m.Remark,
		m.AppendContent,
		m.PublicTime,
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
func (op *jyShenheOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYOtherDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *jyShenheOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update jy_shenhe set %s where 1=1 and id=? ;`

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
func (i *JyShenhe) Delete(){
    _,err := db.JYOtherDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *jyShenheOp) Delete(id int) error {
	return op.DeleteTx(db.JYOtherDB, id)
}

// 根据主键删除相关记录,Tx
func (op *jyShenheOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from jy_shenhe where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *jyShenheOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from jy_shenhe where 1=1 `
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

func (op *jyShenheOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYOtherDB, m)
}

func (op *jyShenheOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from jy_shenhe where 1=1 "
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
