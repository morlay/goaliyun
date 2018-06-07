package rds

import (
	"github.com/morlay/goaliyun"
)

type AllocateInstancePublicConnectionRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *AllocateInstancePublicConnectionRequest) Invoke(client goaliyun.Client) (*AllocateInstancePublicConnectionResponse, error) {
	resp := &AllocateInstancePublicConnectionResponse{}
	err := client.Request("rds", "AllocateInstancePublicConnection", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AllocateInstancePublicConnectionResponse struct {
	RequestId goaliyun.String
}
