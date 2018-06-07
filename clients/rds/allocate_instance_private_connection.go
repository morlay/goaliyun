package rds

import (
	"github.com/morlay/goaliyun"
)

type AllocateInstancePrivateConnectionRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *AllocateInstancePrivateConnectionRequest) Invoke(client goaliyun.Client) (*AllocateInstancePrivateConnectionResponse, error) {
	resp := &AllocateInstancePrivateConnectionResponse{}
	err := client.Request("rds", "AllocateInstancePrivateConnection", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AllocateInstancePrivateConnectionResponse struct {
	RequestId goaliyun.String
}
