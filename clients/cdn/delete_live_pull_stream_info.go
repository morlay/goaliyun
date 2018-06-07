package cdn

import (
	"github.com/morlay/goaliyun"
)

type DeleteLivePullStreamInfoRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLivePullStreamInfoRequest) Invoke(client goaliyun.Client) (*DeleteLivePullStreamInfoResponse, error) {
	resp := &DeleteLivePullStreamInfoResponse{}
	err := client.Request("cdn", "DeleteLivePullStreamInfo", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLivePullStreamInfoResponse struct {
	RequestId goaliyun.String
}
