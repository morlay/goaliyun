package imm

import (
	"github.com/morlay/goaliyun"
)

type DeleteTagByUrlRequest struct {
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	SrcUri   string `position:"Query" name:"SrcUri"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteTagByUrlRequest) Invoke(client goaliyun.Client) (*DeleteTagByUrlResponse, error) {
	resp := &DeleteTagByUrlResponse{}
	err := client.Request("imm", "DeleteTagByUrl", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteTagByUrlResponse struct {
	RequestId goaliyun.String
}
