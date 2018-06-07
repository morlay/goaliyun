package ecs

import (
	"github.com/morlay/goaliyun"
)

type GetInstanceConsoleOutputRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetInstanceConsoleOutputRequest) Invoke(client goaliyun.Client) (*GetInstanceConsoleOutputResponse, error) {
	resp := &GetInstanceConsoleOutputResponse{}
	err := client.Request("ecs", "GetInstanceConsoleOutput", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetInstanceConsoleOutputResponse struct {
	RequestId      goaliyun.String
	InstanceId     goaliyun.String
	ConsoleOutput  goaliyun.String
	LastUpdateTime goaliyun.String
}
