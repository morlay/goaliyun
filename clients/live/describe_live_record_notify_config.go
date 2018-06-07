package live

import (
	"github.com/morlay/goaliyun"
)

type DescribeLiveRecordNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveRecordNotifyConfigRequest) Invoke(client goaliyun.Client) (*DescribeLiveRecordNotifyConfigResponse, error) {
	resp := &DescribeLiveRecordNotifyConfigResponse{}
	err := client.Request("live", "DescribeLiveRecordNotifyConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveRecordNotifyConfigResponse struct {
	RequestId              goaliyun.String
	LiveRecordNotifyConfig DescribeLiveRecordNotifyConfigLiveRecordNotifyConfig
}

type DescribeLiveRecordNotifyConfigLiveRecordNotifyConfig struct {
	DomainName       goaliyun.String
	NotifyUrl        goaliyun.String
	OnDemandUrl      goaliyun.String
	NeedStatusNotify bool
}
