package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBackupsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             int64  `position:"Query" name:"BackupId"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBackupsRequest) Invoke(client goaliyun.Client) (*DescribeBackupsResponse, error) {
	resp := &DescribeBackupsResponse{}
	err := client.Request("r-kvstore", "DescribeBackups", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBackupsResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	Backups    DescribeBackupsBackupList
}

type DescribeBackupsBackup struct {
	BackupId                  goaliyun.Integer
	BackupDBNames             goaliyun.String
	BackupStatus              goaliyun.String
	BackupStartTime           goaliyun.String
	BackupEndTime             goaliyun.String
	BackupType                goaliyun.String
	BackupMode                goaliyun.String
	BackupMethod              goaliyun.String
	BackupDownloadURL         goaliyun.String
	BackupSize                goaliyun.Integer
	EngineVersion             goaliyun.String
	NodeInstanceId            goaliyun.String
	BackupIntranetDownloadURL goaliyun.String
}

type DescribeBackupsBackupList []DescribeBackupsBackup

func (list *DescribeBackupsBackupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupsBackup)
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
