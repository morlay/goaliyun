package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBinlogFilesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBinlogFilesRequest) Invoke(client goaliyun.Client) (*DescribeBinlogFilesResponse, error) {
	resp := &DescribeBinlogFilesResponse{}
	err := client.Request("rds", "DescribeBinlogFiles", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBinlogFilesResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	TotalFileSize    goaliyun.Integer
	Items            DescribeBinlogFilesBinLogFileList
}

type DescribeBinlogFilesBinLogFile struct {
	FileSize             goaliyun.Integer
	LogBeginTime         goaliyun.String
	LogEndTime           goaliyun.String
	DownloadLink         goaliyun.String
	IntranetDownloadLink goaliyun.String
	LinkExpiredTime      goaliyun.String
	Checksum             goaliyun.String
	HostInstanceID       goaliyun.String
}

type DescribeBinlogFilesBinLogFileList []DescribeBinlogFilesBinLogFile

func (list *DescribeBinlogFilesBinLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBinlogFilesBinLogFile)
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
