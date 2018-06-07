package mts

import (
	"github.com/morlay/goaliyun"
)

type CancelJobRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CancelJobRequest) Invoke(client goaliyun.Client) (*CancelJobResponse, error) {
	resp := &CancelJobResponse{}
	err := client.Request("mts", "CancelJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
