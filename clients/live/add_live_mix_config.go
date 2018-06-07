package live

import (
	"github.com/morlay/goaliyun"
)

type AddLiveMixConfigRequest struct {
	Template      string `position:"Query" name:"Template"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *AddLiveMixConfigRequest) Invoke(client goaliyun.Client) (*AddLiveMixConfigResponse, error) {
	resp := &AddLiveMixConfigResponse{}
	err := client.Request("live", "AddLiveMixConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveMixConfigResponse struct {
	RequestId goaliyun.String
}
