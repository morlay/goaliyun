package imm

import (
	"github.com/morlay/goaliyun"
)

type GetTagJobRequest struct {
	JobId    string `position:"Query" name:"JobId"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetTagJobRequest) Invoke(client goaliyun.Client) (*GetTagJobResponse, error) {
	resp := &GetTagJobResponse{}
	err := client.Request("imm", "GetTagJob", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetTagJobResponse struct {
	RequestId  goaliyun.String
	JobId      goaliyun.String
	SetId      goaliyun.String
	SrcUri     goaliyun.String
	Status     goaliyun.String
	Percent    goaliyun.Integer
	CreateTime goaliyun.String
	FinishTime goaliyun.String
}
