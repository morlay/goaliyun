package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InstallCloudAssistantRequest struct {
	ResourceOwnerId      int64                                `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                               `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                               `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                `position:"Query" name:"OwnerId"`
	InstanceIds          *InstallCloudAssistantInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	RegionId             string                               `position:"Query" name:"RegionId"`
}

func (req *InstallCloudAssistantRequest) Invoke(client goaliyun.Client) (*InstallCloudAssistantResponse, error) {
	resp := &InstallCloudAssistantResponse{}
	err := client.Request("ecs", "InstallCloudAssistant", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InstallCloudAssistantResponse struct {
	RequestId goaliyun.String
}

type InstallCloudAssistantInstanceIdList []string

func (list *InstallCloudAssistantInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
