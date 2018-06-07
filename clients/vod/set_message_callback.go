package vod

import (
	"github.com/morlay/goaliyun"
)

type SetMessageCallbackRequest struct {
	CallbackType         string `position:"Query" name:"CallbackType"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	CallbackSwitch       string `position:"Query" name:"CallbackSwitch"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EventTypeList        string `position:"Query" name:"EventTypeList"`
	CallbackURL          string `position:"Query" name:"CallbackURL"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetMessageCallbackRequest) Invoke(client goaliyun.Client) (*SetMessageCallbackResponse, error) {
	resp := &SetMessageCallbackResponse{}
	err := client.Request("vod", "SetMessageCallback", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetMessageCallbackResponse struct {
	RequestId goaliyun.String
}
