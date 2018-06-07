package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDomainGroupListRequest struct {
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	DomainGroupName   string `position:"Query" name:"DomainGroupName"`
	Lang              string `position:"Query" name:"Lang"`
	ShowDeletingGroup string `position:"Query" name:"ShowDeletingGroup"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *QueryDomainGroupListRequest) Invoke(client goaliyun.Client) (*QueryDomainGroupListResponse, error) {
	resp := &QueryDomainGroupListResponse{}
	err := client.Request("domain", "QueryDomainGroupList", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDomainGroupListResponse struct {
	RequestId goaliyun.String
	Data      QueryDomainGroupListDomainGroupList
}

type QueryDomainGroupListDomainGroup struct {
	DomainGroupId     goaliyun.String
	DomainGroupName   goaliyun.String
	TotalNumber       goaliyun.Integer
	CreationDate      goaliyun.String
	ModificationDate  goaliyun.String
	DomainGroupStatus goaliyun.String
	BeingDeleted      bool
}

type QueryDomainGroupListDomainGroupList []QueryDomainGroupListDomainGroup

func (list *QueryDomainGroupListDomainGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDomainGroupListDomainGroup)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
