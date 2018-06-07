package drds

import (
	"github.com/morlay/goaliyun"
)

type DescribeShardDbConnectionInfoRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	SubDbName      string `position:"Query" name:"SubDbName"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeShardDbConnectionInfoRequest) Invoke(client goaliyun.Client) (*DescribeShardDbConnectionInfoResponse, error) {
	resp := &DescribeShardDbConnectionInfoResponse{}
	err := client.Request("drds", "DescribeShardDbConnectionInfo", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeShardDbConnectionInfoResponse struct {
	RequestId      goaliyun.String
	Success        bool
	ConnectionInfo DescribeShardDbConnectionInfoConnectionInfo
}

type DescribeShardDbConnectionInfoConnectionInfo struct {
	InstanceName               goaliyun.String
	InstanceUrl                goaliyun.String
	SubDbName                  goaliyun.String
	DbStatus                   goaliyun.String
	DbType                     goaliyun.String
	MinPoolSize                goaliyun.Integer
	MaxPoolSize                goaliyun.Integer
	IdleTimeOut                goaliyun.Integer
	BlockingTimeout            goaliyun.Integer
	ConnectionProperties       goaliyun.String
	PreparedStatementCacheSize goaliyun.Integer
	UserName                   goaliyun.String
}
