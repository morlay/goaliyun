package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeOssDownloadsForSQLServerRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MigrateTaskId        string `position:"Query" name:"MigrateTaskId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeOssDownloadsForSQLServerRequest) Invoke(client goaliyun.Client) (*DescribeOssDownloadsForSQLServerResponse, error) {
	resp := &DescribeOssDownloadsForSQLServerResponse{}
	err := client.Request("rds", "DescribeOssDownloadsForSQLServer", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeOssDownloadsForSQLServerResponse struct {
	RequestId      goaliyun.String
	DBInstanceName goaliyun.String
	MigrateIaskId  goaliyun.String
	Items          DescribeOssDownloadsForSQLServerOssDownloadList
}

type DescribeOssDownloadsForSQLServerOssDownload struct {
	FileName    goaliyun.String
	CreateTime  goaliyun.String
	CreateTime1 goaliyun.String
	BakType     goaliyun.String
	FileSize    goaliyun.String
	Status      goaliyun.String
	IsAvail     goaliyun.String
	Desc        goaliyun.String
}

type DescribeOssDownloadsForSQLServerOssDownloadList []DescribeOssDownloadsForSQLServerOssDownload

func (list *DescribeOssDownloadsForSQLServerOssDownloadList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOssDownloadsForSQLServerOssDownload)
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
