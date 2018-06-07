package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeFlowNodeInstanceLauncherLogRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Start           int64  `position:"Query" name:"Start"`
	Lines           int64  `position:"Query" name:"Lines"`
	NodeInstanceId  string `position:"Query" name:"NodeInstanceId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeFlowNodeInstanceLauncherLogRequest) Invoke(client goaliyun.Client) (*DescribeFlowNodeInstanceLauncherLogResponse, error) {
	resp := &DescribeFlowNodeInstanceLauncherLogResponse{}
	err := client.Request("emr", "DescribeFlowNodeInstanceLauncherLog", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFlowNodeInstanceLauncherLogResponse struct {
	RequestId goaliyun.String
	Logs      DescribeFlowNodeInstanceLauncherLogLogList
}

type DescribeFlowNodeInstanceLauncherLogLog struct {
	LogEntry DescribeFlowNodeInstanceLauncherLogLogEntry
}

type DescribeFlowNodeInstanceLauncherLogLogEntry struct {
	Content goaliyun.String
}

type DescribeFlowNodeInstanceLauncherLogLogList []DescribeFlowNodeInstanceLauncherLogLog

func (list *DescribeFlowNodeInstanceLauncherLogLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFlowNodeInstanceLauncherLogLog)
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
