package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCasterLayoutsRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	LayoutId string `position:"Query" name:"LayoutId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeCasterLayoutsRequest) Invoke(client goaliyun.Client) (*DescribeCasterLayoutsResponse, error) {
	resp := &DescribeCasterLayoutsResponse{}
	err := client.Request("live", "DescribeCasterLayouts", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCasterLayoutsResponse struct {
	RequestId goaliyun.String
	Total     goaliyun.Integer
	Layouts   DescribeCasterLayoutsLayoutList
}

type DescribeCasterLayoutsLayout struct {
	LayoutId    goaliyun.String
	VideoLayers DescribeCasterLayoutsVideoLayerList
	AudioLayers DescribeCasterLayoutsAudioLayerList
	BlendList   DescribeCasterLayoutsBlendListList
	MixList     DescribeCasterLayoutsMixListList
}

type DescribeCasterLayoutsVideoLayer struct {
	HeightNormalized    goaliyun.Float
	WidthNormalized     goaliyun.Float
	PositionRefer       goaliyun.String
	FixedDelayDuration  goaliyun.Integer
	PositionNormalizeds DescribeCasterLayoutsPositionNormalizedList
}

type DescribeCasterLayoutsAudioLayer struct {
	VolumeRate         goaliyun.Float
	ValidChannel       goaliyun.String
	FixedDelayDuration goaliyun.Integer
}

type DescribeCasterLayoutsLayoutList []DescribeCasterLayoutsLayout

func (list *DescribeCasterLayoutsLayoutList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterLayoutsLayout)
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

type DescribeCasterLayoutsVideoLayerList []DescribeCasterLayoutsVideoLayer

func (list *DescribeCasterLayoutsVideoLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterLayoutsVideoLayer)
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

type DescribeCasterLayoutsAudioLayerList []DescribeCasterLayoutsAudioLayer

func (list *DescribeCasterLayoutsAudioLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterLayoutsAudioLayer)
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

type DescribeCasterLayoutsBlendListList []goaliyun.String

func (list *DescribeCasterLayoutsBlendListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type DescribeCasterLayoutsMixListList []goaliyun.String

func (list *DescribeCasterLayoutsMixListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type DescribeCasterLayoutsPositionNormalizedList []goaliyun.String

func (list *DescribeCasterLayoutsPositionNormalizedList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
