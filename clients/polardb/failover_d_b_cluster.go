package polardb

import (
	"github.com/morlay/goaliyun"
)

type FailoverDBClusterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	TargetDBInstanceId   string `position:"Query" name:"TargetDBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *FailoverDBClusterRequest) Invoke(client goaliyun.Client) (*FailoverDBClusterResponse, error) {
	resp := &FailoverDBClusterResponse{}
	err := client.Request("polardb", "FailoverDBCluster", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FailoverDBClusterResponse struct {
	RequestId goaliyun.String
}
