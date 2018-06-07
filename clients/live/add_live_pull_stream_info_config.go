package live

import (
	"github.com/morlay/goaliyun"
)

type AddLivePullStreamInfoConfigRequest struct {
	SourceUrl     string `position:"Query" name:"SourceUrl"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *AddLivePullStreamInfoConfigRequest) Invoke(client goaliyun.Client) (*AddLivePullStreamInfoConfigResponse, error) {
	resp := &AddLivePullStreamInfoConfigResponse{}
	err := client.Request("live", "AddLivePullStreamInfoConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLivePullStreamInfoConfigResponse struct {
	RequestId goaliyun.String
}
