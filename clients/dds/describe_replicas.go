package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeReplicasRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeReplicasRequest) Invoke(client goaliyun.Client) (*DescribeReplicasResponse, error) {
	resp := &DescribeReplicasResponse{}
	err := client.Request("dds", "DescribeReplicas", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeReplicasResponse struct {
	RequestId        goaliyun.String
	PageNumber       goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Replicas         DescribeReplicasItemsList
}

type DescribeReplicasItems struct {
	ReplicaId          goaliyun.String
	ReplicaDescription goaliyun.String
	ReplicaStatus      goaliyun.String
	DBInstances        DescribeReplicasItems1List
}

type DescribeReplicasItems1 struct {
	DBInstanceId goaliyun.String
	Role         goaliyun.String
}

type DescribeReplicasItemsList []DescribeReplicasItems

func (list *DescribeReplicasItemsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicasItems)
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

type DescribeReplicasItems1List []DescribeReplicasItems1

func (list *DescribeReplicasItems1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicasItems1)
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
