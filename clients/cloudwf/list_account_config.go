package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListAccountConfigRequest struct {
	OrderCol    string `position:"Query" name:"OrderCol"`
	Length      int64  `position:"Query" name:"Length"`
	SearchEmail string `position:"Query" name:"SearchEmail"`
	PageIndex   int64  `position:"Query" name:"PageIndex"`
	OrderDir    string `position:"Query" name:"OrderDir"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ListAccountConfigRequest) Invoke(client goaliyun.Client) (*ListAccountConfigResponse, error) {
	resp := &ListAccountConfigResponse{}
	err := client.Request("cloudwf", "ListAccountConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAccountConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
