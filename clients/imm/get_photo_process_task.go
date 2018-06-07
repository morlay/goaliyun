package imm

import (
	"github.com/morlay/goaliyun"
)

type GetPhotoProcessTaskRequest struct {
	Project  string `position:"Query" name:"Project"`
	TaskId   string `position:"Query" name:"TaskId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetPhotoProcessTaskRequest) Invoke(client goaliyun.Client) (*GetPhotoProcessTaskResponse, error) {
	resp := &GetPhotoProcessTaskResponse{}
	err := client.Request("imm", "GetPhotoProcessTask", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPhotoProcessTaskResponse struct {
	RequestId       goaliyun.String
	TaskId          goaliyun.String
	Status          goaliyun.String
	SrcUri          goaliyun.String
	TgtUri          goaliyun.String
	Style           goaliyun.String
	NotifyTopicName goaliyun.String
	NotifyEndpoint  goaliyun.String
	ExternalID      goaliyun.String
	CreateTime      goaliyun.String
	FinishTime      goaliyun.String
	Percent         goaliyun.Integer
}
