package imm

import (
	"github.com/morlay/goaliyun"
)

type CreateTagSetRequest struct {
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CreateTagSetRequest) Invoke(client goaliyun.Client) (*CreateTagSetResponse, error) {
	resp := &CreateTagSetResponse{}
	err := client.Request("imm", "CreateTagSet", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateTagSetResponse struct {
	RequestId  goaliyun.String
	SetId      goaliyun.String
	Status     goaliyun.String
	Photos     goaliyun.Integer
	CreateTime goaliyun.String
	ModifyTime goaliyun.String
}
