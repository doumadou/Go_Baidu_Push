package pushManager

import (
	"Go_Baidu_Push/config"
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

func (p *PushManager) PushToAll(msg_type, msg, deploy_status string, parameters *map[string]string) {
	parameters["device_type"] = "3"
	parameters["msg_type"] = msg_type
}

func (p *PushManager) PushToSingle(channel_id, msg_type, msg, deploy_status string, parameters *map[string]string) {

}

func (p *PushManager) PushToTag(tag, msg_type, msg, deploy_status string, parameters *map[string]string) {

}

func (p *PushManager) PushToBatchDevices(channel_ids []string, msg_type, msg, topicId string, parameters *map[string]string) {

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
