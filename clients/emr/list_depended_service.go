package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListDependedServiceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListDependedServiceRequest) Invoke(client goaliyun.Client) (*ListDependedServiceResponse, error) {
	resp := &ListDependedServiceResponse{}
	err := client.Request("emr", "ListDependedService", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListDependedServiceResponse struct {
	RequestId     goaliyun.String
	ExistServices bool
	ServiceList   ListDependedServiceServiceList
}

type ListDependedServiceService struct {
	ServiceName        goaliyun.String
	ServiceDisplayName goaliyun.String
}

type ListDependedServiceServiceList []ListDependedServiceService

func (list *ListDependedServiceServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListDependedServiceService)
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
