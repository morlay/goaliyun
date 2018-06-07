package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeOptimizeAdviceOnExcessIndexRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeOptimizeAdviceOnExcessIndexRequest) Invoke(client goaliyun.Client) (*DescribeOptimizeAdviceOnExcessIndexResponse, error) {
	resp := &DescribeOptimizeAdviceOnExcessIndexResponse{}
	err := client.Request("rds", "DescribeOptimizeAdviceOnExcessIndex", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeOptimizeAdviceOnExcessIndexResponse struct {
	RequestId         goaliyun.String
	TotalRecordsCount goaliyun.Integer
	PageNumber        goaliyun.Integer
	PageRecordCount   goaliyun.Integer
	Items             DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndexList
}

type DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndex struct {
	DBName     goaliyun.String
	TableName  goaliyun.String
	IndexCount goaliyun.Integer
}

type DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndexList []DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndex

func (list *DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndexList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndex)
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
