package cdn

import (
	"github.com/morlay/goaliyun"
)

type DeleteFCTriggerRequest struct {
	TriggerARN string `position:"Query" name:"TriggerARN"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeleteFCTriggerRequest) Invoke(client goaliyun.Client) (*DeleteFCTriggerResponse, error) {
	resp := &DeleteFCTriggerResponse{}
	err := client.Request("cdn", "DeleteFCTrigger", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFCTriggerResponse struct {
	RequestId goaliyun.String
}
