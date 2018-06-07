package alidns

import (
	"github.com/morlay/goaliyun"
)

type UpdateDomainGroupRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	GroupId      string `position:"Query" name:"GroupId"`
	GroupName    string `position:"Query" name:"GroupName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *UpdateDomainGroupRequest) Invoke(client goaliyun.Client) (*UpdateDomainGroupResponse, error) {
	resp := &UpdateDomainGroupResponse{}
	err := client.Request("alidns", "UpdateDomainGroup", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateDomainGroupResponse struct {
	RequestId goaliyun.String
	GroupId   goaliyun.String
	GroupName goaliyun.String
}
