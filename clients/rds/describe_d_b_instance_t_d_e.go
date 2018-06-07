package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceTDERequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceTDERequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceTDEResponse, error) {
	resp := &DescribeDBInstanceTDEResponse{}
	err := client.Request("rds", "DescribeDBInstanceTDE", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceTDEResponse struct {
	RequestId goaliyun.String
	TDEStatus goaliyun.String
	Databases DescribeDBInstanceTDEDatabaseList
}

type DescribeDBInstanceTDEDatabase struct {
	DBName    goaliyun.String
	TDEStatus goaliyun.String
}

type DescribeDBInstanceTDEDatabaseList []DescribeDBInstanceTDEDatabase

func (list *DescribeDBInstanceTDEDatabaseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceTDEDatabase)
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
