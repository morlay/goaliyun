package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyCasterLayoutRequest struct {
	BlendLists  *ModifyCasterLayoutBlendListList  `position:"Query" type:"Repeated" name:"BlendList"`
	AudioLayers *ModifyCasterLayoutAudioLayerList `position:"Query" type:"Repeated" name:"AudioLayer"`
	VideoLayers *ModifyCasterLayoutVideoLayerList `position:"Query" type:"Repeated" name:"VideoLayer"`
	CasterId    string                            `position:"Query" name:"CasterId"`
	MixLists    *ModifyCasterLayoutMixListList    `position:"Query" type:"Repeated" name:"MixList"`
	OwnerId     int64                             `position:"Query" name:"OwnerId"`
	LayoutId    string                            `position:"Query" name:"LayoutId"`
	RegionId    string                            `position:"Query" name:"RegionId"`
}

func (req *ModifyCasterLayoutRequest) Invoke(client goaliyun.Client) (*ModifyCasterLayoutResponse, error) {
	resp := &ModifyCasterLayoutResponse{}
	err := client.Request("live", "ModifyCasterLayout", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCasterLayoutAudioLayer struct {
	VolumeRate         float64 `name:"VolumeRate"`
	ValidChannel       string  `name:"ValidChannel"`
	FixedDelayDuration int64   `name:"FixedDelayDuration"`
}

type ModifyCasterLayoutVideoLayer struct {
	HeightNormalized    float64                                   `name:"HeightNormalized"`
	WidthNormalized     float64                                   `name:"WidthNormalized"`
	PositionRefer       string                                    `name:"PositionRefer"`
	PositionNormalizeds *ModifyCasterLayoutPositionNormalizedList `type:"Repeated" name:"PositionNormalized"`
	FixedDelayDuration  int64                                     `name:"FixedDelayDuration"`
}

type ModifyCasterLayoutResponse struct {
	RequestId goaliyun.String
	LayoutId  goaliyun.String
}

type ModifyCasterLayoutBlendListList []string

func (list *ModifyCasterLayoutBlendListList) UnmarshalJSON(data []byte) error {
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

type ModifyCasterLayoutAudioLayerList []ModifyCasterLayoutAudioLayer

func (list *ModifyCasterLayoutAudioLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyCasterLayoutAudioLayer)
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

type ModifyCasterLayoutVideoLayerList []ModifyCasterLayoutVideoLayer

func (list *ModifyCasterLayoutVideoLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyCasterLayoutVideoLayer)
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

type ModifyCasterLayoutMixListList []string

func (list *ModifyCasterLayoutMixListList) UnmarshalJSON(data []byte) error {
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

type ModifyCasterLayoutPositionNormalizedList []float64

func (list *ModifyCasterLayoutPositionNormalizedList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]float64)
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
