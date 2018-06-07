package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SetInstancesProtectionRequest struct {
	InstanceIds          *SetInstancesProtectionInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	ResourceOwnerAccount string                                `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string                                `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64                                 `position:"Query" name:"OwnerId"`
	ProtectedFromScaleIn string                                `position:"Query" name:"ProtectedFromScaleIn"`
	RegionId             string                                `position:"Query" name:"RegionId"`
}

func (req *SetInstancesProtectionRequest) Invoke(client goaliyun.Client) (*SetInstancesProtectionResponse, error) {
	resp := &SetInstancesProtectionResponse{}
	err := client.Request("ess", "SetInstancesProtection", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetInstancesProtectionResponse struct {
	RequestId goaliyun.String
}

type SetInstancesProtectionInstanceIdList []string

func (list *SetInstancesProtectionInstanceIdList) UnmarshalJSON(data []byte) error {
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
