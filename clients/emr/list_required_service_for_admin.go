package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListRequiredServiceForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	EmrVersion      string `position:"Query" name:"EmrVersion"`
	ServiceNameList string `position:"Query" name:"ServiceNameList"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListRequiredServiceForAdminRequest) Invoke(client goaliyun.Client) (*ListRequiredServiceForAdminResponse, error) {
	resp := &ListRequiredServiceForAdminResponse{}
	err := client.Request("emr", "ListRequiredServiceForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListRequiredServiceForAdminResponse struct {
	RequestId        goaliyun.String
	NeedOtherService bool
	ServiceList      ListRequiredServiceForAdminServiceList
}

type ListRequiredServiceForAdminService struct {
	ServiceName        goaliyun.String
	ServiceDisplayName goaliyun.String
}

type ListRequiredServiceForAdminServiceList []ListRequiredServiceForAdminService

func (list *ListRequiredServiceForAdminServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRequiredServiceForAdminService)
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
