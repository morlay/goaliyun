package emr

import (
	"github.com/morlay/goaliyun"
)

type RestartClusterRequest struct {
	RollingRestart        string `position:"Query" name:"RollingRestart"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	UpgradedHostGroupOnly string `position:"Query" name:"UpgradedHostGroupOnly"`
	ClusterId             string `position:"Query" name:"ClusterId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *RestartClusterRequest) Invoke(client goaliyun.Client) (*RestartClusterResponse, error) {
	resp := &RestartClusterResponse{}
	err := client.Request("emr", "RestartCluster", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RestartClusterResponse struct {
	RequestId goaliyun.String
	ClusterId goaliyun.String
}
