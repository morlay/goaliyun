package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeReplicaConflictInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeReplicaConflictInfoRequest) Invoke(client goaliyun.Client) (*DescribeReplicaConflictInfoResponse, error) {
	resp := &DescribeReplicaConflictInfoResponse{}
	err := client.Request("rds", "DescribeReplicaConflictInfo", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeReplicaConflictInfoResponse struct {
	RequestId        goaliyun.String
	ReplicaId        goaliyun.String
	PagNumber        goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	TotalRecordCount goaliyun.Integer
	Items            DescribeReplicaConflictInfoItemsItemList
}

type DescribeReplicaConflictInfoItemsItem struct {
	SourceInstanceId      goaliyun.String
	DestinationInstanceId goaliyun.String
	OccurTime             goaliyun.String
	DetailInfo            goaliyun.String
	ConfictKey            goaliyun.String
	ConfictReason         goaliyun.String
	DatabaseName          goaliyun.String
	RecoveryMode          goaliyun.String
	ConflictGtid          goaliyun.String
}

type DescribeReplicaConflictInfoItemsItemList []DescribeReplicaConflictInfoItemsItem

func (list *DescribeReplicaConflictInfoItemsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaConflictInfoItemsItem)
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
