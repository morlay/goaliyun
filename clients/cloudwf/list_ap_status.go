package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListApStatusRequest struct {
	OrderCol          string `position:"Query" name:"OrderCol"`
	SearchName        string `position:"Query" name:"SearchName"`
	SearchGroupName   string `position:"Query" name:"SearchGroupName"`
	SearchStatus      int64  `position:"Query" name:"SearchStatus"`
	SearchWanIp       string `position:"Query" name:"SearchWanIp"`
	SearchApModelName string `position:"Query" name:"SearchApModelName"`
	Length            int64  `position:"Query" name:"Length"`
	OrderDir          string `position:"Query" name:"OrderDir"`
	SearchBssEquals   int64  `position:"Query" name:"SearchBssEquals"`
	SearchSwVersion   int64  `position:"Query" name:"SearchSwVersion"`
	SearchCompanyName string `position:"Query" name:"SearchCompanyName"`
	SearchMac         string `position:"Query" name:"SearchMac"`
	PageIndex         int64  `position:"Query" name:"PageIndex"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *ListApStatusRequest) Invoke(client goaliyun.Client) (*ListApStatusResponse, error) {
	resp := &ListApStatusResponse{}
	err := client.Request("cloudwf", "ListApStatus", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListApStatusResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
