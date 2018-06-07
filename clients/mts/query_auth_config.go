package mts

import (
	"github.com/morlay/goaliyun"
)

type QueryAuthConfigRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryAuthConfigRequest) Invoke(client goaliyun.Client) (*QueryAuthConfigResponse, error) {
	resp := &QueryAuthConfigResponse{}
	err := client.Request("mts", "QueryAuthConfig", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryAuthConfigResponse struct {
	RequestId goaliyun.String
	Key1      goaliyun.String
	Key2      goaliyun.String
}
