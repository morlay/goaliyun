package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeHaVipsRequest struct {
	Filters              *DescribeHaVipsFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                     `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                    `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                    `position:"Query" name:"OwnerAccount"`
	PageSize             int64                     `position:"Query" name:"PageSize"`
	OwnerId              int64                     `position:"Query" name:"OwnerId"`
	PageNumber           int64                     `position:"Query" name:"PageNumber"`
	RegionId             string                    `position:"Query" name:"RegionId"`
}

func (req *DescribeHaVipsRequest) Invoke(client goaliyun.Client) (*DescribeHaVipsResponse, error) {
	resp := &DescribeHaVipsResponse{}
	err := client.Request("ecs", "DescribeHaVips", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeHaVipsFilter struct {
	Key    string                   `name:"Key"`
	Values *DescribeHaVipsValueList `type:"Repeated" name:"Value"`
}

type DescribeHaVipsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	HaVips     DescribeHaVipsHaVipList
}

type DescribeHaVipsHaVip struct {
	HaVipId                goaliyun.String
	RegionId               goaliyun.String
	VpcId                  goaliyun.String
	VSwitchId              goaliyun.String
	IpAddress              goaliyun.String
	Status                 goaliyun.String
	MasterInstanceId       goaliyun.String
	Description            goaliyun.String
	CreateTime             goaliyun.String
	AssociatedInstances    DescribeHaVipsAssociatedInstanceList
	AssociatedEipAddresses DescribeHaVipsAssociatedEipAddressList
}

type DescribeHaVipsFilterList []DescribeHaVipsFilter

func (list *DescribeHaVipsFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHaVipsFilter)
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

type DescribeHaVipsValueList []string

func (list *DescribeHaVipsValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeHaVipsHaVipList []DescribeHaVipsHaVip

func (list *DescribeHaVipsHaVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHaVipsHaVip)
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

type DescribeHaVipsAssociatedInstanceList []goaliyun.String

func (list *DescribeHaVipsAssociatedInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type DescribeHaVipsAssociatedEipAddressList []goaliyun.String

func (list *DescribeHaVipsAssociatedEipAddressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
