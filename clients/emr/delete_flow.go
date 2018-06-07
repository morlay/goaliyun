package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteFlowRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteFlowRequest) Invoke(client goaliyun.Client) (*DeleteFlowResponse, error) {
	resp := &DeleteFlowResponse{}
	err := client.Request("emr", "DeleteFlow", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFlowResponse struct {
	RequestId goaliyun.String
	Data      bool
}
