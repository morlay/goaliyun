package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateJobRequest struct {
	RunParameter    string `position:"Query" name:"RunParameter"`
	RetryInterval   int64  `position:"Query" name:"RetryInterval"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Type            string `position:"Query" name:"Type"`
	MaxRetry        int64  `position:"Query" name:"MaxRetry"`
	FailAct         string `position:"Query" name:"FailAct"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateJobRequest) Invoke(client goaliyun.Client) (*CreateJobResponse, error) {
	resp := &CreateJobResponse{}
	err := client.Request("emr", "CreateJob", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateJobResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
}
