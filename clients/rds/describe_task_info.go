package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeTaskInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               int64  `position:"Query" name:"TaskId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTaskInfoRequest) Invoke(client goaliyun.Client) (*DescribeTaskInfoResponse, error) {
	resp := &DescribeTaskInfoResponse{}
	err := client.Request("rds", "DescribeTaskInfo", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTaskInfoResponse struct {
	RequestId          goaliyun.String
	TaskId             goaliyun.String
	BeginTime          goaliyun.String
	FinishTime         goaliyun.String
	CreateTime         goaliyun.String
	TaskAction         goaliyun.String
	DBName             goaliyun.String
	TaskErrorCode      goaliyun.String
	Progress           goaliyun.String
	ExpectedFinishTime goaliyun.String
	TaskErrorMessage   goaliyun.String
	ProgressInfo       goaliyun.String
	Status             goaliyun.String
}
