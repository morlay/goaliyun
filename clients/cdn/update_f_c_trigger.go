package cdn

import (
	"github.com/morlay/goaliyun"
)

type UpdateFCTriggerRequest struct {
	Notes      string `position:"Query" name:"Notes"`
	TriggerARN string `position:"Query" name:"TriggerARN"`
	SourceARN  string `position:"Query" name:"SourceARN"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RoleARN    string `position:"Query" name:"RoleARN"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *UpdateFCTriggerRequest) Invoke(client goaliyun.Client) (*UpdateFCTriggerResponse, error) {
	resp := &UpdateFCTriggerResponse{}
	err := client.Request("cdn", "UpdateFCTrigger", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateFCTriggerResponse struct {
	RequestId goaliyun.String
}
