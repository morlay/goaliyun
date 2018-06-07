package push

import (
	"github.com/morlay/goaliyun"
)

type PushRequest struct {
	AndroidNotificationBarType     int64  `position:"Query" name:"AndroidNotificationBarType"`
	SmsSendPolicy                  int64  `position:"Query" name:"SmsSendPolicy"`
	AndroidExtParameters           string `position:"Query" name:"AndroidExtParameters"`
	IOSBadge                       int64  `position:"Query" name:"IOSBadge"`
	IOSBadgeAutoIncrement          string `position:"Query" name:"IOSBadgeAutoIncrement"`
	AndroidOpenType                string `position:"Query" name:"AndroidOpenType"`
	Title                          string `position:"Query" name:"Title"`
	Body                           string `position:"Query" name:"Body"`
	DeviceType                     string `position:"Query" name:"DeviceType"`
	PushTime                       string `position:"Query" name:"PushTime"`
	SmsDelaySecs                   int64  `position:"Query" name:"SmsDelaySecs"`
	SendSpeed                      int64  `position:"Query" name:"SendSpeed"`
	AndroidPopupActivity           string `position:"Query" name:"AndroidPopupActivity"`
	IOSRemindBody                  string `position:"Query" name:"IOSRemindBody"`
	IOSExtParameters               string `position:"Query" name:"IOSExtParameters"`
	AndroidNotifyType              string `position:"Query" name:"AndroidNotifyType"`
	AndroidPopupTitle              string `position:"Query" name:"AndroidPopupTitle"`
	IOSMusic                       string `position:"Query" name:"IOSMusic"`
	IOSApnsEnv                     string `position:"Query" name:"IOSApnsEnv"`
	IOSMutableContent              string `position:"Query" name:"IOSMutableContent"`
	AndroidNotificationBarPriority int64  `position:"Query" name:"AndroidNotificationBarPriority"`
	ExpireTime                     string `position:"Query" name:"ExpireTime"`
	SmsTemplateName                string `position:"Query" name:"SmsTemplateName"`
	AndroidPopupBody               string `position:"Query" name:"AndroidPopupBody"`
	IOSNotificationCategory        string `position:"Query" name:"IOSNotificationCategory"`
	StoreOffline                   string `position:"Query" name:"StoreOffline"`
	IOSSilentNotification          string `position:"Query" name:"IOSSilentNotification"`
	SmsParams                      string `position:"Query" name:"SmsParams"`
	JobKey                         string `position:"Query" name:"JobKey"`
	Target                         string `position:"Query" name:"Target"`
	AndroidOpenUrl                 string `position:"Query" name:"AndroidOpenUrl"`
	AndroidNotificationChannel     string `position:"Query" name:"AndroidNotificationChannel"`
	AndroidRemind                  string `position:"Query" name:"AndroidRemind"`
	AndroidActivity                string `position:"Query" name:"AndroidActivity"`
	AndroidXiaoMiNotifyBody        string `position:"Query" name:"AndroidXiaoMiNotifyBody"`
	IOSSubtitle                    string `position:"Query" name:"IOSSubtitle"`
	SmsSignName                    string `position:"Query" name:"SmsSignName"`
	IOSRemind                      string `position:"Query" name:"IOSRemind"`
	AppKey                         int64  `position:"Query" name:"AppKey"`
	TargetValue                    string `position:"Query" name:"TargetValue"`
	AndroidMusic                   string `position:"Query" name:"AndroidMusic"`
	AndroidXiaoMiActivity          string `position:"Query" name:"AndroidXiaoMiActivity"`
	AndroidXiaoMiNotifyTitle       string `position:"Query" name:"AndroidXiaoMiNotifyTitle"`
	PushType                       string `position:"Query" name:"PushType"`
	RegionId                       string `position:"Query" name:"RegionId"`
}

func (req *PushRequest) Invoke(client goaliyun.Client) (*PushResponse, error) {
	resp := &PushResponse{}
	err := client.Request("push", "Push", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PushResponse struct {
	RequestId goaliyun.String
	MessageId goaliyun.String
}
