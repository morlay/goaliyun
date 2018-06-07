package domain

import (
	"github.com/morlay/goaliyun"
)

type DeleteDomainGroupRequest struct {
	UserClientIp  string `position:"Query" name:"UserClientIp"`
	Lang          string `position:"Query" name:"Lang"`
	DomainGroupId int64  `position:"Query" name:"DomainGroupId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteDomainGroupRequest) Invoke(client goaliyun.Client) (*DeleteDomainGroupResponse, error) {
	resp := &DeleteDomainGroupResponse{}
	err := client.Request("domain", "DeleteDomainGroup", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDomainGroupResponse struct {
	RequestId goaliyun.String
}
