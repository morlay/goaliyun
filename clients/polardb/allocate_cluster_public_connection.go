package polardb

import (
	"github.com/morlay/goaliyun"
)

type AllocateClusterPublicConnectionRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	DBClusterId            string `position:"Query" name:"DBClusterId"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *AllocateClusterPublicConnectionRequest) Invoke(client goaliyun.Client) (*AllocateClusterPublicConnectionResponse, error) {
	resp := &AllocateClusterPublicConnectionResponse{}
	err := client.Request("polardb", "AllocateClusterPublicConnection", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AllocateClusterPublicConnectionResponse struct {
	RequestId        goaliyun.String
	DBInstanceId     goaliyun.String
	ConnectionString goaliyun.String
	IPType           goaliyun.String
	Port             goaliyun.String
}
