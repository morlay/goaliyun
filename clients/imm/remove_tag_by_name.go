package imm

import (
	"github.com/morlay/goaliyun"
)

type RemoveTagByNameRequest struct {
	TagName  string `position:"Query" name:"TagName"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	SrcUri   string `position:"Query" name:"SrcUri"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *RemoveTagByNameRequest) Invoke(client goaliyun.Client) (*RemoveTagByNameResponse, error) {
	resp := &RemoveTagByNameResponse{}
	err := client.Request("imm", "RemoveTagByName", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveTagByNameResponse struct {
	RequestId goaliyun.String
}
