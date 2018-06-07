package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitTerrorismJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	TerrorismConfig      string `position:"Query" name:"TerrorismConfig"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitTerrorismJobRequest) Invoke(client goaliyun.Client) (*SubmitTerrorismJobResponse, error) {
	resp := &SubmitTerrorismJobResponse{}
	err := client.Request("mts", "SubmitTerrorismJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitTerrorismJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
