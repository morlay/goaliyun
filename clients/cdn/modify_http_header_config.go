package cdn

import (
	"github.com/morlay/goaliyun"
)

type ModifyHttpHeaderConfigRequest struct {
	HeaderValue   string `position:"Query" name:"HeaderValue"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	DomainName    string `position:"Query" name:"DomainName"`
	HeaderKey     string `position:"Query" name:"HeaderKey"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ModifyHttpHeaderConfigRequest) Invoke(client goaliyun.Client) (*ModifyHttpHeaderConfigResponse, error) {
	resp := &ModifyHttpHeaderConfigResponse{}
	err := client.Request("cdn", "ModifyHttpHeaderConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyHttpHeaderConfigResponse struct {
	RequestId goaliyun.String
}
