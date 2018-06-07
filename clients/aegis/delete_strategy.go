package aegis

import (
	"github.com/morlay/goaliyun"
)

type DeleteStrategyRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteStrategyRequest) Invoke(client goaliyun.Client) (*DeleteStrategyResponse, error) {
	resp := &DeleteStrategyResponse{}
	err := client.Request("aegis", "DeleteStrategy", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteStrategyResponse struct {
	RequestId goaliyun.String
}
