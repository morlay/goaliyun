package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterServiceQuickLinkForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterServiceQuickLinkForAdminRequest) Invoke(client goaliyun.Client) (*ListClusterServiceQuickLinkForAdminResponse, error) {
	resp := &ListClusterServiceQuickLinkForAdminResponse{}
	err := client.Request("emr", "ListClusterServiceQuickLinkForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterServiceQuickLinkForAdminResponse struct {
	RequestId     goaliyun.String
	QuickLinkList ListClusterServiceQuickLinkForAdminQuickLinkList
}

type ListClusterServiceQuickLinkForAdminQuickLink struct {
	ServiceName        goaliyun.String
	ServiceDisplayName goaliyun.String
	QuickLinkAddress   goaliyun.String
	Protocol           goaliyun.String
}

type ListClusterServiceQuickLinkForAdminQuickLinkList []ListClusterServiceQuickLinkForAdminQuickLink

func (list *ListClusterServiceQuickLinkForAdminQuickLinkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceQuickLinkForAdminQuickLink)
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
