package mts

import (
	"github.com/morlay/goaliyun"
)

type DeleteWaterMarkTemplateRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	WaterMarkTemplateId  string `position:"Query" name:"WaterMarkTemplateId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteWaterMarkTemplateRequest) Invoke(client goaliyun.Client) (*DeleteWaterMarkTemplateResponse, error) {
	resp := &DeleteWaterMarkTemplateResponse{}
	err := client.Request("mts", "DeleteWaterMarkTemplate", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteWaterMarkTemplateResponse struct {
	RequestId           goaliyun.String
	WaterMarkTemplateId goaliyun.String
}
