package emr

import (
	"github.com/morlay/goaliyun"
)

type KillExecutionJobInstanceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	JobInstanceId   string `position:"Query" name:"JobInstanceId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *KillExecutionJobInstanceRequest) Invoke(client goaliyun.Client) (*KillExecutionJobInstanceResponse, error) {
	resp := &KillExecutionJobInstanceResponse{}
	err := client.Request("emr", "KillExecutionJobInstance", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type KillExecutionJobInstanceResponse struct {
	RequestId goaliyun.String
}
