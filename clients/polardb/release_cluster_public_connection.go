package polardb

import (
	"github.com/morlay/goaliyun"
)

type ReleaseClusterPublicConnectionRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId             string `position:"Query" name:"DBClusterId"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *ReleaseClusterPublicConnectionRequest) Invoke(client goaliyun.Client) (*ReleaseClusterPublicConnectionResponse, error) {
	resp := &ReleaseClusterPublicConnectionResponse{}
	err := client.Request("polardb", "ReleaseClusterPublicConnection", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseClusterPublicConnectionResponse struct {
	RequestId goaliyun.String
}
