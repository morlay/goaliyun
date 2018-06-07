package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetReqHeaderConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      int64  `position:"Query" name:"ConfigId"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Value         string `position:"Query" name:"Value"`
	Key           string `position:"Query" name:"Key"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetReqHeaderConfigRequest) Invoke(client goaliyun.Client) (*SetReqHeaderConfigResponse, error) {
	resp := &SetReqHeaderConfigResponse{}
	err := client.Request("cdn", "SetReqHeaderConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetReqHeaderConfigResponse struct {
	RequestId goaliyun.String
}
