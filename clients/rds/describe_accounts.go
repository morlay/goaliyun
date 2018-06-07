package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAccountsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAccountsRequest) Invoke(client goaliyun.Client) (*DescribeAccountsResponse, error) {
	resp := &DescribeAccountsResponse{}
	err := client.Request("rds", "DescribeAccounts", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAccountsResponse struct {
	RequestId goaliyun.String
	Accounts  DescribeAccountsDBInstanceAccountList
}

type DescribeAccountsDBInstanceAccount struct {
	DBInstanceId       goaliyun.String
	AccountName        goaliyun.String
	AccountStatus      goaliyun.String
	AccountType        goaliyun.String
	AccountDescription goaliyun.String
	PrivExceeded       goaliyun.String
	DatabasePrivileges DescribeAccountsDatabasePrivilegeList
}

type DescribeAccountsDatabasePrivilege struct {
	DBName                 goaliyun.String
	AccountPrivilege       goaliyun.String
	AccountPrivilegeDetail goaliyun.String
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

type DescribeAccountsDatabasePrivilegeList []DescribeAccountsDatabasePrivilege

func (list *DescribeAccountsDatabasePrivilegeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccountsDatabasePrivilege)
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
