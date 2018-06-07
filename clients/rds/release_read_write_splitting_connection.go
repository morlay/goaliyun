package rds

import (
	"github.com/morlay/goaliyun"
)

type ReleaseReadWriteSplittingConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReleaseReadWriteSplittingConnectionRequest) Invoke(client goaliyun.Client) (*ReleaseReadWriteSplittingConnectionResponse, error) {
	resp := &ReleaseReadWriteSplittingConnectionResponse{}
	err := client.Request("rds", "ReleaseReadWriteSplittingConnection", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseReadWriteSplittingConnectionResponse struct {
	RequestId goaliyun.String
}
