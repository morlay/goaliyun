package dyvmsapi

import (
	"github.com/morlay/goaliyun"
)

type CancelCallRequest struct {
	CallId               string `position:"Query" name:"CallId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CancelCallRequest) Invoke(client goaliyun.Client) (*CancelCallResponse, error) {
	resp := &CancelCallResponse{}
	err := client.Request("dyvmsapi", "CancelCall", "2017-05-25", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelCallResponse struct {
	RequestId goaliyun.String
	Status    bool
	Code      goaliyun.String
	Message   goaliyun.String
}
