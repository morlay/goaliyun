package afs

import (
	"github.com/morlay/goaliyun"
)

type SetEarlyWarningRequest struct {
	TimeEnd         string `position:"Query" name:"TimeEnd"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	WarnOpen        string `position:"Query" name:"WarnOpen"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Channel         string `position:"Query" name:"Channel"`
	Title           string `position:"Query" name:"Title"`
	TimeOpen        string `position:"Query" name:"TimeOpen"`
	TimeBegin       string `position:"Query" name:"TimeBegin"`
	Frequency       string `position:"Query" name:"Frequency"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *SetEarlyWarningRequest) Invoke(client goaliyun.Client) (*SetEarlyWarningResponse, error) {
	resp := &SetEarlyWarningResponse{}
	err := client.Request("afs", "SetEarlyWarning", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetEarlyWarningResponse struct {
	RequestId goaliyun.String
	BizCode   goaliyun.String
}
