package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeEmrStackVersionRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	MainVersion     string `position:"Query" name:"MainVersion"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeEmrStackVersionRequest) Invoke(client goaliyun.Client) (*DescribeEmrStackVersionResponse, error) {
	resp := &DescribeEmrStackVersionResponse{}
	err := client.Request("emr", "DescribeEmrStackVersion", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeEmrStackVersionResponse struct {
	RequestId    goaliyun.String
	StackVersion goaliyun.String
}
