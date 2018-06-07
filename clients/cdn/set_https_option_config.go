package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetHttpsOptionConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Http2         string `position:"Query" name:"Http.2"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetHttpsOptionConfigRequest) Invoke(client goaliyun.Client) (*SetHttpsOptionConfigResponse, error) {
	resp := &SetHttpsOptionConfigResponse{}
	err := client.Request("cdn", "SetHttpsOptionConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetHttpsOptionConfigResponse struct {
	RequestId goaliyun.String
}
