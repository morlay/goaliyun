package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyNetworkInterfaceAttributeRequest struct {
	ResourceOwnerId      int64                                               `position:"Query" name:"ResourceOwnerId"`
	SecurityGroupIds     *ModifyNetworkInterfaceAttributeSecurityGroupIdList `position:"Query" type:"Repeated" name:"SecurityGroupId"`
	Description          string                                              `position:"Query" name:"Description"`
	NetworkInterfaceName string                                              `position:"Query" name:"NetworkInterfaceName"`
	ResourceOwnerAccount string                                              `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                              `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                               `position:"Query" name:"OwnerId"`
	NetworkInterfaceId   string                                              `position:"Query" name:"NetworkInterfaceId"`
	RegionId             string                                              `position:"Query" name:"RegionId"`
}

func (req *ModifyNetworkInterfaceAttributeRequest) Invoke(client goaliyun.Client) (*ModifyNetworkInterfaceAttributeResponse, error) {
	resp := &ModifyNetworkInterfaceAttributeResponse{}
	err := client.Request("ecs", "ModifyNetworkInterfaceAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyNetworkInterfaceAttributeResponse struct {
	RequestId goaliyun.String
}

type ModifyNetworkInterfaceAttributeSecurityGroupIdList []string

func (list *ModifyNetworkInterfaceAttributeSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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
