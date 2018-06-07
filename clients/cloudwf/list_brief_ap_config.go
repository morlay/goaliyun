package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListBriefApConfigRequest struct {
	SearchScan  int64  `position:"Query" name:"SearchScan"`
	OrderCol    string `position:"Query" name:"OrderCol"`
	SearchName  string `position:"Query" name:"SearchName"`
	Length      int64  `position:"Query" name:"Length"`
	SearchMac   string `position:"Query" name:"SearchMac"`
	PageIndex   int64  `position:"Query" name:"PageIndex"`
	OrderDir    string `position:"Query" name:"OrderDir"`
	SearchModel string `position:"Query" name:"SearchModel"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ListBriefApConfigRequest) Invoke(client goaliyun.Client) (*ListBriefApConfigResponse, error) {
	resp := &ListBriefApConfigResponse{}
	err := client.Request("cloudwf", "ListBriefApConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListBriefApConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
