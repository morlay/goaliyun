package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeScheduledTasksRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ScheduledAction2     string `position:"Query" name:"ScheduledAction.2"`
	ScheduledAction1     string `position:"Query" name:"ScheduledAction.1"`
	ScheduledAction6     string `position:"Query" name:"ScheduledAction.6"`
	ScheduledAction5     string `position:"Query" name:"ScheduledAction.5"`
	ScheduledAction4     string `position:"Query" name:"ScheduledAction.4"`
	ScheduledAction3     string `position:"Query" name:"ScheduledAction.3"`
	ScheduledAction9     string `position:"Query" name:"ScheduledAction.9"`
	ScheduledAction8     string `position:"Query" name:"ScheduledAction.8"`
	ScheduledAction7     string `position:"Query" name:"ScheduledAction.7"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScheduledTaskName20  string `position:"Query" name:"ScheduledTaskName.20"`
	ScheduledTaskName19  string `position:"Query" name:"ScheduledTaskName.19"`
	ScheduledTaskName18  string `position:"Query" name:"ScheduledTaskName.18"`
	ScheduledTaskId20    string `position:"Query" name:"ScheduledTaskId.20"`
	ScheduledTaskName13  string `position:"Query" name:"ScheduledTaskName.13"`
	ScheduledTaskName12  string `position:"Query" name:"ScheduledTaskName.12"`
	ScheduledTaskName11  string `position:"Query" name:"ScheduledTaskName.11"`
	ScheduledTaskName10  string `position:"Query" name:"ScheduledTaskName.10"`
	ScheduledTaskName17  string `position:"Query" name:"ScheduledTaskName.17"`
	ScheduledTaskName16  string `position:"Query" name:"ScheduledTaskName.16"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	ScheduledTaskName15  string `position:"Query" name:"ScheduledTaskName.15"`
	ScheduledTaskName14  string `position:"Query" name:"ScheduledTaskName.14"`
	ScheduledTaskId2     string `position:"Query" name:"ScheduledTaskId.2"`
	ScheduledTaskId1     string `position:"Query" name:"ScheduledTaskId.1"`
	ScheduledTaskId4     string `position:"Query" name:"ScheduledTaskId.4"`
	ScheduledTaskId18    string `position:"Query" name:"ScheduledTaskId.18"`
	ScheduledTaskId3     string `position:"Query" name:"ScheduledTaskId.3"`
	ScheduledTaskId19    string `position:"Query" name:"ScheduledTaskId.19"`
	ScheduledTaskId6     string `position:"Query" name:"ScheduledTaskId.6"`
	ScheduledTaskId5     string `position:"Query" name:"ScheduledTaskId.5"`
	ScheduledTaskId8     string `position:"Query" name:"ScheduledTaskId.8"`
	ScheduledTaskName9   string `position:"Query" name:"ScheduledTaskName.9"`
	ScheduledAction20    string `position:"Query" name:"ScheduledAction.20"`
	ScheduledTaskId7     string `position:"Query" name:"ScheduledTaskId.7"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	ScheduledTaskId12    string `position:"Query" name:"ScheduledTaskId.12"`
	ScheduledTaskName7   string `position:"Query" name:"ScheduledTaskName.7"`
	ScheduledTaskId9     string `position:"Query" name:"ScheduledTaskId.9"`
	ScheduledTaskId13    string `position:"Query" name:"ScheduledTaskId.13"`
	ScheduledTaskName8   string `position:"Query" name:"ScheduledTaskName.8"`
	ScheduledTaskId10    string `position:"Query" name:"ScheduledTaskId.10"`
	ScheduledTaskName5   string `position:"Query" name:"ScheduledTaskName.5"`
	ScheduledTaskId11    string `position:"Query" name:"ScheduledTaskId.11"`
	ScheduledTaskName6   string `position:"Query" name:"ScheduledTaskName.6"`
	ScheduledTaskId16    string `position:"Query" name:"ScheduledTaskId.16"`
	ScheduledTaskName3   string `position:"Query" name:"ScheduledTaskName.3"`
	ScheduledTaskId17    string `position:"Query" name:"ScheduledTaskId.17"`
	ScheduledTaskName4   string `position:"Query" name:"ScheduledTaskName.4"`
	ScheduledTaskId14    string `position:"Query" name:"ScheduledTaskId.14"`
	ScheduledTaskName1   string `position:"Query" name:"ScheduledTaskName.1"`
	ScheduledTaskId15    string `position:"Query" name:"ScheduledTaskId.15"`
	ScheduledTaskName2   string `position:"Query" name:"ScheduledTaskName.2"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ScheduledAction18    string `position:"Query" name:"ScheduledAction.18"`
	ScheduledAction19    string `position:"Query" name:"ScheduledAction.19"`
	ScheduledAction16    string `position:"Query" name:"ScheduledAction.16"`
	ScheduledAction17    string `position:"Query" name:"ScheduledAction.17"`
	ScheduledAction14    string `position:"Query" name:"ScheduledAction.14"`
	ScheduledAction15    string `position:"Query" name:"ScheduledAction.15"`
	ScheduledAction12    string `position:"Query" name:"ScheduledAction.12"`
	ScheduledAction13    string `position:"Query" name:"ScheduledAction.13"`
	ScheduledAction10    string `position:"Query" name:"ScheduledAction.10"`
	ScheduledAction11    string `position:"Query" name:"ScheduledAction.11"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeScheduledTasksRequest) Invoke(client goaliyun.Client) (*DescribeScheduledTasksResponse, error) {
	resp := &DescribeScheduledTasksResponse{}
	err := client.Request("ess", "DescribeScheduledTasks", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeScheduledTasksResponse struct {
	TotalCount     goaliyun.Integer
	PageNumber     goaliyun.Integer
	PageSize       goaliyun.Integer
	RequestId      goaliyun.String
	ScheduledTasks DescribeScheduledTasksScheduledTaskList
}

type DescribeScheduledTasksScheduledTask struct {
	ScheduledTaskId      goaliyun.String
	ScheduledTaskName    goaliyun.String
	Description          goaliyun.String
	ScheduledAction      goaliyun.String
	RecurrenceEndTime    goaliyun.String
	LaunchTime           goaliyun.String
	RecurrenceType       goaliyun.String
	RecurrenceValue      goaliyun.String
	LaunchExpirationTime goaliyun.Integer
	TaskEnabled          bool
}

type DescribeScheduledTasksScheduledTaskList []DescribeScheduledTasksScheduledTask

func (list *DescribeScheduledTasksScheduledTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScheduledTasksScheduledTask)
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
