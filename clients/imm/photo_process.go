package imm

import (
	"github.com/morlay/goaliyun"
)

type PhotoProcessRequest struct {
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	Project         string `position:"Query" name:"Project"`
	ExternalID      string `position:"Query" name:"ExternalID"`
	SrcUri          string `position:"Query" name:"SrcUri"`
	Style           string `position:"Query" name:"Style"`
	TgtUri          string `position:"Query" name:"TgtUri"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *PhotoProcessRequest) Invoke(client goaliyun.Client) (*PhotoProcessResponse, error) {
	resp := &PhotoProcessResponse{}
	err := client.Request("imm", "PhotoProcess", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PhotoProcessResponse struct {
	RequestId  goaliyun.String
	TaskId     goaliyun.String
	TgtLoc     goaliyun.String
	Status     goaliyun.String
	CreateTime goaliyun.String
	Percent    goaliyun.Integer
}
