package imm

import (
	"github.com/morlay/goaliyun"
)

type DeleteTagByNameRequest struct {
	TagName  string `position:"Query" name:"TagName"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	SrcUri   string `position:"Query" name:"SrcUri"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteTagByNameRequest) Invoke(client goaliyun.Client) (*DeleteTagByNameResponse, error) {
	resp := &DeleteTagByNameResponse{}
	err := client.Request("imm", "DeleteTagByName", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteTagByNameResponse struct {
	RequestId goaliyun.String
}
