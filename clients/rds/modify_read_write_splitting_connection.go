package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyReadWriteSplittingConnectionRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	DistributionType       string `position:"Query" name:"DistributionType"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	Weight                 string `position:"Query" name:"Weight"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	MaxDelayTime           string `position:"Query" name:"MaxDelayTime"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *ModifyReadWriteSplittingConnectionRequest) Invoke(client goaliyun.Client) (*ModifyReadWriteSplittingConnectionResponse, error) {
	resp := &ModifyReadWriteSplittingConnectionResponse{}
	err := client.Request("rds", "ModifyReadWriteSplittingConnection", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyReadWriteSplittingConnectionResponse struct {
	RequestId goaliyun.String
}
