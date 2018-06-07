package domain

import (
	"github.com/morlay/goaliyun"
)

type SaveDomainGroupRequest struct {
	UserClientIp    string `position:"Query" name:"UserClientIp"`
	DomainGroupName string `position:"Query" name:"DomainGroupName"`
	Lang            string `position:"Query" name:"Lang"`
	DomainGroupId   int64  `position:"Query" name:"DomainGroupId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *SaveDomainGroupRequest) Invoke(client goaliyun.Client) (*SaveDomainGroupResponse, error) {
	resp := &SaveDomainGroupResponse{}
	err := client.Request("domain", "SaveDomainGroup", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveDomainGroupResponse struct {
	RequestId         goaliyun.String
	DomainGroupId     goaliyun.Integer
	DomainGroupName   goaliyun.String
	TotalNumber       goaliyun.Integer
	CreationDate      goaliyun.String
	ModificationDate  goaliyun.String
	DomainGroupStatus goaliyun.String
	BeingDeleted      bool
}
