package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCasterComponentsRequest struct {
	ComponentId string `position:"Query" name:"ComponentId"`
	CasterId    string `position:"Query" name:"CasterId"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DescribeCasterComponentsRequest) Invoke(client goaliyun.Client) (*DescribeCasterComponentsResponse, error) {
	resp := &DescribeCasterComponentsResponse{}
	err := client.Request("live", "DescribeCasterComponents", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCasterComponentsResponse struct {
	RequestId  goaliyun.String
	Total      goaliyun.Integer
	Components DescribeCasterComponentsComponentList
}

type DescribeCasterComponentsComponent struct {
	ComponentId         goaliyun.String
	ComponentName       goaliyun.String
	LocationId          goaliyun.String
	ComponentType       goaliyun.String
	Effect              goaliyun.String
	ComponentLayer      DescribeCasterComponentsComponentLayer
	TextLayerContent    DescribeCasterComponentsTextLayerContent
	ImageLayerContent   DescribeCasterComponentsImageLayerContent
	CaptionLayerContent DescribeCasterComponentsCaptionLayerContent
}

type DescribeCasterComponentsComponentLayer struct {
	HeightNormalized    goaliyun.Float
	WidthNormalized     goaliyun.Float
	PositionRefer       goaliyun.String
	PositionNormalizeds DescribeCasterComponentsPositionNormalizedList
}

type DescribeCasterComponentsTextLayerContent struct {
	Text                  goaliyun.String
	Color                 goaliyun.String
	FontName              goaliyun.String
	SizeNormalized        goaliyun.Float
	BorderWidthNormalized goaliyun.Float
	BorderColor           goaliyun.String
}

type DescribeCasterComponentsImageLayerContent struct {
	MaterialId goaliyun.String
}

type DescribeCasterComponentsCaptionLayerContent struct {
	LocationId            goaliyun.String
	PtsOffset             goaliyun.Integer
	WordsCount            goaliyun.Integer
	Color                 goaliyun.String
	FontName              goaliyun.String
	SizeNormalized        goaliyun.Float
	BorderWidthNormalized goaliyun.Float
	BorderColor           goaliyun.String
}

type DescribeCasterComponentsComponentList []DescribeCasterComponentsComponent

func (list *DescribeCasterComponentsComponentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterComponentsComponent)
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

type DescribeCasterComponentsPositionNormalizedList []goaliyun.String

func (list *DescribeCasterComponentsPositionNormalizedList) UnmarshalJSON(data []byte) error {
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
