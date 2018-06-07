package emr

import (
	"github.com/morlay/goaliyun"
)

type TerminateClusterOperationRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	OperationId     string `position:"Query" name:"OperationId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *TerminateClusterOperationRequest) Invoke(client goaliyun.Client) (*TerminateClusterOperationResponse, error) {
	resp := &TerminateClusterOperationResponse{}
	err := client.Request("emr", "TerminateClusterOperation", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TerminateClusterOperationResponse struct {
	RequestId goaliyun.String
}
