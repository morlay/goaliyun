package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeOptimizeAdviceOnMissPKRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeOptimizeAdviceOnMissPKRequest) Invoke(client goaliyun.Client) (*DescribeOptimizeAdviceOnMissPKResponse, error) {
	resp := &DescribeOptimizeAdviceOnMissPKResponse{}
	err := client.Request("rds", "DescribeOptimizeAdviceOnMissPK", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeOptimizeAdviceOnMissPKResponse struct {
	RequestId         goaliyun.String
	TotalRecordsCount goaliyun.Integer
	PageNumber        goaliyun.Integer
	PageRecordCount   goaliyun.Integer
	Items             DescribeOptimizeAdviceOnMissPKAdviceOnMissPKList
}

type DescribeOptimizeAdviceOnMissPKAdviceOnMissPK struct {
	DBName    goaliyun.String
	TableName goaliyun.String
}

type DescribeOptimizeAdviceOnMissPKAdviceOnMissPKList []DescribeOptimizeAdviceOnMissPKAdviceOnMissPK

func (list *DescribeOptimizeAdviceOnMissPKAdviceOnMissPKList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnMissPKAdviceOnMissPK)
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
