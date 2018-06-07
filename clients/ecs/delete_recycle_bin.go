package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteRecycleBinRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceIds          string `position:"Query" name:"ResourceIds"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteRecycleBinRequest) Invoke(client goaliyun.Client) (*DeleteRecycleBinResponse, error) {
	resp := &DeleteRecycleBinResponse{}
	err := client.Request("ecs", "DeleteRecycleBin", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteRecycleBinResponse struct {
	RequestId goaliyun.String
}
