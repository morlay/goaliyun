package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteCasterRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteCasterRequest) Invoke(client goaliyun.Client) (*DeleteCasterResponse, error) {
	resp := &DeleteCasterResponse{}
	err := client.Request("live", "DeleteCaster", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCasterResponse struct {
	RequestId goaliyun.String
	CasterId  goaliyun.String
}
