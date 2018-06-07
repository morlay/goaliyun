package alidns

import (
	"github.com/morlay/goaliyun"
)

type AddDomainGroupRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	GroupName    string `position:"Query" name:"GroupName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *AddDomainGroupRequest) Invoke(client goaliyun.Client) (*AddDomainGroupResponse, error) {
	resp := &AddDomainGroupResponse{}
	err := client.Request("alidns", "AddDomainGroup", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddDomainGroupResponse struct {
	RequestId goaliyun.String
	GroupId   goaliyun.String
	GroupName goaliyun.String
}
