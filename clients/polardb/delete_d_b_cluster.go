package polardb

import (
	"github.com/morlay/goaliyun"
)

type DeleteDBClusterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteDBClusterRequest) Invoke(client goaliyun.Client) (*DeleteDBClusterResponse, error) {
	resp := &DeleteDBClusterResponse{}
	err := client.Request("polardb", "DeleteDBCluster", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDBClusterResponse struct {
	RequestId goaliyun.String
}
