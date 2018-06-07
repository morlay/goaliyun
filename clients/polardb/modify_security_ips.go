package polardb

import (
	"github.com/morlay/goaliyun"
)

type ModifySecurityIpsRequest struct {
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId               string `position:"Query" name:"DBClusterId"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	SecurityIps               string `position:"Query" name:"SecurityIps"`
	DBClusterIPArrayName      string `position:"Query" name:"DBClusterIPArrayName"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	DBClusterIPArrayAttribute string `position:"Query" name:"DBClusterIPArrayAttribute"`
	RegionId                  string `position:"Query" name:"RegionId"`
}

func (req *ModifySecurityIpsRequest) Invoke(client goaliyun.Client) (*ModifySecurityIpsResponse, error) {
	resp := &ModifySecurityIpsResponse{}
	err := client.Request("polardb", "ModifySecurityIps", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySecurityIpsResponse struct {
	RequestId goaliyun.String
}
