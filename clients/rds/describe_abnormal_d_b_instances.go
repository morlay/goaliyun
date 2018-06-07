package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAbnormalDBInstancesRequest struct {
	Tag4value            string `position:"Query" name:"Tag.4.value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2key              string `position:"Query" name:"Tag.2.key"`
	Tag5key              string `position:"Query" name:"Tag.5.key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag3key              string `position:"Query" name:"Tag.3.key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5value            string `position:"Query" name:"Tag.5.value"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
	Tag1key              string `position:"Query" name:"Tag.1.key"`
	Tag1value            string `position:"Query" name:"Tag.1.value"`
	Tag2value            string `position:"Query" name:"Tag.2.value"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	Tag4key              string `position:"Query" name:"Tag.4.key"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Tag3value            string `position:"Query" name:"Tag.3.value"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAbnormalDBInstancesRequest) Invoke(client goaliyun.Client) (*DescribeAbnormalDBInstancesResponse, error) {
	resp := &DescribeAbnormalDBInstancesResponse{}
	err := client.Request("rds", "DescribeAbnormalDBInstances", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAbnormalDBInstancesResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeAbnormalDBInstancesInstanceResultList
}

type DescribeAbnormalDBInstancesInstanceResult struct {
	DBInstanceDescription goaliyun.String
	DBInstanceId          goaliyun.String
	AbnormalItems         DescribeAbnormalDBInstancesAbnormalItemList
}

type DescribeAbnormalDBInstancesAbnormalItem struct {
	CheckTime      goaliyun.String
	CheckItem      goaliyun.String
	AbnormalReason goaliyun.String
	AbnormalValue  goaliyun.String
	AbnormalDetail goaliyun.String
	AdviceKey      goaliyun.String
	AdviseValue    DescribeAbnormalDBInstancesAdviseValueList
}

type DescribeAbnormalDBInstancesInstanceResultList []DescribeAbnormalDBInstancesInstanceResult

func (list *DescribeAbnormalDBInstancesInstanceResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAbnormalDBInstancesInstanceResult)
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

type DescribeAbnormalDBInstancesAbnormalItemList []DescribeAbnormalDBInstancesAbnormalItem

func (list *DescribeAbnormalDBInstancesAbnormalItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAbnormalDBInstancesAbnormalItem)
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

type DescribeAbnormalDBInstancesAdviseValueList []goaliyun.String

func (list *DescribeAbnormalDBInstancesAdviseValueList) UnmarshalJSON(data []byte) error {
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
