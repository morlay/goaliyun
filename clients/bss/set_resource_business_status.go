package bss

import (
	"github.com/morlay/goaliyun"
)

type SetResourceBusinessStatusRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	BusinessStatus       string `position:"Query" name:"BusinessStatus"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetResourceBusinessStatusRequest) Invoke(client goaliyun.Client) (*SetResourceBusinessStatusResponse, error) {
	resp := &SetResourceBusinessStatusResponse{}
	err := client.Request("bss", "SetResourceBusinessStatus", "2014-07-14", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetResourceBusinessStatusResponse struct {
	RequestId goaliyun.String
}
