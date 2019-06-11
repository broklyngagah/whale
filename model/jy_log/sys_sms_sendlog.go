package JYLogDB

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

//sys_sms_sendlog
//系统短信日志

// +gen *
type SysSmsSendlog struct {
	Id           int    `db:"id" json:"id"`                       //
	Tel          string `db:"tel" json:"tel"`                     // 手机号码
	Content      string `db:"content" json:"content"`             // 短信内容
	TplId        int    `db:"tpl_id" json:"tpl_id"`               // 短信模板ID
	RCode        string `db:"r_code" json:"r_code"`               // 回调code
	RDesc        string `db:"r_desc" json:"r_desc"`               // 回调内容
	Status       int8   `db:"status" json:"status"`               // 状态  1创建 2发送 3到达
	ChannelId    int    `db:"channel_id" json:"channel_id"`       // 短信通道标识
	ChannelTitle string `db:"channel_title" json:"channel_title"` // 通道名称
	CTime        int64  `db:"c_time" json:"c_time"`               // 创建时间
	STime        int64  `db:"s_time" json:"s_time"`               // 发送时间
	ATime        int64  `db:"a_time" json:"a_time"`               // 到达时间
	VerifyCode   string `db:"verify_code" json:"verify_code"`     // 验证码
	Cksms        int8   `db:"cksms" json:"cksms"`                 // 1已验证通过，2为未使用
	Type         int8   `db:"type" json:"type"`                   // 1手机短信，2语音短信
	Platform     int8   `db:"platform" json:"platform"`           // 平台类型设备(1-ios;2-android;3-wap;4-PC;5微信游戏;6-ios回馈版)
	Ip           string `db:"ip" json:"ip"`                       // IP地址
	IpMd5        string `db:"ip_md5" json:"ip_md5"`               // IP地址的md5
	Imei         string `db:"imei" json:"imei"`                   // 设备标识
	Source       string `db:"source" json:"source"`               // 渠道来源
	AppModel     string `db:"app_model" json:"app_model"`         // 手机型号
	RefererUrl   string `db:"referer_url" json:"referer_url"`     // 发短信接口，上级请求地址
	Remark       string `db:"remark" json:"remark"`               // 备注
	MsgId        string `db:"msg_id" json:"msg_id"`               // 第三方短信发送ID
}

type sysSmsSendlogOp struct{}

var SysSmsSendlogOp = &sysSmsSendlogOp{}
var DefaultSysSmsSendlog = &SysSmsSendlog{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *sysSmsSendlogOp) Get(id int) (*SysSmsSendlog, error) {
	obj := &SysSmsSendlog{}
	sql := "select `id`,`tel`,`content`,`tpl_id`,`r_code`,`r_desc`,`status`,`channel_id`,`channel_title`,`c_time`,`s_time`,`a_time`,`verify_code`,`cksms`,`type`,`platform`,`ip`,`ip_md5`,`imei`,`source`,`app_model`,`referer_url`,`remark`,`msg_id` from sys_sms_sendlog where id=? "
	err := db.JYLogDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *sysSmsSendlogOp) SelectAll() ([]*SysSmsSendlog, error) {
	objList := []*SysSmsSendlog{}
	sql := "select `id`,`tel`,`content`,`tpl_id`,`r_code`,`r_desc`,`status`,`channel_id`,`channel_title`,`c_time`,`s_time`,`a_time`,`verify_code`,`cksms`,`type`,`platform`,`ip`,`ip_md5`,`imei`,`source`,`app_model`,`referer_url`,`remark`,`msg_id` from sys_sms_sendlog"
	err := db.JYLogDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *sysSmsSendlogOp) QueryByMap(m map[string]interface{}) ([]*SysSmsSendlog, error) {
	result := []*SysSmsSendlog{}
	var params []interface{}

	sql := "select `id`,`tel`,`content`,`tpl_id`,`r_code`,`r_desc`,`status`,`channel_id`,`channel_title`,`c_time`,`s_time`,`a_time`,`verify_code`,`cksms`,`type`,`platform`,`ip`,`ip_md5`,`imei`,`source`,`app_model`,`referer_url`,`remark`,`msg_id` from sys_sms_sendlog where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s=? ", k)
		params = append(params, v)
	}
	err := db.JYLogDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *sysSmsSendlogOp) QueryByMapComparison(m map[string]interface{}) ([]*SysSmsSendlog, error) {
	result := []*SysSmsSendlog{}
	var params []interface{}

	sql := "select `id`,`tel`,`content`,`tpl_id`,`r_code`,`r_desc`,`status`,`channel_id`,`channel_title`,`c_time`,`s_time`,`a_time`,`verify_code`,`cksms`,`type`,`platform`,`ip`,`ip_md5`,`imei`,`source`,`app_model`,`referer_url`,`remark`,`msg_id` from sys_sms_sendlog where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s? ", k)
		params = append(params, v)
	}
	err := db.JYLogDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *sysSmsSendlogOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*SysSmsSendlog, error) {
	result := []*SysSmsSendlog{}
	var params []interface{}

	sql := "select `id`,`tel`,`content`,`tpl_id`,`r_code`,`r_desc`,`status`,`channel_id`,`channel_title`,`c_time`,`s_time`,`a_time`,`verify_code`,`cksms`,`type`,`platform`,`ip`,`ip_md5`,`imei`,`source`,`app_model`,`referer_url`,`remark`,`msg_id` from sys_sms_sendlog where 1=1 "
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

	err := db.JYLogDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *sysSmsSendlogOp) GetByMap(m map[string]interface{}) (*SysSmsSendlog, error) {
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
func (op *sysSmsSendlogOp) Insert(m *SysSmsSendlog) (int64, error) {
	return op.InsertTx(db.JYLogDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *sysSmsSendlogOp) InsertTx(ext sqlx.Ext, m *SysSmsSendlog) (int64, error) {
	sql := "insert into sys_sms_sendlog(tel,content,tpl_id,r_code,r_desc,status,channel_id,channel_title,c_time,s_time,a_time,verify_code,cksms,type,platform,ip,ip_md5,imei,source,app_model,referer_url,remark,msg_id) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Tel,
		m.Content,
		m.TplId,
		m.RCode,
		m.RDesc,
		m.Status,
		m.ChannelId,
		m.ChannelTitle,
		m.CTime,
		m.STime,
		m.ATime,
		m.VerifyCode,
		m.Cksms,
		m.Type,
		m.Platform,
		m.Ip,
		m.IpMd5,
		m.Imei,
		m.Source,
		m.AppModel,
		m.RefererUrl,
		m.Remark,
		m.MsgId,
	)
	if err != nil {
		game_error.RaiseError(err)
		return -1, err
	}
	affected, _ := result.RowsAffected()
	return affected, nil
}

/*
func (i *SysSmsSendlog) Update() {
    _,err := db.JYLogDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *sysSmsSendlogOp) Update(m *SysSmsSendlog) error {
	return op.UpdateTx(db.JYLogDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *sysSmsSendlogOp) UpdateTx(ext sqlx.Ext, m *SysSmsSendlog) error {
	sql := `update sys_sms_sendlog set tel=?,content=?,tpl_id=?,r_code=?,r_desc=?,status=?,channel_id=?,channel_title=?,c_time=?,s_time=?,a_time=?,verify_code=?,cksms=?,type=?,platform=?,ip=?,ip_md5=?,imei=?,source=?,app_model=?,referer_url=?,remark=?,msg_id=? where id=?`
	_, err := ext.Exec(sql,
		m.Tel,
		m.Content,
		m.TplId,
		m.RCode,
		m.RDesc,
		m.Status,
		m.ChannelId,
		m.ChannelTitle,
		m.CTime,
		m.STime,
		m.ATime,
		m.VerifyCode,
		m.Cksms,
		m.Type,
		m.Platform,
		m.Ip,
		m.IpMd5,
		m.Imei,
		m.Source,
		m.AppModel,
		m.RefererUrl,
		m.Remark,
		m.MsgId,
		m.Id,
	)

	if err != nil {
		game_error.RaiseError(err)
		return err
	}

	return nil
}

// 用主键做条件，更新map里包含的字段名
func (op *sysSmsSendlogOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYLogDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *sysSmsSendlogOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update sys_sms_sendlog set %s where 1=1 and id=? ;`

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
func (i *SysSmsSendlog) Delete(){
    _,err := db.JYLogDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *sysSmsSendlogOp) Delete(id int) error {
	return op.DeleteTx(db.JYLogDB, id)
}

// 根据主键删除相关记录,Tx
func (op *sysSmsSendlogOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from sys_sms_sendlog where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *sysSmsSendlogOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from sys_sms_sendlog where 1=1 `
	for k, v := range m {
		sql += fmt.Sprintf(" and  %s=? ", k)
		params = append(params, v)
	}
	count := int64(-1)
	err := db.JYLogDB.Get(&count, sql, params...)
	if err != nil {
		game_error.RaiseError(err)
	}
	return count
}

func (op *sysSmsSendlogOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYLogDB, m)
}

func (op *sysSmsSendlogOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from sys_sms_sendlog where 1=1 "
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
