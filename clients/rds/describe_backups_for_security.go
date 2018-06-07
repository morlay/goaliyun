package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBackupsForSecurityRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	TargetAliBid         string `position:"Query" name:"TargetAliBid"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	BackupStatus         string `position:"Query" name:"BackupStatus"`
	BackupLocation       string `position:"Query" name:"BackupLocation"`
	TargetAliUid         string `position:"Query" name:"TargetAliUid"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BackupMode           string `position:"Query" name:"BackupMode"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBackupsForSecurityRequest) Invoke(client goaliyun.Client) (*DescribeBackupsForSecurityResponse, error) {
	resp := &DescribeBackupsForSecurityResponse{}
	err := client.Request("rds", "DescribeBackupsForSecurity", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBackupsForSecurityResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.String
	PageNumber       goaliyun.String
	PageRecordCount  goaliyun.String
	TotalBackupSize  goaliyun.Integer
	Items            DescribeBackupsForSecurityBackupList
}

type DescribeBackupsForSecurityBackup struct {
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

type DescribeBackupsForSecurityBackupList []DescribeBackupsForSecurityBackup

func (list *DescribeBackupsForSecurityBackupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupsForSecurityBackup)
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
