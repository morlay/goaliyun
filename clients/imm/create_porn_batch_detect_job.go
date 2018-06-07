package imm

import (
	"github.com/morlay/goaliyun"
)

type CreatePornBatchDetectJobRequest struct {
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	Project         string `position:"Query" name:"Project"`
	ExternalID      string `position:"Query" name:"ExternalID"`
	SrcUri          string `position:"Query" name:"SrcUri"`
	TgtUri          string `position:"Query" name:"TgtUri"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreatePornBatchDetectJobRequest) Invoke(client goaliyun.Client) (*CreatePornBatchDetectJobResponse, error) {
	resp := &CreatePornBatchDetectJobResponse{}
	err := client.Request("imm", "CreatePornBatchDetectJob", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreatePornBatchDetectJobResponse struct {
	RequestId  goaliyun.String
	JobId      goaliyun.String
	TgtLoc     goaliyun.String
	Status     goaliyun.String
	CreateTime goaliyun.String
	Percent    goaliyun.Integer
}
