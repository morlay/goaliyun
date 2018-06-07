package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMigrateTasksForSQLServerRequest struct {
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

func (req *DescribeMigrateTasksForSQLServerRequest) Invoke(client goaliyun.Client) (*DescribeMigrateTasksForSQLServerResponse, error) {
	resp := &DescribeMigrateTasksForSQLServerResponse{}
	err := client.Request("rds", "DescribeMigrateTasksForSQLServer", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMigrateTasksForSQLServerResponse struct {
	RequestId        goaliyun.String
	DBInstanceID     goaliyun.String
	DBInstanceName   goaliyun.String
	StartTime        goaliyun.String
	EndTime          goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeMigrateTasksForSQLServerMigrateTaskList
}

type DescribeMigrateTasksForSQLServerMigrateTask struct {
	DBName        goaliyun.String
	MigrateIaskId goaliyun.String
	CreateTime    goaliyun.String
	EndTime       goaliyun.String
	TaskType      goaliyun.String
	Status        goaliyun.String
	IsDBReplaced  goaliyun.String
	Desc          goaliyun.String
}

type DescribeMigrateTasksForSQLServerMigrateTaskList []DescribeMigrateTasksForSQLServerMigrateTask

func (list *DescribeMigrateTasksForSQLServerMigrateTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMigrateTasksForSQLServerMigrateTask)
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
