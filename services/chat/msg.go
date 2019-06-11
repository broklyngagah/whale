package chat



type Message struct {
	Id         int    `bson:"id" json:"id"`                   // ID
	Uid        int    `bson:"uid" json:"uid"`                 // 用户ID
	IsClient   int8   `bson:"is_client" json:"is_client"`     // 1大V2用户
	Secret     int8   `bson:"secret" json:"secret"`           // 私密消息1是2否
	RoomId     int    `bson:"room_id" json:"room_id"`         // 房间ID
	Content    string `bson:"content" json:"content"`         // 内容 发图片就存图片地址
	Stocks     string `bson:"stocks" json:"stocks"`           // 股票代码
	IsInteract int8   `bson:"is_interact" json:"is_interact"` // 1直播消息2互动消息
	IsRed      int8   `bson:"is_red" json:"is_red"`           // 是否描红1是2否
	MsgType    int8   `bson:"msg_type" json:"msg_type"`       // 消息类型1图片2文字
	ReplyId    int    `bson:"reply_id" json:"reply_id"`       // 被回复的那条记录ID 0表示非回复
	CTime      int64  `bson:"c_time" json:"c_time"`           // 创建时间
	Nickname   string `bson:"nickname" json:"nickname"`       // 昵称
	Level      int8   `bson:"level" json:"level"`             // 用户等级
	HeadImg    string `bson:"head_img" json:"head_img"`       // 头像地址
	IsSub      int    `bson:"is_sub" json:"is_sub"`           // 是否订阅
	Date       string `bson:"date" json:"date"`               //
	Time       string `bson:"time" json:"time"`               //
}
