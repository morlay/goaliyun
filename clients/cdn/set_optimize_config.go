package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetOptimizeConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetOptimizeConfigRequest) Invoke(client goaliyun.Client) (*SetOptimizeConfigResponse, error) {
	resp := &SetOptimizeConfigResponse{}
	err := client.Request("cdn", "SetOptimizeConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetOptimizeConfigResponse struct {
	RequestId goaliyun.String
}
