package imm

import (
	"github.com/morlay/goaliyun"
)

type DeleteFaceJobRequest struct {
	JobId          string `position:"Query" name:"JobId"`
	Project        string `position:"Query" name:"Project"`
	ClearIndexData string `position:"Query" name:"ClearIndexData"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DeleteFaceJobRequest) Invoke(client goaliyun.Client) (*DeleteFaceJobResponse, error) {
	resp := &DeleteFaceJobResponse{}
	err := client.Request("imm", "DeleteFaceJob", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFaceJobResponse struct {
	RequestId goaliyun.String
}
