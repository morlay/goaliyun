package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterServiceQuickLinkRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterServiceQuickLinkRequest) Invoke(client goaliyun.Client) (*ListClusterServiceQuickLinkResponse, error) {
	resp := &ListClusterServiceQuickLinkResponse{}
	err := client.Request("emr", "ListClusterServiceQuickLink", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterServiceQuickLinkResponse struct {
	RequestId     goaliyun.String
	QuickLinkList ListClusterServiceQuickLinkQuickLinkList
}

type ListClusterServiceQuickLinkQuickLink struct {
	ServiceName        goaliyun.String
	ServiceDisplayName goaliyun.String
	QuickLinkAddress   goaliyun.String
	Protocol           goaliyun.String
	Port               goaliyun.String
	Type               goaliyun.String
}

type ListClusterServiceQuickLinkQuickLinkList []ListClusterServiceQuickLinkQuickLink

func (list *ListClusterServiceQuickLinkQuickLinkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceQuickLinkQuickLink)
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
