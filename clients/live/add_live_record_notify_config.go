package live

import (
	"github.com/morlay/goaliyun"
)

type AddLiveRecordNotifyConfigRequest struct {
	OnDemandUrl      string `position:"Query" name:"OnDemandUrl"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	DomainName       string `position:"Query" name:"DomainName"`
	NotifyUrl        string `position:"Query" name:"NotifyUrl"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	NeedStatusNotify string `position:"Query" name:"NeedStatusNotify"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *AddLiveRecordNotifyConfigRequest) Invoke(client goaliyun.Client) (*AddLiveRecordNotifyConfigResponse, error) {
	resp := &AddLiveRecordNotifyConfigResponse{}
	err := client.Request("live", "AddLiveRecordNotifyConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveRecordNotifyConfigResponse struct {
	RequestId goaliyun.String
}
