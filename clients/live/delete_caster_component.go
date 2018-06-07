package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteCasterComponentRequest struct {
	ComponentId string `position:"Query" name:"ComponentId"`
	CasterId    string `position:"Query" name:"CasterId"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DeleteCasterComponentRequest) Invoke(client goaliyun.Client) (*DeleteCasterComponentResponse, error) {
	resp := &DeleteCasterComponentResponse{}
	err := client.Request("live", "DeleteCasterComponent", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCasterComponentResponse struct {
	RequestId   goaliyun.String
	CasterId    goaliyun.String
	ComponentId goaliyun.String
}
