package vod

import (
	"github.com/morlay/goaliyun"
)

type GetMessageCallbackRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetMessageCallbackRequest) Invoke(client goaliyun.Client) (*GetMessageCallbackResponse, error) {
	resp := &GetMessageCallbackResponse{}
	err := client.Request("vod", "GetMessageCallback", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetMessageCallbackResponse struct {
	RequestId       goaliyun.String
	MessageCallback GetMessageCallbackMessageCallback
}

type GetMessageCallbackMessageCallback struct {
	CallbackSwitch goaliyun.String
	CallbackURL    goaliyun.String
	EventTypeList  goaliyun.String
}
