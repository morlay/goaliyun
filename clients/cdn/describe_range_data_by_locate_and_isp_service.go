package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeRangeDataByLocateAndIspServiceRequest struct {
	IspNames      string `position:"Query" name:"IspNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainNames   string `position:"Query" name:"DomainNames"`
	LocationNames string `position:"Query" name:"LocationNames"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeRangeDataByLocateAndIspServiceRequest) Invoke(client goaliyun.Client) (*DescribeRangeDataByLocateAndIspServiceResponse, error) {
	resp := &DescribeRangeDataByLocateAndIspServiceResponse{}
	err := client.Request("cdn", "DescribeRangeDataByLocateAndIspService", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRangeDataByLocateAndIspServiceResponse struct {
	RequestId  goaliyun.String
	JsonResult goaliyun.String
}
