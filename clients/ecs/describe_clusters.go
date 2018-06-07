package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClustersRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeClustersRequest) Invoke(client goaliyun.Client) (*DescribeClustersResponse, error) {
	resp := &DescribeClustersResponse{}
	err := client.Request("ecs", "DescribeClusters", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClustersResponse struct {
	RequestId goaliyun.String
	Clusters  DescribeClustersClusterList
}

type DescribeClustersCluster struct {
	ClusterId goaliyun.String
}

type DescribeClustersClusterList []DescribeClustersCluster

func (list *DescribeClustersClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClustersCluster)
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
