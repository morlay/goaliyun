package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type GetResultCountRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetResultCountRequest) Invoke(client goaliyun.Client) (*GetResultCountResponse, error) {
	resp := &GetResultCountResponse{}
	err := client.Request("qualitycheck", "GetResultCount", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetResultCountResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.Integer
}
