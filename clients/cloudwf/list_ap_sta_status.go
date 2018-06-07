package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListApStaStatusRequest struct {
	OrderCol       string `position:"Query" name:"OrderCol"`
	SearchProtocal string `position:"Query" name:"SearchProtocal"`
	SearchSsid     string `position:"Query" name:"SearchSsid"`
	SearchIp       string `position:"Query" name:"SearchIp"`
	Length         int64  `position:"Query" name:"Length"`
	SearchUsername string `position:"Query" name:"SearchUsername"`
	SearchMac      string `position:"Query" name:"SearchMac"`
	PageIndex      int64  `position:"Query" name:"PageIndex"`
	Id             int64  `position:"Query" name:"Id"`
	OrderDir       string `position:"Query" name:"OrderDir"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ListApStaStatusRequest) Invoke(client goaliyun.Client) (*ListApStaStatusResponse, error) {
	resp := &ListApStaStatusResponse{}
	err := client.Request("cloudwf", "ListApStaStatus", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListApStaStatusResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
