package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLivePullStreamInfoConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLivePullStreamInfoConfigRequest) Invoke(client goaliyun.Client) (*DeleteLivePullStreamInfoConfigResponse, error) {
	resp := &DeleteLivePullStreamInfoConfigResponse{}
	err := client.Request("live", "DeleteLivePullStreamInfoConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLivePullStreamInfoConfigResponse struct {
	RequestId goaliyun.String
}
