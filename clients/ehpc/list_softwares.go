package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListSoftwaresRequest struct {
	EhpcVersion string `position:"Query" name:"EhpcVersion"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ListSoftwaresRequest) Invoke(client goaliyun.Client) (*ListSoftwaresResponse, error) {
	resp := &ListSoftwaresResponse{}
	err := client.Request("ehpc", "ListSoftwares", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListSoftwaresResponse struct {
	RequestId goaliyun.String
	Softwares ListSoftwaresSoftwareInfoList
}

type ListSoftwaresSoftwareInfo struct {
	EhpcVersion      goaliyun.String
	OsTag            goaliyun.String
	SchedulerType    goaliyun.String
	SchedulerVersion goaliyun.String
	AccountType      goaliyun.String
	AccountVersion   goaliyun.String
	Applications     ListSoftwaresApplicationInfoList
}

type ListSoftwaresApplicationInfo struct {
	Tag      goaliyun.String
	Name     goaliyun.String
	Version  goaliyun.String
	Required bool
}

type ListSoftwaresSoftwareInfoList []ListSoftwaresSoftwareInfo

func (list *ListSoftwaresSoftwareInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSoftwaresSoftwareInfo)
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

type ListSoftwaresApplicationInfoList []ListSoftwaresApplicationInfo

func (list *ListSoftwaresApplicationInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSoftwaresApplicationInfo)
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
