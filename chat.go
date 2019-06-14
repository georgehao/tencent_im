package tencentim

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"time"
)

type msgContent struct {
	Text interface{} `json:"Text"`
}

type msgBody struct {
	MsgType    string     `json:"MsgType"`
	MsgContent msgContent `json:"MsgContent"`
}

// AndroidInfo android 离线消息
type AndroidInfo struct {
	Sound string `json:"Sound"`
}

// ApnsInfo ios离线消息
type ApnsInfo struct {
	Sound     string `json:"Sound"`
	BadgeMode int    `json:"BadgeMode"` // 这个字段缺省或者为 0 表示需要计数，为 1 表示本条消息不需要计数，即右上角图标数字不增加
	Title     string `json:"Title"`
	SubTitle  string `json:"SubTitle"`
	Image     string `json:"Image"`
}

// OfflinePushInfo 离线消息
type OfflinePushInfo struct {
	PushFlag    int         `json:"PushFlag"`
	Desc        string      `json:"Desc"`
	Ext         string      `json:"Ext"`
	AndroidInfo AndroidInfo `json:"AndroidInfo"`
	ApnsInfo    ApnsInfo    `json:"ApnsInfo"`
}

// SendMsg 单发单聊消息
type SendMsg struct {
	QueryStringParam QueryStringParam
	SyncOtherMachine int             `json:"SyncOtherMachine"`
	FromAccount      string          `json:"From_Account,omitempty"` // 管理员指定某一帐号向其它帐号发送消息
	ToAccount        string          `json:"To_Account"`
	MsgLifeTime      int             `json:"MsgLifeTime"`
	MsgRandom        uint32          `json:"MsgRandom"`
	MsgTimeStamp     int64           `json:"MsgTimeStamp"`
	MsgBody          []msgBody       `json:"MsgBody"`
	OfflinePushInfo  OfflinePushInfo `json:"OfflinePushInfo,omitempty"` // 同时设置离线推送信息
}

// NewSendMsg 创建单发单聊消息
func NewSendMsg(adminUserSig, toAccount string, content interface{}) *SendMsg {
	qsp := QueryStringParam{
		AppID:   appID,
		UserSig: adminUserSig,
	}

	msg := msgBody{
		MsgType: TIMTextElemMsgType,
		MsgContent: msgContent{
			Text: content.(string),
		},
	}
	return &SendMsg{
		QueryStringParam: qsp,
		SyncOtherMachine: SyncOtherMachineNoSync,
		ToAccount:        toAccount,
		MsgLifeTime:      60,
		MsgRandom:        rand.Uint32(),
		MsgTimeStamp:     time.Now().Unix(),
		MsgBody:          []msgBody{msg},
	}
}

// SetFromAccount 设置FromAccount
func (sm *SendMsg) SetFromAccount(fromAccount string) {
	sm.FromAccount = fromAccount
}

// SetOfflinePushInfo 设置离线消息
func (sm *SendMsg) SetOfflinePushInfo(sopi *OfflinePushInfo) {
	sm.OfflinePushInfo = *sopi
}

// QueryString 返回Query string
func (sm *SendMsg) QueryString() string {
	return sm.QueryStringParam.BuildQueryString()
}

// Name 获取消息类型
func (sm *SendMsg) Name() string {
	return "SendMsg"
}

// Body 返回SendMsg的msg body
func (sm *SendMsg) Body() (error, *bytes.Buffer) {
	bytesData, err := json.Marshal(sm)
	if err != nil {
		return err, nil
	}
	return nil, bytes.NewBuffer(bytesData)
}