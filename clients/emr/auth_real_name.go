package emr

import (
	"github.com/morlay/goaliyun"
)

type AuthRealNameRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *AuthRealNameRequest) Invoke(client goaliyun.Client) (*AuthRealNameResponse, error) {
	resp := &AuthRealNameResponse{}
	err := client.Request("emr", "AuthRealName", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AuthRealNameResponse struct {
	RequestId goaliyun.String
	Success   bool
}
