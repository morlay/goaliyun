package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetDynamicConfigRequest struct {
	DynamicOrigin       string `position:"Query" name:"DynamicOrigin"`
	StaticType          string `position:"Query" name:"StaticType"`
	SecurityToken       string `position:"Query" name:"SecurityToken"`
	StaticUri           string `position:"Query" name:"StaticUri"`
	DomainName          string `position:"Query" name:"DomainName"`
	StaticPath          string `position:"Query" name:"StaticPath"`
	DynamicCacheControl string `position:"Query" name:"DynamicCacheControl"`
	OwnerId             int64  `position:"Query" name:"OwnerId"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *SetDynamicConfigRequest) Invoke(client goaliyun.Client) (*SetDynamicConfigResponse, error) {
	resp := &SetDynamicConfigResponse{}
	err := client.Request("cdn", "SetDynamicConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetDynamicConfigResponse struct {
	RequestId goaliyun.String
}
