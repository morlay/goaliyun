package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMigrateTasksRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeMigrateTasksRequest) Invoke(client goaliyun.Client) (*DescribeMigrateTasksResponse, error) {
	resp := &DescribeMigrateTasksResponse{}
	err := client.Request("rds", "DescribeMigrateTasks", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMigrateTasksResponse struct {
	RequestId        goaliyun.String
	DBInstanceId     goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeMigrateTasksMigrateTaskList
}

type DescribeMigrateTasksMigrateTask struct {
	DBName        goaliyun.String
	MigrateTaskId goaliyun.String
	CreateTime    goaliyun.String
	EndTime       goaliyun.String
	BackupMode    goaliyun.String
	Status        goaliyun.String
	IsDBReplaced  goaliyun.String
	Description   goaliyun.String
}

type DescribeMigrateTasksMigrateTaskList []DescribeMigrateTasksMigrateTask

func (list *DescribeMigrateTasksMigrateTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMigrateTasksMigrateTask)
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
