package imm

import (
	"github.com/morlay/goaliyun"
)

type GetConvertOfficeFormatTaskRequest struct {
	Project  string `position:"Query" name:"Project"`
	TaskId   string `position:"Query" name:"TaskId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetConvertOfficeFormatTaskRequest) Invoke(client goaliyun.Client) (*GetConvertOfficeFormatTaskResponse, error) {
	resp := &GetConvertOfficeFormatTaskResponse{}
	err := client.Request("imm", "GetConvertOfficeFormatTask", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetConvertOfficeFormatTaskResponse struct {
	RequestId       goaliyun.String
	TgtType         goaliyun.String
	Status          goaliyun.String
	Percent         goaliyun.Integer
	PageCount       goaliyun.Integer
	TaskId          goaliyun.String
	TgtUri          goaliyun.String
	ImageSpec       goaliyun.String
	NotifyTopicName goaliyun.String
	NotifyEndpoint  goaliyun.String
	ExternalID      goaliyun.String
	CreateTime      goaliyun.String
	FinishTime      goaliyun.String
	SrcUri          goaliyun.String
}
