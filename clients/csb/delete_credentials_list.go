package csb

import (
	"github.com/morlay/goaliyun"
)

type DeleteCredentialsListRequest struct {
	Data        string `position:"Body" name:"Data"`
	IgnoreDauth string `position:"Query" name:"IgnoreDauth"`
	Force       string `position:"Query" name:"Force"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DeleteCredentialsListRequest) Invoke(client goaliyun.Client) (*DeleteCredentialsListResponse, error) {
	resp := &DeleteCredentialsListResponse{}
	err := client.Request("csb", "DeleteCredentialsList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCredentialsListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
