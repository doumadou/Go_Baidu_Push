package pushManager

import (
	"Go_Baidu_Push/config"
	"Go_Baidu_Push/util"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"strconv"
	"time"
)

var _pm *PushManager

func SharedPushManager() *PushManager {
	if _pm == nil {
		_pm = &PushManager{
			secretKey: config.SECRET_KEY,
			apiKey:    config.API_KEY,
			userAgent: config.USERAGENT,
		}
	}
	return _pm
}

type PushManager struct {
	secretKey string
	apiKey    string
	userAgent string
}

func (p *PushManager) applyBaseParameters(parameters map[string]string) {
	parameters["apikey"] = p.apiKey
	parameters["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
}

func (p *PushManager) PushToAll(device_type, msg_type, msg, deploy_status string, parameters map[string]string) (resp map[string]interface{}, err error) {
	targetURL := "http://api.tuisong.baidu.com/rest/3.0/push/all"
	dic := make(map[string]string)
	dic["device_type"] = device_type
	dic["msg_type"] = msg_type
	dic["deploy_status"] = deploy_status
	p.applyBaseParameters(dic)
	dic["msg"] = util.BuildMessage(msg, parameters, device_type)
	dic["sign"] = util.GenerateSignature("POST", targetURL, p.secretKey, dic)
	return postURL(targetURL, dic)
}

func (p *PushManager) PushToSingle(device_type, channel_id, msg_type, msg, deploy_status string, parameters map[string]string) (resp map[string]interface{}, err error) {
	targetURL := "http://api.tuisong.baidu.com/rest/3.0/push/single_device"
	dic := make(map[string]string)
	p.applyBaseParameters(dic)
	dic["device_type"] = device_type
	dic["msg_type"] = msg_type
	dic["deploy_status"] = deploy_status
	dic["channel_id"] = channel_id
	dic["msg"] = util.BuildMessage(msg, parameters, device_type)
	dic["sign"] = util.GenerateSignature("POST", targetURL, p.secretKey, dic)
	return postURL(targetURL, dic)
}

func (p *PushManager) PushToTag(device_type, tag, msg_type, msg, deploy_status string, parameters map[string]string) {

}

func (p *PushManager) PushToBatchDevices(device_type, channel_ids []string, msg_type, msg, topicId string, parameters map[string]string) {

}

func (p *PushManager) QueryMsgStatus(msgIds []string) {

}

func (p *PushManager) QueryTimerRecords(timerId, start, limit, rangeStart, rangeEnd string) {

}

func (p *PushManager) QueryTopicRecords(topicId, start, limit, rangeStart, rangeEnd string) {

}

func (p *PushManager) QueryTagDetails(tag string) {

}

func (p *PushManager) QueryTags(start, limit string) {

}

func (p *PushManager) CreateTag(tag string) {

}

func (p *PushManager) DeleteTag(tag string) {

}

func (p *PushManager) AddDevicesToTag(tag string, channelIds []string) {

}

func (p *PushManager) DeleteDevicesFromTag(tag string, channelIds []string) {

}

func (p *PushManager) QueryNumberOfDevicesInTag(tag string) {

}

func (p *PushManager) QueryTimerDetails(timerId string) {

}

func (p *PushManager) QueryTimerList(start, limit string) {

}

func (p *PushManager) CancelTimerTask(timerId string) {

}

func (p *PushManager) QueryTopicList(start, limit string) {

}

func (p *PushManager) QueryDeviceStatistic() {

}

func (p *PushManager) QueryTopicStatistic(topicId string) {

}

func postURL(urlString string, dic map[string]string) (resp map[string]interface{}, err error) {
	form := url.Values{}
	for k, v := range dic {
		form.Set(k, v)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", targetURL, bytes.NewBufferString(form.Encode()))
	if err != nil {
		log.Println(err.Error())
		return
	}
	req.Header.Add("User-Agent", p.userAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	response, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer response.Body.Close()
	bd, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(bd))
	err = json.Unmarshal(bd, &resp)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(resp)
	if resp["error_code"] != nil {
		err = errors.New("Push failed")
	}
	return
}
