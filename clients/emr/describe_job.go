package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeJobRequest) Invoke(client goaliyun.Client) (*DescribeJobResponse, error) {
	resp := &DescribeJobResponse{}
	err := client.Request("emr", "DescribeJob", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeJobResponse struct {
	RequestId     goaliyun.String
	Id            goaliyun.String
	Name          goaliyun.String
	FailAct       goaliyun.String
	Type          goaliyun.String
	MaxRetry      goaliyun.Integer
	RetryInterval goaliyun.Integer
	RunParameter  goaliyun.String
}
