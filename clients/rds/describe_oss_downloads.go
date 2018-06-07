package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeOssDownloadsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MigrateTaskId        string `position:"Query" name:"MigrateTaskId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeOssDownloadsRequest) Invoke(client goaliyun.Client) (*DescribeOssDownloadsResponse, error) {
	resp := &DescribeOssDownloadsResponse{}
	err := client.Request("rds", "DescribeOssDownloads", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeOssDownloadsResponse struct {
	RequestId     goaliyun.String
	DBInstanceId  goaliyun.String
	MigrateTaskId goaliyun.String
	Items         DescribeOssDownloadsOssDownloadList
}

type DescribeOssDownloadsOssDownload struct {
	FileName    goaliyun.String
	CreateTime  goaliyun.String
	EndTime     goaliyun.String
	BackupMode  goaliyun.String
	FileSize    goaliyun.String
	Status      goaliyun.String
	IsAvailable goaliyun.String
	Description goaliyun.String
}

type DescribeOssDownloadsOssDownloadList []DescribeOssDownloadsOssDownload

func (list *DescribeOssDownloadsOssDownloadList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOssDownloadsOssDownload)
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
