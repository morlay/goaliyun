package live

import (
	"github.com/morlay/goaliyun"
)

type StartMultipleStreamMixServiceRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	MixTemplate   string `position:"Query" name:"MixTemplate"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *StartMultipleStreamMixServiceRequest) Invoke(client goaliyun.Client) (*StartMultipleStreamMixServiceResponse, error) {
	resp := &StartMultipleStreamMixServiceResponse{}
	err := client.Request("live", "StartMultipleStreamMixService", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartMultipleStreamMixServiceResponse struct {
	RequestId goaliyun.String
}
