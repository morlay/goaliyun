package imm

import (
	"github.com/morlay/goaliyun"
)

type RemoveTagByUrlRequest struct {
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	SrcUri   string `position:"Query" name:"SrcUri"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *RemoveTagByUrlRequest) Invoke(client goaliyun.Client) (*RemoveTagByUrlResponse, error) {
	resp := &RemoveTagByUrlResponse{}
	err := client.Request("imm", "RemoveTagByUrl", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveTagByUrlResponse struct {
	RequestId goaliyun.String
}
