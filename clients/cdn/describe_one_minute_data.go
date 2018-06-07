package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeOneMinuteDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DataTime             string `position:"Query" name:"DataTime"`
	DomainName           string `position:"Query" name:"DomainName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeOneMinuteDataRequest) Invoke(client goaliyun.Client) (*DescribeOneMinuteDataResponse, error) {
	resp := &DescribeOneMinuteDataResponse{}
	err := client.Request("cdn", "DescribeOneMinuteData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeOneMinuteDataResponse struct {
	RequestId goaliyun.String
	Bps       goaliyun.String
	Qps       goaliyun.String
}
