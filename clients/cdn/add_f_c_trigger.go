package cdn

import (
	"github.com/morlay/goaliyun"
)

type AddFCTriggerRequest struct {
	Notes            string `position:"Query" name:"Notes"`
	EventMetaVersion string `position:"Query" name:"EventMetaVersion"`
	TriggerARN       string `position:"Query" name:"TriggerARN"`
	SourceARN        string `position:"Query" name:"SourceARN"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	RoleARN          string `position:"Query" name:"RoleARN"`
	EventMetaName    string `position:"Query" name:"EventMetaName"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *AddFCTriggerRequest) Invoke(client goaliyun.Client) (*AddFCTriggerResponse, error) {
	resp := &AddFCTriggerResponse{}
	err := client.Request("cdn", "AddFCTrigger", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddFCTriggerResponse struct {
	RequestId goaliyun.String
}
