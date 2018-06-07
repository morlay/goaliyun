package imm

import (
	"github.com/morlay/goaliyun"
)

type RemoveFaceByUrlRequest struct {
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	SrcUri   string `position:"Query" name:"SrcUri"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *RemoveFaceByUrlRequest) Invoke(client goaliyun.Client) (*RemoveFaceByUrlResponse, error) {
	resp := &RemoveFaceByUrlResponse{}
	err := client.Request("imm", "RemoveFaceByUrl", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveFaceByUrlResponse struct {
	RequestId goaliyun.String
}
