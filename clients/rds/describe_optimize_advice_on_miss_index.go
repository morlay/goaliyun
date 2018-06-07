package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeOptimizeAdviceOnMissIndexRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeOptimizeAdviceOnMissIndexRequest) Invoke(client goaliyun.Client) (*DescribeOptimizeAdviceOnMissIndexResponse, error) {
	resp := &DescribeOptimizeAdviceOnMissIndexResponse{}
	err := client.Request("rds", "DescribeOptimizeAdviceOnMissIndex", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeOptimizeAdviceOnMissIndexResponse struct {
	RequestId         goaliyun.String
	DBInstanceId      goaliyun.String
	TotalRecordsCount goaliyun.Integer
	PageNumber        goaliyun.Integer
	PageRecordCount   goaliyun.Integer
	Items             DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList
}

type DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex struct {
	DBName      goaliyun.String
	TableName   goaliyun.String
	QueryColumn goaliyun.String
	SQLText     goaliyun.String
}

type DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList []DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex

func (list *DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex)
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
