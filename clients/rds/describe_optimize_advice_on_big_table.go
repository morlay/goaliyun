package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeOptimizeAdviceOnBigTableRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeOptimizeAdviceOnBigTableRequest) Invoke(client goaliyun.Client) (*DescribeOptimizeAdviceOnBigTableResponse, error) {
	resp := &DescribeOptimizeAdviceOnBigTableResponse{}
	err := client.Request("rds", "DescribeOptimizeAdviceOnBigTable", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeOptimizeAdviceOnBigTableResponse struct {
	RequestId         goaliyun.String
	TotalRecordsCount goaliyun.Integer
	PageNumber        goaliyun.Integer
	PageRecordCount   goaliyun.Integer
	Items             DescribeOptimizeAdviceOnBigTableAdviceOnBigTableList
}

type DescribeOptimizeAdviceOnBigTableAdviceOnBigTable struct {
	DBName    goaliyun.String
	TableName goaliyun.String
	TableSize goaliyun.Integer
	DataSize  goaliyun.Integer
	IndexSize goaliyun.Integer
}

type DescribeOptimizeAdviceOnBigTableAdviceOnBigTableList []DescribeOptimizeAdviceOnBigTableAdviceOnBigTable

func (list *DescribeOptimizeAdviceOnBigTableAdviceOnBigTableList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnBigTableAdviceOnBigTable)
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
