package ehpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteJobTemplatesRequest struct {
	Templates string `position:"Query" name:"Templates"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DeleteJobTemplatesRequest) Invoke(client goaliyun.Client) (*DeleteJobTemplatesResponse, error) {
	resp := &DeleteJobTemplatesResponse{}
	err := client.Request("ehpc", "DeleteJobTemplates", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteJobTemplatesResponse struct {
	RequestId goaliyun.String
}
