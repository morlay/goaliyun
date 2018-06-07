package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetCcConfigRequest struct {
	AllowIps      string `position:"Query" name:"AllowIps"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	BlockIps      string `position:"Query" name:"BlockIps"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetCcConfigRequest) Invoke(client goaliyun.Client) (*SetCcConfigResponse, error) {
	resp := &SetCcConfigResponse{}
	err := client.Request("cdn", "SetCcConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetCcConfigResponse struct {
	RequestId goaliyun.String
}
