package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetGroupApRadioOnoffProgressRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetGroupApRadioOnoffProgressRequest) Invoke(client goaliyun.Client) (*GetGroupApRadioOnoffProgressResponse, error) {
	resp := &GetGroupApRadioOnoffProgressResponse{}
	err := client.Request("cloudwf", "GetGroupApRadioOnoffProgress", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetGroupApRadioOnoffProgressResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
