package live

import (
	"github.com/morlay/goaliyun"
)

type StopMultipleStreamMixServiceRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *StopMultipleStreamMixServiceRequest) Invoke(client goaliyun.Client) (*StopMultipleStreamMixServiceResponse, error) {
	resp := &StopMultipleStreamMixServiceResponse{}
	err := client.Request("live", "StopMultipleStreamMixService", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopMultipleStreamMixServiceResponse struct {
	RequestId goaliyun.String
}
