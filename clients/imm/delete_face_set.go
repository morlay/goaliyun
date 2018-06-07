package imm

import (
	"github.com/morlay/goaliyun"
)

type DeleteFaceSetRequest struct {
	LazyMode   string `position:"Query" name:"LazyMode"`
	Project    string `position:"Query" name:"Project"`
	SetId      string `position:"Query" name:"SetId"`
	CheckEmpty string `position:"Query" name:"CheckEmpty"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeleteFaceSetRequest) Invoke(client goaliyun.Client) (*DeleteFaceSetResponse, error) {
	resp := &DeleteFaceSetResponse{}
	err := client.Request("imm", "DeleteFaceSet", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFaceSetResponse struct {
	RequestId goaliyun.String
}
