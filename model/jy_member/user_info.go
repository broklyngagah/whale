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

//user_info
//用户详情表

// +gen *
type UserInfo struct {
	Id           int    `db:"id" json:"id"`                       // ID
	Uid          int    `db:"uid" json:"uid"`                     // 用户ID user表的id
	Sex          int8   `db:"sex" json:"sex"`                     // 性别1表示男2表示女
	Age          int8   `db:"age" json:"age"`                     // 年龄
	Remark       string `db:"remark" json:"remark"`               // 简介
	RealName     string `db:"real_name" json:"real_name"`         // 用户真实姓名
	IsSubscribe  int8   `db:"is_subscribe" json:"is_subscribe"`   // 是否开启订阅 1开启，2关闭
	RoomTime     int64  `db:"room_time" json:"room_time"`         // 开通直播时间
	IsRoom       int8   `db:"is_room" json:"is_room"`             // 是否开通直播1是2否
	ColIntroduce string `db:"col_introduce" json:"col_introduce"` // 专栏介绍
	AppVersion   string `db:"app_version" json:"app_version"`     // 注册时app的版本号
	StockSort    string `db:"stock_sort" json:"stock_sort"`       // 自选股排序
	Position     string `db:"position" json:"position"`           // 注册时的定位坐标
	Ip           string `db:"ip" json:"ip"`                       // 注册IP
	RegProvince  string `db:"reg_province" json:"reg_province"`   // 注册时的省份
	RegCity      string `db:"reg_city" json:"reg_city"`           // 注册时的市区
	CTime        int64  `db:"c_time" json:"c_time"`               // 创建时间
	UTime        int64  `db:"u_time" json:"u_time"`               // 系统更新时间
	Extend       string `db:"extend" json:"extend"`               // system系统消息   optional自选 comment评论  diggup点赞  subscribe订阅或关注 id记录user_message 里面的最新的那条记录的ID
}

type userInfoOp struct{}

var UserInfoOp = &userInfoOp{}
var DefaultUserInfo = &UserInfo{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userInfoOp) Get(id int) (*UserInfo, error) {
	obj := &UserInfo{}
	sql := "select `id`,`uid`,`sex`,`age`,`remark`,`real_name`,`is_subscribe`,`room_time`,`is_room`,`col_introduce`,`app_version`,`stock_sort`,`position`,`ip`,`reg_province`,`reg_city`,`c_time`,`u_time`,`extend` from user_info where id=? "
	err := db.JYMemberDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userInfoOp) SelectAll() ([]*UserInfo, error) {
	objList := []*UserInfo{}
	sql := "select `id`,`uid`,`sex`,`age`,`remark`,`real_name`,`is_subscribe`,`room_time`,`is_room`,`col_introduce`,`app_version`,`stock_sort`,`position`,`ip`,`reg_province`,`reg_city`,`c_time`,`u_time`,`extend` from user_info"
	err := db.JYMemberDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userInfoOp) QueryByMap(m map[string]interface{}) ([]*UserInfo, error) {
	result := []*UserInfo{}
	var params []interface{}

	sql := "select `id`,`uid`,`sex`,`age`,`remark`,`real_name`,`is_subscribe`,`room_time`,`is_room`,`col_introduce`,`app_version`,`stock_sort`,`position`,`ip`,`reg_province`,`reg_city`,`c_time`,`u_time`,`extend` from user_info where 1=1 "
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

func (op *userInfoOp) QueryByMapComparison(m map[string]interface{}) ([]*UserInfo, error) {
	result := []*UserInfo{}
	var params []interface{}

	sql := "select `id`,`uid`,`sex`,`age`,`remark`,`real_name`,`is_subscribe`,`room_time`,`is_room`,`col_introduce`,`app_version`,`stock_sort`,`position`,`ip`,`reg_province`,`reg_city`,`c_time`,`u_time`,`extend` from user_info where 1=1 "
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

func (op *userInfoOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserInfo, error) {
	result := []*UserInfo{}
	var params []interface{}

	sql := "select `id`,`uid`,`sex`,`age`,`remark`,`real_name`,`is_subscribe`,`room_time`,`is_room`,`col_introduce`,`app_version`,`stock_sort`,`position`,`ip`,`reg_province`,`reg_city`,`c_time`,`u_time`,`extend` from user_info where 1=1 "
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

func (op *userInfoOp) GetByMap(m map[string]interface{}) (*UserInfo, error) {
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
func (op *userInfoOp) Insert(m *UserInfo) (int64, error) {
	return op.InsertTx(db.JYMemberDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userInfoOp) InsertTx(ext sqlx.Ext, m *UserInfo) (int64, error) {
	sql := "insert into user_info(uid,sex,age,remark,real_name,is_subscribe,room_time,is_room,col_introduce,app_version,stock_sort,position,ip,reg_province,reg_city,c_time,u_time,extend) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.Sex,
		m.Age,
		m.Remark,
		m.RealName,
		m.IsSubscribe,
		m.RoomTime,
		m.IsRoom,
		m.ColIntroduce,
		m.AppVersion,
		m.StockSort,
		m.Position,
		m.Ip,
		m.RegProvince,
		m.RegCity,
		m.CTime,
		m.UTime,
		m.Extend,
	)
	if err != nil {
		game_error.RaiseError(err)
		return -1, err
	}
	affected, _ := result.RowsAffected()
	return affected, nil
}

/*
func (i *UserInfo) Update() {
    _,err := db.JYMemberDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userInfoOp) Update(m *UserInfo) error {
	return op.UpdateTx(db.JYMemberDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userInfoOp) UpdateTx(ext sqlx.Ext, m *UserInfo) error {
	sql := `update user_info set uid=?,sex=?,age=?,remark=?,real_name=?,is_subscribe=?,room_time=?,is_room=?,col_introduce=?,app_version=?,stock_sort=?,position=?,ip=?,reg_province=?,reg_city=?,c_time=?,u_time=?,extend=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.Sex,
		m.Age,
		m.Remark,
		m.RealName,
		m.IsSubscribe,
		m.RoomTime,
		m.IsRoom,
		m.ColIntroduce,
		m.AppVersion,
		m.StockSort,
		m.Position,
		m.Ip,
		m.RegProvince,
		m.RegCity,
		m.CTime,
		m.UTime,
		m.Extend,
		m.Id,
	)

	if err != nil {
		game_error.RaiseError(err)
		return err
	}

	return nil
}

// 用主键做条件，更新map里包含的字段名
func (op *userInfoOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYMemberDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userInfoOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update user_info set %s where 1=1 and id=? ;`

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
func (i *UserInfo) Delete(){
    _,err := db.JYMemberDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userInfoOp) Delete(id int) error {
	return op.DeleteTx(db.JYMemberDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userInfoOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from user_info where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userInfoOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_info where 1=1 `
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

func (op *userInfoOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYMemberDB, m)
}

func (op *userInfoOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_info where 1=1 "
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
