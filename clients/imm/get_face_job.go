package imm

import (
	"github.com/morlay/goaliyun"
)

type GetFaceJobRequest struct {
	JobId    string `position:"Query" name:"JobId"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetFaceJobRequest) Invoke(client goaliyun.Client) (*GetFaceJobResponse, error) {
	resp := &GetFaceJobResponse{}
	err := client.Request("imm", "GetFaceJob", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetFaceJobResponse struct {
	RequestId  goaliyun.String
	JobId      goaliyun.String
	SetId      goaliyun.String
	SrcUri     goaliyun.String
	Status     goaliyun.String
	Percent    goaliyun.Integer
	CreateTime goaliyun.String
	FinishTime goaliyun.String
}
