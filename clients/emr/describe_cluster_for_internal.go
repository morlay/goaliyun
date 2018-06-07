package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeClusterForInternalRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterForInternalRequest) Invoke(client goaliyun.Client) (*DescribeClusterForInternalResponse, error) {
	resp := &DescribeClusterForInternalResponse{}
	err := client.Request("emr", "DescribeClusterForInternal", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterForInternalResponse struct {
	RequestId   goaliyun.String
	Id          goaliyun.String
	BizId       goaliyun.String
	Name        goaliyun.String
	StartTime   goaliyun.Integer
	StopTime    goaliyun.Integer
	LogEnable   bool
	LogPath     goaliyun.String
	UserId      goaliyun.String
	RunningTime goaliyun.Integer
	Status      goaliyun.String
	ExpiredTime goaliyun.Integer
	FailReason  DescribeClusterForInternalFailReason
}

type DescribeClusterForInternalFailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}
