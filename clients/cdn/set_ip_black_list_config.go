package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetIpBlackListConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	BlockIps      string `position:"Query" name:"BlockIps"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetIpBlackListConfigRequest) Invoke(client goaliyun.Client) (*SetIpBlackListConfigResponse, error) {
	resp := &SetIpBlackListConfigResponse{}
	err := client.Request("cdn", "SetIpBlackListConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetIpBlackListConfigResponse struct {
	RequestId goaliyun.String
}
