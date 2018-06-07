package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstancesByPerformanceRequest struct {
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
	SortKey              string `position:"Query" name:"SortKey"`
	SortMethod           string `position:"Query" name:"SortMethod"`
	Tag2value            string `position:"Query" name:"Tag.2.value"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	Tag4key              string `position:"Query" name:"Tag.4.key"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Tag3value            string `position:"Query" name:"Tag.3.value"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstancesByPerformanceRequest) Invoke(client goaliyun.Client) (*DescribeDBInstancesByPerformanceResponse, error) {
	resp := &DescribeDBInstancesByPerformanceResponse{}
	err := client.Request("rds", "DescribeDBInstancesByPerformance", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstancesByPerformanceResponse struct {
	RequestId        goaliyun.String
	PageNumber       goaliyun.Integer
	TotalRecordCount goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeDBInstancesByPerformanceDBInstancePerformanceList
}

type DescribeDBInstancesByPerformanceDBInstancePerformance struct {
	CPUUsage              goaliyun.String
	IOPSUsage             goaliyun.String
	DiskUsage             goaliyun.String
	SessionUsage          goaliyun.String
	DBInstanceId          goaliyun.String
	DBInstanceDescription goaliyun.String
}

type DescribeDBInstancesByPerformanceDBInstancePerformanceList []DescribeDBInstancesByPerformanceDBInstancePerformance

func (list *DescribeDBInstancesByPerformanceDBInstancePerformanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesByPerformanceDBInstancePerformance)
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
