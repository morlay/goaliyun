package ehpc

import (
	"github.com/morlay/goaliyun"
)

type StopClusterRequest struct {
	ClusterId string `position:"Query" name:"ClusterId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *StopClusterRequest) Invoke(client goaliyun.Client) (*StopClusterResponse, error) {
	resp := &StopClusterResponse{}
	err := client.Request("ehpc", "StopCluster", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopClusterResponse struct {
	RequestId goaliyun.String
}
