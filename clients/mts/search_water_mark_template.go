package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SearchWaterMarkTemplateRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SearchWaterMarkTemplateRequest) Invoke(client goaliyun.Client) (*SearchWaterMarkTemplateResponse, error) {
	resp := &SearchWaterMarkTemplateResponse{}
	err := client.Request("mts", "SearchWaterMarkTemplate", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SearchWaterMarkTemplateResponse struct {
	RequestId             goaliyun.String
	TotalCount            goaliyun.Integer
	PageNumber            goaliyun.Integer
	PageSize              goaliyun.Integer
	WaterMarkTemplateList SearchWaterMarkTemplateWaterMarkTemplateList
}

type SearchWaterMarkTemplateWaterMarkTemplate struct {
	Id         goaliyun.String
	Name       goaliyun.String
	Width      goaliyun.String
	Height     goaliyun.String
	Dx         goaliyun.String
	Dy         goaliyun.String
	ReferPos   goaliyun.String
	Type       goaliyun.String
	State      goaliyun.String
	Timeline   SearchWaterMarkTemplateTimeline
	RatioRefer SearchWaterMarkTemplateRatioRefer
}

type SearchWaterMarkTemplateTimeline struct {
	Start    goaliyun.String
	Duration goaliyun.String
}

type SearchWaterMarkTemplateRatioRefer struct {
	Dx     goaliyun.String
	Dy     goaliyun.String
	Width  goaliyun.String
	Height goaliyun.String
}

type SearchWaterMarkTemplateWaterMarkTemplateList []SearchWaterMarkTemplateWaterMarkTemplate

func (list *SearchWaterMarkTemplateWaterMarkTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchWaterMarkTemplateWaterMarkTemplate)
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
