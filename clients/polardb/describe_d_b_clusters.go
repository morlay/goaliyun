package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBClustersRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBType               string `position:"Query" name:"DBType"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBClustersRequest) Invoke(client goaliyun.Client) (*DescribeDBClustersResponse, error) {
	resp := &DescribeDBClustersResponse{}
	err := client.Request("polardb", "DescribeDBClusters", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBClustersResponse struct {
	RequestId        goaliyun.String
	PageNumber       goaliyun.Integer
	TotalRecordCount goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeDBClustersDBClusterList
}

type DescribeDBClustersDBCluster struct {
	DBClusterId          goaliyun.String
	DBClusterDescription goaliyun.String
	PayType              goaliyun.String
	DBClusterNetworkType goaliyun.String
	RegionId             goaliyun.String
	ExpireTime           goaliyun.String
	DBClusterStatus      goaliyun.String
	Engine               goaliyun.String
	DBType               goaliyun.String
	DBVersion            goaliyun.String
	LockMode             goaliyun.String
	LockReason           goaliyun.String
	CreateTime           goaliyun.String
	VpcId                goaliyun.String
	VSwitchId            goaliyun.String
}

type DescribeDBClustersDBClusterList []DescribeDBClustersDBCluster

func (list *DescribeDBClustersDBClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClustersDBCluster)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
