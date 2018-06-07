package imm

import (
	"github.com/morlay/goaliyun"
)

type GetFaceSetRequest struct {
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetFaceSetRequest) Invoke(client goaliyun.Client) (*GetFaceSetResponse, error) {
	resp := &GetFaceSetResponse{}
	err := client.Request("imm", "GetFaceSet", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetFaceSetResponse struct {
	RequestId  goaliyun.String
	SetId      goaliyun.String
	Status     goaliyun.String
	Photos     goaliyun.Integer
	CreateTime goaliyun.String
	ModifyTime goaliyun.String
}
