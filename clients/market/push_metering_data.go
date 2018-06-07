package market

import (
	"github.com/morlay/goaliyun"
)

type PushMeteringDataRequest struct {
	Metering string `position:"Query" name:"Metering"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *PushMeteringDataRequest) Invoke(client goaliyun.Client) (*PushMeteringDataResponse, error) {
	resp := &PushMeteringDataResponse{}
	err := client.Request("market", "PushMeteringData", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PushMeteringDataResponse struct {
	RequestId goaliyun.String
	Success   bool
}
