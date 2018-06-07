package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterOperationHostForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	OperationId     string `position:"Query" name:"OperationId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterOperationHostForAdminRequest) Invoke(client goaliyun.Client) (*ListClusterOperationHostForAdminResponse, error) {
	resp := &ListClusterOperationHostForAdminResponse{}
	err := client.Request("emr", "ListClusterOperationHostForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterOperationHostForAdminResponse struct {
	RequestId                goaliyun.String
	TotalCount               goaliyun.Integer
	PageNumber               goaliyun.Integer
	PageSize                 goaliyun.Integer
	ClusterOperationHostList ListClusterOperationHostForAdminClusterOperationHostList
}

type ListClusterOperationHostForAdminClusterOperationHost struct {
	HostId     goaliyun.String
	HostName   goaliyun.String
	Status     goaliyun.String
	Percentage goaliyun.String
}

type ListClusterOperationHostForAdminClusterOperationHostList []ListClusterOperationHostForAdminClusterOperationHost

func (list *ListClusterOperationHostForAdminClusterOperationHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationHostForAdminClusterOperationHost)
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
