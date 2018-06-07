package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAccountsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAccountsRequest) Invoke(client goaliyun.Client) (*DescribeAccountsResponse, error) {
	resp := &DescribeAccountsResponse{}
	err := client.Request("dds", "DescribeAccounts", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAccountsResponse struct {
	RequestId goaliyun.String
	Accounts  DescribeAccountsAccountList
}

type DescribeAccountsAccount struct {
	DBInstanceId       goaliyun.String
	AccountName        goaliyun.String
	AccountStatus      goaliyun.String
	AccountDescription goaliyun.String
}

type DescribeAccountsAccountList []DescribeAccountsAccount

func (list *DescribeAccountsAccountList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccountsAccount)
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
