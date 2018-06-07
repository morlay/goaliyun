package imm

import (
	"github.com/morlay/goaliyun"
)

type DeleteFaceByUrlRequest struct {
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	SrcUri   string `position:"Query" name:"SrcUri"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteFaceByUrlRequest) Invoke(client goaliyun.Client) (*DeleteFaceByUrlResponse, error) {
	resp := &DeleteFaceByUrlResponse{}
	err := client.Request("imm", "DeleteFaceByUrl", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFaceByUrlResponse struct {
	RequestId goaliyun.String
}
