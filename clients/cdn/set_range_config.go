package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetRangeConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetRangeConfigRequest) Invoke(client goaliyun.Client) (*SetRangeConfigResponse, error) {
	resp := &SetRangeConfigResponse{}
	err := client.Request("cdn", "SetRangeConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetRangeConfigResponse struct {
	RequestId goaliyun.String
}
