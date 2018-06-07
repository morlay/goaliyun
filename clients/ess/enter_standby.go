package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type EnterStandbyRequest struct {
	InstanceIds          *EnterStandbyInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	ResourceOwnerAccount string                      `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string                      `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64                       `position:"Query" name:"OwnerId"`
	RegionId             string                      `position:"Query" name:"RegionId"`
}

func (req *EnterStandbyRequest) Invoke(client goaliyun.Client) (*EnterStandbyResponse, error) {
	resp := &EnterStandbyResponse{}
	err := client.Request("ess", "EnterStandby", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EnterStandbyResponse struct {
	RequestId goaliyun.String
}

type EnterStandbyInstanceIdList []string

func (list *EnterStandbyInstanceIdList) UnmarshalJSON(data []byte) error {
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
