package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeReplicaPerformanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SourceDBInstanceId   string `position:"Query" name:"SourceDBInstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	Key                  string `position:"Query" name:"Key"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeReplicaPerformanceRequest) Invoke(client goaliyun.Client) (*DescribeReplicaPerformanceResponse, error) {
	resp := &DescribeReplicaPerformanceResponse{}
	err := client.Request("rds", "DescribeReplicaPerformance", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeReplicaPerformanceResponse struct {
	RequestId       goaliyun.String
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	ReplicaId       goaliyun.String
	PerformanceKeys DescribeReplicaPerformancePerformanceKeys
}

type DescribeReplicaPerformancePerformanceKeys struct {
	PerformanceKey DescribeReplicaPerformancePerformanceKeyItemList
}

type DescribeReplicaPerformancePerformanceKeyItem struct {
	Key               goaliyun.String
	Unit              goaliyun.String
	ValueFormat       goaliyun.String
	PerformanceValues DescribeReplicaPerformancePerformanceValues
}

type DescribeReplicaPerformancePerformanceValues struct {
	PerformanceValue DescribeReplicaPerformancePerformanceValueItemList
}

type DescribeReplicaPerformancePerformanceValueItem struct {
	Value goaliyun.String
	Date  goaliyun.String
}

type DescribeReplicaPerformancePerformanceKeyItemList []DescribeReplicaPerformancePerformanceKeyItem

func (list *DescribeReplicaPerformancePerformanceKeyItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaPerformancePerformanceKeyItem)
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

type DescribeReplicaPerformancePerformanceValueItemList []DescribeReplicaPerformancePerformanceValueItem

func (list *DescribeReplicaPerformancePerformanceValueItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaPerformancePerformanceValueItem)
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
