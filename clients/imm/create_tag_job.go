package imm

import (
	"github.com/morlay/goaliyun"
)

type CreateTagJobRequest struct {
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	Project         string `position:"Query" name:"Project"`
	ExternalID      string `position:"Query" name:"ExternalID"`
	SrcUri          string `position:"Query" name:"SrcUri"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateTagJobRequest) Invoke(client goaliyun.Client) (*CreateTagJobResponse, error) {
	resp := &CreateTagJobResponse{}
	err := client.Request("imm", "CreateTagJob", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateTagJobResponse struct {
	RequestId  goaliyun.String
	JobId      goaliyun.String
	SetId      goaliyun.String
	SrcUri     goaliyun.String
	Status     goaliyun.String
	Percent    goaliyun.Integer
	CreateTime goaliyun.String
	FinishTime goaliyun.String
}
