package cms

import (
	"github.com/morlay/goaliyun"
)

type DeleteMyGroupInstancesRequest struct {
	InstanceIds string `position:"Query" name:"InstanceIds"`
	GroupId     int64  `position:"Query" name:"GroupId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DeleteMyGroupInstancesRequest) Invoke(client goaliyun.Client) (*DeleteMyGroupInstancesResponse, error) {
	resp := &DeleteMyGroupInstancesResponse{}
	err := client.Request("cms", "DeleteMyGroupInstances", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteMyGroupInstancesResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
}
