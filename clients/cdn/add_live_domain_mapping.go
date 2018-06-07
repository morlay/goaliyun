package cdn

import (
	"github.com/morlay/goaliyun"
)

type AddLiveDomainMappingRequest struct {
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *AddLiveDomainMappingRequest) Invoke(client goaliyun.Client) (*AddLiveDomainMappingResponse, error) {
	resp := &AddLiveDomainMappingResponse{}
	err := client.Request("cdn", "AddLiveDomainMapping", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveDomainMappingResponse struct {
	RequestId goaliyun.String
}
