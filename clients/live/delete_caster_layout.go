package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteCasterLayoutRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	LayoutId string `position:"Query" name:"LayoutId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteCasterLayoutRequest) Invoke(client goaliyun.Client) (*DeleteCasterLayoutResponse, error) {
	resp := &DeleteCasterLayoutResponse{}
	err := client.Request("live", "DeleteCasterLayout", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCasterLayoutResponse struct {
	RequestId goaliyun.String
	CasterId  goaliyun.String
	LayoutId  goaliyun.String
}
