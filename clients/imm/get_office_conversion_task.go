package imm

import (
	"github.com/morlay/goaliyun"
)

type GetOfficeConversionTaskRequest struct {
	Project  string `position:"Query" name:"Project"`
	TaskId   string `position:"Query" name:"TaskId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetOfficeConversionTaskRequest) Invoke(client goaliyun.Client) (*GetOfficeConversionTaskResponse, error) {
	resp := &GetOfficeConversionTaskResponse{}
	err := client.Request("imm", "GetOfficeConversionTask", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetOfficeConversionTaskResponse struct {
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
	FailDetail      GetOfficeConversionTaskFailDetail
}

type GetOfficeConversionTaskFailDetail struct {
	Code goaliyun.String
}
