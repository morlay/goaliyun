package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetForceRedirectConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	RedirectType  string `position:"Query" name:"RedirectType"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetForceRedirectConfigRequest) Invoke(client goaliyun.Client) (*SetForceRedirectConfigResponse, error) {
	resp := &SetForceRedirectConfigResponse{}
	err := client.Request("cdn", "SetForceRedirectConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetForceRedirectConfigResponse struct {
	RequestId goaliyun.String
}
