package live

import (
	"github.com/morlay/goaliyun"
)

type DescribeLiveDetectNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveDetectNotifyConfigRequest) Invoke(client goaliyun.Client) (*DescribeLiveDetectNotifyConfigResponse, error) {
	resp := &DescribeLiveDetectNotifyConfigResponse{}
	err := client.Request("live", "DescribeLiveDetectNotifyConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveDetectNotifyConfigResponse struct {
	RequestId              goaliyun.String
	LiveDetectNotifyConfig DescribeLiveDetectNotifyConfigLiveDetectNotifyConfig
}

type DescribeLiveDetectNotifyConfigLiveDetectNotifyConfig struct {
	DomainName goaliyun.String
	NotifyUrl  goaliyun.String
}
