package emr

import (
	"github.com/morlay/goaliyun"
)

type ReleaseClusterRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ForceRelease    string `position:"Query" name:"ForceRelease"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ReleaseClusterRequest) Invoke(client goaliyun.Client) (*ReleaseClusterResponse, error) {
	resp := &ReleaseClusterResponse{}
	err := client.Request("emr", "ReleaseCluster", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseClusterResponse struct {
	RequestId goaliyun.String
}
