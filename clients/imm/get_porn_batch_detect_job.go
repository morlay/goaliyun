package imm

import (
	"github.com/morlay/goaliyun"
)

type GetPornBatchDetectJobRequest struct {
	JobId    string `position:"Query" name:"JobId"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetPornBatchDetectJobRequest) Invoke(client goaliyun.Client) (*GetPornBatchDetectJobResponse, error) {
	resp := &GetPornBatchDetectJobResponse{}
	err := client.Request("imm", "GetPornBatchDetectJob", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPornBatchDetectJobResponse struct {
	RequestId       goaliyun.String
	JobId           goaliyun.String
	SrcUri          goaliyun.String
	TgtUri          goaliyun.String
	NotifyTopicName goaliyun.String
	NotifyEndpoint  goaliyun.String
	ExternalID      goaliyun.String
	Status          goaliyun.String
	CreateTime      goaliyun.String
	FinishTime      goaliyun.String
	Percent         goaliyun.Integer
}
