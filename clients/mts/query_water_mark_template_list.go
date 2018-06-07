package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryWaterMarkTemplateListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	WaterMarkTemplateIds string `position:"Query" name:"WaterMarkTemplateIds"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryWaterMarkTemplateListRequest) Invoke(client goaliyun.Client) (*QueryWaterMarkTemplateListResponse, error) {
	resp := &QueryWaterMarkTemplateListResponse{}
	err := client.Request("mts", "QueryWaterMarkTemplateList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryWaterMarkTemplateListResponse struct {
	RequestId             goaliyun.String
	WaterMarkTemplateList QueryWaterMarkTemplateListWaterMarkTemplateList
	NonExistWids          QueryWaterMarkTemplateListNonExistWidList
}

type QueryWaterMarkTemplateListWaterMarkTemplate struct {
	Id         goaliyun.String
	Name       goaliyun.String
	Width      goaliyun.String
	Height     goaliyun.String
	Dx         goaliyun.String
	Dy         goaliyun.String
	ReferPos   goaliyun.String
	Type       goaliyun.String
	State      goaliyun.String
	Timeline   QueryWaterMarkTemplateListTimeline
	RatioRefer QueryWaterMarkTemplateListRatioRefer
}

type QueryWaterMarkTemplateListTimeline struct {
	Start    goaliyun.String
	Duration goaliyun.String
}

type QueryWaterMarkTemplateListRatioRefer struct {
	Dx     goaliyun.String
	Dy     goaliyun.String
	Width  goaliyun.String
	Height goaliyun.String
}

type QueryWaterMarkTemplateListWaterMarkTemplateList []QueryWaterMarkTemplateListWaterMarkTemplate

func (list *QueryWaterMarkTemplateListWaterMarkTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryWaterMarkTemplateListWaterMarkTemplate)
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

type QueryWaterMarkTemplateListNonExistWidList []goaliyun.String

func (list *QueryWaterMarkTemplateListNonExistWidList) UnmarshalJSON(data []byte) error {
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
