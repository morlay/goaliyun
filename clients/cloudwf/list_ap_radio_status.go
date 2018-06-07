package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListApRadioStatusRequest struct {
	SearchDisabled      int64  `position:"Query" name:"SearchDisabled"`
	OrderCol            string `position:"Query" name:"OrderCol"`
	SearchName          string `position:"Query" name:"SearchName"`
	SearchChannelEquals int64  `position:"Query" name:"SearchChannelEquals"`
	Length              int64  `position:"Query" name:"Length"`
	SearchMac           string `position:"Query" name:"SearchMac"`
	SearchApgroupName   string `position:"Query" name:"SearchApgroupName"`
	PageIndex           int64  `position:"Query" name:"PageIndex"`
	OrderDir            string `position:"Query" name:"OrderDir"`
	SearchApStatus      int64  `position:"Query" name:"SearchApStatus"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *ListApRadioStatusRequest) Invoke(client goaliyun.Client) (*ListApRadioStatusResponse, error) {
	resp := &ListApRadioStatusResponse{}
	err := client.Request("cloudwf", "ListApRadioStatus", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListApRadioStatusResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
