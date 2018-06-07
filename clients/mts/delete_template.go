package mts

import (
	"github.com/morlay/goaliyun"
)

type DeleteTemplateRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TemplateId           string `position:"Query" name:"TemplateId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteTemplateRequest) Invoke(client goaliyun.Client) (*DeleteTemplateResponse, error) {
	resp := &DeleteTemplateResponse{}
	err := client.Request("mts", "DeleteTemplate", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteTemplateResponse struct {
	RequestId  goaliyun.String
	TemplateId goaliyun.String
}
