package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListStaStatusRequest struct {
	OrderCol          string `position:"Query" name:"OrderCol"`
	SearchGroupName   string `position:"Query" name:"SearchGroupName"`
	SearchStatus      int64  `position:"Query" name:"SearchStatus"`
	Length            int64  `position:"Query" name:"Length"`
	SearchUsername    string `position:"Query" name:"SearchUsername"`
	OrderDir          string `position:"Query" name:"OrderDir"`
	SearchProtocal    string `position:"Query" name:"SearchProtocal"`
	SearchSsid        string `position:"Query" name:"SearchSsid"`
	SearchApName      string `position:"Query" name:"SearchApName"`
	SearchIp          string `position:"Query" name:"SearchIp"`
	PageIndex         int64  `position:"Query" name:"PageIndex"`
	SearchMac         string `position:"Query" name:"SearchMac"`
	SearchDescription string `position:"Query" name:"SearchDescription"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *ListStaStatusRequest) Invoke(client goaliyun.Client) (*ListStaStatusResponse, error) {
	resp := &ListStaStatusResponse{}
	err := client.Request("cloudwf", "ListStaStatus", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListStaStatusResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
