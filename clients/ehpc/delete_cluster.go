package ehpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteClusterRequest struct {
	ReleaseInstance string `position:"Query" name:"ReleaseInstance"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteClusterRequest) Invoke(client goaliyun.Client) (*DeleteClusterResponse, error) {
	resp := &DeleteClusterResponse{}
	err := client.Request("ehpc", "DeleteCluster", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteClusterResponse struct {
	RequestId goaliyun.String
}
