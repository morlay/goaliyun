package cdn

import (
	"github.com/morlay/goaliyun"
)

type ModifyDomainCustomLogConfigRequest struct {
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      string `position:"Query" name:"ConfigId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ModifyDomainCustomLogConfigRequest) Invoke(client goaliyun.Client) (*ModifyDomainCustomLogConfigResponse, error) {
	resp := &ModifyDomainCustomLogConfigResponse{}
	err := client.Request("cdn", "ModifyDomainCustomLogConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDomainCustomLogConfigResponse struct {
	RequestId goaliyun.String
}
