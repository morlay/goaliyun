package hpc

import (
	"github.com/morlay/goaliyun"
)

type RevokeSecurityGroupRequest struct {
	TOKEN    string `position:"Query" name:"TOKEN"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Policy   string `position:"Query" name:"Policy"`
	Priority string `position:"Query" name:"Priority"`
	NicType  string `position:"Query" name:"NicType"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *RevokeSecurityGroupRequest) Invoke(client goaliyun.Client) (*RevokeSecurityGroupResponse, error) {
	resp := &RevokeSecurityGroupResponse{}
	err := client.Request("hpc", "RevokeSecurityGroup", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RevokeSecurityGroupResponse struct {
	RequestId goaliyun.String
}
