package imm

import (
	"github.com/morlay/goaliyun"
)

type DeleteTagJobRequest struct {
	JobId          string `position:"Query" name:"JobId"`
	Project        string `position:"Query" name:"Project"`
	ClearIndexData string `position:"Query" name:"ClearIndexData"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DeleteTagJobRequest) Invoke(client goaliyun.Client) (*DeleteTagJobResponse, error) {
	resp := &DeleteTagJobResponse{}
	err := client.Request("imm", "DeleteTagJob", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteTagJobResponse struct {
	RequestId goaliyun.String
}
