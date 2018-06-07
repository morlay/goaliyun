package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAccountsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAccountsRequest) Invoke(client goaliyun.Client) (*DescribeAccountsResponse, error) {
	resp := &DescribeAccountsResponse{}
	err := client.Request("polardb", "DescribeAccounts", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAccountsResponse struct {
	RequestId goaliyun.String
	Accounts  DescribeAccountsDBInstanceAccountList
}

type DescribeAccountsDBInstanceAccount struct {
	DBClusterId        goaliyun.String
	AccountName        goaliyun.String
	AccountStatus      goaliyun.String
	AccountDescription goaliyun.String
	AccountType        goaliyun.String
}

type DescribeAccountsDBInstanceAccountList []DescribeAccountsDBInstanceAccount

func (list *DescribeAccountsDBInstanceAccountList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccountsDBInstanceAccount)
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
