package cdn

import (
	"github.com/morlay/goaliyun"
)

type AddLivePullStreamInfoRequest struct {
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

func (req *AddLivePullStreamInfoRequest) Invoke(client goaliyun.Client) (*AddLivePullStreamInfoResponse, error) {
	resp := &AddLivePullStreamInfoResponse{}
	err := client.Request("cdn", "AddLivePullStreamInfo", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLivePullStreamInfoResponse struct {
	RequestId goaliyun.String
}
