package ehpc

import (
	"github.com/morlay/goaliyun"
)

type StartClusterRequest struct {
	ClusterId string `position:"Query" name:"ClusterId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *StartClusterRequest) Invoke(client goaliyun.Client) (*StartClusterResponse, error) {
	resp := &StartClusterResponse{}
	err := client.Request("ehpc", "StartCluster", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartClusterResponse struct {
	RequestId goaliyun.String
}
