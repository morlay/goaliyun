package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ToggleFeaturesRequest struct {
	DisabledFeaturess *ToggleFeaturesDisabledFeaturesList `position:"Query" type:"Repeated" name:"DisabledFeatures"`
	StoreName         string                              `position:"Query" name:"StoreName"`
	EnabledFeaturess  *ToggleFeaturesEnabledFeaturesList  `position:"Query" type:"Repeated" name:"EnabledFeatures"`
	RegionId          string                              `position:"Query" name:"RegionId"`
}

func (req *ToggleFeaturesRequest) Invoke(client goaliyun.Client) (*ToggleFeaturesResponse, error) {
	resp := &ToggleFeaturesResponse{}
	err := client.Request("cloudphoto", "ToggleFeatures", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ToggleFeaturesResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}

type ToggleFeaturesDisabledFeaturesList []string

func (list *ToggleFeaturesDisabledFeaturesList) UnmarshalJSON(data []byte) error {
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

type ToggleFeaturesEnabledFeaturesList []string

func (list *ToggleFeaturesEnabledFeaturesList) UnmarshalJSON(data []byte) error {
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
