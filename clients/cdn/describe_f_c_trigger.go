package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeFCTriggerRequest struct {
	TriggerARN string `position:"Query" name:"TriggerARN"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeFCTriggerRequest) Invoke(client goaliyun.Client) (*DescribeFCTriggerResponse, error) {
	resp := &DescribeFCTriggerResponse{}
	err := client.Request("cdn", "DescribeFCTrigger", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFCTriggerResponse struct {
	RequestId goaliyun.String
	FCTrigger DescribeFCTriggerFCTrigger
}

type DescribeFCTriggerFCTrigger struct {
	EventMetaName    goaliyun.String
	EventMetaVersion goaliyun.String
	TriggerARN       goaliyun.String
	RoleARN          goaliyun.String
	SourceArn        goaliyun.String
	Notes            goaliyun.String
}
