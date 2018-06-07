package cms

import (
	"github.com/morlay/goaliyun"
)

type UpdateMyGroupInstancesRequest struct {
	Instances string `position:"Query" name:"Instances"`
	GroupId   int64  `position:"Query" name:"GroupId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *UpdateMyGroupInstancesRequest) Invoke(client goaliyun.Client) (*UpdateMyGroupInstancesResponse, error) {
	resp := &UpdateMyGroupInstancesResponse{}
	err := client.Request("cms", "UpdateMyGroupInstances", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateMyGroupInstancesResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
}
