package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateWaterMarkTemplateRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	WaterMarkTemplateId  string `position:"Query" name:"WaterMarkTemplateId"`
	Config               string `position:"Query" name:"Config"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateWaterMarkTemplateRequest) Invoke(client goaliyun.Client) (*UpdateWaterMarkTemplateResponse, error) {
	resp := &UpdateWaterMarkTemplateResponse{}
	err := client.Request("mts", "UpdateWaterMarkTemplate", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateWaterMarkTemplateResponse struct {
	RequestId         goaliyun.String
	WaterMarkTemplate UpdateWaterMarkTemplateWaterMarkTemplate
}

type UpdateWaterMarkTemplateWaterMarkTemplate struct {
	Id         goaliyun.String
	Name       goaliyun.String
	Width      goaliyun.String
	Height     goaliyun.String
	Dx         goaliyun.String
	Dy         goaliyun.String
	ReferPos   goaliyun.String
	Type       goaliyun.String
	State      goaliyun.String
	Timeline   UpdateWaterMarkTemplateTimeline
	RatioRefer UpdateWaterMarkTemplateRatioRefer
}

type UpdateWaterMarkTemplateTimeline struct {
	Start    goaliyun.String
	Duration goaliyun.String
}

type UpdateWaterMarkTemplateRatioRefer struct {
	Dx     goaliyun.String
	Dy     goaliyun.String
	Width  goaliyun.String
	Height goaliyun.String
}
