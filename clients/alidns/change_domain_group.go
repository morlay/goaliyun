package alidns

import (
	"github.com/morlay/goaliyun"
)

type ChangeDomainGroupRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	GroupId      string `position:"Query" name:"GroupId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ChangeDomainGroupRequest) Invoke(client goaliyun.Client) (*ChangeDomainGroupResponse, error) {
	resp := &ChangeDomainGroupResponse{}
	err := client.Request("alidns", "ChangeDomainGroup", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ChangeDomainGroupResponse struct {
	RequestId goaliyun.String
	GroupId   goaliyun.String
	GroupName goaliyun.String
}
