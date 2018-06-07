package alidns

import (
	"github.com/morlay/goaliyun"
)

type DeleteDomainGroupRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	GroupId      string `position:"Query" name:"GroupId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteDomainGroupRequest) Invoke(client goaliyun.Client) (*DeleteDomainGroupResponse, error) {
	resp := &DeleteDomainGroupResponse{}
	err := client.Request("alidns", "DeleteDomainGroup", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDomainGroupResponse struct {
	RequestId goaliyun.String
	GroupName goaliyun.String
}
