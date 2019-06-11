package chat

import (
	"time"
	"go.uber.org/zap"
	"runtime/debug"
	"reflect"
	"fmt"
	"github.com/gorilla/websocket"
	"encoding/json"
	"carp.cn/whale/zaplogger"
	"carp.cn/whale/pkg/network"
	"carp.cn/whale/com"
	"sync/atomic"
	"carp.cn/whale/pkg/rpc"
	"carp.cn/whale/pkg/cerr"
)

const (
	// 16s最多请求256个包
	RPLimit    = 256
	RPInterval = 16
)

type Crypt func([]byte) []byte
type Option func(client *Client)



type Client struct {
	id   uint32          // 连接数
	conn *network.WSConn // 和客户端的链接

	packetCnt           int   // 16s内发了多少个包
	lastUpdateLimitTime int64 // 最后发包限制时间

	Encrypt Crypt // 加密函数
	Decrypt Crypt // 解密函数
	User    *User // 绑定的信息

	closeFlg uint32 // 关闭标记
	rpc      *rpc.RpcHelper
}

func NewClient(conn network.Conn, rpc *rpc.RpcHelper, opts ...Option) *Client {
	c := new(Client)
	c.conn = conn.(*network.WSConn)
	c.id = DefaultCounter.getId()
	c.lastUpdateLimitTime = time.Now().Unix()
	c.rpc = rpc
	c.User = &User{
		loginTime: time.Now().Unix(),
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) CheckRPLimit() bool {
	c.packetCnt++
	if time.Now().Unix()-c.lastUpdateLimitTime > RPInterval {
		c.lastUpdateLimitTime = time.Now().Unix()
		if c.packetCnt > RPLimit {
			zaplogger.Error("WSS2Client CheckRPLimit found exception.",
				zap.Int("userId:", c.User.GetUid()), zap.Int(" packetCnt:", c.packetCnt))
			return false
		}
		zaplogger.Debug("WSS2Client CheckRPLimit cur packet.",
			zap.Int("userId:", c.User.GetUid()), zap.Int(" packetCnt:", c.packetCnt))
		c.packetCnt = 0
	}

	return true
}

func (c *Client) Run() {
	zaplogger.Info("Begin  ClientAgent.Run............", zap.Uint32("Id:", c.id))
	exit := make(chan struct{})
	defer func() {
		if err := recover(); err != nil {
			zaplogger.Error("Client run panic.", zap.Reflect("Error:", err), zap.String("\nstack", string(debug.Stack())))
		}

		// TODO:清除集合
		close(exit)
		c.CloseFlag()
	}()

	// write
	go func() {
		ticker := time.NewTicker(network.PingPeriod)
		defer func() {
			if err := recover(); err != nil {
				zaplogger.Error("Recover from panic:", zap.Reflect("Error:", err))
			}
			ticker.Stop()
			c.CloseFlag()
		}()

		for {
			if c.IsClose() {
				return
			}

			select {
			case b, ok := <-c.conn.WriteChan:
				if err := c.conn.SetWriteDeadline(time.Now().Add(network.WriteWait)); err != nil {
					zaplogger.Error("WSS2Client Run conn.SetWriteDeadline error.",
						zap.Error(err))
				}

				if b == nil || !ok {
					if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
						zaplogger.Error("WSS2Client Run conn.WriteMessage error:", zap.Error(err),
							zap.Int(" UID:", c.User.GetUid()), zap.Uint32(" id:", c.id))
					}
					return
				}

				//zaplogger.Info("Write MSG", zap.Uint32("id:", c.id), zap.Int(" uid:", c.User.GetUid()), zap.String(" [OUT]:", string(b)))

				if c.Encrypt != nil {
					b = c.Encrypt(b)
				}

				err := c.conn.WriteMessage(websocket.TextMessage, b)
				if err != nil {
					zaplogger.Error("write msg err :", zap.String(" err === ", err.Error()))
					if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
						zaplogger.Error("WSS2Client Run conn.WriteMessage err.",
							zap.Error(err))
					}
					return
				}

			case <-ticker.C:
				if err := c.conn.SetWriteDeadline(time.Now().Add(network.WriteWait)); err != nil {
					zaplogger.Error("WSS2Client Run conn.SetWriteDeadline err.",
						zap.Error(err))
				}
				if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					zaplogger.Error("WSS2Client Run conn.WriteMessage err.",
						zap.Error(err))
					return
				}
			case <-exit:
				return
			}
		}
	}()

	// read
	for {
		if c.IsClose() {
			return
		}

		data, err := c.conn.ReadMsg()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				zaplogger.Error("ws unexpected close:", zap.Error(err),
					zap.String(" addr = ", c.conn.RemoteAddr().String()),
					zap.Uint32(" id: ", c.id))
			} else {
				zaplogger.Error("Read message error:", zap.Error(err),
					zap.Uint32(" id:", c.id), zap.Int(" uid:", c.User.GetUid()))
			}
			return
		}

		if !c.CheckRPLimit() {
			return
		}

		if c.Decrypt != nil {
			data = c.Decrypt(data)
		}

		// zaplogger.Info("Read MSG [IN] <= " + string(data))
		zaplogger.Info("Read MSG", zap.Uint32(" id:", c.id), zap.Int(" uid:", c.User.GetUid()), zap.String(" [IN]:", string(data)))

		req := &rpc.Request{}
		jsonErr := json.Unmarshal(data, req)
		if jsonErr != nil {
			zaplogger.Error("json.Unmarshal err: ", zap.Reflect(" msg:", string(data)), zap.String(" Error", jsonErr.Error()))
			return
		}

		if !c.User.isLogin {
			if !c.login(req) {
				if err := c.Send(com.UserKickOut, map[string]interface{}{"state": com.StateKickoutLogin}); err != nil {
					zaplogger.Error("ClientAgent Run a.WriteMsg error.",
						zap.Error(err))
				}
				// close client
				time.Sleep(time.Millisecond * 100) // 保证客户端能收到连接断开的消息
				return
			}

			if checkUserRoomAvailable(c.User) {
				BroadcastChat("UserLogin", c.User.room.Info.Id, c.User.GetUid(), c.User.nickname, c.User.imei)
			}

			continue
		}
		c.HandleFunc(req)
	}
}

func (c *Client) HandleFunc(req *rpc.Request) {
	defer func() {
		if err := recover(); err != nil {
			zaplogger.Error("Recover from panic:", zap.Reflect(" Error:", err), zap.Int(" uid:", c.User.GetUid()), zap.Uint32(" id:", c.id))
		}
	}()

	rpcResp := c.rpc.Handler(req, c)
	cerr.CheckIError(rpcResp.Error)

	if rpcResp.Result == nil {
		return
	}

	// 将方法替换成对象.方法名
	rpcResp.Method = req.Method

	buf, err := json.Marshal(rpcResp)
	cerr.CheckError(err, cerr.ERR_JSON_MARSHAL)
	cerr.CheckError(c.WriteMsg(buf), cerr.ERR_WRITE_TO_CLIENT)

}

func (c *Client) RemoteAddr() string {
	return c.conn.RemoteAddr().String()
}

//回给客户端
func (c *Client) WriteMsg(msg []byte) error {
	defer func() {
		if err := recover(); err != nil {
			zaplogger.Error("WSS2Client WriteMsg failed.", zap.Reflect(" Error:", err),
				zap.String(" Msg:", string(msg)))
		}
	}()
	if c.conn == nil {
		return fmt.Errorf("WriteMsg conn is nil")
	}

	if c.IsClose() {
		return fmt.Errorf("connection is closed")
	}

	err := c.conn.WriteMsg(msg)
	if err != nil {
		zaplogger.Error("write message error:", zap.Reflect(" type =", reflect.TypeOf(msg)), zap.Reflect(" Error:", err))
		return err
	}
	return nil
}

//强杀
func (c *Client) Close() {

	defer func() {
		if err := recover(); err != nil {
			zaplogger.Error("Client close panic", zap.Reflect(" Error:", err),
				zap.Int(" UID:", c.User.GetUid()))
			debug.PrintStack()
		}

	}()

	c.RemoveSelf()

	c.CloseFlag()
}

func (c *Client) IsClose() bool {
	return atomic.LoadUint32(&c.closeFlg) == 1
}

func (c *Client) CloseFlag() {
	zaplogger.Info("client close", zap.Uint32(" id:", c.id))
	atomic.StoreUint32(&c.closeFlg, 1)
}

func (c *Client) GetId() uint32 {
	return c.id
}

// ----------------------------------------------------

func (c *Client) GetCId() uint32 {
	return c.id
}

func (c *Client) SetCId(id uint32) {
	c.id = id
}

func (c *Client) RemoveSelf() {

	DefaultUserSet.RemoveClient(c.User.GetUid())
	DefaultRoomSet.RemoveUser(c)
	zaplogger.Info("ws conn remove self ----------->", zap.Reflect(" uid:", c.User.GetUid()))
	// logout
	// broadcast

	if c.User.isLogin {
		if checkUserRoomAvailable(c.User) {
			BroadcastChat("UserLogout", c.User.room.Info.Id, c.User.GetUid(), c.User.nickname)
		}
	}
	// 解除绑定
	c.rpc = nil
	c.User.room = nil
}

func (c *Client) Send(funcName string, data map[string]interface{}) error {
	resp := &rpc.Response{
		Method: funcName,
		Result: data,
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	return c.WriteMsg(bytes)
}

func (c *Client) login(req *rpc.Request) bool {
	if req.Method != com.UserLogin {
		zaplogger.Error("First method must be login, try again please.", zap.Uint32(" id", c.id))
		return false
	}

	if len(req.Params) < 3 {
		zaplogger.Error("login len(*(req.Params)) < 3  ", zap.Uint32("id", c.id))
		return false
	}

	uid, ok := req.Params[0].(float64)
	if !ok {
		zaplogger.Error("Login Param 0 error.", zap.Reflect("param1:", req.Params[0]))
		return false
	}

	roomId, ok := req.Params[1].(float64)
	if !ok {
		zaplogger.Error("Login Param 1 error.", zap.Reflect("param2:", req.Params[1]))
		return false
	}

	iemi, ok := req.Params[2].(string)
	if !ok {
		zaplogger.Error("Login Param 2 error.", zap.Reflect("param3:", req.Params[2]))
		return false
	}

	return Login(c, int(uid), int(roomId), iemi)
}
