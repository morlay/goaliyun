package hpc

import (
	"github.com/morlay/goaliyun"
)

type RebootInstanceRequest struct {
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *RebootInstanceRequest) Invoke(client goaliyun.Client) (*RebootInstanceResponse, error) {
	resp := &RebootInstanceResponse{}
	err := client.Request("hpc", "RebootInstance", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RebootInstanceResponse struct {
	RequestId goaliyun.String
}
