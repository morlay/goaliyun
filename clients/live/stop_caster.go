package live

import (
	"github.com/morlay/goaliyun"
)

type StopCasterRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *StopCasterRequest) Invoke(client goaliyun.Client) (*StopCasterResponse, error) {
	resp := &StopCasterResponse{}
	err := client.Request("live", "StopCaster", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopCasterResponse struct {
	RequestId goaliyun.String
}
