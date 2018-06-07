package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstancePerformanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Key                  string `position:"Query" name:"Key"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstancePerformanceRequest) Invoke(client goaliyun.Client) (*DescribeDBInstancePerformanceResponse, error) {
	resp := &DescribeDBInstancePerformanceResponse{}
	err := client.Request("polardb", "DescribeDBInstancePerformance", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstancePerformanceResponse struct {
	RequestId       goaliyun.String
	DBInstanceId    goaliyun.String
	Engine          goaliyun.String
	DBType          goaliyun.String
	DBVersion       goaliyun.String
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	PerformanceKeys DescribeDBInstancePerformancePerformanceItemList
}

type DescribeDBInstancePerformancePerformanceItem struct {
	MetricName  goaliyun.String
	Measurement goaliyun.String
	Points      DescribeDBInstancePerformancePerformanceItemValueList
}

type DescribeDBInstancePerformancePerformanceItemValue struct {
	Value     goaliyun.Float
	Timestamp goaliyun.Integer
}

type DescribeDBInstancePerformancePerformanceItemList []DescribeDBInstancePerformancePerformanceItem

func (list *DescribeDBInstancePerformancePerformanceItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceItem)
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

type DescribeDBInstancePerformancePerformanceItemValueList []DescribeDBInstancePerformancePerformanceItemValue

func (list *DescribeDBInstancePerformancePerformanceItemValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceItemValue)
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
