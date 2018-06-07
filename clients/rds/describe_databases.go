package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDatabasesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBStatus             string `position:"Query" name:"DBStatus"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDatabasesRequest) Invoke(client goaliyun.Client) (*DescribeDatabasesResponse, error) {
	resp := &DescribeDatabasesResponse{}
	err := client.Request("rds", "DescribeDatabases", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDatabasesResponse struct {
	RequestId goaliyun.String
	Databases DescribeDatabasesDatabaseList
}

type DescribeDatabasesDatabase struct {
	DBName           goaliyun.String
	DBInstanceId     goaliyun.String
	Engine           goaliyun.String
	DBStatus         goaliyun.String
	CharacterSetName goaliyun.String
	DBDescription    goaliyun.String
	Accounts         DescribeDatabasesAccountPrivilegeInfoList
}

type DescribeDatabasesAccountPrivilegeInfo struct {
	Account                goaliyun.String
	AccountPrivilege       goaliyun.String
	AccountPrivilegeDetail goaliyun.String
}

type DescribeDatabasesDatabaseList []DescribeDatabasesDatabase

func (list *DescribeDatabasesDatabaseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDatabasesDatabase)
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

type DescribeDatabasesAccountPrivilegeInfoList []DescribeDatabasesAccountPrivilegeInfo

func (list *DescribeDatabasesAccountPrivilegeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDatabasesAccountPrivilegeInfo)
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
