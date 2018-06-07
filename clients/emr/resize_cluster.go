package emr

import (
	"github.com/morlay/goaliyun"
)

type ResizeClusterRequest struct {
	NewMasterInstances int64  `position:"Query" name:"NewMasterInstances"`
	Period             int64  `position:"Query" name:"Period"`
	AutoRenew          string `position:"Query" name:"AutoRenew"`
	CoreInstanceType   string `position:"Query" name:"CoreInstanceType"`
	NewCoreInstances   int64  `position:"Query" name:"NewCoreInstances"`
	NewTaskInstances   int64  `position:"Query" name:"NewTaskInstances"`
	ClusterId          string `position:"Query" name:"ClusterId"`
	ChargeType         string `position:"Query" name:"ChargeType"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *ResizeClusterRequest) Invoke(client goaliyun.Client) (*ResizeClusterResponse, error) {
	resp := &ResizeClusterResponse{}
	err := client.Request("emr", "ResizeCluster", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResizeClusterResponse struct {
	RequestId   goaliyun.String
	ClusterId   goaliyun.String
	EmrOrderId  goaliyun.String
	CoreOrderId goaliyun.String
}
