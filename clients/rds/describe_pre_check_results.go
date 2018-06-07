package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePreCheckResultsRequest struct {
	PreCheckId           string `position:"Query" name:"PreCheckId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribePreCheckResultsRequest) Invoke(client goaliyun.Client) (*DescribePreCheckResultsResponse, error) {
	resp := &DescribePreCheckResultsResponse{}
	err := client.Request("rds", "DescribePreCheckResults", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePreCheckResultsResponse struct {
	RequestId    goaliyun.String
	DBInstanceId goaliyun.String
	Items        DescribePreCheckResultsPreCheckResultList
}

type DescribePreCheckResultsPreCheckResult struct {
	PreCheckName   goaliyun.String
	PreCheckResult goaliyun.String
	FailReasion    goaliyun.String
	RepairMethod   goaliyun.String
}

type DescribePreCheckResultsPreCheckResultList []DescribePreCheckResultsPreCheckResult

func (list *DescribePreCheckResultsPreCheckResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePreCheckResultsPreCheckResult)
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
