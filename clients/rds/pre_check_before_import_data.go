package rds

import (
	"github.com/morlay/goaliyun"
)

type PreCheckBeforeImportDataRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ImportDataType         string `position:"Query" name:"ImportDataType"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken            string `position:"Query" name:"ClientToken"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	SourceDatabaseDBNames  string `position:"Query" name:"SourceDatabaseDBNames"`
	SourceDatabaseIp       string `position:"Query" name:"SourceDatabaseIp"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	SourceDatabasePassword string `position:"Query" name:"SourceDatabasePassword"`
	SourceDatabasePort     string `position:"Query" name:"SourceDatabasePort"`
	SourceDatabaseUserName string `position:"Query" name:"SourceDatabaseUserName"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *PreCheckBeforeImportDataRequest) Invoke(client goaliyun.Client) (*PreCheckBeforeImportDataResponse, error) {
	resp := &PreCheckBeforeImportDataResponse{}
	err := client.Request("rds", "PreCheckBeforeImportData", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PreCheckBeforeImportDataResponse struct {
	RequestId  goaliyun.String
	PreCheckId goaliyun.String
}
