package imm

import (
	"github.com/morlay/goaliyun"
)

type DeleteTagSetRequest struct {
	LazyMode   string `position:"Query" name:"LazyMode"`
	Project    string `position:"Query" name:"Project"`
	SetId      string `position:"Query" name:"SetId"`
	CheckEmpty string `position:"Query" name:"CheckEmpty"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeleteTagSetRequest) Invoke(client goaliyun.Client) (*DeleteTagSetResponse, error) {
	resp := &DeleteTagSetResponse{}
	err := client.Request("imm", "DeleteTagSet", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteTagSetResponse struct {
	RequestId goaliyun.String
}
