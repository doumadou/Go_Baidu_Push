package pushManager

var _pm *PushManager

func SharedPushManager() *PushManager {
	if _pm == nil {
		_pm = &PushManager{}
	}
	return _pm
}

type PushManager struct {
}

func (p *PushManager) PushToAll() {

}

func (p *PushManager) PushToSingle() {

}

func (p *PushManager) PushToTag() {

}

func (p *PushManager) PushToBatchDevices() {

}

func (p *PushManager) QueryMsgStatus() {

}

func (p *PushManager) QueryTimerRecords() {

}

func (p *PushManager) QueryTopicRecords() {

}

func (p *PushManager) QueryTags() {

}

func (p *PushManager) CreateTag() {

}

func (p *PushManager) DeleteTag() {

}

func (p *PushManager) AddDevicesToTag() {

}

func (p *PushManager) DeleteDevicesFromTag() {

}

func (p *PushManager) QueryNumberOfDevicesInTag() {

}

func (p *PushManager) QueryTimerList() {

}

func (p *PushManager) CancelTimerTask() {

}

func (p *PushManager) QueryTopicList() {

}

func (p *PushManager) QueryDeviceStatistic() {

}

func (p *PushManager) QueryTopicStatistic() {

}
