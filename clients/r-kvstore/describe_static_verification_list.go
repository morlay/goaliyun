package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeStaticVerificationListRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	ReplicaId             string `position:"Query" name:"ReplicaId"`
	DestinationInstanceId string `position:"Query" name:"DestinationInstanceId"`
	SourceInstanceId      string `position:"Query" name:"SourceInstanceId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *DescribeStaticVerificationListRequest) Invoke(client goaliyun.Client) (*DescribeStaticVerificationListResponse, error) {
	resp := &DescribeStaticVerificationListResponse{}
	err := client.Request("r-kvstore", "DescribeStaticVerificationList", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeStaticVerificationListResponse struct {
	RequestId              goaliyun.String
	ReplicaId              goaliyun.String
	SourceInstanceId       goaliyun.String
	SourceDBNumber         goaliyun.Integer
	SourceTableNumber      goaliyun.Integer
	SourceCountNumber      goaliyun.Integer
	SourceDBSize           goaliyun.Integer
	DestinationInstanceId  goaliyun.String
	DestinationDBNumber    goaliyun.Integer
	DestinationTableNumber goaliyun.Integer
	DestinationCountNumber goaliyun.Integer
	DestinationDBSize      goaliyun.Integer
	ConsistencyPercent     goaliyun.String
	Items                  DescribeStaticVerificationListItemsItemList
}

type DescribeStaticVerificationListItemsItem struct {
	AbnormalType      goaliyun.String
	SourceDetail      goaliyun.String
	DestinationDetail goaliyun.String
}

type DescribeStaticVerificationListItemsItemList []DescribeStaticVerificationListItemsItem

func (list *DescribeStaticVerificationListItemsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStaticVerificationListItemsItem)
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
