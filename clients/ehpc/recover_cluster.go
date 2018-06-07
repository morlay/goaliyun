package ehpc

import (
	"github.com/morlay/goaliyun"
)

type RecoverClusterRequest struct {
	ClusterId string `position:"Query" name:"ClusterId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *RecoverClusterRequest) Invoke(client goaliyun.Client) (*RecoverClusterResponse, error) {
	resp := &RecoverClusterResponse{}
	err := client.Request("ehpc", "RecoverCluster", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RecoverClusterResponse struct {
	RequestId goaliyun.String
}
