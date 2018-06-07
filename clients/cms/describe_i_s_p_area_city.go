package cms

import (
	"github.com/morlay/goaliyun"
)

type DescribeISPAreaCityRequest struct {
	City     string `position:"Query" name:"City"`
	Isp      string `position:"Query" name:"Isp"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeISPAreaCityRequest) Invoke(client goaliyun.Client) (*DescribeISPAreaCityResponse, error) {
	resp := &DescribeISPAreaCityResponse{}
	err := client.Request("cms", "DescribeISPAreaCity", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeISPAreaCityResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Success   goaliyun.String
	RequestId goaliyun.String
	Data      goaliyun.String
}
