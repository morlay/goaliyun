package cms

import (
	"github.com/morlay/goaliyun"
)

type QuerySystemEventCountRequest struct {
	QueryJson string `position:"Query" name:"QueryJson"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *QuerySystemEventCountRequest) Invoke(client goaliyun.Client) (*QuerySystemEventCountResponse, error) {
	resp := &QuerySystemEventCountResponse{}
	err := client.Request("cms", "QuerySystemEventCount", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QuerySystemEventCountResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
	RequestId goaliyun.String
	Success   goaliyun.String
}
