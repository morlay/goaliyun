package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstanceAutoRenewalAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceAutoRenewalAttributeRequest) Invoke(client goaliyun.Client) (*DescribeInstanceAutoRenewalAttributeResponse, error) {
	resp := &DescribeInstanceAutoRenewalAttributeResponse{}
	err := client.Request("rds", "DescribeInstanceAutoRenewalAttribute", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceAutoRenewalAttributeResponse struct {
	RequestId        goaliyun.String
	PageNumber       goaliyun.Integer
	TotalRecordCount goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeInstanceAutoRenewalAttributeItemList
}

type DescribeInstanceAutoRenewalAttributeItem struct {
	DBInstanceId goaliyun.String
	RegionId     goaliyun.String
	Duration     goaliyun.Integer
	Status       goaliyun.String
	AutoRenew    goaliyun.String
}

type DescribeInstanceAutoRenewalAttributeItemList []DescribeInstanceAutoRenewalAttributeItem

func (list *DescribeInstanceAutoRenewalAttributeItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAutoRenewalAttributeItem)
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
