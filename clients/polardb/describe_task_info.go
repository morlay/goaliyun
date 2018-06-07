package polardb

import (
	"github.com/morlay/goaliyun"
)

type DescribeTaskInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               string `position:"Query" name:"TaskId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTaskInfoRequest) Invoke(client goaliyun.Client) (*DescribeTaskInfoResponse, error) {
	resp := &DescribeTaskInfoResponse{}
	err := client.Request("polardb", "DescribeTaskInfo", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTaskInfoResponse struct {
	RequestId          goaliyun.String
	TaskId             goaliyun.String
	BeginTime          goaliyun.String
	FinishTime         goaliyun.String
	ExpectedFinishTime goaliyun.String
	TaskAction         goaliyun.String
	Progress           goaliyun.Integer
	DBName             goaliyun.String
	ProgressInfo       goaliyun.String
}
