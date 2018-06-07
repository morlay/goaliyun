package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddCasterLayoutRequest struct {
	BlendLists  *AddCasterLayoutBlendListList  `position:"Query" type:"Repeated" name:"BlendList"`
	AudioLayers *AddCasterLayoutAudioLayerList `position:"Query" type:"Repeated" name:"AudioLayer"`
	VideoLayers *AddCasterLayoutVideoLayerList `position:"Query" type:"Repeated" name:"VideoLayer"`
	CasterId    string                         `position:"Query" name:"CasterId"`
	MixLists    *AddCasterLayoutMixListList    `position:"Query" type:"Repeated" name:"MixList"`
	OwnerId     int64                          `position:"Query" name:"OwnerId"`
	RegionId    string                         `position:"Query" name:"RegionId"`
}

func (req *AddCasterLayoutRequest) Invoke(client goaliyun.Client) (*AddCasterLayoutResponse, error) {
	resp := &AddCasterLayoutResponse{}
	err := client.Request("live", "AddCasterLayout", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCasterLayoutAudioLayer struct {
	VolumeRate         float64 `name:"VolumeRate"`
	ValidChannel       string  `name:"ValidChannel"`
	FixedDelayDuration int64   `name:"FixedDelayDuration"`
}

type AddCasterLayoutVideoLayer struct {
	HeightNormalized    float64                                `name:"HeightNormalized"`
	WidthNormalized     float64                                `name:"WidthNormalized"`
	PositionRefer       string                                 `name:"PositionRefer"`
	PositionNormalizeds *AddCasterLayoutPositionNormalizedList `type:"Repeated" name:"PositionNormalized"`
	FixedDelayDuration  int64                                  `name:"FixedDelayDuration"`
}

type AddCasterLayoutResponse struct {
	RequestId goaliyun.String
	LayoutId  goaliyun.String
}

type AddCasterLayoutBlendListList []string

func (list *AddCasterLayoutBlendListList) UnmarshalJSON(data []byte) error {
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

type AddCasterLayoutAudioLayerList []AddCasterLayoutAudioLayer

func (list *AddCasterLayoutAudioLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterLayoutAudioLayer)
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

type AddCasterLayoutVideoLayerList []AddCasterLayoutVideoLayer

func (list *AddCasterLayoutVideoLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterLayoutVideoLayer)
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

type AddCasterLayoutMixListList []string

func (list *AddCasterLayoutMixListList) UnmarshalJSON(data []byte) error {
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

type AddCasterLayoutPositionNormalizedList []float64

func (list *AddCasterLayoutPositionNormalizedList) UnmarshalJSON(data []byte) error {
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
