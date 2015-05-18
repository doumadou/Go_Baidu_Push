package Model

import (
	. "Go_Baidu_Push/Config"
	. "Go_Baidu_Push/Util"
	"log"
	"strconv"
	"time"
)

const (
	//可选推送类型
	PUSH_TO_USER = "1"
	PUSH_TO_TAG  = "2"
	PUSH_TO_ALL  = "3"

	//Channel 错误常量
	CHANNEL_SDK_SYS                                = 1
	CHANNEL_SDK_INIT_FAIL                          = 2
	CHANNEL_SDK_PARAM                              = 3
	CHANNEL_SDK_HTTP_STATUS_ERROR_AND_RESULT_ERROR = 4
	CHANNEL_SDK_HTTP_STATUS_OK_BUT_RESULT_ERROR    = 5
)

type Channel struct {
	ChannelId       string
	EXPIRES         string
	VERSION         string
	USER_ID         string
	USER_TYPE       string
	DEVICE_TYPE     string
	START           string
	LIMIT           string
	MESSAGES        string
	MSG_IDS         string
	MSG_KEYS        string
	MESSAGE_TYPE    string
	MESSAGE_EXPIRES string
	TAG_NAME        string
	TAG_INFO        string
	TAG_ID          string
	APPID           string
	API_KEY         string
	SECRET_KEY      string

	SIGN         string
	METHOD       string
	HOST         string
	PRODUCT      string
	DEFAULT_HOST string
	NAME         string
	DESCRIPTION  string
	CERT         string
	RELEASE_CERT string
	DEV_CERT     string
	PUSH_TYPE    string
}

func (channel *Channel) queryBindList() {

}

func (channel *Channel) verifyBind() {

}

func (channel *Channel) fetchMessage() {

}

func (channel *Channel) fetchMessageCount() {

}

func (channel *Channel) deleteMessage() {

}

func (channel *Channel) setTag() {

}

func (channel *Channel) fetchTag() {

}

func (channel *Channel) queryUserTag() {

}

func (channel *Channel) queryDeviceType() {

}

func (channel *Channel) deleteTag() {

}

func (channel *Channel) PushMessage() {
	resource := "channel"
	targetURL := BASEURL + resource
	parameters := make(map[string]string)
	parameters["method"] = "push_msg"
	parameters["apikey"] = API_KEY
	parameters["push_type"] = channel.PUSH_TYPE
	parameters["channel_id"] = channel.ChannelId
	parameters["tag"] = channel.TAG_ID
	parameters["deice_type"] = channel.DEVICE_TYPE
	parameters["message_type"] = channel.MESSAGE_TYPE
	parameters["messages"] = channel.MESSAGES
	//parameters["msg_keys"] = channel.mess
	parameters["timestamp"] = strconv.Itoa(time.Now().Unix())
	resp, err := GetURLWithParameters(targetURL, parameters)
	if err != nil {
		log.Println(err.Error())
	}
}
