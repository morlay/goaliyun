package polardb

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBClusterDescriptionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBClusterDescription string `position:"Query" name:"DBClusterDescription"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBClusterDescriptionRequest) Invoke(client goaliyun.Client) (*ModifyDBClusterDescriptionResponse, error) {
	resp := &ModifyDBClusterDescriptionResponse{}
	err := client.Request("polardb", "ModifyDBClusterDescription", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBClusterDescriptionResponse struct {
	RequestId goaliyun.String
}
