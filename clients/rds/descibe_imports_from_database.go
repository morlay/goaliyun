package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescibeImportsFromDatabaseRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	ImportId             int64  `position:"Query" name:"ImportId"`
	Engine               string `position:"Query" name:"Engine"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescibeImportsFromDatabaseRequest) Invoke(client goaliyun.Client) (*DescibeImportsFromDatabaseResponse, error) {
	resp := &DescibeImportsFromDatabaseResponse{}
	err := client.Request("rds", "DescibeImportsFromDatabase", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescibeImportsFromDatabaseResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescibeImportsFromDatabaseImportResultFromDBList
}

type DescibeImportsFromDatabaseImportResultFromDB struct {
	ImportId                    goaliyun.Integer
	ImportDataType              goaliyun.String
	ImportDataStatus            goaliyun.String
	ImportDataStatusDescription goaliyun.String
	IncrementalImportingTime    goaliyun.String
}

type DescibeImportsFromDatabaseImportResultFromDBList []DescibeImportsFromDatabaseImportResultFromDB

func (list *DescibeImportsFromDatabaseImportResultFromDBList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescibeImportsFromDatabaseImportResultFromDB)
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
