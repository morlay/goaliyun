package imm

import (
	"github.com/morlay/goaliyun"
)

type CreateFaceJobRequest struct {
	Project  string `position:"Query" name:"Project"`
	SrcUri   string `position:"Query" name:"SrcUri"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CreateFaceJobRequest) Invoke(client goaliyun.Client) (*CreateFaceJobResponse, error) {
	resp := &CreateFaceJobResponse{}
	err := client.Request("imm", "CreateFaceJob", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateFaceJobResponse struct {
	RequestId  goaliyun.String
	JobId      goaliyun.String
	SetId      goaliyun.String
	SrcUri     goaliyun.String
	Percent    goaliyun.Integer
	CreateTime goaliyun.String
	FinishTime goaliyun.String
	Status     goaliyun.String
}
