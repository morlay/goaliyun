package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveMixConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveMixConfigRequest) Invoke(client goaliyun.Client) (*DeleteLiveMixConfigResponse, error) {
	resp := &DeleteLiveMixConfigResponse{}
	err := client.Request("live", "DeleteLiveMixConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveMixConfigResponse struct {
	RequestId goaliyun.String
}
