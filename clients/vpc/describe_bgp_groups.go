package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBgpGroupsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBgpGroupsRequest) Invoke(client goaliyun.Client) (*DescribeBgpGroupsResponse, error) {
	resp := &DescribeBgpGroupsResponse{}
	err := client.Request("vpc", "DescribeBgpGroups", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBgpGroupsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	BgpGroups  DescribeBgpGroupsBgpGroupList
}

type DescribeBgpGroupsBgpGroup struct {
	Name        goaliyun.String
	Description goaliyun.String
	BgpGroupId  goaliyun.String
	PeerAsn     goaliyun.String
	AuthKey     goaliyun.String
	RouterId    goaliyun.String
	Status      goaliyun.String
	Keepalive   goaliyun.String
	LocalAsn    goaliyun.String
	Hold        goaliyun.String
	IsFake      goaliyun.String
	RouteLimit  goaliyun.String
	RegionId    goaliyun.String
}

type DescribeBgpGroupsBgpGroupList []DescribeBgpGroupsBgpGroup

func (list *DescribeBgpGroupsBgpGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBgpGroupsBgpGroup)
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
