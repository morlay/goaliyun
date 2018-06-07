package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type SPCleanUpRequest struct {
	UserId   int64  `position:"Query" name:"UserId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *SPCleanUpRequest) Invoke(client goaliyun.Client) (*SPCleanUpResponse, error) {
	resp := &SPCleanUpResponse{}
	err := client.Request("phoenixsp-inner", "SPCleanUp", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SPCleanUpResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
}
