package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetWafConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetWafConfigRequest) Invoke(client goaliyun.Client) (*SetWafConfigResponse, error) {
	resp := &SetWafConfigResponse{}
	err := client.Request("cdn", "SetWafConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetWafConfigResponse struct {
	RequestId goaliyun.String
}
