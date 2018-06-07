package emr

import (
	"github.com/morlay/goaliyun"
)

type RefreshClusterResourcePoolRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ResourcePoolId  int64  `position:"Query" name:"ResourcePoolId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *RefreshClusterResourcePoolRequest) Invoke(client goaliyun.Client) (*RefreshClusterResourcePoolResponse, error) {
	resp := &RefreshClusterResourcePoolResponse{}
	err := client.Request("emr", "RefreshClusterResourcePool", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RefreshClusterResourcePoolResponse struct {
	RequestId          goaliyun.String
	WorkFlowInstanceId goaliyun.String
	OperationId        goaliyun.String
}
