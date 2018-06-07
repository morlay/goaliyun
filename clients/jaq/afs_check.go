package jaq

import (
	"github.com/morlay/goaliyun"
)

type AfsCheckRequest struct {
	CallerName string `position:"Query" name:"CallerName"`
	Session    string `position:"Query" name:"Session"`
	Token      string `position:"Query" name:"Token"`
	Sig        string `position:"Query" name:"Sig"`
	Platform   int64  `position:"Query" name:"Platform"`
	Scene      string `position:"Query" name:"Scene"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *AfsCheckRequest) Invoke(client goaliyun.Client) (*AfsCheckResponse, error) {
	resp := &AfsCheckResponse{}
	err := client.Request("jaq", "AfsCheck", "2016-11-23", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AfsCheckResponse struct {
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Data      bool
}
