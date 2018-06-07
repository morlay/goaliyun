package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainGroupsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainGroupsRequest) Invoke(client goaliyun.Client) (*DescribeDomainGroupsResponse, error) {
	resp := &DescribeDomainGroupsResponse{}
	err := client.Request("alidns", "DescribeDomainGroups", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainGroupsResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	DomainGroups DescribeDomainGroupsDomainGroupList
}

type DescribeDomainGroupsDomainGroup struct {
	GroupId   goaliyun.String
	GroupName goaliyun.String
}

type DescribeDomainGroupsDomainGroupList []DescribeDomainGroupsDomainGroup

func (list *DescribeDomainGroupsDomainGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainGroupsDomainGroup)
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
