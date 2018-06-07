package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBackupsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             string `position:"Query" name:"BackupId"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	BackupStatus         string `position:"Query" name:"BackupStatus"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	BackupMode           string `position:"Query" name:"BackupMode"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBackupsRequest) Invoke(client goaliyun.Client) (*DescribeBackupsResponse, error) {
	resp := &DescribeBackupsResponse{}
	err := client.Request("polardb", "DescribeBackups", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBackupsResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.String
	PageNumber       goaliyun.String
	PageRecordCount  goaliyun.String
	TotalBackupSize  goaliyun.Integer
	Items            DescribeBackupsBackupList
}

type DescribeBackupsBackup struct {
	BackupId                  goaliyun.String
	DBInstanceId              goaliyun.String
	BackupStatus              goaliyun.String
	BackupStartTime           goaliyun.String
	BackupEndTime             goaliyun.String
	BackupType                goaliyun.String
	BackupMode                goaliyun.String
	BackupMethod              goaliyun.String
	BackupDownloadURL         goaliyun.String
	BackupIntranetDownloadURL goaliyun.String
	BackupLocation            goaliyun.String
	BackupExtractionStatus    goaliyun.String
	BackupScale               goaliyun.String
	BackupDBNames             goaliyun.String
	TotalBackupSize           goaliyun.Integer
	BackupSize                goaliyun.Integer
	HostInstanceID            goaliyun.String
	StoreStatus               goaliyun.String
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
