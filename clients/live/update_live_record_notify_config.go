package live

import (
	"github.com/morlay/goaliyun"
)

type UpdateLiveRecordNotifyConfigRequest struct {
	OnDemandUrl      string `position:"Query" name:"OnDemandUrl"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	DomainName       string `position:"Query" name:"DomainName"`
	NotifyUrl        string `position:"Query" name:"NotifyUrl"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	NeedStatusNotify string `position:"Query" name:"NeedStatusNotify"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *UpdateLiveRecordNotifyConfigRequest) Invoke(client goaliyun.Client) (*UpdateLiveRecordNotifyConfigResponse, error) {
	resp := &UpdateLiveRecordNotifyConfigResponse{}
	err := client.Request("live", "UpdateLiveRecordNotifyConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateLiveRecordNotifyConfigResponse struct {
	RequestId goaliyun.String
}
