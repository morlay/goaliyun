package emr

import (
	"github.com/morlay/goaliyun"
)

type KillFlowRequest struct {
	FlowInstanceId  string `position:"Query" name:"FlowInstanceId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *KillFlowRequest) Invoke(client goaliyun.Client) (*KillFlowResponse, error) {
	resp := &KillFlowResponse{}
	err := client.Request("emr", "KillFlow", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type KillFlowResponse struct {
	RequestId goaliyun.String
	Data      bool
}
