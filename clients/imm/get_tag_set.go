package imm

import (
	"github.com/morlay/goaliyun"
)

type GetTagSetRequest struct {
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetTagSetRequest) Invoke(client goaliyun.Client) (*GetTagSetResponse, error) {
	resp := &GetTagSetResponse{}
	err := client.Request("imm", "GetTagSet", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetTagSetResponse struct {
	RequestId  goaliyun.String
	SetId      goaliyun.String
	Status     goaliyun.String
	CreateTime goaliyun.String
	ModifyTime goaliyun.String
	Photos     goaliyun.Integer
}
