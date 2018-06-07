package mts

import (
	"github.com/morlay/goaliyun"
)

type SetAuthConfigRequest struct {
	Key1                 string `position:"Query" name:"Key.1"`
	Key2                 string `position:"Query" name:"Key.2"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetAuthConfigRequest) Invoke(client goaliyun.Client) (*SetAuthConfigResponse, error) {
	resp := &SetAuthConfigResponse{}
	err := client.Request("mts", "SetAuthConfig", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetAuthConfigResponse struct {
	RequestId goaliyun.String
	Key1      goaliyun.String
	Key2      goaliyun.String
}
