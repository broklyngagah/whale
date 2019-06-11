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

//user_auth
//用户申请大V认证表

// +gen *
type UserAuth struct {
	Id             int    `db:"id" json:"id"`                           // id
	Uid            int    `db:"uid" json:"uid"`                         // 用户UID
	CardId         string `db:"card_id" json:"card_id"`                 // 身份证号码
	RealName       string `db:"real_name" json:"real_name"`             // 真实姓名
	CardImg        string `db:"card_img" json:"card_img"`               // 身份证照片
	CredentialsImg string `db:"credentials_img" json:"credentials_img"` // 资质材料
	MediaData      string `db:"media_data" json:"media_data"`           // 自媒材料
	Email          string `db:"email" json:"email"`                     // 邮箱
	AlipayNo       string `db:"alipay_no" json:"alipay_no"`             // 支付宝账号
	MsgContent     string `db:"msg_content" json:"msg_content"`         // 审核失败返回内容
	Status         int8   `db:"status" json:"status"`                   // 审核状态 1审核中，2审核驳回 5审核通过
	OptTime        int64  `db:"opt_time" json:"opt_time"`               // 认证时间
	CTime          int64  `db:"c_time" json:"c_time"`                   // 创建时间
	UTime          int64  `db:"u_time" json:"u_time"`                   // 改修时间
}

type userAuthOp struct{}

var UserAuthOp = &userAuthOp{}
var DefaultUserAuth = &UserAuth{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userAuthOp) Get(id int) (*UserAuth, error) {
	obj := &UserAuth{}
	sql := "select `id`,`uid`,`card_id`,`real_name`,`card_img`,`credentials_img`,`media_data`,`email`,`alipay_no`,`msg_content`,`status`,`opt_time`,`c_time`,`u_time` from user_auth where id=? "
	err := db.JYMemberDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userAuthOp) SelectAll() ([]*UserAuth, error) {
	objList := []*UserAuth{}
	sql := "select `id`,`uid`,`card_id`,`real_name`,`card_img`,`credentials_img`,`media_data`,`email`,`alipay_no`,`msg_content`,`status`,`opt_time`,`c_time`,`u_time` from user_auth"
	err := db.JYMemberDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userAuthOp) QueryByMap(m map[string]interface{}) ([]*UserAuth, error) {
	result := []*UserAuth{}
	var params []interface{}

	sql := "select `id`,`uid`,`card_id`,`real_name`,`card_img`,`credentials_img`,`media_data`,`email`,`alipay_no`,`msg_content`,`status`,`opt_time`,`c_time`,`u_time` from user_auth where 1=1 "
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

func (op *userAuthOp) QueryByMapComparison(m map[string]interface{}) ([]*UserAuth, error) {
	result := []*UserAuth{}
	var params []interface{}

	sql := "select `id`,`uid`,`card_id`,`real_name`,`card_img`,`credentials_img`,`media_data`,`email`,`alipay_no`,`msg_content`,`status`,`opt_time`,`c_time`,`u_time` from user_auth where 1=1 "
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

func (op *userAuthOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserAuth, error) {
	result := []*UserAuth{}
	var params []interface{}

	sql := "select `id`,`uid`,`card_id`,`real_name`,`card_img`,`credentials_img`,`media_data`,`email`,`alipay_no`,`msg_content`,`status`,`opt_time`,`c_time`,`u_time` from user_auth where 1=1 "
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

func (op *userAuthOp) GetByMap(m map[string]interface{}) (*UserAuth, error) {
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
func (op *userAuthOp) Insert(m *UserAuth) (int64, error) {
	return op.InsertTx(db.JYMemberDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userAuthOp) InsertTx(ext sqlx.Ext, m *UserAuth) (int64, error) {
	sql := "insert into user_auth(uid,card_id,real_name,card_img,credentials_img,media_data,email,alipay_no,msg_content,status,opt_time,c_time,u_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.CardId,
		m.RealName,
		m.CardImg,
		m.CredentialsImg,
		m.MediaData,
		m.Email,
		m.AlipayNo,
		m.MsgContent,
		m.Status,
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
func (i *UserAuth) Update() {
    _,err := db.JYMemberDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userAuthOp) Update(m *UserAuth) error {
	return op.UpdateTx(db.JYMemberDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userAuthOp) UpdateTx(ext sqlx.Ext, m *UserAuth) error {
	sql := `update user_auth set uid=?,card_id=?,real_name=?,card_img=?,credentials_img=?,media_data=?,email=?,alipay_no=?,msg_content=?,status=?,opt_time=?,c_time=?,u_time=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.CardId,
		m.RealName,
		m.CardImg,
		m.CredentialsImg,
		m.MediaData,
		m.Email,
		m.AlipayNo,
		m.MsgContent,
		m.Status,
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
func (op *userAuthOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYMemberDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userAuthOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update user_auth set %s where 1=1 and id=? ;`

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
func (i *UserAuth) Delete(){
    _,err := db.JYMemberDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userAuthOp) Delete(id int) error {
	return op.DeleteTx(db.JYMemberDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userAuthOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from user_auth where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userAuthOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_auth where 1=1 `
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

func (op *userAuthOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYMemberDB, m)
}

func (op *userAuthOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_auth where 1=1 "
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
