package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ExitStandbyRequest struct {
	InstanceIds          *ExitStandbyInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	ResourceOwnerAccount string                     `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string                     `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64                      `position:"Query" name:"OwnerId"`
	RegionId             string                     `position:"Query" name:"RegionId"`
}

func (req *ExitStandbyRequest) Invoke(client goaliyun.Client) (*ExitStandbyResponse, error) {
	resp := &ExitStandbyResponse{}
	err := client.Request("ess", "ExitStandby", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ExitStandbyResponse struct {
	RequestId goaliyun.String
}

type ExitStandbyInstanceIdList []string

func (list *ExitStandbyInstanceIdList) UnmarshalJSON(data []byte) error {
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
