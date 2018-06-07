package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstanceAutoRenewAttributeRequest struct {
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

func (req *DescribeInstanceAutoRenewAttributeRequest) Invoke(client goaliyun.Client) (*DescribeInstanceAutoRenewAttributeResponse, error) {
	resp := &DescribeInstanceAutoRenewAttributeResponse{}
	err := client.Request("rds", "DescribeInstanceAutoRenewAttribute", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceAutoRenewAttributeResponse struct {
	RequestId        goaliyun.String
	PageNumber       goaliyun.Integer
	TotalRecordCount goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeInstanceAutoRenewAttributeItemList
}

type DescribeInstanceAutoRenewAttributeItem struct {
	DBInstanceId goaliyun.String
	RegionId     goaliyun.String
	Duration     goaliyun.Integer
	Status       goaliyun.String
	AutoRenew    goaliyun.String
}

type DescribeInstanceAutoRenewAttributeItemList []DescribeInstanceAutoRenewAttributeItem

func (list *DescribeInstanceAutoRenewAttributeItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAutoRenewAttributeItem)
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
