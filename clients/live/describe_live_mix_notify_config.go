package live

import (
	"github.com/morlay/goaliyun"
)

type DescribeLiveMixNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveMixNotifyConfigRequest) Invoke(client goaliyun.Client) (*DescribeLiveMixNotifyConfigResponse, error) {
	resp := &DescribeLiveMixNotifyConfigResponse{}
	err := client.Request("live", "DescribeLiveMixNotifyConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveMixNotifyConfigResponse struct {
	RequestId goaliyun.String
	NotifyUrl goaliyun.String
}
