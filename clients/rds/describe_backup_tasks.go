package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBackupTasksRequest struct {
	BackupJobId          string `position:"Query" name:"BackupJobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Flag                 string `position:"Query" name:"Flag"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BackupMode           string `position:"Query" name:"BackupMode"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackupJobStatus      string `position:"Query" name:"BackupJobStatus"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBackupTasksRequest) Invoke(client goaliyun.Client) (*DescribeBackupTasksResponse, error) {
	resp := &DescribeBackupTasksResponse{}
	err := client.Request("rds", "DescribeBackupTasks", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBackupTasksResponse struct {
	RequestId goaliyun.String
	Items     DescribeBackupTasksBackupJobList
}

type DescribeBackupTasksBackupJob struct {
	BackupProgressStatus goaliyun.String
	JobMode              goaliyun.String
	Process              goaliyun.String
	TaskAction           goaliyun.String
	BackupjobId          goaliyun.String
}

type DescribeBackupTasksBackupJobList []DescribeBackupTasksBackupJob

func (list *DescribeBackupTasksBackupJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupTasksBackupJob)
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
