package cdn

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveDomainMappingRequest struct {
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveDomainMappingRequest) Invoke(client goaliyun.Client) (*DeleteLiveDomainMappingResponse, error) {
	resp := &DeleteLiveDomainMappingResponse{}
	err := client.Request("cdn", "DeleteLiveDomainMapping", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveDomainMappingResponse struct {
	RequestId goaliyun.String
}
