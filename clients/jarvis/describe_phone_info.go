package jarvis

import (
	"github.com/morlay/goaliyun"
)

type DescribePhoneInfoRequest struct {
	SourceIp   string `position:"Query" name:"SourceIp"`
	PhoneNum   string `position:"Query" name:"PhoneNum"`
	Lang       string `position:"Query" name:"Lang"`
	SourceCode string `position:"Query" name:"SourceCode"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribePhoneInfoRequest) Invoke(client goaliyun.Client) (*DescribePhoneInfoResponse, error) {
	resp := &DescribePhoneInfoResponse{}
	err := client.Request("jarvis", "DescribePhoneInfo", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePhoneInfoResponse struct {
	RequestId  goaliyun.String
	Module     goaliyun.String
	PhoneNum   goaliyun.Integer
	RiskLevel  goaliyun.Integer
	DetectTime goaliyun.String
}
