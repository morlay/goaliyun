package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type RegisterTagRequest struct {
	StoreName string `position:"Query" name:"StoreName"`
	Text      string `position:"Query" name:"Text"`
	TagKey    string `position:"Query" name:"TagKey"`
	Lang      string `position:"Query" name:"Lang"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *RegisterTagRequest) Invoke(client goaliyun.Client) (*RegisterTagResponse, error) {
	resp := &RegisterTagResponse{}
	err := client.Request("cloudphoto", "RegisterTag", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RegisterTagResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}
