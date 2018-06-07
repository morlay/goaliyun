package mts

import (
	"github.com/morlay/goaliyun"
)

type AddWaterMarkTemplateRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Config               string `position:"Query" name:"Config"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddWaterMarkTemplateRequest) Invoke(client goaliyun.Client) (*AddWaterMarkTemplateResponse, error) {
	resp := &AddWaterMarkTemplateResponse{}
	err := client.Request("mts", "AddWaterMarkTemplate", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddWaterMarkTemplateResponse struct {
	RequestId         goaliyun.String
	WaterMarkTemplate AddWaterMarkTemplateWaterMarkTemplate
}

type AddWaterMarkTemplateWaterMarkTemplate struct {
	Id         goaliyun.String
	Name       goaliyun.String
	Width      goaliyun.String
	Height     goaliyun.String
	Dx         goaliyun.String
	Dy         goaliyun.String
	ReferPos   goaliyun.String
	Type       goaliyun.String
	State      goaliyun.String
	Timeline   AddWaterMarkTemplateTimeline
	RatioRefer AddWaterMarkTemplateRatioRefer
}

type AddWaterMarkTemplateTimeline struct {
	Start    goaliyun.String
	Duration goaliyun.String
}

type AddWaterMarkTemplateRatioRefer struct {
	Dx     goaliyun.String
	Dy     goaliyun.String
	Width  goaliyun.String
	Height goaliyun.String
}
