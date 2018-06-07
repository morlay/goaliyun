package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeOptimizeAdviceOnStorageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeOptimizeAdviceOnStorageRequest) Invoke(client goaliyun.Client) (*DescribeOptimizeAdviceOnStorageResponse, error) {
	resp := &DescribeOptimizeAdviceOnStorageResponse{}
	err := client.Request("rds", "DescribeOptimizeAdviceOnStorage", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeOptimizeAdviceOnStorageResponse struct {
	RequestId         goaliyun.String
	DBInstanceId      goaliyun.String
	TotalRecordsCount goaliyun.Integer
	PageNumber        goaliyun.Integer
	PageRecordCount   goaliyun.Integer
	Items             DescribeOptimizeAdviceOnStorageAdviceOnStorageList
}

type DescribeOptimizeAdviceOnStorageAdviceOnStorage struct {
	DBName        goaliyun.String
	TableName     goaliyun.String
	CurrentEngine goaliyun.String
	AdviseEngine  goaliyun.String
}

type DescribeOptimizeAdviceOnStorageAdviceOnStorageList []DescribeOptimizeAdviceOnStorageAdviceOnStorage

func (list *DescribeOptimizeAdviceOnStorageAdviceOnStorageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnStorageAdviceOnStorage)
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
