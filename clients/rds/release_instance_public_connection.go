package rds

import (
	"github.com/morlay/goaliyun"
)

type ReleaseInstancePublicConnectionRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *ReleaseInstancePublicConnectionRequest) Invoke(client goaliyun.Client) (*ReleaseInstancePublicConnectionResponse, error) {
	resp := &ReleaseInstancePublicConnectionResponse{}
	err := client.Request("rds", "ReleaseInstancePublicConnection", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseInstancePublicConnectionResponse struct {
	RequestId goaliyun.String
}
