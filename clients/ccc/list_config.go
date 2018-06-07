package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListConfigRequest struct {
	InstanceId  string                    `position:"Query" name:"InstanceId"`
	ConfigItems *ListConfigConfigItemList `position:"Query" type:"Repeated" name:"ConfigItem"`
	RegionId    string                    `position:"Query" name:"RegionId"`
}

func (req *ListConfigRequest) Invoke(client goaliyun.Client) (*ListConfigResponse, error) {
	resp := &ListConfigResponse{}
	err := client.Request("ccc", "ListConfig", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListConfigResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	ConfigItems    ListConfigConfigItemList
}

type ListConfigConfigItem struct {
	Name  goaliyun.String
	Value goaliyun.String
}

type ListConfigConfigItemList []ListConfigConfigItem

func (list *ListConfigConfigItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListConfigConfigItem)
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
