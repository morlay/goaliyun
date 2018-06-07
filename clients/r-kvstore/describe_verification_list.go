package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVerificationListRequest struct {
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

func (req *DescribeVerificationListRequest) Invoke(client goaliyun.Client) (*DescribeVerificationListResponse, error) {
	resp := &DescribeVerificationListResponse{}
	err := client.Request("r-kvstore", "DescribeVerificationList", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVerificationListResponse struct {
	RequestId        goaliyun.String
	ReplicaId        goaliyun.String
	PagNumber        goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	TotalRecordCount goaliyun.Integer
	Items            DescribeVerificationListItemsItemList
}

type DescribeVerificationListItemsItem struct {
	InstanceIdA        goaliyun.String
	InstanceIdB        goaliyun.String
	Key                goaliyun.String
	KeyType            goaliyun.String
	InconsistentType   goaliyun.String
	OccurTime          goaliyun.String
	Schema             goaliyun.String
	InconsistentFields goaliyun.String
}

type DescribeVerificationListItemsItemList []DescribeVerificationListItemsItem

func (list *DescribeVerificationListItemsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVerificationListItemsItem)
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
