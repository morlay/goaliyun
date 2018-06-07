package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstancePerformanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ReplicaSetRole       string `position:"Query" name:"ReplicaSetRole"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	NodeId               string `position:"Query" name:"NodeId"`
	Key                  string `position:"Query" name:"Key"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstancePerformanceRequest) Invoke(client goaliyun.Client) (*DescribeDBInstancePerformanceResponse, error) {
	resp := &DescribeDBInstancePerformanceResponse{}
	err := client.Request("dds", "DescribeDBInstancePerformance", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstancePerformanceResponse struct {
	RequestId       goaliyun.String
	DBInstanceId    goaliyun.String
	Engine          goaliyun.String
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	PerformanceKeys DescribeDBInstancePerformancePerformanceKeyList
}

type DescribeDBInstancePerformancePerformanceKey struct {
	Key               goaliyun.String
	Unit              goaliyun.String
	ValueFormat       goaliyun.String
	PerformanceValues DescribeDBInstancePerformancePerformanceValueList
}

type DescribeDBInstancePerformancePerformanceValue struct {
	Value goaliyun.String
	Date  goaliyun.String
}

type DescribeDBInstancePerformancePerformanceKeyList []DescribeDBInstancePerformancePerformanceKey

func (list *DescribeDBInstancePerformancePerformanceKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceKey)
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

type DescribeDBInstancePerformancePerformanceValueList []DescribeDBInstancePerformancePerformanceValue

func (list *DescribeDBInstancePerformancePerformanceValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceValue)
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
