package imm

import (
	"github.com/morlay/goaliyun"
)

type CreateFaceSetRequest struct {
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CreateFaceSetRequest) Invoke(client goaliyun.Client) (*CreateFaceSetResponse, error) {
	resp := &CreateFaceSetResponse{}
	err := client.Request("imm", "CreateFaceSet", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateFaceSetResponse struct {
	RequestId  goaliyun.String
	SetId      goaliyun.String
	Status     goaliyun.String
	Photos     goaliyun.Integer
	CreateTime goaliyun.String
	ModifyTime goaliyun.String
}
