package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyJobRequest struct {
	RunParameter    string `position:"Query" name:"RunParameter"`
	RetryInterval   int64  `position:"Query" name:"RetryInterval"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Id              string `position:"Query" name:"Id"`
	Type            string `position:"Query" name:"Type"`
	MaxRetry        int64  `position:"Query" name:"MaxRetry"`
	FailAct         string `position:"Query" name:"FailAct"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyJobRequest) Invoke(client goaliyun.Client) (*ModifyJobResponse, error) {
	resp := &ModifyJobResponse{}
	err := client.Request("emr", "ModifyJob", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyJobResponse struct {
	RequestId goaliyun.String
}
