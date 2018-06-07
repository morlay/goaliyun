package cbn

import (
	"github.com/morlay/goaliyun"
)

type DeleteCenRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteCenRequest) Invoke(client goaliyun.Client) (*DeleteCenResponse, error) {
	resp := &DeleteCenResponse{}
	err := client.Request("cbn", "DeleteCen", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCenResponse struct {
	RequestId goaliyun.String
}
