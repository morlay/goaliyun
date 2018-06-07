package cms

import (
	"github.com/morlay/goaliyun"
)

type QuerySystemEventDetailRequest struct {
	QueryJson string `position:"Query" name:"QueryJson"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *QuerySystemEventDetailRequest) Invoke(client goaliyun.Client) (*QuerySystemEventDetailResponse, error) {
	resp := &QuerySystemEventDetailResponse{}
	err := client.Request("cms", "QuerySystemEventDetail", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QuerySystemEventDetailResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
	RequestId goaliyun.String
	Success   goaliyun.String
}
