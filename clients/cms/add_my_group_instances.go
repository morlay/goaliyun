package cms

import (
	"github.com/morlay/goaliyun"
)

type AddMyGroupInstancesRequest struct {
	Instances string `position:"Query" name:"Instances"`
	GroupId   int64  `position:"Query" name:"GroupId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *AddMyGroupInstancesRequest) Invoke(client goaliyun.Client) (*AddMyGroupInstancesResponse, error) {
	resp := &AddMyGroupInstancesResponse{}
	err := client.Request("cms", "AddMyGroupInstances", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddMyGroupInstancesResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
}
