package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListRequiredServiceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	EmrVersion      string `position:"Query" name:"EmrVersion"`
	ServiceNameList string `position:"Query" name:"ServiceNameList"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListRequiredServiceRequest) Invoke(client goaliyun.Client) (*ListRequiredServiceResponse, error) {
	resp := &ListRequiredServiceResponse{}
	err := client.Request("emr", "ListRequiredService", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListRequiredServiceResponse struct {
	RequestId        goaliyun.String
	NeedOtherService bool
	ServiceList      ListRequiredServiceServiceList
}

type ListRequiredServiceService struct {
	ServiceName        goaliyun.String
	ServiceDisplayName goaliyun.String
}

type ListRequiredServiceServiceList []ListRequiredServiceService

func (list *ListRequiredServiceServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRequiredServiceService)
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
