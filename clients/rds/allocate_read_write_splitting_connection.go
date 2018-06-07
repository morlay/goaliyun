package rds

import (
	"github.com/morlay/goaliyun"
)

type AllocateReadWriteSplittingConnectionRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	Weight                 string `position:"Query" name:"Weight"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	IPType                 string `position:"Query" name:"IPType"`
	Port                   string `position:"Query" name:"Port"`
	DistributionType       string `position:"Query" name:"DistributionType"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	MaxDelayTime           string `position:"Query" name:"MaxDelayTime"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *AllocateReadWriteSplittingConnectionRequest) Invoke(client goaliyun.Client) (*AllocateReadWriteSplittingConnectionResponse, error) {
	resp := &AllocateReadWriteSplittingConnectionResponse{}
	err := client.Request("rds", "AllocateReadWriteSplittingConnection", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AllocateReadWriteSplittingConnectionResponse struct {
	RequestId goaliyun.String
}
