package cms

import (
	"github.com/morlay/goaliyun"
)

type QueryCustomEventDetailRequest struct {
	QueryJson string `position:"Query" name:"QueryJson"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *QueryCustomEventDetailRequest) Invoke(client goaliyun.Client) (*QueryCustomEventDetailResponse, error) {
	resp := &QueryCustomEventDetailResponse{}
	err := client.Request("cms", "QueryCustomEventDetail", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCustomEventDetailResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
	RequestId goaliyun.String
	Success   goaliyun.String
}
