package hpc

import (
	"github.com/morlay/goaliyun"
)

type AuthorizeSecurityGroupRequest struct {
	TOKEN    string `position:"Query" name:"TOKEN"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Policy   string `position:"Query" name:"Policy"`
	Priority string `position:"Query" name:"Priority"`
	NicType  string `position:"Query" name:"NicType"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *AuthorizeSecurityGroupRequest) Invoke(client goaliyun.Client) (*AuthorizeSecurityGroupResponse, error) {
	resp := &AuthorizeSecurityGroupResponse{}
	err := client.Request("hpc", "AuthorizeSecurityGroup", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AuthorizeSecurityGroupResponse struct {
	RequestId goaliyun.String
}
