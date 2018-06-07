package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeHpcClustersRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	HpcClusterIds        string `position:"Query" name:"HpcClusterIds"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeHpcClustersRequest) Invoke(client goaliyun.Client) (*DescribeHpcClustersResponse, error) {
	resp := &DescribeHpcClustersResponse{}
	err := client.Request("ecs", "DescribeHpcClusters", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeHpcClustersResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	HpcClusters DescribeHpcClustersHpcClusterList
}

type DescribeHpcClustersHpcCluster struct {
	HpcClusterId goaliyun.String
	Name         goaliyun.String
	Description  goaliyun.String
}

type DescribeHpcClustersHpcClusterList []DescribeHpcClustersHpcCluster

func (list *DescribeHpcClustersHpcClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHpcClustersHpcCluster)
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
