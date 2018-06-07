package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSQLLogFilesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	FileName             string `position:"Query" name:"FileName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSQLLogFilesRequest) Invoke(client goaliyun.Client) (*DescribeSQLLogFilesResponse, error) {
	resp := &DescribeSQLLogFilesResponse{}
	err := client.Request("rds", "DescribeSQLLogFiles", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSQLLogFilesResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeSQLLogFilesLogFileList
}

type DescribeSQLLogFilesLogFile struct {
	FileID         goaliyun.String
	LogStatus      goaliyun.String
	LogDownloadURL goaliyun.String
	LogSize        goaliyun.String
	LogStartTime   goaliyun.String
	LogEndTime     goaliyun.String
}

type DescribeSQLLogFilesLogFileList []DescribeSQLLogFilesLogFile

func (list *DescribeSQLLogFilesLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogFilesLogFile)
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
