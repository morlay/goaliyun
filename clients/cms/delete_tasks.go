package cms

import (
	"github.com/morlay/goaliyun"
)

type DeleteTasksRequest struct {
	IsDeleteAlarms int64  `position:"Query" name:"IsDeleteAlarms"`
	TaskIds        string `position:"Query" name:"TaskIds"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DeleteTasksRequest) Invoke(client goaliyun.Client) (*DeleteTasksResponse, error) {
	resp := &DeleteTasksResponse{}
	err := client.Request("cms", "DeleteTasks", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteTasksResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Success   goaliyun.String
	RequestId goaliyun.String
	Data      goaliyun.String
}
