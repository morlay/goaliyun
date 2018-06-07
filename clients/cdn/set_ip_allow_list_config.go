package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetIpAllowListConfigRequest struct {
	AllowIps      string `position:"Query" name:"AllowIps"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetIpAllowListConfigRequest) Invoke(client goaliyun.Client) (*SetIpAllowListConfigResponse, error) {
	resp := &SetIpAllowListConfigResponse{}
	err := client.Request("cdn", "SetIpAllowListConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetIpAllowListConfigResponse struct {
	RequestId goaliyun.String
}
