package live

import (
	"github.com/morlay/goaliyun"
)

type RemoveMultipleStreamMixServiceRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	MixStreamName string `position:"Query" name:"MixStreamName"`
	MixDomainName string `position:"Query" name:"MixDomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	MixAppName    string `position:"Query" name:"MixAppName"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *RemoveMultipleStreamMixServiceRequest) Invoke(client goaliyun.Client) (*RemoveMultipleStreamMixServiceResponse, error) {
	resp := &RemoveMultipleStreamMixServiceResponse{}
	err := client.Request("live", "RemoveMultipleStreamMixService", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveMultipleStreamMixServiceResponse struct {
	RequestId goaliyun.String
}
