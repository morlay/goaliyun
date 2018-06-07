package jaq

import (
	"github.com/morlay/goaliyun"
)

type AfsAppCheckRequest struct {
	CallerName string `position:"Query" name:"CallerName"`
	Session    string `position:"Query" name:"Session"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *AfsAppCheckRequest) Invoke(client goaliyun.Client) (*AfsAppCheckResponse, error) {
	resp := &AfsAppCheckResponse{}
	err := client.Request("jaq", "AfsAppCheck", "2016-11-23", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AfsAppCheckResponse struct {
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Data      AfsAppCheckData
}

type AfsAppCheckData struct {
	SecondCheckResult goaliyun.Integer
}
