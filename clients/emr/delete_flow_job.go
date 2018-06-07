package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteFlowJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteFlowJobRequest) Invoke(client goaliyun.Client) (*DeleteFlowJobResponse, error) {
	resp := &DeleteFlowJobResponse{}
	err := client.Request("emr", "DeleteFlowJob", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFlowJobResponse struct {
	RequestId goaliyun.String
	Data      bool
}
