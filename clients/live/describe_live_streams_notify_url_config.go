package live

import (
	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamsNotifyUrlConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamsNotifyUrlConfigRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamsNotifyUrlConfigResponse, error) {
	resp := &DescribeLiveStreamsNotifyUrlConfigResponse{}
	err := client.Request("live", "DescribeLiveStreamsNotifyUrlConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamsNotifyUrlConfigResponse struct {
	RequestId               goaliyun.String
	LiveStreamsNotifyConfig DescribeLiveStreamsNotifyUrlConfigLiveStreamsNotifyConfig
}

type DescribeLiveStreamsNotifyUrlConfigLiveStreamsNotifyConfig struct {
	DomainName goaliyun.String
	NotifyUrl  goaliyun.String
}
